package client

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/client"
)

var (
	timeout   = 1 * time.Minute
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36"
)

type req struct {
	*client.Request
}

// Request ...
func Request(url string, method ...string) req {
	m := fiber.MethodGet
	if len(method) > 0 {
		m = method[0]
	}

	request := client.AcquireRequest()
	request.SetTimeout(timeout).SetUserAgent(userAgent)
	request.SetURL(url).SetMethod(m)
	return req{request}
}

// Result ...
func (r req) Result(v ...any) (res []byte, err error) {
	response, err := r.Send()
	if err != nil {
		return
	}

	defer response.Close()
	if response.StatusCode() != fiber.StatusOK {
		err = fiber.NewError(response.StatusCode(), response.Status())
		return
	}

	if len(v) > 0 {
		switch string(response.RawResponse.Header.ContentType()) {
		case fiber.MIMEApplicationJSON, fiber.MIMEApplicationJSONCharsetUTF8:
			err = response.JSON(v[0])
		case fiber.MIMETextXML, fiber.MIMEApplicationXML, fiber.MIMEApplicationXMLCharsetUTF8:
			err = response.XML(v[0])
		default:
			res = response.Body()
		}
	} else {
		res = response.Body()
	}

	return
}
