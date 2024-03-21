package main

import (
	"fmt"

	"github.com/go4itx/go-fiber-rest/response"
)

func main() {
	fmt.Println(response.HandleResult("hello world"))
}
