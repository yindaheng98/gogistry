package gogistery

import (
	"github.com/yindaheng98/gogistry/protocol"
	"github.com/yindaheng98/gogistry/registrant"
	"github.com/yindaheng98/gogistry/registry"
)

//NewRegistry returns the pointer of an "registry".
func NewRegistry(
	Info protocol.RegistryInfo,
	maxRegistrants uint64,
	timeoutController registry.TimeoutController,
	ResponseProto protocol.ResponseProtocol) *registry.Registry {
	return registry.New(Info, maxRegistrants, timeoutController, ResponseProto)
}

//NewRegistry returns the pointer of an "registrant".
func NewRegistrant(
	Info protocol.RegistrantInfo,
	registryN uint64,
	CandidateList registrant.RegistryCandidateList,
	retryNController registrant.WaitTimeoutRetryNController,
	RequestProto protocol.RequestProtocol) *registrant.Registrant {
	return registrant.New(Info, registryN, CandidateList, retryNController, RequestProto)
}
