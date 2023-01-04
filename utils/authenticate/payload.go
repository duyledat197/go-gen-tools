package authenticate

import (
	"encoding/json"
	"fmt"
	"time"
)

type Payload struct {
	UserID    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	Ip        string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
	ExpiredAt time.Time `json:"expired_at"`
	IssueAt   time.Time `json:"issue_at"`
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return fmt.Errorf("")
	}
	return nil
}

func (p *Payload) ToJSONString() (string, error) {
	m, err := json.Marshal(p)
	if err != nil {
		return "", nil
	}
	return string(m), nil
}

func (p *Payload) AddExpired(expirationTime time.Duration) {
	t := time.Now()
	p.ExpiredAt = t.Add(expirationTime)

}

func JSONStringToPayload(str string) (*Payload, error) {
	payload := &Payload{}
	if err := json.Unmarshal([]byte(str), payload); err != nil {
		return nil, err
	}
	return payload, nil
}
