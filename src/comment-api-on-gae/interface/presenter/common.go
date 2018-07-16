package presenter

import "comment-api-on-gae/usecase"

type responseJson struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func RenderJson(json interface{}, result *usecase.Result) interface{} {
	return &responseJson{
		Message: result.Message(),
		Data:    json,
	}
}
