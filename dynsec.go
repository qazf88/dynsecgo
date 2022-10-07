package dynsecgo

import (
	"fmt"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type DynSec struct {
	mc          mqtt.Client
	command     *dynSecCommand
	always      bool
	timeOut     time.Duration //Millisecond
	subResponse chan []byte
}

type dynSecCommand struct {
}

func NewDynSecCommand() *dynSecCommand {
	return &dynSecCommand{}
}

// SetTimeout Millisecond
func (ds *DynSec) SetTimeout(timeout time.Duration) {
	ds.timeOut = timeout
}

func NewDynSecInternalClient(user, password, clientID, host string, port int, runAlways bool) *DynSec {

	if len(clientID) < 1 {
		clientID = randString(8)
	}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("mqtt://%s:%d", host, port))
	opts.SetClientID(clientID)
	opts.ConnectTimeout = 500 * time.Millisecond
	opts.WriteTimeout = 500 * time.Millisecond

	if len(user) > 0 {
		opts.SetUsername(user)
	}

	if len(password) > 0 {
		opts.SetPassword(password)
	}

	if runAlways {
		opts.CleanSession = false

	}

	mc := mqtt.NewClient(opts)

	return &DynSec{
		mc:          mc,
		command:     NewDynSecCommand(),
		always:      runAlways,
		timeOut:     500,
		subResponse: make(chan []byte),
	}
}

const (
	charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	dinSecPubTopic = "$CONTROL/dynamic-security/v1"
	dinSecSubTopic = "$CONTROL/dynamic-security/v1/response/#"
)

// subscribe
func (ds *DynSec) subscribe() error {

	if token := ds.mc.Subscribe(dinSecSubTopic, 0, func(client mqtt.Client, msg mqtt.Message) {
		ds.subResponse <- msg.Payload()
	}); token.Wait() && token.Error() != nil {

		return token.Error()
	}
	return nil
}

// publishCommand
func (ds *DynSec) sendCommand(request []byte) ([]byte, error) {

	var payload []byte
	flagTimeout := false

	if !ds.mc.IsConnected() {
		if token := ds.mc.Connect(); token.Wait() && token.Error() != nil {
			return nil, token.Error()
		}

		if err := ds.subscribe(); err != nil {
			return nil, err
		}
	}

	go func() {
		if token := ds.mc.Publish(dinSecPubTopic, 0, false, request); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			return
		}
	}()

loop2:
	for {
		select {
		case <-time.After(ds.timeOut * time.Millisecond):
			flagTimeout = true
			break loop2
		case payload = <-ds.subResponse:
			break loop2
		}
	}

	if !ds.always {
		if token := ds.mc.Unsubscribe(dinSecSubTopic); token.Wait() && token.Error() != nil {
			return nil, token.Error()
		}
		ds.mc.Disconnect(0)
	}

	if flagTimeout {
		return nil, fmt.Errorf("timeout wait response")
	}

	return payload, nil
}

// randString
func randString(length int) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand().Intn(len(charset))]
	}

	return string(b)
}

// seededRand
func seededRand() *rand.Rand {
	return rand.New(
		rand.NewSource(time.Now().UnixNano()))
}
