package dynsecgo

import (
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type DynSec struct {
	mc mqtt.Client
}

type acl struct {
	Acltype string `json:"acltype"`
	Topic   string `json:"topic"`
	Allow   bool   `json:"allow"`
}

type group struct {
	Groupname string `json:"groupname"`
	Priority  int    `json:"priority"`
}

type role struct {
	Rolename string `json:"rolename"`
	Priority int    `json:"priority"`
}

type commands struct {
	Commands interface{} `json:"commands"`
}

type response struct {
	Responses []result `json:"responses"`
}
type result struct {
	Error string `json:"error"`
}

type command struct {
	Command         string  `json:"command"`
	Username        string  `json:"username,omitempty"`
	Password        string  `json:"password,omitempty"`
	Clientid        string  `json:"clientid,omitempty"`
	Roles           []role  `json:"roles,omitempty"`
	Role            string  `json:"role,omitempty"`
	Groups          []group `json:"groups,omitempty"`
	Group           string  `json:"group,omitempty"`
	Clients         string  `json:"clients,omitempty"`
	Rolename        string  `json:"rolename,omitempty"`
	Groupname       string  `json:"groupname,omitempty"`
	Textname        string  `json:"textname,omitempty"`
	Priority        int     `json:"priority,omitempty"`
	Textdescription string  `json:"textdescription,omitempty"`
	Verbose         bool    `json:"verbose,omitempty"`
	Acls            []acl   `json:"acls,omitempty"`
}

func NewDinSec(mc mqtt.Client) *DynSec {
	return &DynSec{
		mc: mc,
	}
}

const (
	dinSecPubTopic = "$CONTROL/dynamic-security/v1"
	dinSecSubTopic = "$CONTROL/dynamic-security/v1/response/#"
)

// publishCommand
func publishCommand(mc mqtt.Client, request []byte) ([]byte, error) {

	var payload []byte
	ok := make(chan bool)

	if token := mc.Connect(); token.Wait() && token.Error() != nil {

		return nil, token.Error()
	}

	if token := mc.Subscribe(dinSecSubTopic, 0, func(client mqtt.Client, msg mqtt.Message) { payload = msg.Payload() }); token.Wait() && token.Error() != nil {

		return nil, token.Error()
	}

	if token := mc.Publish(dinSecPubTopic, 0, false, request); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
loop:
	for {
		select {
		case <-time.After(1500 * time.Millisecond):
			break loop
		case <-ok:
			break loop
		}
	}

	if token := mc.Unsubscribe(dinSecSubTopic); token.Wait() && token.Error() != nil {

		return nil, token.Error()
	}
	mc.Disconnect(100)
	return payload, nil
}
