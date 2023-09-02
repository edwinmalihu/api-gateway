package controller

import (
	"auth-services/middleware"
	"auth-services/model"
	"auth-services/repository"
	"auth-services/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserSkkController interface {
	Adduserskk(*gin.Context)
	Listuserskk(*gin.Context)
	Detailuserskk(*gin.Context)
	DetailuserskkProfil(*gin.Context)
	Updateuserskk(*gin.Context)
	Loginuserskk(*gin.Context)
	TypeUserRole(*gin.Context)
	ChangePasswordSKk(*gin.Context)
	UpdatePasswordSKk(*gin.Context)
	SendOtp(*gin.Context)
	UpdateProfil(*gin.Context)

	// New Login
	NewLoginuserskk(*gin.Context)
	NewSendOtp(*gin.Context)

	// New Forgot Pass
	SendOtpHard(*gin.Context)
	SendOtpSoa(*gin.Context)
	ValidasiOtpHard(*gin.Context)
	ValidasiOtpSoa(*gin.Context)
}

type userSkkController struct {
	userRepo   repository.UserSkkRepository
	policyRepo repository.PolicySkkRepository
}

// SendOtpHard implements UserSkkController
func (h userSkkController) SendOtpHard(ctx *gin.Context) {
	var req model.CheckUser
	response := &model.Response{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, resp, err := h.userRepo.SendOtpForgotHard(req)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// SendOtpSoa implements UserSkkController
func (h userSkkController) SendOtpSoa(ctx *gin.Context) {
	var req model.CheckUser
	response := &model.Response{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	user, resp, err := h.userRepo.SendOtpForgotSoa(req)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// ValidasiOtpHard implements UserSkkController
func (h userSkkController) ValidasiOtpHard(ctx *gin.Context) {
	var req model.PostUsername
	response := &model.Response{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, resp, err := h.userRepo.ValidasiOtpForgotHard(req)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// ValidasiOtpSoa implements UserSkkController
func (h userSkkController) ValidasiOtpSoa(ctx *gin.Context) {
	var req model.PostUsername
	response := &model.Response{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	user, resp, err := h.userRepo.ValidasiOtpForgotSoa(req)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// NewLoginuserskk implements UserSkkController
func (h userSkkController) NewLoginuserskk(ctx *gin.Context) {
	var req model.LoginSkk
	response := &model.Response{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, resp, err := h.userRepo.NewLoginUserSkk(req)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// NewSendOtp implements UserSkkController
func (h userSkkController) NewSendOtp(ctx *gin.Context) {
	var login model.PostUsername
	response := &model.Response{}
	if err := ctx.ShouldBindJSON(&login); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, resp, err := h.userRepo.NewSendOtp(login)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if resp.StatusCode() == 401 {
		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(resp.String()), &jsonMap)
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": jsonMap["msg"]}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	roles, resp, err := h.policyRepo.GetRoleSkk(model.RequestGetRoleSkk{Username: user.Username})
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	token := utils.GenerateToken(user.Username)
	user.Role = roles.Data
	user.Token = token

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// UpdateProfil implements UserSkkController
func (h userSkkController) UpdateProfil(ctx *gin.Context) {
	var req model.UpdateProfil
	response := &model.Response{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	CekUser := middleware.GetUserToken(ctx)

	fmt.Println("ini user get:", CekUser)

	if CekUser != req.Username {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not Valid Token"})
		return
	}

	user, resp, err := h.userRepo.UpdateProfil(req)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// SendOtp implements UserSkkController
func (h userSkkController) SendOtp(ctx *gin.Context) {
	var req model.PostUsername
	idle_timeout, _ := strconv.Atoi(os.Getenv("IDLE_TIMEOUT"))
	response := &model.Response{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	user, resp, err := h.userRepo.SendOtp(req)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() == 401 {
		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(resp.String()), &jsonMap)
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": jsonMap["msg"]}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	roles, resp, err := h.policyRepo.GetRoleSkk(model.RequestGetRoleSkk{Username: user.Username})
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	token := utils.GenerateToken(user.Username)
	user.Role = roles.Data
	user.Token = token
	user.Idle = idle_timeout

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// UpdatePasswordSKk implements UserSkkController
func (h userSkkController) UpdatePasswordSKk(ctx *gin.Context) {
	var req model.UpdatePassword
	response := &model.Response{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	user, resp, err := h.userRepo.UpdatePasswordSkk(req)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// ChangePasswordSKk implements UserSkkController
func (h userSkkController) ChangePasswordSKk(ctx *gin.Context) {
	var req model.ChangePassword
	response := &model.Response{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	user, resp, err := h.userRepo.ChangePasswordUserSkk(req)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// Loginuserskk implements UserSkkController
func (h userSkkController) Loginuserskk(ctx *gin.Context) {
	var login model.LoginSkk
	response := &model.Response{}
	if err := ctx.ShouldBindJSON(&login); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, resp, err := h.userRepo.LoginUserSkk(login)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if resp.StatusCode() == 401 {
		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(resp.String()), &jsonMap)
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": jsonMap["msg"]}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	roles, resp, err := h.policyRepo.GetRoleSkk(model.RequestGetRoleSkk{Username: user.Username})
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	token := utils.GenerateToken(user.Username)
	user.Role = roles.Data
	user.Token = token

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// Updateuserskk implements UserSkkController
func (h userSkkController) Updateuserskk(ctx *gin.Context) {
	var req model.RequestUpdateUserSkk
	response := &model.Response{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	CekUser := middleware.GetUserToken(ctx)
	if CekUser != req.Userlogin {
		ctx.JSON(http.StatusBadRequest, "Not Valid Token")
		return
	}

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	user, resp, err := h.userRepo.UpdateUserSkk(req)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// Detailuserskk implements UserSkkController
func (h userSkkController) Detailuserskk(ctx *gin.Context) {

	var req model.RequestDetailUserSkkConten

	response := &model.Response{}
	if err := ctx.ShouldBind(&req); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	CekUser := middleware.GetUserToken(ctx)

	if CekUser != req.Userlogin {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not Valid Token"})
		return
	}

	user, resp, err := h.userRepo.DetailUserSkk(req)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// DetailuserskkProfil implements UserSkkController
func (h userSkkController) DetailuserskkProfil(ctx *gin.Context) {

	var request model.RequestDetailUserSkk

	response := &model.Response{}
	if err := ctx.ShouldBind(&request); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println("ini user req:", request.Username)

	usercek := middleware.GetUserToken(ctx)
	fmt.Println("ini token name:", usercek)

	if usercek != request.Username {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not Valid Token"})
		return
	}

	user, resp, err := h.userRepo.DetailuserskkProfil(request)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// Listuserskk implements UserSkkController
func (h userSkkController) Listuserskk(ctx *gin.Context) {
	var req model.RequestParamUser
	ctx.ShouldBind(&req)

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	CekUser := middleware.GetUserToken(ctx)

	fmt.Println("ini user get:", CekUser)

	if CekUser != req.Username {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not Valid Token"})
		return
	}
	req.Username = CekUser
	response := &model.Response{}
	user, resp, err := h.userRepo.ListUserSkk(req)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// Adduserskk implements UserSkkController
func (h userSkkController) Adduserskk(ctx *gin.Context) {
	var req model.RequestAddUserSkk
	response := &model.Response{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Success = false
		response.Status = http.StatusBadRequest
		response.ErrorCode = "9902"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct != true {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	CekUser := middleware.GetUserToken(ctx)

	if CekUser != req.Userlogin {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not Valid Token"})
		return
	}

	user, resp, err := h.userRepo.AddUserSkk(req)
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = user
	ctx.JSON(http.StatusOK, response)
}

// TypeUserRole implements PolicySkkController
func (h userSkkController) TypeUserRole(ctx *gin.Context) {
	response := &model.Response{}
	role, resp, err := h.userRepo.TypeUserRole()
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = role
	ctx.JSON(http.StatusOK, response)
}

// NewUserController -> returns new user controller
func NewUserSkkController(repo repository.UserSkkRepository, policyRepo repository.PolicySkkRepository) UserSkkController {
	return userSkkController{
		userRepo:   repo,
		policyRepo: policyRepo,
	}
}
