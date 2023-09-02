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

type asrontroller struct {
	asrRepo repository.AsrRepository
}

// TrenBunga implements AsrController
func (a asrontroller) TrenBunga(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req model.RequsestAsr
	ctx.ShouldBind(&req)
	// json.Unmarshal([]byte(request.Payload), &req)

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

	// 			var signature_payload model.RequsestAsr
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req) {
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

	var user model.ReqValidasiUsername
	ctx.ShouldBind(&user)

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	validasistructUser := utils.ValidasiStruct(user)

	if validasistructUser == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	CekUser := middleware.GetUserToken(ctx)

	if CekUser != user.Username {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not Valid Token"})
		return
	}

	asr, resp, err := a.asrRepo.TrenBungaNettoPph(req, user)
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
	response.RespData = asr
	ctx.JSON(http.StatusOK, response)
}

// TrenPenerimaanPengeluaran implements AsrController
func (a asrontroller) TrenPenerimaanPengeluaran(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req model.RequsestAsr
	ctx.ShouldBind(&req)
	// json.Unmarshal([]byte(request.Payload), &req)

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

	// 			var signature_payload model.RequsestAsr
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req) {
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

	var user model.ReqValidasiUsername
	ctx.ShouldBind(&user)

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	validasistructUser := utils.ValidasiStruct(user)

	if validasistructUser == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	CekUser := middleware.GetUserToken(ctx)

	if CekUser != user.Username {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not Valid Token"})
		return
	}

	asr, resp, err := a.asrRepo.PenerimaanPengeluaran(req, user)
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
	response.RespData = asr
	ctx.JSON(http.StatusOK, response)
}

// GrafikGiroDepo implements AsrController
func (a asrontroller) GrafikGiroDepo(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req model.RequsestAsr
	ctx.ShouldBind(&req)
	// json.Unmarshal([]byte(request.Payload), &req)

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

	// 			var signature_payload model.RequsestAsr
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req) {
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

	var user model.ReqValidasiUsername
	ctx.ShouldBind(&user)

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	validasistructUser := utils.ValidasiStruct(user)

	if validasistructUser == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	CekUser := middleware.GetUserToken(ctx)

	if CekUser != user.Username {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not Valid Token"})
		return
	}

	asr, resp, err := a.asrRepo.GrafikGiroDepo(req, user)
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
	response.RespData = asr
	ctx.JSON(http.StatusOK, response)
}

// GrafikTotalDanaAsr implements AsrController
func (a asrontroller) GrafikTotalDanaAsr(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req model.RequsestAsr
	ctx.ShouldBind(&req)
	// json.Unmarshal([]byte(request.Payload), &req)

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

	// 			var signature_payload model.RequsestAsr
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req) {
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

	var user model.ReqValidasiUsername
	ctx.ShouldBind(&user)

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	validasistructUser := utils.ValidasiStruct(user)

	if validasistructUser == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	CekUser := middleware.GetUserToken(ctx)

	if CekUser != user.Username {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not Valid Token"})
		return
	}

	asr, resp, err := a.asrRepo.TotalDanaAsr(req, user)
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
	response.RespData = asr
	ctx.JSON(http.StatusOK, response)
}

// CardDashboard implements AsrController
func (a asrontroller) CardDashboard(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req model.RequsestAsr
	ctx.ShouldBind(&req)
	// json.Unmarshal([]byte(request.Payload), &req)

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

	// 			var signature_payload model.RequsestAsr
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req) {
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

	var user model.ReqValidasiUsername
	ctx.ShouldBind(&user)

	validasistruct := utils.ValidasiStruct(req)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	validasistructUser := utils.ValidasiStruct(user)

	if validasistructUser == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	CekUser := middleware.GetUserToken(ctx)

	if CekUser != user.Username {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not Valid Token"})
		return
	}

	asr, resp, err := a.asrRepo.CardDashboard(req, user)
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
	response.RespData = asr
	ctx.JSON(http.StatusOK, response)
}

type AsrController interface {
	CardDashboard(*gin.Context)
	GrafikTotalDanaAsr(*gin.Context)
	GrafikGiroDepo(*gin.Context)
	TrenPenerimaanPengeluaran(*gin.Context)
	TrenBunga(*gin.Context)
}

func NewAsrController(repo repository.AsrRepository) AsrController {
	return asrontroller{
		asrRepo: repo,
	}
}
