package mumbleclient

import "mumbletest/mumbleclient/mumbleprotocol"

func (mc *MumbleClient) OnMessageVersion(callback func(mumbleprotocol.Version)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageVersionCallbacks = append(mc.messageVersionCallbacks, callback)
}

func (mc *MumbleClient) OnMessageUDPTunnel(callback func(mumbleprotocol.UDPTunnel)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageUDPTunnelCallbacks = append(mc.messageUDPTunnelCallbacks, callback)
}

func (mc *MumbleClient) OnMessageAuthenticate(callback func(mumbleprotocol.Authenticate)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageAuthenticateCallbacks = append(mc.messageAuthenticateCallbacks, callback)
}

func (mc *MumbleClient) OnMessagePing(callback func(mumbleprotocol.Ping)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messagePingCallbacks = append(mc.messagePingCallbacks, callback)
}

func (mc *MumbleClient) OnMessageReject(callback func(mumbleprotocol.Reject)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageRejectCallbacks = append(mc.messageRejectCallbacks, callback)
}

func (mc *MumbleClient) OnMessageServerSync(callback func(mumbleprotocol.ServerSync)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageServerSyncCallbacks = append(mc.messageServerSyncCallbacks, callback)
}

func (mc *MumbleClient) OnMessageChannelRemove(callback func(mumbleprotocol.ChannelRemove)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageChannelRemoveCallbacks = append(mc.messageChannelRemoveCallbacks, callback)
}

func (mc *MumbleClient) OnMessageChannelState(callback func(mumbleprotocol.ChannelState)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageChannelStateCallbacks = append(mc.messageChannelStateCallbacks, callback)
}

func (mc *MumbleClient) OnMessageUserRemove(callback func(mumbleprotocol.UserRemove)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageUserRemoveCallbacks = append(mc.messageUserRemoveCallbacks, callback)
}

func (mc *MumbleClient) OnMessageUserState(callback func(mumbleprotocol.UserState)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageUserStateCallbacks = append(mc.messageUserStateCallbacks, callback)
}

func (mc *MumbleClient) OnMessageBanList(callback func(mumbleprotocol.BanList)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageBanListCallbacks = append(mc.messageBanListCallbacks, callback)
}

func (mc *MumbleClient) OnMessageTextMessage(callback func(mumbleprotocol.TextMessage)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageTextMessageCallbacks = append(mc.messageTextMessageCallbacks, callback)
}

func (mc *MumbleClient) OnMessagePermissionDenied(callback func(mumbleprotocol.PermissionDenied)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messagePermissionDeniedCallbacks = append(mc.messagePermissionDeniedCallbacks, callback)
}

func (mc *MumbleClient) OnMessageACL(callback func(mumbleprotocol.ACL)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageACLCallbacks = append(mc.messageACLCallbacks, callback)
}

func (mc *MumbleClient) OnMessageQueryUsers(callback func(mumbleprotocol.QueryUsers)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageQueryUsersCallbacks = append(mc.messageQueryUsersCallbacks, callback)
}

func (mc *MumbleClient) OnMessageCryptSetup(callback func(mumbleprotocol.CryptSetup)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageCryptSetupCallbacks = append(mc.messageCryptSetupCallbacks, callback)
}

func (mc *MumbleClient) OnMessageContextActionModify(callback func(mumbleprotocol.ContextActionModify)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageContextActionModifyCallbacks = append(mc.messageContextActionModifyCallbacks, callback)
}

func (mc *MumbleClient) OnMessageContextAction(callback func(mumbleprotocol.ContextAction)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageContextActionCallbacks = append(mc.messageContextActionCallbacks, callback)
}

func (mc *MumbleClient) OnMessageUserList(callback func(mumbleprotocol.UserList)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageUserListCallbacks = append(mc.messageUserListCallbacks, callback)
}

func (mc *MumbleClient) OnMessageVoiceTarget(callback func(mumbleprotocol.VoiceTarget)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageVoiceTargetCallbacks = append(mc.messageVoiceTargetCallbacks, callback)
}

func (mc *MumbleClient) OnMessagePermissionQuery(callback func(mumbleprotocol.PermissionQuery)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messagePermissionQueryCallbacks = append(mc.messagePermissionQueryCallbacks, callback)
}

func (mc *MumbleClient) OnMessageCodecVersion(callback func(mumbleprotocol.CodecVersion)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageCodecVersionCallbacks = append(mc.messageCodecVersionCallbacks, callback)
}

func (mc *MumbleClient) OnMessageUserStats(callback func(mumbleprotocol.UserStats)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageUserStatsCallbacks = append(mc.messageUserStatsCallbacks, callback)
}

func (mc *MumbleClient) OnMessageRequestBlob(callback func(mumbleprotocol.RequestBlob)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageRequestBlobCallbacks = append(mc.messageRequestBlobCallbacks, callback)
}

