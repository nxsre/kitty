package common

type Result struct {
	Message   string `json:"message"`
	Success   bool `json:"success"`
	Content   interface{} `json:"content"`
	ErrorType int `json:"errorType"`

}
