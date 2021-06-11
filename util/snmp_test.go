package util

import (
    "reflect"
    "testing"
)

func TestGet(t *testing.T) {
    type args struct {
        ip string
    }
    tests := []struct {
        name    string
        args    args
        want    *PrinterStatus
        wantErr bool
    }{
        // TODO: Add test cases.
        {name: "1", args: args{ip: "192.168.1.167"}, want: &PrinterStatus{
            Code:   0,
            Msg:    "ok",
            Detail: "",
        }},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Get(tt.args.ip)
            if (err != nil) != tt.wantErr {
                t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Get() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_match(t *testing.T) {
    type args struct {
        data map[int]string
    }
    tests := []struct {
        name string
        args args
        want status
    }{
        // TODO: Add test cases.
        {name: "1", args: args{data: map[int]string{9: "1", 10: "1", 19: "0", 20: "0", 23: "0"}},want: STATUS_UNDEFINED},
        {name: "3", args: args{data: map[int]string{9: "1", 10: "0", 19: "0", 20: "0", 23: "0"}},want: STATUS_PRINTING},
        {name: "2", args: args{data: map[int]string{9: "0", 10: "1", 19: "0", 20: "0", 23: "0"}},want: STATUS_LID_OPEND},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := match(tt.args.data); got != tt.want {
                t.Errorf("match() = %v, want %v", got, tt.want)
            }
        })
    }
}