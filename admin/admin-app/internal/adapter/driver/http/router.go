package http

import "github.com/gin-gonic/gin"

func (s *Server) registerRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/signin", s.HttpHandlers.SignIn)
	r.POST("/signup/:link", s.HttpHandlers.SignUp)
	r.POST("/admin/update-password", s.HttpHandlers.UpdatePassword)
	r.POST("/admin/invite-link", s.HttpHandlers.GenerateInvitLink)
	return r
}
