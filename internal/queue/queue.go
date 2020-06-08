package queue

// Queue queue
type Queue interface {
	// Push message to back
	Push(message.Message) error
	// Pop message
	Pop() (message.Message, error)
}

type QueueTopic interface {
	GetQueue(topic string) Queue
}

var (
	queueManager map[string]QueueTopic
)

func Registry(typ string, queueTopic QueueTopic) {
	queueManager[typ] = queueTopic
}

// GetQueue
func GetQueueTopic(typ string) QueueTopic {
	return queueManager[typ]
}
