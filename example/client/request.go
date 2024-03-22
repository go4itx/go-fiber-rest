package main

import (
	"fmt"

	"github.com/go4itx/go-fiber-rest/client"
	"github.com/go4itx/go-fiber-rest/response"
	"github.com/gofiber/fiber/v3"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	// 例子1：
	bytes, err := client.Request("http://127.0.0.1:8080").Result()
	if err != nil {
		fmt.Println(bytes)
		return
	}

	fmt.Println(string(bytes))

	// 例子2：
	var res response.Result
	request := client.Request("http://127.0.0.1:8080/login", fiber.MethodPost)
	request.SetJSON(User{Name: "test0322", Password: "123456"})
	if _, err := request.Result(&res); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("xxxxxxxxxxxxxxxx")
	fmt.Println(res)
}
