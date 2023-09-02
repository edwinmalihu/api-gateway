package controller

import (
	"auth-services/model"
	"auth-services/repository"
	"auth-services/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HoldingController interface {
	AddHolding(*gin.Context)
	ListHolding(*gin.Context)
	DetailHolding(*gin.Context)
	UpdateHolding(*gin.Context)
	TotalHolding(*gin.Context)
}

type holdingController struct {
	holdingRepo repository.HolodngRepository
}

func NewHoldingController(repo repository.HolodngRepository) HoldingController {
	return holdingController{
		holdingRepo: repo,
	}
}

// TotalHolding implements HoldingController
func (h holdingController) TotalHolding(ctx *gin.Context) {
	//var request model.SignatureRequest
	//ctx.ShouldBind(&request)

	// var holding model.RequestFilterDate
	// ctx.ShouldBind(&holding)
	//json.Unmarshal([]byte(request.Payload), &pengajuan)

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

	// 			var signature_payload model.RequestFilterDate
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

	resholding, resp, err := h.holdingRepo.TotalHolding()
	// if len(resholding) < 1 {
	// 	resholding = make([]model.ResListHoldng, 0)
	// }
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
	response.RespData = resholding
	ctx.JSON(http.StatusOK, response)
}

func (h holdingController) AddHolding(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var holding model.RequestHolding
	ctx.ShouldBind(&holding)
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

	// 			var signature_payload model.RequestPengajuan
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			fmt.Printf("json_request: %s\n, json_payload: %s\n", request, pengajuan)
	// 			fmt.Printf("signature_request: %s\n, signature_payload: %s\n", signature_request, signature_payload)
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

	validasistruct := utils.ValidasiStruct(holding)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	resholding, resp, err := h.holdingRepo.AddHolding(holding)
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
	response.RespData = resholding
	ctx.JSON(http.StatusOK, response)
}

func (h holdingController) DetailHolding(ctx *gin.Context) {
	// var request model.SignatureRequest
	//ctx.ShouldBind(&request)

	var holding model.GetHolding
	ctx.ShouldBind(&holding)
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

	// 			var signature_payload model.RequestByNoPengajuan
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
	validasistruct := utils.ValidasiStruct(holding)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	resholding, resp, err := h.holdingRepo.DetailHolding(holding)
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
	response.RespData = resholding
	ctx.JSON(http.StatusOK, response)
}

func (h holdingController) ListHolding(ctx *gin.Context) {
	//var request model.SignatureRequest
	//ctx.ShouldBind(&request)

	//json.Unmarshal([]byte(request.Payload), &pengajuan)

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

	// 			var signature_payload model.RequestFilterDate
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

	resholding, resp, err := h.holdingRepo.ListHolding()
	if len(resholding) < 1 {
		resholding = make([]model.ResListHoldng, 0)
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
	response.RespData = resholding
	ctx.JSON(http.StatusOK, response)
}

func (h holdingController) UpdateHolding(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var holding model.UpdateHolding
	ctx.ShouldBind(&holding)
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

	// 			var signature_payload model.RequestUpdate
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

	validasistruct := utils.ValidasiStruct(holding)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	resholding, resp, err := h.holdingRepo.UpdateHolding(holding)
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
	response.RespData = resholding
	ctx.JSON(http.StatusOK, response)
}
