package http

import (
	"admin/internal/core/dto"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *HttpHandlers) SignIn(req *gin.Context) {

	auth := dto.SignInRequest{}
	req.BindJSON(&auth)

	fmt.Println(auth)

	token, err := s.AuthUsecase.SignIn(auth)
	if err != nil {
		req.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		req.JSON(200, token)
	}
}

func (s *HttpHandlers) SignUp(req *gin.Context) {

	data := dto.SignUpRequest{}
	req.ShouldBindUri(&data)
	req.BindJSON(&data)
	fmt.Println(data)

	token, err := s.AuthUsecase.SignUp(data)

	fmt.Println(err)

	if err != nil {
		req.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		req.JSON(200, token)
	}
}

// func ForgetPassword(w http.ResponseWriter, r *http.Request) {}
