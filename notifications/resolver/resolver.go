package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"main/notifications/message"
)

type EventsResolver struct {
	Redis *redis.Client
}

func (er EventsResolver) StartResolve() {
	ctx := context.Background()

	pubsub := er.Redis.Subscribe(ctx, "events")

	ch := pubsub.Channel()

	for msg := range ch {
		var event message.EventMessage
		err := json.Unmarshal([]byte(msg.Payload), &event)

		if err != nil {
			log.Print("error while unmarshaling", err)
			continue
		}

		er.send2telegram(ctx, event)

	}
}

func (er EventsResolver) send2telegram(ctx context.Context, msg message.EventMessage) {
	userIdsForTg := make([]int, 0, 10)

	for _, userId := range msg.UserIds {
		status, err := er.Redis.Get(ctx, fmt.Sprintf("%d:userNotificationStatus", userId)).Result()
		if err != nil {
			log.Print(err)
			continue
		}
		if status == "active" {
			userIdsForTg = append(userIdsForTg, userId)
		}
	}

	msg.UserIds = userIdsForTg
	jsonMsg, err := json.Marshal(&msg)
	if err != nil {
		log.Print(err)
	}

	er.Redis.Publish(ctx, "telegram", string(jsonMsg))
}
