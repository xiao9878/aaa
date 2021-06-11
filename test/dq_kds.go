package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net"
    "strconv"
    "strings"
    "time"
)

// Socket接口数据样式
type NotifyS struct {
    StoreCode string `json:"store_code"`
    OrderNum  string `json:"order_num"`
    TakeNum   string `json:"take_num"`
    OrderDate string `json:"order_date"`
    Status    string `json:"status"`
}

func main() {
    fmt.Println("本地ip")
    GetLocalIP()
    var (
        portStr string
        n    int
        err  error
    )
    for  {
        log.Print("请输入端口号:")
        n, err = fmt.Scan(&portStr)
        if _, err = strconv.Atoi(portStr); err != nil {
            log.Printf("请输入正确的端口号:%v\n", err)
            continue
        }
        if n > 0 {
            break
        }
    }
    listen, err := net.Listen("tcp", ":"+portStr)
    if err != nil {
        log.Fatalf("端口监听失败:%v", err)
    }
    log.Println("监听端口:",portStr)
    defer listen.Close()
    for {
        fmt.Println("-------------------------------")
        conn, err := listen.Accept()
        if err != nil {
            log.Fatal(err)
        }
        r := make([]byte, 0)
        buf := make([]byte, 20)
        for {
            n, err := conn.Read(buf)
            if n == 0 || err == io.EOF {
                break
            }
            if err != nil {
                fmt.Println(err)
                continue
            }
            for i := 0; i < n; i++ {
                r = append(r, buf[i])
            }
            //for _, item := range buf {
            //    r = append(r,item)
            //}
        }
        fmt.Printf("当前时间：%v从客户端接收的数据：%v\n", time.Now().Format("2006-01-02 15:04:05"), string(r))
        var info NotifyS
        if err = json.Unmarshal(r, &info); err != nil {
            fmt.Printf("Unmarshal fail:%v", err)
        }
        fmt.Printf("门店代码:%v\t取餐号:%v\n下单时间:%v\t订单状态:%v\n", info.StoreCode, info.TakeNum, info.OrderDate, IF(info.Status == "0", "制作中", "待取餐"))
    }
}

func IF(b bool, v1, v2 string) string {
    if b {
        return v1
    }
    return v2
}

func GetLocalIP() {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        fmt.Println(err)
    }
    for _, address := range addrs {
        // 检查ip地址判断是否回环地址
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil && strings.Contains(ipnet.IP.String(), "192.168.") {
                fmt.Println(ipnet.IP.String())
            }
        }
    }
}
