## go-fiber-rest（升级旧项目go-fiber-api，搭建项目更快捷）
### 1、使用的是golang 1.24.0，fiber v3（注：目前fiber v3官方未正式发布，估计4月底发布）

### 2、可快速搭建一个基于fiber的restful api的骨架，轻松插入中间件，已具备jwt验证, 集成github.com/go-playground/validator参数校验，返回结果封装，统一返回结果格式

#### 快速上手：
```
server.New(func(app *fiber.App) {
		app.All("/", func(ctx fiber.Ctx) error {
			return response.New(ctx).JSON("Hello, World!")
		})
	})
```
#### 测试：curl http://127.0.0.1:8080/
```
{
  "code": 200,
  "msg": "OK",
  "Timestamp": 1715302453664,
  "data": "Hello, World!"
}
```
#### 返回值统一结构：
```
type Result struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Timestamp  int64       `json:"Timestamp"`
	Data       interface{} `json:"data"`
}
```

```
{
    "code": 400,
    "msg": "Name为必填字段",
    "Timestamp": 1711016877188,
    "data": ""
}
```

### 3、完整使用例子，参考example/server/main.go
```
package main

import (
	"github.com/go4itx/go-fiber-rest/jwt"
	"github.com/go4itx/go-fiber-rest/response"
	"github.com/go4itx/go-fiber-rest/server"
	"github.com/gofiber/fiber/v3"
)

func main() {
	server.New(router)
}

// router ...
func router(app *fiber.App) {
	app.All("/", func(ctx fiber.Ctx) error {
		return response.New(ctx).JSON("Hello, World!")
	})

	app.Post("/login", func(ctx fiber.Ctx) error {
		var user User
		if err := ctx.Bind().Body(&user); err != nil {
			return err
		}

		exp := jwt.GetExpTime(7) // 设置token过期时间为7天
		return response.New(ctx).JSON(jwt.CreateToken(user.Name, exp, secret))
	})

	// 中间件，开启验证token，下面的请求需要Authorization
	app.Use(jwt.New(secret))
	// 如：curl --header "Authorization: Bearer token"  http://localhost:8080/user
	app.Get("/user", func(ctx fiber.Ctx) error {
		// jwt验证通过后，会把jwt.MapClaims放入Locals
		userInfo := ctx.Locals("user")
		return response.New(ctx).JSON(userInfo)
	})
}

// jwt密钥
const secret = "my-jwt-secret"

type User struct {
	Name     string `validate:"required,min=5,max=20"`
	Password string `validate:"required,min=5,max=20"`
}
```
启动程序
```
go run example/server/main.go
```
```
 / ____(_) /_  ___  _____
  / /_  / / __ \/ _ \/ ___/
 / __/ / / /_/ /  __/ /    
/_/   /_/_.___/\___/_/          v3.0.0-beta.1
--------------------------------------------------
INFO Server started on:         http://127.0.0.1:8080 (bound on host 0.0.0.0 and port 8080)
INFO Total handlers count:      14
INFO Prefork:                   Disabled
INFO PID:                       40450
INFO Total process count:       1
```
#### 测试登录-接口
```
curl --location --request POST 'http://127.0.0.1:8080/login' \
--form 'name="test0321"' \
--form 'password="123456"'
```
结果
```
{
    "code": 200,
    "msg": "OK",
    "Timestamp": 1711015369236,
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ0ZXN0MDMyMSIsImV4cCI6MTcxMTYyMDE2OSwiaWF0IjoxNzExMDE1MzY5fQ.daG6qmHn_ZWwe0HmY49PSr0yJsQrYzq2884PvTG_ze4"
}
```

#### 测试获取当前用户信息-接口
```
curl --location --request GET 'http://127.0.0.1:8080/user' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ0ZXN0MDMyMSIsImV4cCI6MTcxMTYyMDE2OSwiaWF0IjoxNzExMDE1MzY5fQ.daG6qmHn_ZWwe0HmY49PSr0yJsQrYzq2884PvTG_ze4'
```

结果
```
{
    "code": 200,
    "msg": "OK",
    "Timestamp": 1711015572915,
    "data": {
        "aud": "test0321",
        "exp": 1711620169,
        "iat": 1711015369
    }
}
```