package utils

type RequestError struct {
	Code    int                   `json:"code" xml:"code" example:"422"`
	Message string                `json:"message" xml:"message" example:"Invalid email address"`
	Fields  []DataValidationError `json:"fields" xml:"fields"`
}

func (re RequestError) Error() string {
	return re.Message
}

type DataValidationError struct {
	Field   string `json:"field" xml:"field" example:"email"`
	Message string `json:"message" xml:"message" example:"Invalid email address"`
}

type GlobalError struct {
	Message string `json:"message" xml:"message" example:"invalid name"`
}
