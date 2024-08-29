package main

import (
	"flag"
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"log"
	"sync"
)

type chatServer struct {
	gnet.BuiltinEventEngine

	// 연결된 커넥션을 보관하는 맵
	cliMap sync.Map
}

func (cs *chatServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("client connected. address:%s", c.RemoteAddr().String())
	// 새로운 연결이 되면 커넥셔능ㄹ 맵에 보관한다.
	cs.cliMap.Store(c, true)
	return nil, gnet.None
}

func (cs *chatServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	log.Printf("client disconnected. address:%s", c.RemoteAddr().String())
	// 연결이 해제되면 맵에서 삭제한다.
	if _, ok := cs.cliMap.LoadAndDelete(c); ok {
		log.Printf("connection removed")
	}
	return gnet.None
}

func (cs *chatServer) OnBoot(eng gnet.Engine) gnet.Action {
	log.Printf("chat server is listening\n")
	return gnet.None
}

func (cs *chatServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	// 데이터 수신 시 모든 커넥션에 데이터를 전송한다.
	cs.cliMap.Range(func(key, value any) bool {
		if conn, ok := key.(gnet.Conn); ok {
			conn.AsyncWrite(buf, nil)
		}
		return true
	})
	return gnet.None
}

func main() {
	var port int
	var multicore bool

	// Example command: go run server.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")

	chat := &chatServer{}
	log.Fatal(gnet.Run(chat, fmt.Sprintf("tcp://:%d", port), gnet.WithMulticore(multicore)))
}
