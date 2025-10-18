package response

type Result struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"ts"`
	Data      any    `json:"data"`
}

type PaginationData struct {
	Items any `json:"items"`
	Count any `json:"count"`
	Other any `json:"other,omitempty"`
}
