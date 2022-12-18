package main

import (
	"math/rand"
	"os"
	"strconv"

	nlibgo "github.com/borerer/nlib-go"
)

func ping(in map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{"message": "pong"}
}

func random(in map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"int": rand.Int(),
	}
}

func toFloat(in interface{}) float64 {
	switch v := in.(type) {
	case string:
		res, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0
		}
		return res
	case float64:
		return v
	}
	return 0
}

func add(in map[string]interface{}) map[string]interface{} {
	a := toFloat(in["a"])
	b := toFloat(in["b"])
	return map[string]interface{}{
		"a":   a,
		"b":   b,
		"a+b": a + b,
	}
}

func wait() {
	ch := make(chan bool)
	<-ch
}

func main() {
	nlib := nlibgo.NewClient(os.Getenv("NLIB_SERVER"), "nlib-go-example")
	nlib.RegisterFunction("ping", ping)
	nlib.RegisterFunction("random", random)
	nlib.RegisterFunction("add", add)
	wait()
}
