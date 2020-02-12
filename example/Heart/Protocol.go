package Heart

import (
	"errors"
	"fmt"
	"gogistery/Protocol"
	"math/rand"
	"sync/atomic"
	"time"
)

var src = rand.NewSource(10)
var failRate int32 = 30

type Response struct {
	ID string
}

type Request struct {
	ID string
}

func (r Response) String() string {
	return fmt.Sprintf("Response{id:%s}", r.ID)
}
func (r Request) String() string {
	return fmt.Sprintf("Request{id:%s}", r.ID)
}

type RequestSendOption struct {
	ID   string
	Addr string
}

func (o RequestSendOption) String() string {
	return fmt.Sprintf("RequestSendOption{id:%s,addr:%s}", o.ID, o.Addr)
}

type ResponseSendOption struct {
	ID string
}

func (o ResponseSendOption) String() string {
	return fmt.Sprintf("ResponseSendOption{id:%s}", o.ID)
}

type RequestBeatProtocol struct {
	src       *rand.Source
	failRate  int32
	responseN uint32
}

func NewRequestBeatProtocol() *RequestBeatProtocol {
	return &RequestBeatProtocol{&src, failRate, 0}
}

func (t *RequestBeatProtocol) Request(requestChan <-chan Protocol.TobeSendRequest, responseChan chan<- Protocol.ReceivedResponse) {
	atomic.AddUint32(&t.responseN, 1)
	protoRequest := <-requestChan
	request, option := protoRequest.Request.(Request), protoRequest.Option.(RequestSendOption)
	s := "\n------RequestBeatProtocol------>"
	s += fmt.Sprintf("It was sending attempt %02d in protocol. %s is sending with %s. ",
		t.responseN, request.String(), option.String())
	timeout := time.Duration(rand.Int63n(1e3) * 1e6)
	s += fmt.Sprintf("Response will arrived in %d. ", timeout)
	defer func() {
		if recover() != nil {
			fmt.Print(s + "This Sending was timeout.")
		}
	}()
	r := rand.New(*t.src).Int31n(100)
	if r < t.failRate {
		fmt.Print(s + "This Sending was failed.")
		responseChan <- Protocol.ReceivedResponse{Error: errors.New(fmt.Sprintf(
			"Your fail rate is %d%%, but this random output is %02d, so failed.", t.failRate, r))}
		return
	}
	time.Sleep(timeout)
	responseChan <- Protocol.ReceivedResponse{Response: Response{fmt.Sprintf("%02d", t.responseN)}}
	fmt.Print(s + "This Sending was success.")
}

type ResponseBeatProtocol struct {
	src      *rand.Source
	failRate int32
	id       string
}

func NewResponseBeatProtocol(id string) *ResponseBeatProtocol {
	return &ResponseBeatProtocol{&src, failRate, id}
}

func (t ResponseBeatProtocol) Response(requestChan chan<- Protocol.ReceivedRequest, responseChan <-chan Protocol.TobeSendResponse) {
	time.Sleep(time.Duration(rand.Int31n(1e3) * 1e3))
	request := Request{t.id}
	s := "\n------ResponseBeatProtocol------>"
	s += fmt.Sprintf("A request %s arrived in protocol. ", request.String())

	r := rand.New(*t.src).Int31n(100)
	if r < t.failRate {
		s += "This Receiving was failed. "
		requestChan <- Protocol.ReceivedRequest{Error: errors.New(fmt.Sprintf(
			"Your fail rate is %d%%, but this random output is %02d, so failed.", t.failRate, r))}
	} else {
		requestChan <- Protocol.ReceivedRequest{Request: request}
		s += "This Receiving was success. "
	}

	protoResponse, ok := <-responseChan
	if !ok {
		fmt.Print(s + "But the Response was timeouted.")
	} else {
		response, option := protoResponse.Response.(Response), protoResponse.Option.(ResponseSendOption)
		fmt.Print(s + fmt.Sprintf("And the Response is %s, with the option %s",
			response.String(), option.String()))
	}
}
