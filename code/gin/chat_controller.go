package main

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"strconv"
	"time"
)

var recordCh = make(chan int)

var ShareCache string

type ChatRsp struct {
	Answer   string `json:"answer"`
	RecordID int64  `json:"record_id"`
}

var ErrCodeStop = 900000

type RPCResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

const (
	eol       = "$$END$$"
	keepalive = "$$KEEPALIVE$$"
	stop      = "stop"
)

var CtxStreamOutput = "STREAM-OUTPUT"
var CtxCancelStreamOutput = "CANCEL-STREAM-OUTPUT"

var CtxIsStreamV2 = "IS-STREAM-V2"

var StreamV2Middleware = func(c *gin.Context) {
	c.Set(CtxIsStreamV2, true)
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	c.Next()
}

func StopStream(c *gin.Context) {
	ShareCache = "hit"
	c.JSON(200, "OK")
	return
}

func ChatStream(c *gin.Context) {
	byteBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("ReadAll error:%s\n", err.Error())
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(byteBody))
	fmt.Printf("ChatStream Req:%s\n", byteBody)
	v2 := c.GetBool(CtxIsStreamV2)
	v, _ := c.Get("clientChan")
	clientChan := v.(ClientChan)

	ch := make(chan RPCResponse, 1024)
	cancelCh := make(chan RPCResponse, 1024)
	c.Set(CtxStreamOutput, ch)
	c.Set(CtxCancelStreamOutput, cancelCh)
	g := errgroup.Group{}

	ticker := time.NewTicker(2 * time.Second)            // 定时
	timer := time.NewTimer(5 * time.Minute)              // 超时
	stopTicker := time.NewTicker(250 * time.Millisecond) // 定时

	g.Go(func() error {
		Chat(c)
		return nil
	})

	g.Go(func() error {
		defer func() { clientChan <- eol }()
		defer ticker.Stop()
		defer stopTicker.Stop()
	stop:
		for {
			select {
			case <-stopTicker.C:
				cancelV, ok := <-cancelCh
				if ok {
					fmt.Printf("stopTicker receive...%+v \n", cancelV)

					stopStream := checkIsStop()
					if stopStream {
						fmt.Printf("check stop...%t \n", stopStream)
						cancelV.Code = ErrCodeStop
						s, _ := jsoniter.MarshalToString(cancelV)
						clientChan <- s
						break stop
					}
				}

			case <-ticker.C:
				clientChan <- keepalive
			case <-timer.C:
				fmt.Printf("time is up")
				break stop
			case v, ok := <-ch:
				if !ok {
					break stop
				}

				var r any = v
				if !v2 {
					if v.Code != 0 && v.Code != ErrCodeStop {
						r = map[string]string{"data": v.Message} // 错误处理
					} else if rsp, ok := v.Data.(ChatRsp); ok {
						r = map[string]any{"data": rsp.Answer, "record_id": strconv.Itoa(int(rsp.RecordID))} // 正常处理
					}
				}
				s, _ := jsoniter.MarshalToString(r)
				clientChan <- s
			}
		}
		return nil
	})

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-clientChan; ok {
			fmt.Printf("ChatStream msg: %v \n", msg)
			if msg == eol {
				fmt.Printf("ChatStream msg eol: %v \n", msg)
				return false
			}
			if msg == stop {
				fmt.Printf("ChatStream msg stop: %v \n", msg)
				return false
			}
			if msg == keepalive {
				w.Write([]byte(": keepalive\n"))
			} else {
				c.Render(-1, sse.Event{Data: msg, Retry: 30000})
			}
			return true
		}
		return false
	})

	waitErr := g.Wait()
	fmt.Printf("wait done:%+v\n", waitErr)
}

func checkIsStop() bool {
	return ShareCache == "hit"
}

func (stream *Event) listen() {
	for {
		select {
		case client := <-stream.NewClients:
			stream.TotalClients[client] = true
		case client := <-stream.ClosedClients:
			delete(stream.TotalClients, client)
			close(client)
		}
	}
}

func Chat(c *gin.Context) {
	var res string
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		res += cast.ToString(i)
		output(c, RPCResponse{Code: 0, Message: "OK", Data: res}, false)
	}
	output(c, RPCResponse{Code: 0, Message: "Done", Data: "Done"}, true)
}

func output(c *gin.Context, r RPCResponse, final bool) {
	v, _ := c.Get(CtxStreamOutput)
	v2, _ := c.Get(CtxCancelStreamOutput)

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	fmt.Printf("output final: %v, rsp: %v, v: %T\n", final, r, v)
	if ch, ok := v.(chan RPCResponse); ok {
		if final && r.Code == 0 {
			r.Code = ErrCodeStop
		}
		ch <- r
		if final {
			fmt.Printf("output final rsp: %v", r)
			close(ch)
		}
	} else if final {
		c.JSON(http.StatusOK, r)
	}

	if ch, ok := v2.(chan RPCResponse); ok {
		if final && r.Code == 0 {
			r.Code = ErrCodeStop
		}
		ch <- r
		if final {
			fmt.Printf("output final rsp: %v", r)
			close(ch)
		}
	}
}

type ClientChan chan string
type CancelChan chan string

func (stream *Event) serve() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientChan := make(ClientChan)
		stream.NewClients <- clientChan
		defer func() {
			stream.ClosedClients <- clientChan
		}()
		cancelChan := make(CancelChan)

		c.Set("clientChan", clientChan)
		c.Set("cancelChan", cancelChan)
		c.Next()
	}
}

type Event struct {
	NewClients    chan chan string     // New client connections
	ClosedClients chan chan string     // Closed client connections
	TotalClients  map[chan string]bool // Total client connections
}
