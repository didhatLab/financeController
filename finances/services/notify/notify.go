package notify

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"main/notifications/message"
	"time"
)

type Notifier interface {
	Notify(ctx context.Context, msg string, userIds []int) error
}

type NotifierImp struct {
	rdb *redis.Client
}

func NewNotifier(rdb *redis.Client) NotifierImp {
	return NotifierImp{rdb: rdb}
}

func (ni NotifierImp) Notify(ctx context.Context, msg string, userIds []int) error {
	msgForNotify := message.EventMessage{
		Type:    "common",
		Time:    time.Now(),
		Message: msg,
		UserIds: userIds,
	}
	jsonMsg, err := json.Marshal(&msgForNotify)
	if err != nil {
		log.Print(err)
		return err
	}

	ni.rdb.Publish(ctx, "events", string(jsonMsg))

	return nil
}
