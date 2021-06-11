package main

import (
    "bytes"
    "compress/gzip"
    "fmt"
    mp "github.com/vmihailenco/msgpack"

    "io"
    "log"
)

func gUnzipData(data []byte) (resData []byte, err error) {
    b := bytes.NewBuffer(data)

    var r io.Reader
    r, err = gzip.NewReader(b)
    if err != nil {
        return
    }

    var resB bytes.Buffer
    _, err = resB.ReadFrom(r)
    if err != nil {
        return
    }

    resData = resB.Bytes()

    return
}

func gZipData(data []byte) (compressedData []byte, err error) {
    var b bytes.Buffer
    gz := gzip.NewWriter(&b)

    _, err = gz.Write(data)
    if err != nil {
        return
    }

    if err = gz.Flush(); err != nil {
        return
    }

    if err = gz.Close(); err != nil {
        return
    }

    compressedData = b.Bytes()

    return
}

const name = "msgpack_gzip"

// Codec that encodes to and decodes from msgpack.
var Codec = new(msgpackCodec)

type msgpackCodec int

func (m msgpackCodec) Marshal(v interface{}) ([]byte, error) {
    tmp, err := mp.Marshal(v)
    if err != nil {
        return nil, err
    }
    return gZipData(tmp)
}

func (m msgpackCodec) Unmarshal(b []byte, v interface{}) error {
    tmp, err := gUnzipData(b)
    if err != nil {
        return err
    }
    return mp.Unmarshal(tmp, v)
}

func (m msgpackCodec) Name() string {
    return name
}

func main() {
    // define original data
    data := []byte(`MyzYrIyMLyNqwDSTBqSwM2D6KD9sA8S/d3Vyy6ldE+oRVdWyqNQrjTxQ6uG3XBOS0P4GGaIMJEPQ/gYZogwkQ+A0/gSU03fRJvdhIGQ1AMARVdWyqNQrjRFV1bKo1CuNEVXVsqjUK40RVdWyqNQrjRFV1bKo1CuNPmQF870PPsnSNeKI1U/MrOA0/gSU03fRb2A3OsnORNIruhCUYTIrOMTNU7JuGb5RSYJxa6PiMHdiRmFtXLNoY+GVmTD7aOV/K1yo4y0dR7Q=`)
    fmt.Printf("original data:%s,\n", data)
    fmt.Println("original data len:", len(data))

    // compress data
    compressedData, compressedDataErr := gZipData(data)
    if compressedDataErr != nil {
        log.Fatal(compressedDataErr)
    }

    fmt.Printf("compressed data:%s\n", compressedData)
    fmt.Println("compressed data len:", len(compressedData))

    // uncompress data
    uncompressedData, uncompressedDataErr := gUnzipData(compressedData)
    if uncompressedDataErr != nil {
        log.Fatal(uncompressedDataErr)
    }

    fmt.Printf("uncompressed data:%s\n", uncompressedData)
    fmt.Println("uncompressed data len:", len(uncompressedData))
}