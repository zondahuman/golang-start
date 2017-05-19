package im

import (
	"fmt"
	"golang-start/common"
)

func Han() {
	fmt.Println("i am han")
	han1()
}

func han1() {
	var httpUrl = "http://localhost:8300/team/call1"
	//var httpUrl = "http://172.16.2.133:8300/team/call1"
	map1 := make(map[string]string)
	map1["teamName"] = "haha"
	var result = common.HttpDoi(map1, httpUrl)
	fmt.Println(result)
}
