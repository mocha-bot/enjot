package dto

type defaultResponse struct {
	Message string `json:"message"`
}

func ParseToDefaultReponse() defaultResponse {
	return defaultResponse{
		Message: "OK",
	}
}
