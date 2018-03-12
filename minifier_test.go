package main

import (
	"reflect"
	"testing"

	"github.com/go-redis/redis"
)

func setup() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})
	if _, err := redisClient.FlushDB().Result(); err != nil {
		panic(err)
	}
}

func TestRedisKey(t *testing.T) {
	input := &urlEntry{
		Minified: "foobar",
	}
	actual := input.redisKey()
	expected := "u-foobar"
	if actual != expected {
		t.Fatalf("%+v redisKey(): expected %s, actual %s", input, expected, actual)
	}
}

func TestGenerateID(t *testing.T) {
	var idTests = []struct {
		input    int
		expected string
	}{
		{0, "a"},
		{62, "aa"},
		{100, "Ma"},
		{1000, "ip"},
		{10000, "sKb"},
	}

	for _, pt := range idTests {
		actual := generateID(pt.input)
		if actual != pt.expected {
			t.Fatalf("generateID(%+v): expected %s, actual %s", pt.input, pt.expected, actual)
		}
	}
}

func TestCreateGetMinifiedURL(t *testing.T) {
	setup()
	input := "http://example.com"
	actual, err := createMinifiedURL(input)
	if err != nil {
		t.Fatal(err)
	}
	expected := &urlEntry{
		URL:      "http://example.com",
		Minified: "a",
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("createMinifiedURL(%+v): expected %s, actual %s", input, expected, actual)
	}
	entry, err := getMinifiedURL("a")
	if err != nil {
		t.Fatal(err)
	}
	if entry.URL != expected.URL {
		t.Fatalf("getMinifiedURL(%+v): expected %s, actual %s", "a", expected, actual)
	}
}

func TestGetMinifiedURLNotFound(t *testing.T) {
	setup()
	input := "z"
	actual, err := getMinifiedURL(input)
	if err != nil {
		t.Fatal(err)
	}
	if actual != nil {
		t.Fatalf("getMinifiedURL(%+v): expected nil, actual %s", input, actual)
	}
}
