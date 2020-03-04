package requester

import (
	"context"
	"github.com/yindaheng98/gogistry/protocol"
	"time"
)

type Heart struct {
	beater      HeartBeater
	requester   *requester
	Handlers    *handlers
	startedChan chan bool //启动时关闭此通道
}

func NewHeart(beater HeartBeater, RequestProto protocol.RequestProtocol) *Heart {
	heart := &Heart{beater,
		newRequester(RequestProto),
		newEvents(),
		make(chan bool, 1)}
	heart.requester.RetryHandler = func(request protocol.TobeSendRequest, err error) {
		heart.Handlers.RetryHandler(request, err)
	}
	return heart
}

//开始心跳，直到最后由协议主动停止心跳或出错才返回
func (h *Heart) RunBeating(ctx context.Context,
	initRequest protocol.TobeSendRequest, initTimeout time.Duration, initRetryN uint64) error {
	defer func() { recover() }()
	close(h.startedChan) //如果此通道被关闭，则说明已有一个RunBeating在运行，会直接退出
	request, Timeout, RetryN := initRequest, initTimeout, initRetryN
	established := false
	var lastResponse protocol.Response
	var err error = nil
	defer func() {
		if established {
			h.Handlers.DisconnectionHandler(lastResponse, err)
		}
	}()
	run := true
	for run {
		response, err, timeout, retryN := h.requester.Send(ctx, request, Timeout, RetryN)
		if err != nil {
			return err
		}
		lastResponse = response
		if established { //如果已经达成过连接就触发更新事件
			h.Handlers.UpdateConnectionHandler(response)
		}
		run = false
		h.beater.Beat(response, timeout, retryN,
			func(requestB protocol.TobeSendRequest, timeoutB time.Duration, retryNB uint64) {
				request, Timeout, RetryN = requestB, timeoutB, retryNB
				run = true
			})
		if run { //只有上级协议判定可以继续进行接下来的连接才能视为连接达成
			if !established { //此时可以触发新建连接事件
				h.Handlers.NewConnectionHandler(response)
			}
			established = true //并且设置连接达成标记
		}
	}
	return nil
}
