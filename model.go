package dynsecgo

import (
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	UserName        string  `json:"username"`
	Textname        string  `json:"textname,omitempty" example:"1"`
	Textdescription string  `json:"textdescription,omitempty" example:"2"`
	Disabled        bool    `json:"disabled"`
	Roles           []role  `json:"roles,omitempty"`
	Groups          []group `json:"groups,omitempty"`
}

type DynSec struct {
	mc          mqtt.Client
	command     *DynSecCommand
	always      bool
	timeOut     time.Duration //Millisecond
	subResponse chan []byte
}

type Acl struct {
	Acltype  string `json:"acltype"`
	Topic    string `json:"topic"`
	Priority int    `json:"priority"`
	Allow    bool   `json:"allow"`
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
	Error   *string `json:"error,omitempty"`
	Command string  `json:"command,omitempty"`
	Data    *data   `json:"data,omitempty"`
}

type command struct {
	Command         string  `json:"command,omitempty"`
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
	Disabled        bool    `json:"disabled,omitempty"`
	Acls            []Acl   `json:"acls,omitempty"`
}

type data struct {
	Client command `json:"client,omitempty"`
}
