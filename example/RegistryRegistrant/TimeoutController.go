package RegistryRegistrant

import (
	"gogistery/protocol"
	"math"
	"time"
)

type TimeoutController struct {
	minT time.Duration //最小Timeout
	maxT time.Duration //最大Timeout
	cT   float64       //从最小到最大的增长系数
	tMap map[string]time.Duration
}

func NewTimeoutController(minT time.Duration, maxT time.Duration, cT float64) *TimeoutController {
	return &TimeoutController{minT, maxT, cT,
		make(map[string]time.Duration)}
}

func (p TimeoutController) TimeoutForNew(request protocol.Request) time.Duration {
	p.tMap[request.RegistrantInfo.GetRegistrantID()] = p.minT
	return p.minT

}
func (p TimeoutController) TimeoutForUpdate(request protocol.Request) time.Duration {
	t := p.tMap[request.RegistrantInfo.GetRegistrantID()]
	t += time.Duration(math.Floor(float64(p.maxT-t) / p.cT))
	p.tMap[request.RegistrantInfo.GetRegistrantID()] = t
	return t

}