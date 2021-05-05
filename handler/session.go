package handler

import (
	"context"
	"rpccode/protofunc"
	"rpccode/protomsg"
	"time"
)

/**
 * Created by @CaomaoBoy on 2021/5/4.
 *  email:<115882934@qq.com>
 */

var SessionManger map[string]*Session

func init() {
	SessionManger = make(map[string]*Session)
}

type Session struct {
	protofunc.ClientMsg_PushMsgStreamServer
	*User
	ctx context.Context
}

func (s *Session) SendMsg(msg *protomsg.RpcResponseMsg) error {
	return s.Send(msg)
}

func NewSession(clientMsg_PushMsgStreamServer protofunc.ClientMsg_PushMsgStreamServer, userValid time.Duration) *Session {
	se := &Session{ClientMsg_PushMsgStreamServer: clientMsg_PushMsgStreamServer}
	se.User = &User{}
	se.LoginTime = time.Now().UnixNano() / 1e6
	ctx, _ := context.WithTimeout(context.Background(), userValid)
	se.ctx = ctx
	return se
}

type User struct {
	ID        string
	LoginTime int64
	Coin      int
}
