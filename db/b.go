package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "gorm.io/driver/postgres"
    "log"
    "net/http"

    _ "github.com/lib/pq"
    "github.com/olekukonko/tablewriter"
    "github.com/robfig/cron/v3"
    "github.com/spf13/cast"
    "gorm.io/gorm"
)


func Work1() {

    fmt.Printf("Database Connected:%v\r\n", db)

    var ret []map[string]interface{}
    db.Raw("select * from metadata_sync_log where sync_status = 5 order by created desc limit 50").Find(&ret)
    if ret == nil {
        return
    }
    //m := make(map[interface{}]interface{})
    //for _, item := range ret {
    //    m[item["id"]] = item
    //}
    //time.Sleep(time.Minute)
    //
    //var ret1 []map[string]interface{}
    //db.Raw("select * from metadata_sync_log where sync_status = 0 order by created desc limit 50").Find(&ret1)
    //if ret == nil {
    //    return
    //}
    //var result []map[string]interface{}
    //for _, item := range ret1 {
    //    if _, b := m[item["id"]]; b {
    //        result = append(result, item)
    //    }
    //}
    var data [][]string = [][]string{}

    for _, row := range ret {
        item := []string{
            cast.ToString(row["record_type"]),
            cast.ToString(row["org_type"]),
            cast.ToString(row["action"]),
            //cast.ToString(row["created"]),
            cast.ToTime(row["created"]).Format("2006-01-02 15:04:05"),
            //cast.ToString(row["sync_status"]),
            cast.ToString(row["err_msg"]),
        }
        data = append(data, item)
    }
    var buf bytes.Buffer
    table := tablewriter.NewWriter(&buf)
    table.SetHeader([]string{"修改类型", "表名", "操作", "同步时间", "错误信息"})
    table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
    table.SetCenterSeparator("|")
    table.AppendBulk(data) // Add Bulk Data
    table.Render()

    url := "https://oapi.dingtalk.com/robot/send?access_token=25a43dce1879907c1d80f330b5f67317cfb23a516253558110e28dec4bd64716"
    tb := string(buf.Bytes())
    var body = map[string]interface{}{
        "msgtype": "markdown",
        "markdown": map[string]interface{}{
            "title": "主档同步告警",
            "text":  fmt.Sprintf("# 异常数据 \r\n%v", tb),
        },
    }
    jsonstr, _ := json.Marshal(body)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonstr))
    // req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}
func ConnectDb() (*gorm.DB, error) {
    conn := "postgresql://lelecha_pos_prod:GBZQzRe2Eddjas9osuqj@pc-uf6ijoy6gvu5uem85.pg.polardb.rds.aliyuncs.com:1921/lelecha_pos_metadata_entity"

    return gorm.Open(postgres.Open(conn))
}
var (
    db *gorm.DB
)

func main() {
    // 连接数据库
    var err error
    db, err = ConnectDb()
    if err != nil || db == nil {
        log.Fatal(err)
    }
    cron := cron.New()
    _, err = cron.AddFunc("* * * * ?", Work1) // 0 0 8,10,12,14,16,18,20,22 * ?
    if err != nil {
        log.Fatalf("corn err:%v\n", err)
    }
    cron.Start()
    // 健康监查接口
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("pong"))
    })
    if err := http.ListenAndServe(":8888", nil); err != nil {
        log.Fatalf("listen fail:%v\n", err)
        return
    }
}
