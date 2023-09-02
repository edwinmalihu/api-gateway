package controller

import (
	"auth-services/model"
	"auth-services/repository"
	"auth-services/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KkksSkkController interface {
	AddKkks(*gin.Context)
	List(*gin.Context)
	UpdateKkks(*gin.Context)
	DetailAccount(*gin.Context)
	EditKkks(*gin.Context)
	TotalKkks(*gin.Context)
	UpdateAccountKkks(*gin.Context)
}

type kkksSkkController struct {
	skkRepo repository.KkksRepository
}

// UpdateAccountKkks implements KkksSkkController
func (a kkksSkkController) UpdateAccountKkks(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var reqAccount model.ReqDetailKkks
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

	validasistruct := utils.ValidasiStruct(reqAccount)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	skk, resp, err := a.skkRepo.UpdateAccountKkks(reqAccount)
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

// AddKkks implements KkksSkkController
func (a kkksSkkController) AddKkks(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var reqAccount model.ReqDetailKkks
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

	validasistruct := utils.ValidasiStruct(reqAccount)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	skk, resp, err := a.skkRepo.AddKkks(reqAccount)
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

// DetailAccount implements KkksSkkController
func (a kkksSkkController) DetailAccount(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var reqAccount model.GetDetailK3s
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

	validasistruct := utils.ValidasiStruct(reqAccount)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	skk, resp, err := a.skkRepo.DetailKkks(reqAccount)
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

// EditKkks implements KkksSkkController
func (a kkksSkkController) EditKkks(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var reqAccount model.GetDetailK3s
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

	skk, resp, err := a.skkRepo.EditKkks(reqAccount)
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

// List implements KkksSkkController
func (a kkksSkkController) List(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	// var reqAccount model.ReqUpdate
	// ctx.ShouldBind(&reqAccount)
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

	skk, resp, err := a.skkRepo.ListKkks()
	if len(skk) < 1 {
		skk = make([]model.DataListKkks, 0)
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

// TotalKkks implements KkksSkkController
func (a kkksSkkController) TotalKkks(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	// var reqAccount model.ReqDetailKkks
	// ctx.ShouldBind(&reqAccount)
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

	skk, resp, err := a.skkRepo.TotalKkks()
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

// UpdateKkks implements KkksSkkController
func (a kkksSkkController) UpdateKkks(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var reqAccount model.ReqUpdate
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

	validasistruct := utils.ValidasiStruct(reqAccount)

	if validasistruct == false {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	skk, resp, err := a.skkRepo.UpdateKkks(reqAccount)
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

func NewKkksSkkController(repo repository.KkksRepository) KkksSkkController {
	return kkksSkkController{
		skkRepo: repo,
	}
}
