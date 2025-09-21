package handlers

import (
	"log"
	"net/http"

	"github.com/tumeraltunbas/go-blog/config"
	"github.com/tumeraltunbas/go-blog/constants/errors"
	"github.com/tumeraltunbas/go-blog/models/dtos/request"
	"github.com/tumeraltunbas/go-blog/services"
	"github.com/tumeraltunbas/go-blog/utils"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value("dto").(request.RegisterReqDto)

	existingUser, err := h.userService.GetUserByEmail(r.Context(), dto.Email)

	if err != nil {
		log.Println("Auth handler - register - GetUserByEmail", err)
		errors.NewInternalServerError(w)
		return
	}

	if existingUser != nil {
		errors.NewBusinessRuleError(errors.UserAlreadyExist, w)
		return
	}

	securityConfig := config.Cfg.Security

	hash, err := utils.HashPassword(dto.Password, securityConfig.HashCost)

	if err != nil {
		log.Println("Auth handler - register - HashPassword", err)
		errors.NewInternalServerError(w)
		return
	}

	insertUserErr := h.userService.InsertUser(r.Context(), dto.Email, string(hash))

	if insertUserErr != nil {
		log.Println("Auth handler - register - InsertUser", err)
		errors.NewInternalServerError(w)
		return
	}
}
