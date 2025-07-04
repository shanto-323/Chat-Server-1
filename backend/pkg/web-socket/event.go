package websocket

import (
	"encoding/json"

	"chat_app/backend/logger"

	"go.uber.org/zap"
)

const (
	TYPE_INFO  = "info"
	TYPE_CHAT  = "chat"
	TYPE_LIST  = "list"
	TYPE_ALIVE = "ping"
	TYPE_CLOSE = "close"
)

type Event struct {
	client *Client
	logger *logger.ZapLogger
}

func NewEvent(c *Client, logger *logger.ZapLogger) *Event {
	return &Event{
		client: c,
		logger: logger,
	}
}

func (e *Event) CreateMessage(payload []byte) *IncommingMessage {
	var message IncommingMessage
	if err := json.Unmarshal(payload, &message); err != nil {
		e.logger.Error("Error marshaling payload!")
		return nil
	}
	return &message
}

func (e *Event) ChatEvent(receiverClient *Client, message *IncommingMessage) {
	select {
	case receiverClient.msgPool <- *message:
		e.logger.Info("Message", zap.String("TYPE", TYPE_CHAT))
	default:
		e.logger.Error("Buffer is full!")
	}
}

func (e *Event) ListEvent(message *IncommingMessage) {
	c := e.client
	var list []string
	for i := range c.manager.clients {
		if c.manager.clients[i].id == c.id {
			continue
		}
		list = append(list, c.manager.clients[i].id)
	}

	payload, _ := json.Marshal(&UserList{IdList: list})
	message.Payload = payload

	select {
	case c.msgPool <- *message:
		e.logger.Info("Message", zap.String("TYPE", TYPE_LIST))
	default:
		e.logger.Error("Buffer is full!")
	}
}

func (e *Event) InfoEvent() {
	c := e.client
	message := IncommingMessage{
		MsgType:    TYPE_INFO,
		SenderId:   c.id,
		ReceiverId: c.id,
	}
	payload := UserModel{
		Id:       c.id,
		ConnAddr: c.conn.RemoteAddr().String(),
	}
	payloadJson, _ := json.Marshal(&payload)
	message.Payload = payloadJson

	select {
	case c.msgPool <- message:
		e.logger.Info("Info about client.")
	default:
		e.logger.Error("Buffer is full!")
	}
}
