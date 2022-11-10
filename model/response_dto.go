package model

// Country struct for db table - country
type ResponseDto struct {
	Message    string      `json:"message"`
	IsSuccess  bool        `json:"isSuccess"`
	Payload    interface{} `json:"payload"`
	StatusCode int         `json:"statusCode"`
}
