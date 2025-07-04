package ui

import "encoding/json"

const (
	TYPE_INFO  = "info"
	TYPE_CHAT  = "chat"
	TYPE_LIST  = "list"
	TYPE_ALIVE = "ping"
	TYPE_CLOSE = "close"
)

type Message struct {
	MsgType  string          `json:"type"`
	SenderId string          `json:"sender_id,omitempty"`
	Payload  json.RawMessage `json:"payload,omitempty"`
}

type OutgoingMessage struct {
	MsgType    string `json:"type"`
	ReceiverId string `json:"receiver_id,omitempty"`
	Payload    string `json:"payload,omitempty"`
	SenderId   string `json:"sender_id,omitempty"`
}

type UserList struct {
	IdList []string `json:"id_list"`
}

type UserModel struct {
	Id       string `json:"id"`
	ConnAddr string `json:"conn_addr"`
}
