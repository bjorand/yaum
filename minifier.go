package main

import (
	"fmt"
	"math"

	"github.com/go-redis/redis"
)

type urlEntry struct {
	URL      string
	Minified string
}

func generateID(id int) string {
	var code []byte
	chars := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	for id > len(chars)-1 {
		key := id % (len(chars))
		id = int(math.Floor(float64(id)/(float64(len(chars)))) - 1)
		code = append(code, []byte(chars[key : key+1])[0])
	}
	return string(append(code, chars[id]))
}

func (u *urlEntry) redisKey() string {
	return fmt.Sprintf("u-%s", u.Minified)
}

func (u *urlEntry) save() error {
	err := redisClient.Set(u.redisKey(), u.URL, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func createMinifiedURL(url string) (*urlEntry, error) {
	lastInsertID, err := redisClient.Incr("lastInsertID").Result()
	if err != nil {
		return nil, err
	}
	id := generateID(int(lastInsertID - 1))
	entry := &urlEntry{
		URL:      url,
		Minified: id,
	}
	if err := entry.save(); err != nil {
		return nil, err
	}
	return entry, nil
}

func getMinifiedURL(value string) (*urlEntry, error) {
	entry := &urlEntry{
		Minified: value,
	}
	url, err := redisClient.Get(entry.redisKey()).Result()
	switch {
	case err == redis.Nil:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		entry.URL = url
		return entry, nil
	}
}
