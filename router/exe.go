package router

import (
	"reflect"
	"rpccode/handler"
	"rpccode/protomsg"
)

/**
 * Created by @CaomaoBoy on 2021/5/5.
 *  email:<115882934@qq.com>
 */

func HandlerRequest(msg *protomsg.RpcRequestMsg, session *handler.Session) {
	_, ok := session.Context().Deadline()
	if !ok {
		return
	}
	oServer := Routers[int(msg.Msgtype)]
	element := reflect.TypeOf(oServer).Elem()
	nServer := reflect.New(element).Interface().(Server)
	nServer.Init(session, []byte(msg.Data))
	err := nServer.Handler()
	if err != nil {
		panic(err)
	}
}
