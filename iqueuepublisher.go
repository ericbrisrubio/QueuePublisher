package queuepublisher

type IQueuePublisher interface {
    publishToQueue(string, string, map[string]interface{}) error
}
