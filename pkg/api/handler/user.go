package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	domain "github.com/SethukumarJ/go-gin-clean-arch/pkg/domain"
	"github.com/SethukumarJ/go-gin-clean-arch/pkg/response"
	services "github.com/SethukumarJ/go-gin-clean-arch/pkg/usecase/interface"
	utils "github.com/SethukumarJ/go-gin-clean-arch/pkg/utils"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

type Response struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// @title Go + Gin Workey API
// @version 1.0
// @description This is a sample server Job Portal server. You can visit the GitHub repository at https://github.com/fazilnbr/Job_Portal_Project

// @contact.name API Support
// @contact.url https://fazilnbr.github.io/mypeosolal.web.portfolio/
// @contact.email fazilkp2000@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /
// @query.collection.format multi

// FindAll godoc
// @summary Get all users
// @description Get all users
// @tags users
// @id FindAll
// @produce json
// @Router /api/users [get]
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
func (cr *UserHandler) FindAll(c *gin.Context) {
	users, err := cr.userUseCase.FindAll(c.Request.Context())

	if err != nil {
		response := response.ErrorResponse("error while getting users from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Listed All Users", users)
	utils.ResponseJSON(*c, response)
}

// FindOne godoc
// @summary Get one users
// @description Get one users
// @tags users
// @id FindOne
// @produce json
// @Param        userId   query      string  true  "User Id : "
// @Router /users [get]
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
func (cr *UserHandler) FindByID(c *gin.Context) {
	userId := c.Query("userId")
	fmt.Println(userId)

	user, err := cr.userUseCase.FindByID(c.Request.Context(), userId)
	fmt.Printf("\n\nuser  : %v\n\nerr  %v\n\n", user, err)

	if err != nil {
		response := response.ErrorResponse("FAILL", err.Error(), nil)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", user)
	utils.ResponseJSON(*c, response)

}

// FindAll godoc
// @summary Get all users
// @description Save user
// @tags users
// @id Save
// @param RegisterAdmin body domain.Users{} true "admin signup with username, phonenumber email ,password"
// @produce json
// @Router /api/users [Post]
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
func (cr *UserHandler) Save(c *gin.Context) {
	var newUser domain.Users

	if err := c.Bind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := cr.userUseCase.Save(c.Request.Context(), newUser)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	user.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// DeleteOne godoc
// @summary Delete one users
// @description Delete one users
// @tags users
// @id DeleteOne
// @produce json
// @Param        userId   query      string  true  "User Id : "
// @Router /users [delete]
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
func (cr *UserHandler) Delete(c *gin.Context) {
	userId := c.Query("userId")

	ctx := c.Request.Context()
	user, err := cr.userUseCase.FindByID(ctx, userId)

	if err != nil {
		response :=response.ErrorResponse("FAILL",err.Error(),nil)
		utils.ResponseJSON(*c,response)
		return
	}

	if user == (domain.Users{}) {
		response :=response.ErrorResponse("FAILL","There is no users with your id check id",nil)
		utils.ResponseJSON(*c,response)
		return
	}

	err =cr.userUseCase.Delete(ctx, userId)
	if err!=nil {
		response :=response.ErrorResponse("FAILL",err.Error(),nil)
		utils.ResponseJSON(*c,response)
		return
	}

	response :=response.SuccessResponse(true,"SUCCESS",nil)
	utils.ResponseJSON(*c,response)
}
