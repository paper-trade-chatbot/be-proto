package websocketMsgType

type WebsocketMsgType int

const (
	Msg_Error        WebsocketMsgType = -1
	Msg_Chat         WebsocketMsgType = 1
	Msg_Order        WebsocketMsgType = 2
	Msg_Wallet       WebsocketMsgType = 3
	Msg_SystemNotify WebsocketMsgType = 4
	Msg_RiskWarning  WebsocketMsgType = 5
	Msg_RiskStopLoss WebsocketMsgType = 6
)
