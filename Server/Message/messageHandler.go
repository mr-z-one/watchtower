package Message

import "encoding/json"

func CreateErrorMessage(message string, statusCode int) map[string]interface{} {

	return map[string]interface{}{
		"statusCode": statusCode,
		"data": map[string]string{
			"message": message,
		},
	}
}

func CreateErrorMessageJson(message string, statusCode int) []byte {
	b, _ := json.Marshal(CreateErrorMessage(message, statusCode))
	return b

}

func CreateMessageJson(message string, statusCode int) []byte {
	b, _ := json.Marshal(CreateErrorMessage(message, statusCode))
	return b

}
