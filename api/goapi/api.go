package goapi

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func Client() {
	client := resty.New()
	_ = client
	resp, err := client.R().
		EnableTrace().
		Get("https://swapi.dev/api/people/1/")

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
}
