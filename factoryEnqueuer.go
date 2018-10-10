package queuepublisher

import "github.com/fiorix/go-redis/redis"

func FactoryEnqueuer (enqueuer string, dbhost string, dbport string, dbname string, dbpass string, secureConn bool) IQueuePublisher{
    switch enqueuer {
    default: // case "resque"
        client := redis.New(dbhost+":"+dbport+" "+dbname) // Create new Redis client to use for enqueuing
        clientName := "redis-go"
        return &ResquePublisher{client, clientName}
        break
    }
    return nil
}