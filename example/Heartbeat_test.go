package example

import (
	"fmt"
	"gogistery/Heartbeat"
	"math/rand"
	"testing"
	"time"
)

var src = rand.NewSource(10)

func testReq(i uint64, logger func(string)) {
	s := "------TestRequester------>"
	requester := Heartbeat.NewRequester(&RequestBeatProtocol{&src, 30, 0})
	requester.Events.Retry.AddHandler(func(o Heartbeat.TobeSendRequest, err error) {
		logger(s + fmt.Sprintf("An retry was occured. error: %s", err.Error()))
	})
	requester.Events.Retry.Enable()
	response, err := requester.Send(Heartbeat.TobeSendRequest{
		Request: Request{fmt.Sprintf("%02d", i)},
		Option: RequestOption{
			fmt.Sprintf("%02d", i),
			fmt.Sprintf("%02d.%02d.%02d.%02d", i, i, i, i)}},
		time.Duration(5e8), /*********将该值调低可模拟超时情况**********/
		10)
	if err != nil {
		logger(s + fmt.Sprintf("No.%02d test failed. err is %s", i, err.Error()))
		return
	}
	logger(s + fmt.Sprintf("No.%02d sending test succeed. response is %s", i, response.(Response).Print()))
}

//单次Heartbeat
func TestRequester(t *testing.T) {
	for i := uint64(0); i < 10; i++ {
		testReq(i, func(s string) {
			t.Log(s)
		})
	}
	time.Sleep(1e9)
}

func testRes(i uint64, logger func(string)) {
	s := "------TestResponser------>"
	responser := Heartbeat.NewResponser(ResponseBeatProtocol{&src, 30, fmt.Sprintf("%d", i)})
	request, err, responseFunc := responser.Recv()
	d := time.Duration(rand.Int31n(1e3) * 1e3)
	if err != nil {
		logger(s + err.Error())
		time.Sleep(d)
		responseFunc(Heartbeat.TobeSendResponse{Response: Response{fmt.Sprintf("error%02d", i)},
			Option: ResponseOption{fmt.Sprintf("error%02d", i)}})
	} else {
		logger(s + fmt.Sprintf("A request %s arrived. Response will be sent back in %d",
			request.(Request).Print(), d))
		time.Sleep(d)
		responseFunc(Heartbeat.TobeSendResponse{Response: Response{fmt.Sprintf("%02d", i)},
			Option: ResponseOption{fmt.Sprintf("%02d", i)}})
	}
}

//单次Heartbeat
func TestResponser(t *testing.T) {
	for i := uint64(0); i < 30; i++ {
		testRes(i, func(s string) {
			t.Log(s)
		})
	}
	time.Sleep(1e6)
}
