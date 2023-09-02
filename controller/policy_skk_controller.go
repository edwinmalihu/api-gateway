package controller

import (
	"auth-services/middleware"
	"auth-services/model"
	"auth-services/repository"
	"auth-services/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PolicySkkController interface {
	AddRoleSkk(*gin.Context)
	UpdateRoleSkk(*gin.Context)
	DeleteRoleSkk(*gin.Context)
	ListRoleSkk(*gin.Context)
	DetailRoleSkk(*gin.Context)
}

type policySkkController struct {
	policyRepo repository.PolicySkkRepository
}

// AddRoleSkk implements PolicySkkController
func (h policySkkController) AddRoleSkk(ctx *gin.Context) {
	var req model.RequestData
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

	role, resp, err := h.policyRepo.AddRoleSkk(req)
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

// DeleteRoleSkk implements PolicySkkController
func (h policySkkController) DeleteRoleSkk(ctx *gin.Context) {
	var req model.DetailtGetRole
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

	role, resp, err := h.policyRepo.DeleteRoleSkk(req.Role)
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

// DetailRoleSkk implements PolicySkkController
func (h policySkkController) DetailRoleSkk(ctx *gin.Context) {
	var req model.DetailtGetRole
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

	role, resp, err := h.policyRepo.DetailRoleSkk(req)
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

// ListRoleSkk implements PolicySkkController
func (h policySkkController) ListRoleSkk(ctx *gin.Context) {
	response := &model.Response{}
	role, resp, err := h.policyRepo.ListRoleSkk()
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

// UpdateRoleSkk implements PolicySkkController
func (h policySkkController) UpdateRoleSkk(ctx *gin.Context) {
	var req model.RequestData
	rolename := ctx.Query("role")
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

	role, resp, err := h.policyRepo.UpdateRoleSkk(req, rolename)
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

func NewPolicySkkController(repo repository.PolicySkkRepository) PolicySkkController {
	return policySkkController{
		policyRepo: repo,
	}
}
