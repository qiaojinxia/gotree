package handler

//import (
//	"github.com/golang/protobuf/proto"
//	"reflect"
//	"rpccode/protomsg"
//	"rpccode/utils"
//)
//
///**
// * Created by @CaomaoBoy on 2021/5/4.
// *  email:<115882934@qq.com>
// */
//
//type Codec interface {
//	Encode([]byte)  (proto.Message,error)
//	Decode() ([]byte,error)
//	Register(proto.Message)
//}
//
//type CodeNo1 struct {
//	MsgTag map[string]proto.Message
//}
//
//func (c CodeNo1) Register(msg proto.Message) {
//	c.MsgTag[msg.String()] = msg
//}
//
//func (c CodeNo1) Decode(bytes []byte) (proto.Message,error){
//	msgTag := int32(utils.BytesToInt(bytes[:2]))
//	tmp :=  c.MsgTag
//	structPointer := tmp[protomsg.MsgBodyTag_name[msgTag]]
//	element := reflect.TypeOf(structPointer).Elem()
//	x := reflect.New(element).Interface().(proto.Message)
//	err := proto.Unmarshal(bytes[2:],x)
//	return structPointer,err
//}
//
//func (c CodeNo1) Encode(msg proto.Message) ([]byte, error) {
//	data := make([]byte,0,10)
//	msgTag :=  utils.IntToBytes(int(protomsg.MsgBodyTag_value[msg.String()]))
//	data = append(data,msgTag...)
//	msgData,err := proto.Marshal(msg)
//	if err != nil{
//		return nil,err
//	}
//	data = append(data, msgData...)
//	return data,nil
//}
