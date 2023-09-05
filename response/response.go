package response

import "net/http"

type response struct {
	Data   interface{} `json:"data"`
	Status int32       `json:"status"`
}

func NewResponse() response {
	return response{}
}

func (r response) GetResponse(obj interface{}) response {
	resp := NewResponse()

	resp.Data = obj
	resp.Status = http.StatusOK

	return resp
}
