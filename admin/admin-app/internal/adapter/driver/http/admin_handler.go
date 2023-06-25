package http

import (
	"admin/internal/core/dto"
	"admin/pkg/auth/authenticator"

	"github.com/gin-gonic/gin"

)

func (s *HttpHandlers) UpdatePassword(req *gin.Context) {

	request := dto.UpdatePasswordRequest{}
	req.BindJSON(&request)
	request.AccessToken = authenticator.GetTokenString(req.GetHeader("Authorization"))

	err := s.AdminUsecase.UpdatePassword(request)
	if err != nil {
		req.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		req.JSON(200, gin.H{
			"message": "ok",
		})
	}
}

func (s *HttpHandlers) GenerateInvitLink(req *gin.Context) {

	request := dto.InviteLinkRequest{}
	req.BindJSON(&request)
	request.AccessToken = authenticator.GetTokenString(req.GetHeader("Authorization"))

	err := s.AdminUsecase.GenerateInviteLink(request)
	if err != nil {
		req.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		req.JSON(200, gin.H{
			"message": "ok",
		})
	}

}
