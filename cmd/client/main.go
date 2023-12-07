package main

import (
	"context"
	"fmt"

	"github.com/narumiruna/requestgen-str-example/pkg/api"
)

func main() {
	client := api.NewClient()
	req, err := client.NewRootRequest().Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(*req)
}
