package response

import (
	"net/http"
)

type response struct {
	Data   interface{} `json:"data"`
	Status int32       `json:"status"`
}

func NewResponse(obj interface{}) response {
	resp := response{}

	resp.Data = obj
	resp.Status = http.StatusOK

	return resp
}
