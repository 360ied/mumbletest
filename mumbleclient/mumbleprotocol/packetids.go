package mumbleprotocol

type MumblePacketID uint16

const (
	IDVersion MumblePacketID = iota
	IDUDPTunnel
	IDAuthenticate
	IDPing
	IDReject
	IDServerSync
	IDChannelRemove
	IDChannelState
	IDUserRemove
	IDUserState
	IDBanList
	IDTextMessage
	IDPermissionDenied
	IDACL
	IDQueryUsers
	IDCryptSetup
	IDContextActionModify
	IDContextAction
	IDUserList
	IDVoiceTarget
	IDPermissionQuery
	IDCodecVersion
	IDUserStats
	IDRequestBlob
	IDServerConfig
	IDSuggestConfig
	IDPluginDataTransmission
)
