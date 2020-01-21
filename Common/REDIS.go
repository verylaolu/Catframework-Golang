package Common

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)
var Redis *redis.Client
func SetRedis(client *redis.Client) {
	Redis = client
}
func ConnectToRedis() *redis.Client{
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")


	client := redis.NewClient(&redis.Options{
		Addr:     host+":"+port,
		Password: password, // no password set
		DB:       0,  // use default DB
	})
	return client
}