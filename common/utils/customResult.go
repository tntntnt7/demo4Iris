package utils

type result struct {
	Success bool			`json:"success"`
	Data 	interface{}		`json:"data"`
	Error	string			`json:"error"`
}

func HandleSuccessfulResult(data interface{}) result {
	return result{true, data, "no error"}
}

func HandleErrorResult(err error) result {
	return result{false, nil, err.Error()}
}