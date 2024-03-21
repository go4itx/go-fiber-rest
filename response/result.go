package response

type Result struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	ServerTime int64       `json:"serverTime"`
	Data       interface{} `json:"data"`
}

type PaginationData struct {
	Items interface{} `json:"items"`
	Count interface{} `json:"count"`
	Other interface{} `json:"other,omitempty"`
}
