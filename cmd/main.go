package main

import (
	"crawler/util"
	"log"
)

const (
	url  = "https://www.lagou.com/"
	path = "./data.txt"
)

func main() {
	data, err := util.Get(url)
	if err != nil {
		log.Fatalf("util.Get error(%v)", err)
		return
	}
	util.SaveAsFile(path, data)
}
