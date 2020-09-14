package main

import (
	"Go-times/cmd"
	"log"
)

/**
* @Author: super
* @Date: 2020-09-14 22:17
* @Description:
**/

func main() {
	err := cmd.Execute()
	if err != nil{
		log.Fatalf("cmd.Execute err: %v", err)
	}
}