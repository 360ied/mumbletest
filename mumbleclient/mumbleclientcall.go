package mumbleclient

import (
	"fmt"
	"mumbletest/bufferhelpers"
	"mumbletest/bufferpool"
	"mumbletest/mumbleclient/mumbleprotocol"

	"google.golang.org/protobuf/proto"
)

// should be run in its own goroutine
// error returned is always non-nil
func (mc *MumbleClient) receivePackets() error {
	for {
		if err := receivePacket(mc); err != nil {
			return err
		}
	}
}

// should only be called by receivePackets
func receivePacket(mc *MumbleClient) error {
	// https://mumble-protocol.readthedocs.io/en/latest/protocol_stack_tcp.html

	const maxPacketLength = 8*1024*1024 - 1

	packetID, err := bufferhelpers.ReadBEUint16(mc.connReader)
	if err != nil {
		return err
	}

	packetLength, err := bufferhelpers.ReadBEUint32(mc.connReader)
	if err != nil {
		return err
	}

	if packetLength > maxPacketLength {
		return fmt.Errorf("mc.receivePackets: Invalid Packet Length: %d > Maximum Packet Length %d", packetLength, maxPacketLength)
	}

	buf, err := bufferhelpers.ReadLenIntoBuf(mc.connReader, int64(packetLength))
	if err != nil {
		return err
	}
	defer bufferpool.PutBuffer(buf)

	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()

	switch mumbleprotocol.MumblePacketID(packetID) {
	case mumbleprotocol.IDVersion:
		message := mumbleprotocol.Version{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageVersionCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageVersionCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageVersionCallbacks = nil
	case mumbleprotocol.IDUDPTunnel:
		message := mumbleprotocol.UDPTunnel{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageUDPTunnelCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageUDPTunnelCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageUDPTunnelCallbacks = nil
	case mumbleprotocol.IDAuthenticate:
		message := mumbleprotocol.Authenticate{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageAuthenticateCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageAuthenticateCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageAuthenticateCallbacks = nil
	case mumbleprotocol.IDPing:
		message := mumbleprotocol.Ping{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messagePingCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessagePingCallbacks {
			go callback(&message)
		}
		mc.singleCallMessagePingCallbacks = nil
	case mumbleprotocol.IDReject:
		message := mumbleprotocol.Reject{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageRejectCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageRejectCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageRejectCallbacks = nil
	case mumbleprotocol.IDServerSync:
		message := mumbleprotocol.ServerSync{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageServerSyncCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageServerSyncCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageServerSyncCallbacks = nil
	case mumbleprotocol.IDChannelRemove:
		message := mumbleprotocol.ChannelRemove{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageChannelRemoveCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageChannelRemoveCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageChannelRemoveCallbacks = nil
	case mumbleprotocol.IDChannelState:
		message := mumbleprotocol.ChannelState{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageChannelStateCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageChannelStateCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageChannelStateCallbacks = nil
	case mumbleprotocol.IDUserRemove:
		message := mumbleprotocol.UserRemove{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageUserRemoveCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageUserRemoveCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageUserRemoveCallbacks = nil
	case mumbleprotocol.IDUserState:
		message := mumbleprotocol.UserState{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageUserStateCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageUserStateCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageUserStateCallbacks = nil
	case mumbleprotocol.IDBanList:
		message := mumbleprotocol.BanList{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageBanListCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageBanListCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageBanListCallbacks = nil
	case mumbleprotocol.IDTextMessage:
		message := mumbleprotocol.TextMessage{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageTextMessageCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageTextMessageCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageTextMessageCallbacks = nil
	case mumbleprotocol.IDPermissionDenied:
		message := mumbleprotocol.PermissionDenied{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messagePermissionDeniedCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessagePermissionDeniedCallbacks {
			go callback(&message)
		}
		mc.singleCallMessagePermissionDeniedCallbacks = nil
	case mumbleprotocol.IDACL:
		message := mumbleprotocol.ACL{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageACLCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageACLCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageACLCallbacks = nil
	case mumbleprotocol.IDQueryUsers:
		message := mumbleprotocol.QueryUsers{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageQueryUsersCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageQueryUsersCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageQueryUsersCallbacks = nil
	case mumbleprotocol.IDCryptSetup:
		message := mumbleprotocol.CryptSetup{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageCryptSetupCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageCryptSetupCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageCryptSetupCallbacks = nil
	case mumbleprotocol.IDContextActionModify:
		message := mumbleprotocol.ContextActionModify{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageContextActionModifyCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageContextActionModifyCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageContextActionModifyCallbacks = nil
	case mumbleprotocol.IDContextAction:
		message := mumbleprotocol.ContextAction{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageContextActionCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageContextActionCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageContextActionCallbacks = nil
	case mumbleprotocol.IDUserList:
		message := mumbleprotocol.UserList{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageUserListCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageUserListCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageUserListCallbacks = nil
	case mumbleprotocol.IDVoiceTarget:
		message := mumbleprotocol.VoiceTarget{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageVoiceTargetCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageVoiceTargetCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageVoiceTargetCallbacks = nil
	case mumbleprotocol.IDPermissionQuery:
		message := mumbleprotocol.PermissionQuery{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messagePermissionQueryCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessagePermissionQueryCallbacks {
			go callback(&message)
		}
		mc.singleCallMessagePermissionQueryCallbacks = nil
	case mumbleprotocol.IDCodecVersion:
		message := mumbleprotocol.CodecVersion{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageCodecVersionCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageCodecVersionCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageCodecVersionCallbacks = nil
	case mumbleprotocol.IDUserStats:
		message := mumbleprotocol.UserStats{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageUserStatsCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageUserStatsCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageUserStatsCallbacks = nil
	case mumbleprotocol.IDRequestBlob:
		message := mumbleprotocol.RequestBlob{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageRequestBlobCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageRequestBlobCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageRequestBlobCallbacks = nil
	case mumbleprotocol.IDServerConfig:
		message := mumbleprotocol.ServerConfig{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageServerConfigCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageServerConfigCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageServerConfigCallbacks = nil
	case mumbleprotocol.IDSuggestConfig:
		message := mumbleprotocol.SuggestConfig{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messageSuggestConfigCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessageSuggestConfigCallbacks {
			go callback(&message)
		}
		mc.singleCallMessageSuggestConfigCallbacks = nil
	case mumbleprotocol.IDPluginDataTransmission:
		message := mumbleprotocol.PluginDataTransmission{}
		if err := proto.Unmarshal(buf.Bytes(), &message); err != nil {
			return err
		}

		for _, callback := range mc.messagePluginDataTransmissionCallbacks {
			go callback(&message)
		}

		for _, callback := range mc.singleCallMessagePluginDataTransmissionCallbacks {
			go callback(&message)
		}
		mc.singleCallMessagePluginDataTransmissionCallbacks = nil
	}

	return nil
}
