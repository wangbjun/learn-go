package main

import "github.com/tidwall/gjson"

var Name = "Plugin Name"

func GetName() string {
	err := gjson.ValidBytes([]byte("222222"))
	println(err)
	return Name
}
