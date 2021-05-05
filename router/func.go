package router

import (
	"github.com/golang/protobuf/proto"
	"rpccode/handler"
	"rpccode/protomsg"
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
	return int(protomsg.RpcRequestMsg_AddMony)
}
