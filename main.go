package main

import (
    "aaa/util"
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net"
    "strconv"
    "sync"
    "time"
)

type Extend struct {
    Status         string `json:"status"`
    RefundTicketId string `json:"refund_ticket_id"`
    RefundId       string `json:"refund_id"`
    Extend         struct {
        TakeawayInfo struct {
            DeliveryName string `json:"delivery_name"`
            Source       string `json:"source"`
            OrderStatus  int    `json:"order_status"`
            TpOrderId    string `json:"tp_order_id"`
            TpPriOrderId int    `json:"tp_pri_order_id"`
            TakeMealSn   string `json:"take_meal_sn"`
        } `json:"takeaway_info"`
    } `json:"extend"`
}

var m sync.RWMutex

func GO() {
    go func() {
        for {
            fmt.Println("111")
            time.Sleep(time.Second * 10)
        }
    }()
    for {
        fmt.Println("222")
        time.Sleep(time.Second)
    }
}
func F() {
    str := `{"status":"1","extend":{"takeaway_info":{"delivery_name":"awdw1","tp_order_id":"1231","tp_pri_order_id":231}}}`
    var to Extend
    if err := json.Unmarshal([]byte(str), &to); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%v", to)
}

type AAA struct {
    Data []byte
}
type NotifyS struct {
    StoreCode string `json:"store_code"`
    OrderNum  string `json:"order_num"`
    TakeNum   string `json:"take_num"`
    OrderDate string `json:"order_date"`
    Status    string `json:"status"`
}

func GetUpdateSQLByID(tableName, fieldName string, data map[int]interface{}) string {
    caseWhen := ""
    for k, v := range data {
        caseWhen += fmt.Sprintf(` WHEN id = %v THEN '%s'`, k, v)
    }
    return fmt.Sprintf(`UPDATE %v SET "%v" = CASE %v ELSE '' END`, tableName, fieldName, caseWhen)
}

type A struct {
    M map[string]interface{}
}

func MN() {
    l := []int{1, 2, 3}
    for _, v := range l {
        l = append(l, v+3)
        fmt.Printf("elem: %v, len: %v, cap: %v\n", v+3, len(l), cap(l))
    }
    fmt.Printf("len: %v", len(l))

}
func Getb() *B {
    return nil
}

type B struct {
    A int
}

func Sleep(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():

        }
    }

}
func main() {
    data, err := util.Get("192.168.1.167")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("data:%#v", data)
    //m := make(map[int]int, 0)
    //l := make([]int, 0)
    //for i := 0; i < 1000; i++ {
    //    m[i] = 0
    //    l = append(l, 0)
    //}
    //n := time.Now()
    //for _, _ := range m {
    //}
    //t1 := time.Since(n)
    //n = time.Now()
    //for _, _ := range l {
    //}
    //t2 := time.Since(n)
    //fmt.Printf("t1:%v,t2:%v\n", t1, t2)

    //GO()
    //MN()
    //t1 := time.Now()
    //time.Sleep(time.Second*2)
    //t2 := time.Now()
    //fmt.Printf("t1:%v,t2:%v\n",t1,t2)
    //fmt.Println((t1.Sub(t2)).Seconds())
    //var c float64
    //c = 2
    //MapTest()
    //var t NotifyS
    //m := map[string]interface{}{"store_code":"1"}
    //b, _ := json.Marshal(m)
    //_ = json.Unmarshal(b, &t)
    //fmt.Printf("%#v", t)
    //m := map[int]interface{}{
    //	1:"zhsnag",
    //	2:"ls",
    //	3:"dwad",
    //}
    //fmt.Println(GetUpdateSQLByID("awaw", "ame", m))
    //fmt.Println(fmt.Sprintf("%v",math.Pow(1.1,365)))
    //fmt.Println(time.Now().Unix())
}
func Send() {
    n := NotifyS{
        StoreCode: "112",
        OrderNum:  "2313123",
        TakeNum:   "0001",
        OrderDate: time.Now().Format(time.RFC850),
        Status:    "0",
    }
    conn, err := net.Dial("tcp", "localhost:9999")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    bytes, _ := json.Marshal(n)
    _, _ = conn.Write(bytes)
}

type Map map[string]interface{}

func (this Map) Put(key string, value interface{}) {
    ioutil.ReadAll(nil)
    this[key] = value
}
func Add(v *int) int {
    *v ++
    return *v
}
func GOTO() {
    defer func() {
        fmt.Println("2321312231231")
    }()
    for i := 0; i < 10; i++ {
        fmt.Println(strconv.Itoa(i) + "111")
        if i == 5 {
            goto AAA
        }
        fmt.Println(strconv.Itoa(i) + "222")
    AAA:
        fmt.Println(strconv.Itoa(i) + "333")
    }
    fmt.Println("success")
}
func test1() {
    a := "aaa"
    b := "bbb"
    s := ""
    st := time.Now()
    for i := 0; i < 10000000; i++ {
        s = fmt.Sprintf("%v.%v", a, b)
    }
    fmt.Println(s)
    fmt.Printf("t:%v\n", time.Since(st))
}

var count = 1000000

func test2() {
    a := "aaa"
    b := "bbb"
    s := ""
    st := time.Now()
    for i := 0; i < 10000000; i++ {
        s = a + "." + b
    }
    fmt.Println(s)
    fmt.Printf("t:%v\n", time.Since(st))
}
func DBOption() {
    defer func() {
        fmt.Println("defer")
    }()
    if true {
        log.Fatal("111")
        return
    }
    fmt.Println("111")
}
func Word(ctx context.Context) {
    select {
    case <-ctx.Done():
        return
    default:
        fmt.Println("232131231")
        time.Sleep(time.Second * 10)
        fmt.Println("============")
    }
    //select {
    //case <-ctx.Done():
    //    fmt.Println("goroutine exit")
    //    return
    //case <-time.After(time.Second):
    //    return
    //}
}
func ContextT() {

}

func GR() {
    var wg sync.WaitGroup
    wg.Add(10)
    for i := 0; i < 10; i++ {
        go func(id int) {
            time.Sleep(time.Second)
            fmt.Printf("g%v\n", id)
            wg.Done()
        }(i)
    }
    wg.Wait()
}
func MapTest() {
    m := map[int]int{
        1: 2,
        2: 2,
        3: 3,
    }
    for k, v := range m {
        if v == 2 {
            delete(m, k)
        }
    }
    fmt.Println(m)
}

func Contans() {

}
func Switch() {

}
