package main

import (
    "fmt"
    g "github.com/gosnmp/gosnmp"
    "log"
    "strconv"
)

func main() {
    // Default is a pointer to a GoSNMP struct that contains sensible defaults
    // eg port 161, community public, etc
    g.Default.Target = "192.168.1.167"
    err := g.Default.Connect()
    if err != nil {
        log.Fatalf("Connect() err: %v", err)
    }
    defer g.Default.Conn.Close()

    oids := []string{"1.3.6.1.4.1.2435.3.3.9.1.6.1.0"}
    result, err2 := g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
    if err2 != nil {
        log.Fatalf("Get() err: %v", err2)
    }

    for i, variable := range result.Variables {
        fmt.Printf("%d: oid: %s ", i, variable.Name)

        // the Value of each variable returned by Get() implements
        // interface{}. You could do a type switch...
        switch variable.Type {
        case g.OctetString:
            bytes := variable.Value.([]byte)
            l := make([]string, 0)
            for _, item := range bytes {
               s := strconv.FormatInt(int64(item), 16)
               l = append(l, s)
            }
            fmt.Printf("string: %v\n", bytes)
            fmt.Printf("string: %v\n", l)
        default:
            // ... or often you're just interested in numeric values.
            // ToBigInt() will return the Value as a BigInt, for plugging
            // into your calculations.
            fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
        }
    }

}
