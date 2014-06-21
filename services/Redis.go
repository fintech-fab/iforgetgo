package services

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

type Redis struct {
	Host     string
	Port     string

	connection redis.Conn
}

func (this *Redis) Connect() {
	connection, error := redis.Dial("tcp", this.Host+":"+this.Port)

	if (error != nil) {
		log.Fatal(error)
	}

	this.connection = connection
}

func (this *Redis) Delete(key interface{}) {
	this.connection.Send("DEL", key)
	this.connection.Flush()
}

func (this *Redis) Get(key string) []string {
	result, error := this.connection.Do("HGETALL", key)

	if (error != nil) {
		log.Fatal(error)
	}

	value, error := redis.Strings(result, error)

	if (error != nil) {
		log.Fatal(error)
	}

	return value;
}

func (this *Redis) GetRangeByScore(name string, min int, max int) []string {

	result, error := this.connection.Do("ZRANGEBYSCORE", name, min, max)

	if(error != nil) {
		log.Fatal(error)
	}

	results, error := redis.Strings(result, error)


	return results
}

func (this *Redis) SearchKeys(query string) []string {

	result, error := this.connection.Do("KEYS", query)

	if (error != nil) {
		log.Fatal(error)
	}

	results, error := redis.Strings(result, error)

	if (error != nil) {
		log.Fatal(error)
	}

	return results
}
