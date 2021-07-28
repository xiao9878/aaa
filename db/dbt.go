package main

import "C"
import (
    jj "encoding/json"
    "fmt"
    "github.com/asdine/storm"
    "github.com/asdine/storm/codec/json"
    "log"
    "sync"
    "sync/atomic"
    "time"
)




type Pos struct {
    ID         string `json:"id"`
    Code       string `json:"code"`
    DeviceId   string `json:"device_id"`
    DeviceCode string `json:"device_code"`
}
type Operator struct {
    ID        string `json:"id"`
    LoginId   string `json:"loginId"`
    Name      string `json:"name"`
    Code      string `json:"code"`
    LoginTime string `json:"login_time"`
}
type Amount struct {
    TaxAmount        float64 `json:"taxAmount"`
    GrossAmount      float64 `json:"gross_amount"`
    PayAmount        float64 `json:"pay_amount"`
    NetAmount        float64 `json:"net_amount"`
    DiscountAmount   float64 `json:"discount_amount"`
    RemovezeroAmount float64 `json:"removezero_amount"`
    Rounding         float64 `json:"rounding"`
    OverflowAmount   float64 `json:"overflow_amount"`
    ChangeAmount     float64 `json:"changeAmount"`
    ServiceFee       float64 `json:"serviceFee"`
    Tip              float64 `json:"tip"`
}
type RefundInfo struct {
    RefundId     string `json:"refund_id"`
    RefundNo     string `json:"refund_no"`
    RefTicketId  string `json:"ref_ticket_id"`
    RefTicketNo  string `json:"ref_ticket_no"`
    RefundReason string `json:"refund_reason"`
}

type Channel struct {
    Source       string `json:"source"`
    DeviceType   string `json:"deviceType"`
    OrderType    string `json:"orderType"`
    DeliveryType string `json:"deliveryType"`
    TpName       string `json:"tpName"`
    Code         string `json:"code"`
}

func Channel2Code(c Channel) string {
    return ""
}

// ticket 中包含的product
type ProductInTicket struct {
    ID             string             `json:"id"`
    Name           string             `json:"name"`
    Code           string             `json:"code"`
    SeqId          int                `json:"seq_id"`
    Price          float64            `json:"price"`
    Amount         float64            `json:"amount"`
    Qty            int                `json:"qty"`
    DiscountAmount float64            `json:"discount_amount"`
    Type           string             `json:"type"`
    Accesssories   []*ProductInTicket `json:"accessories"`
    ComboItems     []*ProductInTicket `json:"combo_items"`
    Remark         string             `json:"remark"`
    SkuRemarks     []SkuRemark        `json:"skuRemark"`
    TaxAmount      float64            `json:"taxAmount"`
    NetAmount      float64            `json:"net_amount"`
}

func (tp ProductInTicket) GetProductCount() int {
    if len(tp.ComboItems) == 0 {
        return tp.Qty
    }
    count := 0
    for _, cp := range tp.ComboItems {
        count += cp.GetProductCount()
    }
    return tp.Qty * count
}

type SkuRemark struct {
    Name  SkuName  `json:"name"`
    Value SkuValue `json:"values"`
}

type SkuName struct {
    ID   string `json:"id"`
    Code string `json:"code"`
    Name string `json:"name"`
}
type SkuValue struct {
    Code string `json:"code"`
    Name string `json:"name"`
}
type Payment struct {
    ID              string  `json:"id"`
    SeqId           string  `json:"seq_id"`
    PayAmount       float64 `json:"pay_amount"`
    Receivable      float64 `json:"receivable"`
    Change          float64 `json:"change"`
    Overflow        float64 `json:"overflow"`
    Rounding        float64 `json:"rounding"`
    PayTime         string  `json:"pay_time"`
    TransCode       string  `json:"trans_code"`
    Name            string  `json:"name"`
    TpTransactionNo string  `json:"tpTransactionNo"`
}

type Promotion struct {
    PromotionInfo PromotionInfo      `json:"promotionInfo"`
    Source        PromotionSource    `json:"source"`
    Product       []PromotionProduct `json:"products"`
}

type PromotionInfo struct {
    Type               string  `json:"type"`
    DiscountType       string  `json:"discount_type"`
    Name               string  `json:"name"`
    PromotionId        string  `json:"promotion_id"`
    PromotionCode      string  `json:"promotion_code"`
    PromotionType      string  `json:"promotion_type"`
    AllowOverlap       bool    `json:"allow_overlap"`
    TriggerTimesCustom bool    `json:"trigger_times_custom"`
    TicketDisplay      string  `json:"ticket_display"`
    MaxDiscount        float64 `json:"max_discount"`
}

type PromotionSource struct {
    Trigger  int      `json:"trigger"`
    Discount float64  `json:"discount"`
    Fired    []string `json:"fired"`
}

