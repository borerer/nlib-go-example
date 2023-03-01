package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	nlibgo "github.com/borerer/nlib-go"
)

func ping(req *nlibgo.Request) *nlibgo.Response {
	return nlibgo.Text("pong")
}

func random(req *nlibgo.Request) *nlibgo.Response {
	return nlibgo.Text(fmt.Sprintf("%d", rand.Int()))
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

func add(req *nlibgo.Request) *nlibgo.Response {
	a := toFloat(nlibgo.GetQuery(req, "a"))
	b := toFloat(nlibgo.GetQuery(req, "b"))
	return nlibgo.Text(fmt.Sprintf("%f", a+b))
}

func main() {
	nlibgo.SetEndpoint(os.Getenv("NLIB_SERVER"))
	nlibgo.SetAppID("nlib-go-example")
	nlibgo.Must(nlibgo.Connect())
	nlibgo.RegisterFunction("ping", ping)
	nlibgo.RegisterFunction("random", random)
	nlibgo.RegisterFunction("add", add)
	nlibgo.Wait()
}
