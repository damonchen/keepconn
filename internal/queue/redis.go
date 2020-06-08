package queue

import (
	"github.com/go-redis/redis/v8"
)

type RedisQueueTopic struct {
}

func (t *RedisQueueTopic) GetQueue(topic string) Queue {
	return
}

type RedisQueue struct {
	redisClient *redis.Client
}

// Push message to back
func (q *RedisQueue) Push(message.Message) error {

}

// Pop message
func (q *RedisQueue) Pop() (message.Message, error) {

}
