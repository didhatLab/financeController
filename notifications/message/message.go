package message

import "time"

type EventMessage struct {
	Type    string
	Time    time.Time
	Message string
	UserIds []int
}
