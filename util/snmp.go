package util

import (
    g "github.com/gosnmp/gosnmp"
    "github.com/pkg/errors"
    "reflect"
    "strconv"
    "strings"
)

const (
    PRINTER_STATUS = "1.3.6.1.4.1.2435.3.3.9.1.6.1.0"
)

type status int

const (
    STATUS_OK           = iota //正常-待打印
    STATUS_CONNECT_FAIL        //连接异常
    STATUS_PRINTING            //正常-打印中
    STATUS_LID_OPEND           //盖子被打开
    STATUS_NOPAPER             //没纸
    STATUS_UNDEFINED           // 未知错误
)

var (
    statusMap = map[status]string{
        STATUS_OK:           `ok`,
        STATUS_CONNECT_FAIL: `Abnormal connection`,
        STATUS_PRINTING:     `Printing`,
        STATUS_LID_OPEND:    `Cover open error`,
        STATUS_NOPAPER:      `Error occurred`,
        STATUS_UNDEFINED:    `unknown error`,
    }
    // NORMAL_STATUS 正常状态的响应
    normalStatus = []string{"80", "20", "42", "35", "36", "30", "4", "0", "0", "0", "28", "4b", "0", "0", "3f", "0", "0", "3c", "0", "0", "0", "0", "0", "0", "0"}
    needWatch    = []int{9, 10, 19, 20, 23}
    // 打印机状态列举
    m = map[status]map[int]string{
        STATUS_OK:        {9: "0", 10: "0", 19: "0", 20: "0", 23: "0"},  // 正常情况
        STATUS_PRINTING:  {9: "0", 10: "0", 19: "0", 20: "1", 23: "0"},  // 正在打印
        STATUS_LID_OPEND: {9: "0", 10: "10", 19: "2", 20: "0", 23: "0"}, // 开盖
        STATUS_NOPAPER:   {9: "2", 10: "0", 19: "2", 20: "1", 23: "0"},  // TODO 没纸暂时不好判断
    }
)

func getPrinterStatus(code status, l []string) *PrinterStatus {
    return &PrinterStatus{
        Code:   code,
        Msg:    code.msg(),
        Detail: strings.Join(l, " "),
    }
}

// PrinterStatus 打印机状态检查
type PrinterStatus struct {
    Code   status `json:"code"`
    Msg    string `json:"msg"`
    Detail string `json:"detail"`
}

func (s status) msg() string {
    if v, ok := statusMap[s]; ok {
        return v
    }
    return statusMap[STATUS_UNDEFINED]
}

// Get 获取打印机状态
func Get(ip string) (*PrinterStatus, error) {
    g.Default.Target = ip
    err := g.Default.Connect()
    if err != nil {
        return nil, err
    }
    defer g.Default.Conn.Close()

    result, err := g.Default.Get([]string{PRINTER_STATUS}) // Get() accepts up to g.MAX_OIDS
    if err != nil {
        return nil, err
    }

    if len(result.Variables) != 1 {
        return nil, errors.New("get msg error")
    }
    var l []string
    if _, ok := result.Variables[0].Value.([]byte); !ok {
        return nil, errors.New("trans err")
    }
    // 转16进制
    for _, item := range result.Variables[0].Value.([]byte) {
        s := strconv.FormatInt(int64(item), 16)
        l = append(l, s)
    }
    l = l[:len(normalStatus)] // 只取25位
    data, err := parseData(l)
    if err != nil {
        return nil, err
    }
    return getPrinterStatus(match(data), l), nil
}

// parseData 与正常返回相比较，返回不一致的位置和数据
func parseData(l []string) (result map[int]string, err error) {
    result = make(map[int]string)
    for _, item := range needWatch {
        idx := item - 1
        if idx < 0 {
            result = nil
            err = errors.New("idx must more than zero")
            return
        }
        result[item] = l[idx]
    }
    return
}

// 规则匹配
func match(data map[int]string) status {
    for code, item := range m {
        if reflect.DeepEqual(data, item) {
            return code
        }
    }
    return STATUS_UNDEFINED
}
