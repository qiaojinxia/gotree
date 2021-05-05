package main

import (
	"context"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	pb "rpccode/protofunc"
	msg "rpccode/protomsg"
	"time"
)

/**
 * Created by @CaomaoBoy on 2021/5/5.
 *  email:<115882934@qq.com>
 */

const (
	address = "localhost:7777"
)

func main() {
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewClientMsgClient(conn)
	// 1秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	r, err := c.PushMsgStream(ctx)

	go func() {
		for {
			data, _ := r.Recv()
			if data == nil {
				continue
			}
			msgx := &msg.AddMonyBodyRsp{}
			err := proto.Unmarshal([]byte(data.Data), msgx)
			if err != nil {
				panic(err)
			}
			log.Printf("消息类型化 %s 返回数据 %v ", data.Msgtype, msgx)
		}
	}()

	go func() {
		msgBody := &msg.AddMonyBody{
			Amount: 10,
		}
		data, err := proto.Marshal(msgBody)
		if err != nil {
			panic(err)
		}
		for i := 0; i < 2; i++ {
			r.Send(&msg.RpcRequestMsg{
				Msgtype: msg.RpcRequestMsg_AddMony,
				Data:    string(data),
			})
			time.Sleep(time.Second)
		}
	}()

	select {}

}