type PromotionProduct struct {
    Price    float64  `json:"price"`
    Amt      float64  `json:"amt"`
    AccAmt   float64  `json:"accAmt"`
    Qty      float64  `json:"qty"`
    KeyId    string   `json:"key_id"`
    Accies   []string `json:"accies"`
    Type     string   `json:"type"`
    Discount float64  `json:"discount"`
    FreeAmt  float64  `json:"free_amt"`
    Method   string   `json:"method"`
}

type Member struct {
    MemberCode string `json:"member_code"` //会员号
    Mobile     string `json:"mobile"`      //会员手机号
    Name       string `json:"name"`        //会员姓名
    Greetings  string `json:"greetings"`   //会员问候语
}

type Coupon struct {
    IsOnline       bool    `json:"is_online"`   //是否是线上券
    Id             string  `json:"id"`          //卡券唯一编码
    Name           string  `json:"name"`        //卡券名称
    Code           string  `json:"code"`        //会员券类型
    Type           int     `json:"type"`        //是否是会员券
    PromotionValue float64 `json:"par_value"`   //卡券优惠金额
    SequenceId     string  `json:"sequence_id"` //序列号
}

type Table struct {
    ID        string `json:"id"`
    ZoneID    string `json:"zone_id"`
    TableNo   string `json:"tableNo"`
    People    int    `json:"people"`
    Temporary bool   `json:"temporary"`
}

type Tax struct {
    Amount   float64 `json:"amount"`
    SubTotal float64 `json:"subTotal"`
    Code     string  `json:"code"`
    Name     string  `json:"name"`
    Rate     float64 `json:"rate"`
}

type Fee struct {
    Name   string  `json:"name"`
    Price  float64 `json:"price"`
    Qty    int     `json:"qty"`
    Amount float64 `json:"amount"`
    Type   string  `json:"type"` //用于表示喜茶的各种费用类型，目前是有外卖单有费用和费用类型，D:第三方配送费 E:门店自配送费用 P:打包费 B:平台佣金 T:平台补贴 G: 第三方赠品 S: 商家优惠
}

type Takeaway struct {
    OrderMethod         string      `json:"order_method"`
    TpOrderId           string      `json:"tp_order_id"`
    OrderTime           string      `json:"order_time"`
    OrdersTimestamp     int64       `json:"ordersTimestamp"`
    ExpectTimestamp     int64       `json:"expectTimestamp"`
    DeliverTime         string      `json:"deliver_time"`
    Description         string      `json:"description"`
    Consignee           interface{} `json:"consignee"`
    PhoneList           []string    `json:"phone_list"`
    Tp                  string      `json:"tp"`
    Source              string      `json:"source"`
    SourceOrderId       string      `json:"source_order_id"`
    DaySeq              string      `json:"day_seq"`
    DeliveryType        int         `json:"delivery_type"`
    DeliveryName        string      `json:"delivery_name"`
    DeliveryAddress     string      `json:"delivery_poi_address"`
    InvoiceTitle        string      `json:"invoice_title"`
    WaitingTime         string      `json:"waiting_time"`
    TablewareNum        int         `json:"tableware_num"`
    SendFee             interface{} `json:"send_fee"`
    PackageFee          float64     `json:"package_fee"`
    DeliveryTime        string      `json:"delivery_time"`
    TakeMealSn          string      `json:"take_meal_sn"`
    NeedInvoice         bool        `json:"needInvoice"`
    PartnerPlatformId   int32       `json:"partnerPlatformId"`   //喜茶独有
    PartnerPlatformName string      `json:"partnerPlatformName"` //喜茶独有
    WxName              string      `json:"wxName"`              //点单用户的微信昵称，需要打印在杯贴上，喜茶独有逻辑
    IsHighPriority      bool        `json:"isHighPriority"`      //是否是优先券
    TakeoutType         string      `json:"takeoutType"`         //外卖类型，分为正常单和部分退款单
    OriginalOrderNo     string      `json:"originalOrderNo"`     //部分退款时的外卖原单号
}
type Ticket struct {
    Id          string      `json:"ticket_id" storm:"id"`
    EndTime     string      `json:"end_time" storm:"index"`
    BusDate     string      `json:"bus_date" storm:"index"`
    Status      string      `json:"status" storm:"index"`
    ShiftNumber int         `json:"shiftNumber"`
    Extend      interface{} `json:"extend"`
    OrderSource string      `json:"order_source"`
    IsUpload    bool        `json:"isUpload"`

    TakemealsNumber string    `json:"takemealsNumber"`          // 取餐号
    FlowStatus      uint8     `json:"flowStatus" storm:"index"` // 订单状态 初始状态 》排队中》制作中》制作完成》已取餐
    UpdateTime      time.Time `json:"updateTime" storm:"index"` // flowStatus 状态更新时间
    TotalQty        int32     `json:"totalQty"`                 //该订单总杯数
    MakeTime        float64   `json:"makeTime"`                 //该订单制作时间
}

