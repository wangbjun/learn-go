package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
)

func main() {
	p := &Person{
		Id:    1,
		Name:  "jun",
		Age:   25,
		Money: 24.5,
		Car:   []string{"car1", "car2"},
		Phone: &Person_Phone{Number: "0551-12323232", Type: "1"},
		Sex:   Person_female,
	}

	//Marshal
	out, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	file, _ := os.OpenFile("out", os.O_CREATE|os.O_WRONLY, 0666)
	_, _ = file.Write(out)
	_ = file.Close()

	//unMarshal
	in, _ := os.Open("out")
	bytes, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}
	p1 := &Person{}
	err = proto.Unmarshal(bytes, p1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", p1.String())
	fmt.Printf("%d\n", p1.GetId())
}
