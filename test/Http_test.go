package test

import "golang-start/common"

func main() {
	map1 := make(map[string]string)
	map1["teamName"] = "haha"
	common.HttpDoi(map1, "")
}
