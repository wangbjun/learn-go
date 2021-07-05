package main

type Instance struct{}

var myInstance = &Instance{}

func GetInstance() *Instance {
	return myInstance
}
