package control

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

// Sink Available Actions
const (
	SinkStart   = "StartSink"
	SinkStop    = "StopSink"
	SinkRestart = "RestartSink"
	SinkStatus  = "GetSinkStatus"
	SinkMeta    = "GetSinkMeta"
	SinkSchema  = "GetSinkSchema"
)

// Sink Possible Satuses
const (
	SinkStatusOnError = "ON_ERROR"
	SinkStatusRunning = "RUNNING"
	SinkStatusWaiting = "WAITING"
)

// Sink Message
// This message is used for agent control
type Sink struct {
	Action  string `json:"action"`
	Name    string `json:"name"`
	Token   string `json:"token"`
	Payload []byte `json:"payload"`
}

func (a Sink) GetAction() string {
	return a.Action
}

func (a Sink) GetName() string {
	return a.Name
}

func (a Sink) GetToken() string {
	return a.Token
}

func (a Sink) GetTypeName() string {
	return "SINK_CONTROL"
}

func (a Sink) NewMessage(tenantToken string, uuid string, action string) *Sink {
	return &Sink{
		Token:  tenantToken,
		Name:   uuid,
		Action: action,
	}
}

func (a *Sink) WithPayload(i interface{}) *Sink {
	b, err := json.Marshal(i)
	if err != nil {
		logrus.Fatal(err)
	}
	a.Payload = b
	return a
}
