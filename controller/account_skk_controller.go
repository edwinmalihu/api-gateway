package controller

import (
	"auth-services/model"
	"auth-services/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountSkkController interface {
	AddAccount(*gin.Context)
	List(*gin.Context)
	UpdateAccount(*gin.Context)
	DetailAccount(*gin.Context)
}

type accountSkkController struct {
	skkRepo repository.AccountSkkRepository
}

// DetailAccount implements AccountSkkController
func (a accountSkkController) DetailAccount(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var reqAccount model.ReqDetail
	ctx.ShouldBind(&reqAccount)
	// json.Unmarshal([]byte(request.Payload), &pengajuan)

	// signature := ctx.GetHeader("Signature")
	// if signature == "" {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "No Signature header found"})
	// 	return
	// }
	// if token, err := utils.ValidateToken(signature); err != nil {
	// 	fmt.Println("signature ", signature, err.Error())
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Not Valid Signature"})
	// 	return
	// } else {
	// 	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Failed parse Signature"})
	// 		return
	// 	} else {
	// 		if token.Valid {
	// 			data_string := claims["data"].(string)
	// 			var signature_request model.SignatureRequest
	// 			json.Unmarshal([]byte(data_string), &signature_request)

	// 			var signature_payload model.RequestFilterKKP
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)
	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != pengajuan) {
	// 				ctx.JSON(http.StatusBadRequest, gin.H{
	// 					"error": "request tampered"})
	// 				return
	// 			}
	// 		} else {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{
	// 				"error": "Not Valid Signature"})
	// 			return
	// 		}
	// 	}
	// }

	skk, resp, err := a.skkRepo.DetailAccount(reqAccount)
	response := &model.Response{}
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
	response.RespData = skk
	ctx.JSON(http.StatusOK, response)
}

// UpdateAccount implements AccountSkkController
func (a accountSkkController) UpdateAccount(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var reqAccount model.ReqUpdateAccount
	ctx.ShouldBind(&reqAccount)
	// json.Unmarshal([]byte(request.Payload), &pengajuan)

	// signature := ctx.GetHeader("Signature")
	// if signature == "" {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "No Signature header found"})
	// 	return
	// }
	// if token, err := utils.ValidateToken(signature); err != nil {
	// 	fmt.Println("signature ", signature, err.Error())
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Not Valid Signature"})
	// 	return
	// } else {
	// 	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Failed parse Signature"})
	// 		return
	// 	} else {
	// 		if token.Valid {
	// 			data_string := claims["data"].(string)
	// 			var signature_request model.SignatureRequest
	// 			json.Unmarshal([]byte(data_string), &signature_request)

	// 			var signature_payload model.RequestFilterKKP
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)
	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != pengajuan) {
	// 				ctx.JSON(http.StatusBadRequest, gin.H{
	// 					"error": "request tampered"})
	// 				return
	// 			}
	// 		} else {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{
	// 				"error": "Not Valid Signature"})
	// 			return
	// 		}
	// 	}
	// }

	skk, resp, err := a.skkRepo.UpdateAccount(reqAccount)
	response := &model.Response{}
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
	response.RespData = skk
	ctx.JSON(http.StatusOK, response)
}

// AddAccount implements AccountSkkController
func (a accountSkkController) AddAccount(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var reqAccount model.RequestAdd
	ctx.ShouldBind(&reqAccount)
	// json.Unmarshal([]byte(request.Payload), &pengajuan)

	// signature := ctx.GetHeader("Signature")
	// if signature == "" {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "No Signature header found"})
	// 	return
	// }
	// if token, err := utils.ValidateToken(signature); err != nil {
	// 	fmt.Println("signature ", signature, err.Error())
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Not Valid Signature"})
	// 	return
	// } else {
	// 	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Failed parse Signature"})
	// 		return
	// 	} else {
	// 		if token.Valid {
	// 			data_string := claims["data"].(string)
	// 			var signature_request model.SignatureRequest
	// 			json.Unmarshal([]byte(data_string), &signature_request)

	// 			var signature_payload model.RequestFilterKKP
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)
	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != pengajuan) {
	// 				ctx.JSON(http.StatusBadRequest, gin.H{
	// 					"error": "request tampered"})
	// 				return
	// 			}
	// 		} else {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{
	// 				"error": "Not Valid Signature"})
	// 			return
	// 		}
	// 	}
	// }

	skk, resp, err := a.skkRepo.AddAccount(reqAccount)
	response := &model.Response{}
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
	response.RespData = skk
	ctx.JSON(http.StatusOK, response)
}

// List implements AccountSkkController
func (a accountSkkController) List(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	// var pengajuan model.RequestFilterKKP
	// json.Unmarshal([]byte(request.Payload), &pengajuan)

	// signature := ctx.GetHeader("Signature")
	// if signature == "" {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "No Signature header found"})
	// 	return
	// }
	// if token, err := utils.ValidateToken(signature); err != nil {
	// 	fmt.Println("signature ", signature, err.Error())
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Not Valid Signature"})
	// 	return
	// } else {
	// 	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Failed parse Signature"})
	// 		return
	// 	} else {
	// 		if token.Valid {
	// 			data_string := claims["data"].(string)
	// 			var signature_request model.SignatureRequest
	// 			json.Unmarshal([]byte(data_string), &signature_request)

	// 			var signature_payload model.RequestFilterKKP
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)
	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != pengajuan) {
	// 				ctx.JSON(http.StatusBadRequest, gin.H{
	// 					"error": "request tampered"})
	// 				return
	// 			}
	// 		} else {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{
	// 				"error": "Not Valid Signature"})
	// 			return
	// 		}
	// 	}
	// }

	skk, resp, err := a.skkRepo.List()
	if len(skk) < 1 {
		skk = make([]model.ResListAccount, 0)
	}
	response := &model.Response{}
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
	response.RespData = skk
	ctx.JSON(http.StatusOK, response)
}

func NewAccountSkkController(repo repository.AccountSkkRepository) AccountSkkController {
	return accountSkkController{
		skkRepo: repo,
	}
}
