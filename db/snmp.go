package main

import (
    "fmt"
    g "github.com/gosnmp/gosnmp"
    "log"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {
    oids := []string{".1.3.6.1.2.1.1.5.0"}
    s := NewSnmpClientV2("192.168.1.167", "public", uint16(9100))

    err1 := s.client.Connect()
    if err1 != nil {
        fmt.Println("GetModeV2 snmp connect failure,err: %v", err1.Error())
    }
    defer s.client.Conn.Close()

    res, err2 := s.client.Get(oids)
    if err2 != nil {
        fmt.Println("GetModeV2 snmp conn switch v3, info: %v", err2)
    } else {
        for _, v := range res.Variables {
            switch v.Type {
            case g.ObjectIdentifier:
                fmt.Println(v.Value.(string), "ObjectIdentifier")
            case g.OctetString:
                tempStr := strings.Split(string(v.Value.([]uint8)), "\r")
                fmt.Println(tempStr[0], "OctetString")
            case g.Integer:
                temp := v.Value.(int)
                fmt.Println(strconv.Itoa(temp), "g.Integer")
            }
        }
    }
}

func NewSnmpClientV2(target, community string, port uint16) *SnmpClient {
    return &SnmpClient{
        g.GoSNMP{
            Target:         target,
            Community:      community,
            Port:           port,
            Version:        g.Version2c,
            Timeout:        time.Duration(2) * time.Second,
            Retries:        1,
            MaxOids:        1,
            MaxRepetitions: 2,
            Logger:         g.NewLogger(log.New(os.Stdout, "", 0)),
        },
    }
}

type SnmpClient struct {
    client g.GoSNMP
}
