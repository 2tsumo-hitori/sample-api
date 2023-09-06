package model

type AutoCompleteResponse struct {
	MovieNm string `json:"movieNm"`
}

type MovieRequest struct {
	MovieNm string `json:"movieNm"`
}
