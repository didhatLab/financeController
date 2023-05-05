package resolver

import (
	"github.com/redis/go-redis/v9"
	"main/notifications/message"
)

type EventsResolver struct {
	redis redis.Client
}

func (er EventsResolver) Resolve(event message.EventMessage) {

}
