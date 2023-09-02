package controller

import (
	"auth-services/middleware"
	"auth-services/model"
	"auth-services/repository"
	"auth-services/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type daftartrxController struct {
	reksalRepo repository.DaftarTrxRepository
}

// GetDropdownWk implements DaftarTrxController
func (d daftartrxController) GetDropdownWk(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req_reks model.RequestRekeningSaldo
	ctx.ShouldBind(&req_reks)
	// json.Unmarshal([]byte(request.Payload), &req_reks)

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

	// 			var signature_payload model.GetIdRate
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req_reks) {
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

	// validasistruct := utils.ValidasiStruct(req_reks)

	// if validasistruct == false {
	// 	ctx.JSON(http.StatusBadRequest, "bad request")
	// 	return
	// }

	var user model.ReqValidasiUsername

	ctx.ShouldBind(&user)

	CekUser := middleware.GetUserToken(ctx)

	user.Username = CekUser

	rate, resp, err := d.reksalRepo.GetDropdownWk(req_reks, user)
	if len(rate) < 0 {
		rate = make([]model.ResDropdownWK, 0)
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
	response.RespData = rate
	ctx.JSON(http.StatusOK, response)
}

// ExportCsv implements DaftarTrxController
func (d daftartrxController) ExportCsv(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req_reks model.RequestDaftarTrx
	ctx.ShouldBind(&req_reks)
	// json.Unmarshal([]byte(request.Payload), &req_reks)

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

	// 			var signature_payload model.RequestDaftarTrx
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req_reks) {
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

	rate, resp, err := d.reksalRepo.ExportCsv(req_reks)
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
	response.RespData = rate
	ctx.JSON(http.StatusOK, response)
}

// ExportExcel implements DaftarTrxController
func (d daftartrxController) ExportExcel(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req_reks model.RequestDaftarTrx
	ctx.ShouldBind(&req_reks)
	// json.Unmarshal([]byte(request.Payload), &req_reks)

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

	// 			var signature_payload model.RequestDaftarTrx
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req_reks) {
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

	rate, resp, err := d.reksalRepo.ExportExcel(req_reks)
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
	response.RespData = rate
	ctx.JSON(http.StatusOK, response)
}

// CardBunga implements DaftarTrxController
func (d daftartrxController) CardBunga(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req_reks model.RequestDaftarTrx
	ctx.ShouldBind(&req_reks)
	// json.Unmarshal([]byte(request.Payload), &req_reks)

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

	// 			var signature_payload model.RequestDaftarTrx
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req_reks) {
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

	// validasistruct := utils.ValidasiStruct(req_reks)

	// if validasistruct == false {
	// 	ctx.JSON(http.StatusBadRequest, "bad request")
	// 	return
	// }

	validateDate := utils.ValidateBackDate(req_reks.StartDate, req_reks.EndDate)

	if !validateDate {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	startDate, err := time.Parse("2006-01-02", req_reks.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(startDate)

	endDate, err := time.Parse("2006-01-02", req_reks.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(endDate)

	rate, resp, err := d.reksalRepo.CardBunga(req_reks)
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
	response.RespData = rate
	ctx.JSON(http.StatusOK, response)
}

// CardLainLain implements DaftarTrxController
func (d daftartrxController) CardLainLain(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req_reks model.RequestDaftarTrx
	ctx.ShouldBind(&req_reks)
	// json.Unmarshal([]byte(request.Payload), &req_reks)

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

	// 			var signature_payload model.RequestDaftarTrx
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req_reks) {
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

	// validasistruct := utils.ValidasiStruct(req_reks)

	// if validasistruct == false {
	// 	ctx.JSON(http.StatusBadRequest, "bad request")
	// 	return
	// }

	validateDate := utils.ValidateBackDate(req_reks.StartDate, req_reks.EndDate)

	if !validateDate {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	startDate, err := time.Parse("2006-01-02", req_reks.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(startDate)

	endDate, err := time.Parse("2006-01-02", req_reks.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(endDate)

	rate, resp, err := d.reksalRepo.CardLainLain(req_reks)
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
	response.RespData = rate
	ctx.JSON(http.StatusOK, response)
}

// CardPenerimaan implements DaftarTrxController
func (d daftartrxController) CardPenerimaan(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req_reks model.RequestDaftarTrx
	ctx.ShouldBind(&req_reks)
	// json.Unmarshal([]byte(request.Payload), &req_reks)

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

	// 			var signature_payload model.RequestDaftarTrx
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req_reks) {
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

	// validasistruct := utils.ValidasiStruct(req_reks)

	// if validasistruct == false {
	// 	ctx.JSON(http.StatusBadRequest, "bad request")
	// 	return
	// }

	validateDate := utils.ValidateBackDate(req_reks.StartDate, req_reks.EndDate)

	if !validateDate {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	startDate, err := time.Parse("2006-01-02", req_reks.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(startDate)

	endDate, err := time.Parse("2006-01-02", req_reks.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(endDate)

	rate, resp, err := d.reksalRepo.CardPenerimaan(req_reks)
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
	response.RespData = rate
	ctx.JSON(http.StatusOK, response)
}

// CardPengeluaran implements DaftarTrxController
func (d daftartrxController) CardPengeluaran(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req_reks model.RequestDaftarTrx
	ctx.ShouldBind(&req_reks)
	// json.Unmarshal([]byte(request.Payload), &req_reks)

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

	// 			var signature_payload model.RequestDaftarTrx
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req_reks) {
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

	// validasistruct := utils.ValidasiStruct(req_reks)

	// if validasistruct == false {
	// 	ctx.JSON(http.StatusBadRequest, "bad request")
	// 	return
	// }

	validateDate := utils.ValidateBackDate(req_reks.StartDate, req_reks.EndDate)

	if !validateDate {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	startDate, err := time.Parse("2006-01-02", req_reks.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(startDate)

	endDate, err := time.Parse("2006-01-02", req_reks.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(endDate)

	rate, resp, err := d.reksalRepo.CardPengeluaran(req_reks)
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
	response.RespData = rate
	ctx.JSON(http.StatusOK, response)
}

// CardSaldo implements DaftarTrxController
func (d daftartrxController) CardSaldo(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req_reks model.RequestDaftarTrx
	ctx.ShouldBind(&req_reks)
	// json.Unmarshal([]byte(request.Payload), &req_reks)

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

	// 			var signature_payload model.RequestDaftarTrx
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req_reks) {
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

	// validasistruct := utils.ValidasiStruct(req_reks)

	// if validasistruct == false {
	// 	ctx.JSON(http.StatusBadRequest, "bad request")
	// 	return
	// }

	validateDate := utils.ValidateBackDate(req_reks.StartDate, req_reks.EndDate)

	if !validateDate {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	startDate, err := time.Parse("2006-01-02", req_reks.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(startDate)

	endDate, err := time.Parse("2006-01-02", req_reks.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(endDate)

	rate, resp, err := d.reksalRepo.CardSaldo(req_reks)
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
	response.RespData = rate
	ctx.JSON(http.StatusOK, response)
}

// ListTableTrx implements DaftarTrxController
func (d daftartrxController) ListTableTrx(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req_reks model.RequestDaftarTrx
	ctx.ShouldBind(&req_reks)
	// json.Unmarshal([]byte(request.Payload), &req_reks)

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

	// 			var signature_payload model.RequestDaftarTrx
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req_reks) {
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

	// validasistruct := utils.ValidasiStruct(req_reks)

	// if validasistruct == false {
	// 	ctx.JSON(http.StatusBadRequest, "bad request")
	// 	return
	// }

	// Parse the date strings into time.Time values

	validateDate := utils.ValidateBackDate(req_reks.StartDate, req_reks.EndDate)

	if !validateDate {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	startDate, err := time.Parse("2006-01-02", req_reks.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(startDate)

	endDate, err := time.Parse("2006-01-02", req_reks.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(endDate)

	rate, resp, err := d.reksalRepo.ListTable(req_reks)
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
	response.RespData = rate
	ctx.JSON(http.StatusOK, response)
}

// GetDropdownAccount implements DaftarTrxController
func (d daftartrxController) GetDropdownAccount(ctx *gin.Context) {
	// var request model.SignatureRequest
	// ctx.ShouldBind(&request)

	var req_reks model.RequestDaftarTrx
	ctx.ShouldBind(&req_reks)
	// json.Unmarshal([]byte(request.Payload), &req_reks)

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

	// 			var signature_payload model.RequestDaftarTrx
	// 			json.Unmarshal([]byte(signature_request.Payload), &signature_payload)

	// 			if err != nil {
	// 				log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// 			}
	// 			if (request.Host != signature_request.Host) || (request.Path != signature_request.Path) || (signature_payload != req_reks) {
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

	// validasistruct := utils.ValidasiStruct(req_reks)

	// if validasistruct == false {
	// 	ctx.JSON(http.StatusBadRequest, "bad request")
	// 	return
	// }

	validateDate := utils.ValidateBackDate(req_reks.StartDate, req_reks.EndDate)

	if !validateDate {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	startDate, err := time.Parse("2006-01-02", req_reks.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(startDate)

	endDate, err := time.Parse("2006-01-02", req_reks.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(endDate)

	rate, resp, err := d.reksalRepo.GetDropdownAccount(req_reks)
	if len(rate) < 0 {
		rate = make([]model.ResDropAccountName, 0)
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
	response.RespData = rate
	ctx.JSON(http.StatusOK, response)
}

type DaftarTrxController interface {
	GetDropdownAccount(*gin.Context)
	ListTableTrx(*gin.Context)
	CardSaldo(*gin.Context)
	CardBunga(*gin.Context)
	CardPenerimaan(*gin.Context)
	CardPengeluaran(*gin.Context)
	CardLainLain(*gin.Context)
	ExportExcel(*gin.Context)
	ExportCsv(*gin.Context)
	GetDropdownWk(*gin.Context)
}

func NewDaftarTrxController(repo repository.DaftarTrxRepository) DaftarTrxController {
	return daftartrxController{
		reksalRepo: repo,
	}
}
