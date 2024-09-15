package user

import (
	"net/http"
	"pdv/database"
	"pdv/schemas"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CompanyId int    `json:"companyId"`
}

func CreateUserHandler(ctx *gin.Context) {
	db := database.GetDB()
	request := UserRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := schemas.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		CompanyId: request.CompanyId,
	}

	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
