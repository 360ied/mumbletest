package mumblebot

import (
	"crypto/tls"
	"log"
	"mumbletest/mumbleclient"
	"mumbletest/mumbleclient/mumbleprotocol"
)

type MumbleBot struct {
	Logger    *log.Logger
	ServerIP  string
	TLSConfig *tls.Config
	Username  string
	Password  string
}

func (mb *MumbleBot) Start() error {
	if mb.Logger == nil {
		mb.Logger = log.Default()
	}

	tlsConn, err := tls.Dial("tcp", mb.ServerIP, mb.TLSConfig)
	if err != nil {
		return err
	}
	defer tlsConn.Close()
	mb.Logger.Printf("Dialed %s", mb.ServerIP)

	return mumbleclient.ConnectMumble(tlsConn, mb.Username, mb.Password, nil, func(mc *mumbleclient.MumbleClient) {
		mb.Logger.Printf("Connected to %s as %s", mb.ServerIP, mb.Username)
		mc.OnMessageTextMessage(func(tm *mumbleprotocol.TextMessage) {
			mb.Logger.Printf("[MESSAGE] <%d> %s", tm.GetActor(), tm.GetMessage())
		})
	})
}
