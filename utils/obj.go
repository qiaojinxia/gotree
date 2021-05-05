package utils

import (
	"github.com/golang/protobuf/proto"
	"log"
	"rpccode/config"
	"sync"
	"sync/atomic"
)

/**
 * Created by @CaomaoBoy on 2021/5/4.
 *  email:<115882934@qq.com>
 */

var total = make(chan struct{}, 100)

func ObjectTostring(msg proto.Message) (string, error) {
	data, err := proto.Marshal(msg)
	return string(data), err
}

var poolChan chan func()
var poolGoCount int32
var waitAll sync.WaitGroup
var goid uint32
var gocount int32

func init() {
	poolChan = make(chan func(), 1)
	waitAll = sync.WaitGroup{}
}

var stopChanForGo = make(chan struct{})

func Go(fn func()) {
	pc := config.PoolSize + 1
	select {
	case poolChan <- fn:
		return
	default:
		pc = atomic.AddInt32(&poolGoCount, 1)
		if pc > config.PoolSize {
			atomic.AddInt32(&poolGoCount, -1)
		}
	}
	waitAll.Add(1)
	//id := atomic.AddUint32(&goid, 1)
	c := atomic.AddInt32(&gocount, 1)
	go func() {
		Try(fn, nil)
		for pc <= config.PoolSize {
			select {
			case <-stopChanForGo:
				pc = config.PoolSize + 1
			case nfn := <-poolChan:
				Try(nfn, nil)
			}
		}
		waitAll.Done()
		c = atomic.AddInt32(&gocount, -1)
	}()
}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			if handler == nil {
				log.Print("error catch:%v", err)
			} else {
				handler(err)
			}
			atomic.AddInt32(&config.PanicCount, 1)
		}
	}()
	fun()
}
