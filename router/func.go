package router

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	uuid "github.com/satori/go.uuid"
	"rpccode/handler"
	"rpccode/protomsg"
	"rpccode/redisClient"
	"rpccode/utils"
)

/**
 * Created by @CaomaoBoy on 2021/5/4.
 *  email:<115882934@qq.com>
 */

type RoutersManger map[int]Server

var Routers RoutersManger

func init() {
	Routers = make(map[int]Server)
	Routers.Register(&Server1{})
	Routers.Register(&ServerRegister{})
}

func (r RoutersManger) Register(s Server) {
	r[s.Alias()] = s
}

type Server interface {
	Handler() error
	Alias() int
	Init(*handler.Session, []byte)
}

type Server1 struct {
	msg []byte
	*handler.Session
}

func (s *Server1) Init(se *handler.Session, data []byte) {
	s.msg = data
	s.Session = se
}

func (s *Server1) Handler() error {
	msgBody := &protomsg.AddMonyBody{}
	err := proto.Unmarshal(s.msg, msgBody)
	if err != nil {
		panic(err)
	}
	s.User.Coin += int(msgBody.Amount)

	msg := &protomsg.AddMonyBodyRsp{
		Amount:  int32(s.User.Coin),
		RespMsg: "add Success!",
	}
	msgS, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	s.Send(&protomsg.RpcResponseMsg{
		Msgtype: "加钱消息!",
		Data:    string(msgS),
		Error:   "",
	})
	return nil
}

func (s Server1) Alias() int {
	return int(protomsg.MsgBodyTag_AddMony)
}

type ServerRegister struct {
	msg []byte
	*handler.Session
}

func (s *ServerRegister) Init(se *handler.Session, data []byte) {
	s.msg = data
	s.Session = se
}

func (s ServerRegister) Alias() int {
	return int(protomsg.MsgBodyTag_RegisterReq)
}

func (s *ServerRegister) Handler() error {
	msgBody := &protomsg.RegisterBodyReq{}
	err := proto.Unmarshal(s.msg, msgBody)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	ID := uuid.NewV4().String()

	type Register struct {
		ID       string `json:"id"`
		UserName string `json:"name"`
		Passwd   string `json:"passwd"`
		NickName string `json:"nick_name"`
	}
	a := &Register{}
	a.ID = ID
	a.NickName = msgBody.NickName
	a.UserName = msgBody.UserName
	a.Passwd = msgBody.PassWd
	n_a, err := utils.ToMap(a, "json")
	if err != nil {
		panic(err)
	}
	err = redisClient.Instance.HSet(ctx, fmt.Sprintf("UserName:%s", ID), n_a).Err()
	if err != nil {
		panic(err)
	}
	msg := &protomsg.RegisterBodyRsp{
		ID:       ID,
		UserName: msgBody.UserName,
		NickName: msgBody.NickName,
	}
	msgS, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	s.Send(&protomsg.RpcResponseMsg{
		Msgtype: "加钱消息!",
		Data:    string(msgS),
		Error:   "",
	})
	return nil
}
