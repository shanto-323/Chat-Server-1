package websocket

import "encoding/json"

type IncommingMessage struct {
	MsgType    string          `json:"type"`
	ReceiverId string          `json:"receiver_id,omitempty"`
	Payload    json.RawMessage `json:"payload,omitempty"`
	SenderId   string          `json:"sender_id,omitempty"`
}

type OutgoingMessage struct {
	MsgType  string          `json:"type"`
	SenderId string          `json:"sender_id,omitempty"`
	Payload  json.RawMessage `json:"payload,omitempty"`
}

type UserList struct {
	IdList []string `json:"id_list"`
}

type UserModel struct {
	Id       string `json:"id"`
	ConnAddr string `json:"conn_addr"`
}
