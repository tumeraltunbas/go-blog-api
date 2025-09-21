package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator"
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

	//TODO: Move them to a centeralízed func.
	var registerReqDto request.RegisterReqDto

	if err := json.NewDecoder(r.Body).Decode(&registerReqDto); err != nil {
		log.Fatalln("Error")
	}

	validate := validator.New()
	if err := validate.Struct(registerReqDto); err != nil {
		errors.NewInternalServerError(w) //TODO: ValidationError
		return
	}

	existingUser, err := h.userService.GetUserByEmail(r.Context(), registerReqDto.Email)

	if err != nil {
		errors.NewInternalServerError(w)
	}

	if existingUser == nil {
		errors.NewBusinessRuleError(errors.UserNotFound, w)
	}

	securityConfig := config.Cfg.Security

	hash, err := utils.HashPassword(registerReqDto.Password, securityConfig.HashCost)

	if err != nil {
		errors.NewInternalServerError(w)
	}

	insertUserErr := h.userService.InsertUser(r.Context(), registerReqDto.Email, string(hash))

	if insertUserErr != nil {
		errors.NewInternalServerError(w)
	}

}
