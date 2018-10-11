package queuepublisher

type IQueuePublisher interface {
    PublishToQueue(queue string,class string,jobdata map[string]interface{}) error
}
