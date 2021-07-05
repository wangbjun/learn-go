package main

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"os"
	"testing"
)

func BenchmarkProto(b *testing.B) {
	g := &Group{}
	for i := 0; i < 10; i++ {
		p := &Person{
			Id:    int32(i),
			Name:  "测试名称",
			Age:   int32(25 * i),
			Money: 240000.5,
			Car:   []string{"car1", "car2", "car3", "car4", "car5", "car7", "car6", "car21", "car22"},
			Phone: &Person_Phone{Number: "0551-12323232", Type: "1"},
			Sex:   Person_female,
		}
		g.Person = append(g.Person, p)
	}

	out, _ := proto.Marshal(g)
	file, _ := os.OpenFile("out", os.O_CREATE|os.O_WRONLY, 0666)
	_, _ = file.Write(out)
	_ = file.Close()

	//
	//b.ResetTimer()
	//b.N = 1000
	//for i := 0; i < b.N; i++ {
	//	out, err := proto.Marshal(g)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	g1 := &Group{}
	//	err = proto.Unmarshal(out, g1)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
}

func BenchmarkJson(b *testing.B) {
	g := &Group{}
	for i := 0; i < 10; i++ {
		p := &Person{
			Id:    int32(i),
			Name:  "测试名称",
			Age:   int32(25 * i),
			Money: 240000.5,
			Car:   []string{"car1", "car2", "car3", "car4", "car5", "car7", "car6", "car21", "car22"},
			Phone: &Person_Phone{Number: "0551-12323232", Type: "1"},
			Sex:   Person_female,
		}
		g.Person = append(g.Person, p)
	}
	out, _ := json.Marshal(g)
	file, _ := os.OpenFile("out", os.O_CREATE|os.O_WRONLY, 0666)
	_, _ = file.Write(out)
	_ = file.Close()
	//b.ResetTimer()
	//b.N = 1000
	//for i := 0; i < b.N; i++ {
	//	out, err := json.Marshal(g)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	g1 := &Group{}
	//	err = json.Unmarshal(out, g1)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
}
