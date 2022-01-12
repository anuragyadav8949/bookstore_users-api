package users

import (
	"net/http"

	"github.com/anuragyadav8949/bookstore_users-api/domain/users"
	"github.com/anuragyadav8949/bookstore_users-api/services"
	"github.com/anuragyadav8949/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetWordCount(rq *gin.Context) {
	var stringData users.BookText

	//Std way to do above error handling
	if err := rq.ShouldBindJSON(&stringData); err != nil {

		//Handle JSON Error, issrecpective like above
		restErr := errors.RestError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
			Error:   "Bad Request",
		}
		rq.JSON(restErr.Status, restErr)
		// fmt.Println(err)
	}
	result, saveErr := services.CheckForWord(stringData)
	if saveErr != nil {
		//Handle User Creation Error
		return
	}
	rq.JSON(http.StatusCreated, result) //Returning JSON response
}
