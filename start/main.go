package main

import "fmt"
import "golang-start/common"
import "golang-start/im"

func main() {
	fmt.Println(111)
	fmt.Println("df")
	im.Han()
	var httpUrl = "http://localhost:7100/team/call1"
	map1 := make(map[string]string)
	map1["teamName"] = "haha"
	//result = common.HttpDoi(map1, httpUrl)
	common.HttpDoi(map1, httpUrl)
}
