package queuepublisher

type IQueuePublisher interface {
    PublishToQueue(string, string, map[string]interface{}) error
}
