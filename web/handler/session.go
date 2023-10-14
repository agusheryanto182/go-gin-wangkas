package handler

import (
	"net/http"

	"github.com/agusheryanto182/go-wangkas/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	New(c *gin.Context)
	Create(c *gin.Context)
	Destroy(c *gin.Context)
}

type SessionHandler struct {
	userService user.Service
}

func NewSessionHandler(userService user.Service) *SessionHandler {
	return &SessionHandler{userService}
}

func (h *SessionHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "session_new.html", nil)
}

func (h *SessionHandler) Create(c *gin.Context) {
	var input user.User

	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	user, err := h.userService.LoginUser(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Set("userName", user.Username)
	session.Save()

	c.Redirect(http.StatusFound, "/transactions")
}

func (h *SessionHandler) Destroy(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}
