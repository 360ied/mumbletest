package mumbleclient

import (
	"mumbletest/mumbleclient/mumbleprotocol"
	"time"
)

func (mc *MumbleClient) loopPing() error {
	for {
		time.Sleep(5 * time.Second)
		unix := uint64(time.Now().Unix())

		if err := mc.SendPacket(mumbleprotocol.IDPing, &mumbleprotocol.Ping{
			Timestamp: &unix,
		}); err != nil {
			return err
		}
	}
}
