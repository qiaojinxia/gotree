package main

import (
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"rpccode/handler"
	"rpccode/protofunc"
	"rpccode/router"
	"rpccode/utils"
	"runtime/debug"
	"time"
)

/**
 * Created by @CaomaoBoy on 2021/5/4.
 *  email:<115882934@qq.com>
 */

const (
	port = ":7777"
)

type client struct{} //服务对象

func (c *client) PushMsgStream(server protofunc.ClientMsg_PushMsgStreamServer) error {
	uuid := uuid.NewV4().String()
	client := handler.NewSession(server, time.Second*60)
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
		delete(handler.SessionManger, uuid)
	}()
	handler.SessionManger[uuid] = client
	for {
		msg, err := client.Recv()
		if err != nil {
			panic(err)
		}
		//待改进 协程池
		utils.Go(func() {
			router.HandlerRequest(msg, client)
		})
	}

}

//func(c *client) PushMsg(ctx context.Context, msg *protomsg.RpcMessageReq) (rsp *protomsg.RpcMessageRsp, err error){
//	xmsg := msg.Msg
//	xmsg.ID = uuid.NewV4().String()
//	msgStr,err := utils.ObjectTostring(msg)
//	if err != nil{
//		panic(error1.RPCERROR)
//	}
//	if err := redisClient.Instance.LPush(ctx,fmt.Sprintf("rpc:%s",xmsg.NodeID),msgStr).Err();err != nil{
//		panic(err)
//	}
//	dataStr,err := redisClient.Instance.BRPop(ctx, time.Duration(xmsg.Timeout),fmt.Sprintf("rpc:%s:%s",xmsg.NodeID,xmsg.ID)).Result()
//	if err != nil{
//		log.Println("RpcClientFunc error:",err.Error())
//		debug.PrintStack()
//		return
//	}
//	if len(dataStr) != 0{
//
//	}
//	return nil,nil
//}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() //起一个服务

	protofunc.RegisterClientMsgServer(s, &client{})
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
