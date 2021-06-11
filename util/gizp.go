package util

import (
    "bytes"
    "compress/gzip"
    "github.com/pkg/errors"
    mp "github.com/vmihailenco/msgpack"
    "io/ioutil"
)

//Marshal(v interface{}) ([]byte, error)
//Unmarshal(b []byte, v interface{}) error
//// name of this codec
//Name() string

type msgpackGzipCodec int

var Codec = new(msgpackGzipCodec)

const name = "msgpackGzip"

func (m msgpackGzipCodec) Marshal(v interface{}) ([]byte, error) {
    bs, err := mp.Marshal(v)
    if err != nil {
        return nil, err
    }
    l, err := gzipWriter(bs)
    if err != nil {
        return nil, err
    }
    return l, nil
}
func (m msgpackGzipCodec) Unmarshal(b []byte, v interface{}) error {
    to, err := gzipReader(b)
    if err != nil {
        return err
    }
    return mp.Unmarshal(to, v)
}
func (m msgpackGzipCodec) Name() string {
    return name
}

func gzipWriter(content []byte) ([]byte, error) {
    var b bytes.Buffer
    w := gzip.NewWriter(&b)
    if _, err := w.Write(content); err != nil {
        return nil, errors.Wrap(err, "gzipWriter 写入")
    }
    defer w.Close()
    //if err := w.Close(); err != nil {
    //    return nil, errors.Wrap(err, "gzipWriter 关闭")
    //}

    return b.Bytes(), nil
}

func gzipReader(input []byte) ([]byte, error) {
    fz, err := gzip.NewReader(bytes.NewReader(input))
    if err != nil {
        return nil, err
    }
    defer fz.Close()

    s, err := ioutil.ReadAll(fz)
    if err := fz.Close(); err != nil {
       return nil, errors.Wrap(err, "gzipWriter 关闭")
    }
    return s, nil
}
