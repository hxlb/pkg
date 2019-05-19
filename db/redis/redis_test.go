package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
)

func TestRedis(t *testing.T) {
	example1()
}

func example1() {
	client := DefaultClient()
	defer client.Close()

	//fmt.Println(client.Set("name", "Raed Shomali")) // true <nil>
	//fmt.Println(client.SetNx("name", "Hello"))      // false <nil>
	fmt.Println(client.SetEx("name", "Raed Shomali", 100)) // true <nil>
	//fmt.Println(client.Expire("name", 1))           // true <nil>
	//fmt.Println(client.Expire("unknown", 1))        // false <nil>
	//fmt.Println(client.Keys("*"))                   // [id name] <nil>
	//fmt.Println(client.Get("name"))                 // "Raed Shomali" true <nil>
	//fmt.Println(client.Exists("name"))              // true <nil>
	//fmt.Println(client.Del("name"))                 // 1 <nil>
	//fmt.Println(client.Exists("name"))              // false <nil>
	//fmt.Println(client.Get("name"))                 // "" false <nil>
	//fmt.Println(client.Del("name"))                 // 0 <nil>
	//fmt.Println(client.Append("name", "a"))         // 1 <nil>
	//fmt.Println(client.Append("name", "b"))         // 2 <nil>
	//fmt.Println(client.Append("name", "c"))         // 3 <nil>
	//fmt.Println(client.Get("name"))                 // "abc" true <nil>
	//fmt.Println(client.GetRange("name", 0 , 1))     // "ab" <nil>
	//fmt.Println(client.SetRange("name", 2, "xyz"))  // 5 <nil>
	//fmt.Println(client.Get("name"))                 // "abxyz" <nil>
	//fmt.Println(client.Scan(0, "*"))                // 0 [name id] <nil>
	//fmt.Println(client.Del("id", "name"))           // 2 <nil>
}

func example2() {
	options := &Options{
		Host: "localhost",
		Port: 6379,
	}

	client := SetupClient(options)
	defer client.Close()

	fmt.Println(client.Ping()) // PONG <nil>
}

func example3() {
	options := &SentinelOptions{
		Addresses:  []string{"localhost:26379"},
		MasterName: "master",
	}

	client := SetupSentinelClient(options)
	defer client.Close()

	fmt.Println(client.Ping()) // PONG <nil>
}

func example4() {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	client := NewClient(pool)
	defer client.Close()

	fmt.Println(client.Ping()) // PONG <nil>
}
