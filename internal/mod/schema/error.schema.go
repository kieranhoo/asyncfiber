package schema

type Error struct {
	Success bool   `json:"status"`
	Msg     string `json:"msg"`
}

type Response struct {
	Success bool   `json:"success" validate:"required"`
	Msg     string `json:"msg" validate:"required"`
}

type DataResponse struct {
	Success bool        `json:"success" validate:"required"`
	Data    interface{} `json:"data" validate:"required"`
}
