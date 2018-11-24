package redis

import (
  "github.com/go-redis/redis"
)

var (
  // redisClient is the connection handle for the redis
  redisClient *redis.Client
)

// SetupRedis connection
func SetupRedis(hostname string, password string, db int) error {
  redisClient = redis.NewClient(&redis.Options{
    Addr:     hostname,
    Password: password,
    DB:       db,
  })

  _, err := redisClient.Ping().Result()
  if err != nil {
    return err
  }

  return nil
}

// GetHandle get redis handle
func GetHandle() *redis.Client {
  return redisClient
}
