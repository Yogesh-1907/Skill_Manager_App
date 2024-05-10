package middleware

import (
	"SkillManager_API/utility"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var logger = utility.Log

// ErrorHandler is a middleware that handles errors
func ErrorHandler(ctx *gin.Context) {

	//handle panic
	defer func() {
		if err := recover(); err != nil {
			msg, _ := utility.ResourceManager{}.GetProperty("General.ERROR_MESSAGE")
			errorInfo := utility.NewErrorInfo(http.StatusInternalServerError, msg, time.Now())
			ctx.JSON(http.StatusInternalServerError, errorInfo)
			logger.Fatal(err)
		}
	}()

	ctx.Next()

	// Check if an error occurred during the request
	if len(ctx.Errors) > 0 {
		err := ctx.Errors.Last().Err
		if err != nil {
			// Handle specific errors or provide a generic error message
			switch err := err.(type) {
			case utility.SkillzManagerError:
				msg, _ := utility.ResourceManager{}.GetProperty(err.Error())
				errorInfo := utility.NewErrorInfo(statusCode(err.Error()), msg, time.Now())
				ctx.JSON(statusCode(err.Error()), errorInfo)
			case error:
				msg, _ := utility.ResourceManager{}.GetProperty("General.ERROR_MESSAGE")
				errorInfo := utility.NewErrorInfo(http.StatusInternalServerError, msg, time.Now())
				ctx.JSON(http.StatusInternalServerError, errorInfo)
			}
			// Log the error for debugging purposes
			logger.Error(err)
		}
	}
}

func statusCode(msg string) int {
	if strings.Contains(msg, "Service") {
		return http.StatusNotFound
	}
	if strings.Contains(msg, "API") || strings.Contains(msg, "Validator") {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
