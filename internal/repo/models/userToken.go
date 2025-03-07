package models

import "encoding/json"

type UserToken struct {
	UserID    uint   `json:"user_id"`    // 用户ID
	Account   string `json:"account"`    // 账号
	Nickname  string `json:"nickname"`   // 昵称
	ExpiredAt int64  `json:"expired_at"` // 过期时间，时间戳
	Ticket    string `json:"ticket"`     // 用户登录票据
}

// MarshalBinary Redis序列化
func (obj *UserToken) MarshalBinary() ([]byte, error) {
	return json.Marshal(obj)
}

// UnmarshalBinary Redis反序列化
func (obj *UserToken) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &obj)
}
