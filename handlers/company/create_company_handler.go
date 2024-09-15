package company

import (
	"net/http"
	"pdv/database"
	"pdv/schemas"

	"github.com/gin-gonic/gin"
)



type CompanyRequest struct {
  Name string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func CreateCompanyHandler(ctx *gin.Context) {
	db := database.GetDB()
	request := CompanyRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
    return
	}

	company := schemas.Company{
		Name:  request.Name,
    Phone: request.Phone,
    Email: request.Email,
    Password: request.Password,
	}

	if err := db.Create(&company).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
    return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"company": company,
	})
}