func (mc *MumbleClient) OnMessageServerConfig(callback func(mumbleprotocol.ServerConfig)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageServerConfigCallbacks = append(mc.messageServerConfigCallbacks, callback)
}

func (mc *MumbleClient) OnMessageSuggestConfig(callback func(mumbleprotocol.SuggestConfig)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messageSuggestConfigCallbacks = append(mc.messageSuggestConfigCallbacks, callback)
}

func (mc *MumbleClient) OnMessagePluginDataTransmission(callback func(mumbleprotocol.PluginDataTransmission)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.messagePluginDataTransmissionCallbacks = append(mc.messagePluginDataTransmissionCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageVersion(callback func(mumbleprotocol.Version)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageVersionCallbacks = append(mc.singleCallMessageVersionCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageUDPTunnel(callback func(mumbleprotocol.UDPTunnel)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageUDPTunnelCallbacks = append(mc.singleCallMessageUDPTunnelCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageAuthenticate(callback func(mumbleprotocol.Authenticate)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageAuthenticateCallbacks = append(mc.singleCallMessageAuthenticateCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessagePing(callback func(mumbleprotocol.Ping)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessagePingCallbacks = append(mc.singleCallMessagePingCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageReject(callback func(mumbleprotocol.Reject)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageRejectCallbacks = append(mc.singleCallMessageRejectCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageServerSync(callback func(mumbleprotocol.ServerSync)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageServerSyncCallbacks = append(mc.singleCallMessageServerSyncCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageChannelRemove(callback func(mumbleprotocol.ChannelRemove)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageChannelRemoveCallbacks = append(mc.singleCallMessageChannelRemoveCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageChannelState(callback func(mumbleprotocol.ChannelState)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageChannelStateCallbacks = append(mc.singleCallMessageChannelStateCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageUserRemove(callback func(mumbleprotocol.UserRemove)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageUserRemoveCallbacks = append(mc.singleCallMessageUserRemoveCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageUserState(callback func(mumbleprotocol.UserState)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageUserStateCallbacks = append(mc.singleCallMessageUserStateCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageBanList(callback func(mumbleprotocol.BanList)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageBanListCallbacks = append(mc.singleCallMessageBanListCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageTextMessage(callback func(mumbleprotocol.TextMessage)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageTextMessageCallbacks = append(mc.singleCallMessageTextMessageCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessagePermissionDenied(callback func(mumbleprotocol.PermissionDenied)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessagePermissionDeniedCallbacks = append(mc.singleCallMessagePermissionDeniedCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageACL(callback func(mumbleprotocol.ACL)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageACLCallbacks = append(mc.singleCallMessageACLCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageQueryUsers(callback func(mumbleprotocol.QueryUsers)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageQueryUsersCallbacks = append(mc.singleCallMessageQueryUsersCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageCryptSetup(callback func(mumbleprotocol.CryptSetup)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageCryptSetupCallbacks = append(mc.singleCallMessageCryptSetupCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageContextActionModify(callback func(mumbleprotocol.ContextActionModify)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageContextActionModifyCallbacks = append(mc.singleCallMessageContextActionModifyCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageContextAction(callback func(mumbleprotocol.ContextAction)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageContextActionCallbacks = append(mc.singleCallMessageContextActionCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageUserList(callback func(mumbleprotocol.UserList)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageUserListCallbacks = append(mc.singleCallMessageUserListCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageVoiceTarget(callback func(mumbleprotocol.VoiceTarget)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageVoiceTargetCallbacks = append(mc.singleCallMessageVoiceTargetCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessagePermissionQuery(callback func(mumbleprotocol.PermissionQuery)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessagePermissionQueryCallbacks = append(mc.singleCallMessagePermissionQueryCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageCodecVersion(callback func(mumbleprotocol.CodecVersion)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageCodecVersionCallbacks = append(mc.singleCallMessageCodecVersionCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageUserStats(callback func(mumbleprotocol.UserStats)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageUserStatsCallbacks = append(mc.singleCallMessageUserStatsCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageRequestBlob(callback func(mumbleprotocol.RequestBlob)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageRequestBlobCallbacks = append(mc.singleCallMessageRequestBlobCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageServerConfig(callback func(mumbleprotocol.ServerConfig)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageServerConfigCallbacks = append(mc.singleCallMessageServerConfigCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessageSuggestConfig(callback func(mumbleprotocol.SuggestConfig)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessageSuggestConfigCallbacks = append(mc.singleCallMessageSuggestConfigCallbacks, callback)
}

func (mc *MumbleClient) SingleCallOnMessagePluginDataTransmission(callback func(mumbleprotocol.PluginDataTransmission)) {
	mc.messageCallbacksMu.Lock()
	defer mc.messageCallbacksMu.Unlock()
	mc.singleCallMessagePluginDataTransmissionCallbacks = append(mc.singleCallMessagePluginDataTransmissionCallbacks, callback)
}
