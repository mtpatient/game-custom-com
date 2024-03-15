package service

type IMessage interface {
}

var localMessage IMessage

func RegisterMessage(i IMessage) {
	localMessage = i
}

func Message() IMessage {
	return localMessage
}
