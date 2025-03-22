package controller

func response(message string, result any, status int) map[string]any {
	response := make(map[string]any)
	response["message"] = message
	response["result"] = result
	response["status"] = status

	return response
}
