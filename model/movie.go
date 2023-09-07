package model

type Response interface {
	AutoCompleteResponse | SearchResponse
}

type AutoCompleteResponse struct {
	MovieNm string `json:"movieNm"`
}

type SearchResponse struct {
	MovieNm     string      `json:"movieNm"`
	MovieNmEn   string      `json:"movieNmEn"`
	RepNationNm string      `json:"repNationNm"`
	Directors   []directors `json:"directors"`
	NationAlt   string      `json:"nationAlt"`
	GenreAlt    interface{} `json:"genreAlt"`
}

type directors struct {
	PeopleNm string
}

type MovieRequest struct {
	MovieNm string `json:"movieNm"`
}
