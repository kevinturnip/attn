package util

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func GenerateResp(data interface{}, message string) (resp Response) {
	resp.Data = data
	resp.Message = message
	return
}
