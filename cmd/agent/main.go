package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

const pollInterval = 2 * time.Second
const pushInterval = 10 * time.Second

func main() {
	var memory runtime.MemStats
	client := resty.New()
	urls := make([]string, 29)
	start := time.Now()
	i := 1
	for {
		startWork := time.Now()
		u := 0
		urls[u] = fmt.Sprintf("http://127.0.0.1:8080/update/gauge/RandomValue/%d", rand.Intn(10000))
		u++
		urls[u] = fmt.Sprintf("http://127.0.0.1:8080/update/counter/PollCount/%d", i)
		u++
		i++
		runtime.ReadMemStats(&memory)
		v := reflect.ValueOf(memory)
		typeOfV := v.Type()

		for n := 0; n < v.NumField(); n++ {
			val := 0.0
			if !v.Field(n).CanUint() && !v.Field(n).CanFloat() {
				continue
			} else if !v.Field(n).CanUint() {
				val = v.Field(n).Float()
			} else {
				val = float64(v.Field(n).Uint())
			}

			name := typeOfV.Field(n).Name
			urls[u] = fmt.Sprintf("http://127.0.0.1:8080/update/gauge/%s/%f", name, val)
			u++
		}
		time.Sleep(pollInterval - time.Since(startWork))
		if pushInterval-time.Since(start) <= 0 {
			start = time.Now()
			send(urls, client)
		}
	}

}

func send(urls []string, client *resty.Client) {
	fmt.Println("Send")
	for _, url := range urls {
		go client.R().Post(url)
	}
}