type Ticket1 struct {
    TicketId             string             `json:"ticket_id" storm:"id"`
    TicketNo             string             `json:"ticket_no"`
    TicketUno            string             `json:"ticketUno"`
    StartTime            string             `json:"start_time"`
    EndTime              string             `json:"end_time" storm:"index"`
    BusDate              string             `json:"bus_date" storm:"index"`
    UploadTime           string             `json:"uploadTime"`
    UploadFlag           bool               `json:"upload_flag"` // 这个数据是为了避免在 升级到这个版本之后，历史数据满足没有上传的条件 会一下子触发上传，
    UploadFailed         int                `json:"upload_failed"`
    Pos                  Pos                `json:"pos"`
    Operator             Operator           `json:"operator"`
    Amount               Amount             `json:"amounts"`
    TakemealNumber       string             `json:"takemealNumber"`
    Qty                  int                `json:"qty"`
    Status               string             `json:"status"`
    RefundInfo           RefundInfo         `json:"refundInfo"`
    Channel              Channel            `json:"channel"`
    Products             []*ProductInTicket `json:"products"`
    Payments             []Payment          `json:"payments"`
    Promotions           []Promotion        `json:"promotions"`
    Members              []Member           `json:"members"`
    Table                Table              `json:"table"`
    Coupons              []Coupon           `json:"coupons"`
    People               int                `json:"people"`
    RoomNo               string             `json:"room_no"`
    Remark               string             `json:"remark"`
    HouseAc              bool               `json:"house_ac"`
    OrderTimeType        string             `json:"order_time_type"`
    ShiftNumber          int             `json:"shiftNumber" storm:"index"`
    TaxList              []Tax              `json:"taxList"`
    TakeawayInfo         Takeaway           `json:"takeaway_info"`
    Fees                 []Fee              `json:"fees"`
    TimeZone             string             `json:"timeZone"`
    DiscountProportioned bool               `json:"discount_proportioned"`
}
type ProductionAddition struct {
    Id string
    SortCode float64
}

func GetTicket()  {
    db, err := storm.Open("db/storm_data", storm.Codec(json.Codec))
    //db, err := storm.Open("db/storm_data", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    if db == nil {
        log.Fatal("db is nil")
    }
    var t Ticket
    err = db.One("Id", "155b70a1d160486bbd415113560b30bd", &t)
    if err != nil && err != storm.ErrNotFound {
        panic(err)
    }
    if err == storm.ErrNotFound {
        fmt.Println("没有相关数据")
        return
    }
    b, _ := jj.Marshal(t)
    fmt.Printf("%s",b)
}
func BoltQuery() {
    db, err := storm.Open("db/storm_data", storm.Codec(json.Codec))
    //db, err := storm.Open("db/storm_data", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    if db == nil {
        log.Fatal("db is nil")
    }
    var to []Ticket
    err = db.Find("BusDate", "2021-07-27", &to)
    //var t Ticket
    //err = db.One("Id", "9ecc0da32bbf41fe92e817bdfb317f42", &t)
    if err != nil && err != storm.ErrNotFound {
       panic(err)
    }
    for _, ticket := range to {
        if !ticket.IsUpload {
            fmt.Println(1)
        }
    }
    b, _ := jj.Marshal(to)
    fmt.Printf("%s",b)
    //var u User
    //err = db.Update(func(tx *bolt.Tx) error {
    //    b := tx.DeleteBucket([]byte("User"))
    //    fmt.Println(b)
    //    if b != nil {
    //        return b
    //    }
    //    //err := b.Put([]byte("answer"), []byte("42"))
    //    return nil
    //})
    //fmt.Println(err)
}
func B()  {
    go func() {
        for  {
            fmt.Println(1)
            time.Sleep(time.Second)
        }
    }()
}
var a bool
func A()  {
    if !a {
        a = true
        return
    }
    var i uint64 = 1
    fmt.Println(atomic.AddUint64(&i, 1))
}

var  o sync.Once
func main() {

    GetTicket()
    //m := map[string]interface{}{
    //    "name":"一年级1班",
    //    "student": []map[string]interface{}{
    //        {"no":1,"name":"zs","age":12},
    //        {"no":2,"name":"ls","age":13},
    //        {"no":3,"name":"ww","age":14},
    //        {"no":4,"name":"zl","age":15},
    //    },
    //}
    //var to interface{}
    //fmt.Printf("orgin:%v\n",m)
    //mpb, err := msgpack.Marshal(m)
    //if err != nil {
    //    log.Fatal(err)
    //}
    //fmt.Printf("msgpack:%v\n",mpb)
    //err = msgpack.Unmarshal(mpb, &to)
    //if err != nil {
    //    log.Fatal(err)
    //}
    //fmt.Printf("ummarshal:%v\n", to)
    //jsonb, err := json.Marshal(m)
    //if err != nil {
    //    log.Fatal(err)
    //}
    //fmt.Printf("json:%v\n",string(jsonb))

    //ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
    //go func() {
    //    Work()
    //}()
    //time.Sleep(time.Second*1)
    //select {
    //case <- ctx.Done():
    //    fmt.Println("time out!")
    //}
}
func String(v interface{}) string {
    return fmt.Sprintf("%v", v)
}

func Work() {
    fmt.Println("start!")
    time.Sleep(time.Second * 10)
    fmt.Println("end !")
}
