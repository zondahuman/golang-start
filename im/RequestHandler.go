package im

import (
	"golang-start/im/model"
	"github.com/golang/protobuf/proto"
	"log"
	"fmt"
	"encoding/json"
)

func Handler(params string) string {
	//print(params)
	//request := &model.RequestVo{}
	//err := proto.Unmarshal(params, request)
	//if err != nil {
	//	log.Fatal("unmarshaling error: ", err)
	//}
	//fmt.Println(request)
	return "bbbb"
}

func CreateMessage() string {
	request  := &model.RequestVo{
		UserId : proto.String("lee"),
		Source : model.RequestVo_WEB.Enum(),
		Content : proto.String("hi, guys"),
		SourceId : proto.String("asd"),
	}
	requestPacket , err := proto.Marshal(request)
	if err != nil {
		log.Fatal("marshal error:", err)
	}
	var result = string(requestPacket)
	print("result : ", result)
	return result;
}

func CycleMessage()  {
	request  := &model.RequestVo{
		UserId : proto.String("lee"),
		Source : model.RequestVo_WEB.Enum(),
		Content : proto.String("hi, guys"),
		SourceId : proto.String("asd"),
	}
	requestPacket , err := proto.Marshal(request)
	if err != nil {
		log.Fatal("marshal error:", err)
	}
	result, err2 := json.Marshal(requestPacket)
	if err2 != nil {
		log.Fatal("json marshal error:", err2)
	}
	print("result : ", result)

	response := &model.RequestVo{}
	err1 := proto.Unmarshal(requestPacket, response)
	if err1 != nil {
		log.Fatal("unmarshaling error: ", err1)
	}
	fmt.Println("response= ", response)
}


func CycleMessage1()  {
	request  := &model.RequestVo{
		UserId : proto.String("lee"),
		Source : model.RequestVo_WEB.Enum(),
		Content : proto.String("hi, guys"),
		SourceId : proto.String("asd"),
	}
	requestPacket , err := proto.Marshal(request)
	if err != nil {
		log.Fatal("marshal error:", err)
	}
	result, err2 := json.Marshal(requestPacket)
	if err2 != nil {
		log.Fatal("json marshal error:", err2)
	}
	print("result : ", result)

	jsonObj := &model.RequestVo{}
	if err := json.Unmarshal([]byte(string(result)), &jsonObj); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(jsonObj)
		fmt.Println(jsonObj.Content)
	}
	fmt.Println("jsonObj=",jsonObj)
	fmt.Println("jsonObj.Content=", jsonObj.Content)

	//response := &model.RequestVo{}
	//err1 := proto.Unmarshal(jsonObj, response)
	//if err1 != nil {
	//	log.Fatal("unmarshaling error: ", err1)
	//}
	//fmt.Println("response= ", response)
}