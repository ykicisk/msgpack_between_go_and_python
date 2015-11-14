package main

import (
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
	dic := map[string]string{"gokey1": "goval1", "gokey2": "goval2"}
	ls_int := []int{1, 3, 5, 8}
	ls_float := []float64{2.3, 6.4, 5.5}
	{
		msg, msgpack_err := msgpack.Marshal(dic)
		check(msgpack_err)
		check(redis_client.Set("go_map", msg, 0).Err())
	}
	{
		msg, msgpack_err := msgpack.Marshal(ls_int)
		check(msgpack_err)
		check(redis_client.Set("go_ls_int", msg, 0).Err())
	}
	{
		msg, msgpack_err := msgpack.Marshal(ls_float)
		check(msgpack_err)
		check(redis_client.Set("go_ls_float", msg, 0).Err())
	}
}
