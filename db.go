package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "os"
    "strings"
    "time"
)


func getLogger() logger.Interface {
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
        logger.Config{
            SlowThreshold: time.Second, // Slow SQL threshold
            LogLevel:      logger.Info, // Log level
            //IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
            Colorful: false, // Disable color
        },
    )
    return newLogger
}
func getDSN() gorm.Dialector {
    result := mysql.New(mysql.Config{
        DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
            "xiao9878",
            "Xiao9878",
            "rm-bp1i6i0kh0x0p7srano.mysql.rds.aliyuncs.com",
            "3306",
            "manager"),
    })
    return result
}

type User struct {
    ID int `json:"id"`
    UserName string `json:"user_name"`
    Password string `json:"password"`
}

func ADw()  {
    for  {
        select {
        default:
            fmt.Println(111)
        }
    }
}

type AAA1 struct {
    Id int `json:"-"`
    Age int `json:"age"`
}
func main() {
    fmt.Println(strings.Contains("dwada","a"))
    //var u User
    //err := json.Unmarshal(nil, &u)
    //if err != nil {
    //    log.Fatal(err)
    //}
    //fmt.Println(u)
    //conn, err := net.Listen("tcp", ":8888")
    //if err != nil {
    //    log.Fatal(err)
    //}
    //time.Sleep(time.Second*20)
    //if err := conn.Close(); err != nil {
    //    log.Fatal(err)
    //}
    //log.Println("close success")
    //db, err := gorm.Open(getDSN(), &gorm.Config{Logger: getLogger()})
    //if err != nil {
    //    log.Fatal(err)
    //}
    //users := make([]User,0)
    //tx := db.Table("sys_user_copy1").Find(&users)
    //if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
    //    log.Fatal(tx.Error.Error())
    //}
    //log.Printf("user:%#v",users)
    //for _, user := range users {
    //    user.Password = "123"
    //}
    //log.Printf("user:%#v",users)
    //tx := db.Raw("UPDATE `sys_user_copy1` SET `age` = 110 WHERE `id` = 2;").Scan(nil)
    //if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
    //    log.Fatal(tx.Error.Error())
    //}
    //fmt.Println(tx.RowsAffected)
}
