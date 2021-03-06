package control

import (
	"encoding/json"
	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
)

// Agent Available Actions
const (
	AgentStart           = "StartAgent"
	AgentStop            = "StopAgent"
	AgentRestart         = "RestartAgent"
	AgentStatus          = "StatusAgent"
	AgentConfigure       = "ConfigAgent"
	AgentAvailableAction = "AvailableActionAgent"
)

// Agent Possible Satuses
const (
	AgentStatusStarting       = "STARTING"
	AgentStatusWaitingForConf = "WAITING_FOT_CONFIGURATION"
	AgentStatusRegistred      = "REGISTRED"
	AgentStatusUnRegistred    = "UNREGISTRED"
	AgentStatusOnline         = "ONLINE"
	AgentStatusOffline        = "OFFLINE"
	AgentStatusOnError        = "ON_ERROR"
)

// Agent Message
// This message is used for agent control
type Agent struct {
	Action  string `json:"Action"`
	UUID    string `json:"UUID"`
	Token   string `json:"Token"`
	Payload []byte `json:"Payload"`
}

func (a *Agent) GetAction() string {
	return a.Action
}

func (a *Agent) GetUUID() string {
	return a.UUID
}

func (a *Agent) GetToken() string {
	return a.Token
}

func (a *Agent) GetTypeName() string {
	return TypeAgent
}

func (a *Agent) NewMessage(tenantToken string, uuid string, action string) *Agent {
	a.Token = tenantToken
	a.UUID = uuid
	a.Action = action

	return a
}

func (a *Agent) WithPayload(i interface{}) *Agent {
	b, err := json.Marshal(i)
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.ErrorStack(err),
		}).Fatal("Error running agent")
		return nil
	}
	a.Payload = b
	return a
}
