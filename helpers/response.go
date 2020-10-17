package helpers

import (
	"github.com/gin-gonic/gin"
	viper "github.com/spf13/viper"
)

var respData map[string]interface{}

func initialResponse() {
	respData = map[string]interface{}{
		"version": viper.GetString("app.version"),
	}
}

// ResponseErrValidation ...
func ResponseErrValidation(errors []map[string]interface{}) gin.H {
	initialResponse()
	resp := gin.H(
		MergeMap(respData, map[string]interface{}{
			"message": nil,
			"errors":  errors,
		}),
	)

	return resp
}

// ResponseErr ...
func ResponseErr(message string) gin.H {
	initialResponse()
	resp := gin.H(
		MergeMap(respData, map[string]interface{}{
			"message": message,
			"errors":  nil,
		}),
	)

	return resp
}

// ResponseMany ...
func ResponseMany(data []map[string]interface{}) gin.H {
	initialResponse()
	resp := gin.H(
		MergeMap(respData, map[string]interface{}{
			"items": data,
		}),
	)
	return resp
}

// ResponseManyWithPagination ...
func ResponseManyWithPagination(data []map[string]interface{}, pagination []map[string]interface{}) gin.H {
	initialResponse()
	resp := gin.H(
		MergeMap(respData, map[string]interface{}{
			"data":       data,
			"pagination": pagination,
		}),
	)

	return resp
}

// ResponseOne ...
func ResponseOne(data map[string]interface{}) gin.H {
	initialResponse()
	resp := gin.H(
		MergeMap(respData, map[string]interface{}{
			"data": data,
		}),
	)

	return resp
}

// Response ...
func Response(data map[string]interface{}) gin.H {
	initialResponse()
	resp := gin.H(MergeMap(respData, data))
	return resp
}
