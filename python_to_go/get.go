package main

import (
	"fmt"
	"gopkg.in/redis.v3"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	redis_client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var unpacked_map map[string]string
	var unpacked_ls_int []int
	var unpacked_ls_float []float64
	{
		msg, msgpack_err := redis_client.Get("py_map").Bytes()
		check(msgpack_err)
		msgpack.Unmarshal(msg, &unpacked_map)
	}
	{
		msg, msgpack_err := redis_client.Get("py_ls_int").Bytes()
		check(msgpack_err)
		msgpack.Unmarshal(msg, &unpacked_ls_int)
	}
	{
		msg, msgpack_err := redis_client.Get("py_ls_float").Bytes()
		check(msgpack_err)
		msgpack.Unmarshal(msg, &unpacked_ls_float)
	}
	fmt.Println(unpacked_map)
	fmt.Println(unpacked_ls_int)
	fmt.Println(unpacked_ls_float)
}
