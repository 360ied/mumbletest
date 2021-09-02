package mumbleclient

import (
	"bufio"
	"mumbletest/bufferhelpers"
	"mumbletest/bufferpool"
	"mumbletest/mumbleclient/mumbleprotocol"
	"net"
	"sync"

	"google.golang.org/protobuf/proto"
)

type MumbleClient struct {
	conn          net.Conn      // TLS conn
	connReader    *bufio.Reader // TLS conn reader
	connWriteLock sync.Mutex

	messageCallbacksMu sync.Mutex

	messageVersionCallbacks                []func(mumbleprotocol.Version)
	messageUDPTunnelCallbacks              []func(mumbleprotocol.UDPTunnel)
	messageAuthenticateCallbacks           []func(mumbleprotocol.Authenticate)
	messagePingCallbacks                   []func(mumbleprotocol.Ping)
	messageRejectCallbacks                 []func(mumbleprotocol.Reject)
	messageServerSyncCallbacks             []func(mumbleprotocol.ServerSync)
	messageChannelRemoveCallbacks          []func(mumbleprotocol.ChannelRemove)
	messageChannelStateCallbacks           []func(mumbleprotocol.ChannelState)
	messageUserRemoveCallbacks             []func(mumbleprotocol.UserRemove)
	messageUserStateCallbacks              []func(mumbleprotocol.UserState)
	messageBanListCallbacks                []func(mumbleprotocol.BanList)
	messageTextMessageCallbacks            []func(mumbleprotocol.TextMessage)
	messagePermissionDeniedCallbacks       []func(mumbleprotocol.PermissionDenied)
	messageACLCallbacks                    []func(mumbleprotocol.ACL)
	messageQueryUsersCallbacks             []func(mumbleprotocol.QueryUsers)
	messageCryptSetupCallbacks             []func(mumbleprotocol.CryptSetup)
	messageContextActionModifyCallbacks    []func(mumbleprotocol.ContextActionModify)
	messageContextActionCallbacks          []func(mumbleprotocol.ContextAction)
	messageUserListCallbacks               []func(mumbleprotocol.UserList)
	messageVoiceTargetCallbacks            []func(mumbleprotocol.VoiceTarget)
	messagePermissionQueryCallbacks        []func(mumbleprotocol.PermissionQuery)
	messageCodecVersionCallbacks           []func(mumbleprotocol.CodecVersion)
	messageUserStatsCallbacks              []func(mumbleprotocol.UserStats)
	messageRequestBlobCallbacks            []func(mumbleprotocol.RequestBlob)
	messageServerConfigCallbacks           []func(mumbleprotocol.ServerConfig)
	messageSuggestConfigCallbacks          []func(mumbleprotocol.SuggestConfig)
	messagePluginDataTransmissionCallbacks []func(mumbleprotocol.PluginDataTransmission)

	singleCallMessageVersionCallbacks                []func(mumbleprotocol.Version)
	singleCallMessageUDPTunnelCallbacks              []func(mumbleprotocol.UDPTunnel)
	singleCallMessageAuthenticateCallbacks           []func(mumbleprotocol.Authenticate)
	singleCallMessagePingCallbacks                   []func(mumbleprotocol.Ping)
	singleCallMessageRejectCallbacks                 []func(mumbleprotocol.Reject)
	singleCallMessageServerSyncCallbacks             []func(mumbleprotocol.ServerSync)
	singleCallMessageChannelRemoveCallbacks          []func(mumbleprotocol.ChannelRemove)
	singleCallMessageChannelStateCallbacks           []func(mumbleprotocol.ChannelState)
	singleCallMessageUserRemoveCallbacks             []func(mumbleprotocol.UserRemove)
	singleCallMessageUserStateCallbacks              []func(mumbleprotocol.UserState)
	singleCallMessageBanListCallbacks                []func(mumbleprotocol.BanList)
	singleCallMessageTextMessageCallbacks            []func(mumbleprotocol.TextMessage)
	singleCallMessagePermissionDeniedCallbacks       []func(mumbleprotocol.PermissionDenied)
	singleCallMessageACLCallbacks                    []func(mumbleprotocol.ACL)
	singleCallMessageQueryUsersCallbacks             []func(mumbleprotocol.QueryUsers)
	singleCallMessageCryptSetupCallbacks             []func(mumbleprotocol.CryptSetup)
	singleCallMessageContextActionModifyCallbacks    []func(mumbleprotocol.ContextActionModify)
	singleCallMessageContextActionCallbacks          []func(mumbleprotocol.ContextAction)
	singleCallMessageUserListCallbacks               []func(mumbleprotocol.UserList)
	singleCallMessageVoiceTargetCallbacks            []func(mumbleprotocol.VoiceTarget)
	singleCallMessagePermissionQueryCallbacks        []func(mumbleprotocol.PermissionQuery)
	singleCallMessageCodecVersionCallbacks           []func(mumbleprotocol.CodecVersion)
	singleCallMessageUserStatsCallbacks              []func(mumbleprotocol.UserStats)
	singleCallMessageRequestBlobCallbacks            []func(mumbleprotocol.RequestBlob)
	singleCallMessageServerConfigCallbacks           []func(mumbleprotocol.ServerConfig)
	singleCallMessageSuggestConfigCallbacks          []func(mumbleprotocol.SuggestConfig)
	singleCallMessagePluginDataTransmissionCallbacks []func(mumbleprotocol.PluginDataTransmission)
}

// synchronous function
func (mc *MumbleClient) SendPacket(packetID mumbleprotocol.MumblePacketID, message proto.Message) error {
	b, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	buf := bufferpool.GetBuffer()
	defer bufferpool.PutBuffer(buf)

	bufferhelpers.WriteBEUint16(buf, uint16(packetID))
	bufferhelpers.WriteBEUint32(buf, uint32(len(b)))
	buf.Write(b)

	_, err = mc.conn.Write(buf.Bytes())
	return err
}

func ConnectMumble(tlsConn net.Conn, username string, password string, tokens []string) (mc *MumbleClient, err error) {
	mc = new(MumbleClient)
	mc.conn = tlsConn
	mc.connReader = bufio.NewReader(tlsConn)

	clientReleaseName := "mumbletest"
	err = mc.SendPacket(mumbleprotocol.IDVersion, &mumbleprotocol.Version{
		Version: (*uint32)(&CurrentMumbleVersion),
		Release: &clientReleaseName,
	})
	if err != nil {
		return
	}

}
