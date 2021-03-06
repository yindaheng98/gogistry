package protocol

import (
	"fmt"
	"github.com/yindaheng98/gogistry/protocol"
	"time"
)

type ResponseSendOption struct {
	Timestamp time.Time
}

func (o ResponseSendOption) String() string {
	return fmt.Sprintf(`{"type":"github.com/yindaheng98/gogistry/example/protocol.ResponseSendOption",
	"Timestamp":"%s"}`, o.Timestamp)
}

type RegistrantInfo struct {
	ID     string
	Type   string
	Option ResponseSendOption
}

func (info RegistrantInfo) GetRegistrantID() string {
	return info.ID
}
func (info RegistrantInfo) GetServiceType() string {
	return info.Type
}
func (info RegistrantInfo) GetResponseSendOption() protocol.ResponseSendOption {
	return info.Option
}
func (info RegistrantInfo) String() string {
	return fmt.Sprintf(`{"type":"github.com/yindaheng98/gogistry/example/protocol.RegistrantInfo",
	"ID":"%s","Type":"%s","Option":%s}`, info.ID, info.Type, info.Option.String())
}

type RequestSendOption struct {
	RequestAddr string
	Timestamp   time.Time
}

func (o RequestSendOption) String() string {
	return fmt.Sprintf(`{"type":"github.com/yindaheng98/gogistry/example/protocol.RequestSendOption",
	"RequestAddr":"%s","Timestamp":"%s"}`, o.RequestAddr, o.Timestamp)
}

type RegistryInfo struct {
	ID         string
	Type       string
	Option     RequestSendOption
	Candidates []protocol.RegistryInfo
}

func (info RegistryInfo) GetRegistryID() string {
	return info.ID
}
func (info RegistryInfo) GetServiceType() string {
	return info.Type
}
func (info RegistryInfo) GetRequestSendOption() protocol.RequestSendOption {
	return info.Option
}
func (info RegistryInfo) GetCandidates() []protocol.RegistryInfo {
	return info.Candidates
}
func (info RegistryInfo) String() string {
	Candidates := ""
	for _, RegistryInfo := range info.Candidates {
		Candidates += RegistryInfo.String() + ",\n\t"
	}
	if len(Candidates) >= 3 {
		Candidates = Candidates[0 : len(Candidates)-3]
	}
	Candidates = "[\n\t" + Candidates + "\n]"
	return fmt.Sprintf(`{"type":"github.com/yindaheng98/gogistry/example/protocol.RegistryInfo",
	"ID":"%s","Type":"%s","Option":%s,"Candidates":%s}`,
		info.ID, info.Type, info.Option.String(), Candidates)
}
