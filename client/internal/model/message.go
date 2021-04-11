package model

import (
	"encoding/json"
	"time"
)

type Message struct {
	AccountID int64           `json:"account_id"`
	Data      json.RawMessage `json:"data"`
	Timestamp time.Time       `json:"timestamp"`
}
