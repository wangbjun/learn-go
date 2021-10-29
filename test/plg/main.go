package main

var Name = "Plugin Name"

type NameStruct struct{}

func NewName() NameStruct {
	return NameStruct{}
}

func (NameStruct) GetName() string {
	return Name
}
