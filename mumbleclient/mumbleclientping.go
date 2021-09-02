package mumbleclient

import (
	"mumbletest/mumbleclient/mumbleprotocol"
	"time"
)

func (mc *MumbleClient) loopPing() error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for timestamp := <-ticker.C; true; {
		unix := uint64(timestamp.Unix())

		if err := mc.SendPacket(mumbleprotocol.IDPing, &mumbleprotocol.Ping{
			Timestamp: &unix,
		}); err != nil {
			return err
		}
	}

	panic("unreachable")
}
