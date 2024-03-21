package response

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

type resp struct {
	ctx fiber.Ctx
}

// New
func New(ctx fiber.Ctx) resp {
	return resp{ctx: ctx}
}

func (r resp) JSON(data ...any) error {
	return r.ctx.JSON(HandleResult(data...))
}

func HandleResult(params ...any) Result {
	var (
		data any
		fe   *fiber.Error
	)
	length := len(params)
	if length > 0 {
		lastIndex := length - 1
		lastParam := params[lastIndex]

		if err, ok := lastParam.(*fiber.Error); ok {
			fe = err
		} else { // golang error to fiber  Error
			if err, ok := lastParam.(error); ok && err != nil {
				fe = fiber.NewError(fiber.StatusInternalServerError, err.Error())
			}
		}

		if fe == nil {
			if lastParam == nil {
				length = lastIndex
			}

			switch length {
			case 1:
				data = params[0]
			case 2:
				data = PaginationData{
					Items: params[0],
					Count: params[1],
				}
			case 3:
				data = PaginationData{
					Items: params[0],
					Count: params[1],
					Other: params[2],
				}
			}
		}
	}

	if fe == nil {
		fe = fiber.NewError(fiber.StatusOK)
	}

	if data == nil {
		data = ""
	}

	return Result{
		Code:       fe.Code,
		Msg:        fe.Message,
		ServerTime: time.Now().UnixNano() / 1e6,
		Data:       data,
	}
}
