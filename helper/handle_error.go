package helper

import (
	"edge/data/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpResponse struct {
	Message     string
	Status      int
	Description string
}

func ErrorHandler(c *gin.Context, e any) {
	//goErr := errors.Wrap(err, 2)
	//httpResponse := HttpResponse{Message: "Internal server error", Status: 500, Description: goErr.Error()}
	//c.AbortWithStatusJSON(500, httpResponse)

	switch err := e.(type) {
	case CustomException: // Handle the custom error type
		res := response.Response{
			Code:   err.Code,
			Status: err.Message,
			Data:   gin.H{"error": "Customer error"},
		}
		c.JSON(http.StatusForbidden, res)
	case ValidateException: // Validate the custom error type
		res := response.Response{
			Code:   err.Code,
			Status: err.Message,
			Data:   gin.H{"error": "Validate error"},
		}
		c.JSON(http.StatusBadRequest, res)
	case error:
		// Handle other generic errors
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		res := response.Response{
			Code:   -1,
			Status: err.Error(),
			Data:   gin.H{"error": "Internal server error"},
		}
		c.JSON(http.StatusInternalServerError, res)
	default:
		// Handle unexpected types
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown Error"})
	}

}
