package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"time"

	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/json"
)



type User struct {
	ID   int    `json:"product_id"  storm:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Select()  {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second*6)
		ch <- 1
	}()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <- ch:
		 fmt.Println("exit")
			ticker.Stop()
			return
		case <-ticker.C:
			fmt.Println(1111)
			time.Sleep(time.Second)
		}
	}
}
func TestA()  {
	// gobDb, _ := storm.Open("gob.db", storm.Codec(gob.Codec))
	jsonDb, _ := storm.Open("db/jsondb", storm.Codec(json.Codec))
	//msgpackDb, _ := storm.Open("db/msgpackdb", storm.Codec(msgpack.Codec))
	//gzipDb, _ := storm.Open("db/gzipdb", storm.Codec(util.Codec))
	// serealDb, _ := storm.Open("sereal.db", storm.Codec(sereal.Codec))
	// protobufDb, _ := storm.Open("protobuf.db", storm.Codec(protobuf.Codec))
	defer jsonDb.Close()
	//defer msgpackDb.Close()
	//defer gzipDb.Close()
	for i := 50000; i < 60010 ; i++ {
		u := User{
			ID:   i,
			Name: fmt.Sprint(i),
			Age:  i,
		}
		// protobufDb.Save(&u)
		// gobDb.Save(&u)
		// serealDb.Save(&u)
		//if err := gzipDb.Save(&u); err != nil {
		//	log.Fatal("gzip", err)
		//}
		//if err := msgpackDb.Save(&u); err != nil {
		//	log.Fatal("msgpack", err)
		//}
		if err := jsonDb.Save(&u); err != nil {
			log.Fatal("json", err)
		}
		fmt.Println(i)
	}

	// fmt.Printf("%T\n", gobDb.Codec())
	fmt.Printf("%T\n", jsonDb.Codec())
	//fmt.Printf("%T\n", msgpackDb.Codec())
	//fmt.Printf("%T\n", gzipDb.Codec())
	// fmt.Printf("%T\n", serealDb.Codec())
	// fmt.Printf("%T\n", protobufDb.Codec())
}
func main() {
	TestA()
	//zip()
//Select()
}

func zip() {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()

	err := ioutil.WriteFile("hello_world.txt.gz", b.Bytes(), 0666)
	fmt.Println(err)
}

func gzipWriter(content []byte) ([]byte, error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	if _, err := w.Write(content); err != nil {
		return nil, errors.Wrap(err, "gzipWriter 写入")
	}
	if err := w.Close(); err != nil {
		return nil, errors.Wrap(err, "gzipWriter 关闭")
	}
	return b.Bytes(), nil
}

func gzipReader(input []byte) ([]byte, error) {
	fz, err := gzip.NewReader(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}
	defer fz.Close()

	s, err := ioutil.ReadAll(fz)
	if err != nil {
		return nil, errors.Wrap(err, "gzipReader read")
	}
	return s, nil
}

