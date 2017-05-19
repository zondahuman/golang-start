package im

import (
	"golang-start/im/model"
	"github.com/golang/protobuf/proto"
	"log"
	"fmt"
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
	var result = string(requestPacket)
	print("result : ", result)


	response := &model.RequestVo{}
	err1 := proto.Unmarshal(requestPacket, response)
	if err1 != nil {
		log.Fatal("unmarshaling error: ", err1)
	}
	fmt.Println(request)
}