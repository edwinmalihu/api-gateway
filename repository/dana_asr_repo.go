package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type asrRepository struct {
	client *resty.Client
}

// TrenBungaNettoPph implements AsrRepository
func (a asrRepository) TrenBungaNettoPph(req model.RequsestAsr, user model.ReqValidasiUsername) (model.ResTrenBunga, *resty.Response, error) {
	var result model.ResTrenBunga
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"tipe_user":  req.Company,
		"id_holding": req.IdHolding,
		"id_kkks":    req.IdKkks,
		"id_wk":      req.IdWk,
		"username":   user.Username,
	}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DANA_ASR_SERVICE"), "/api/tren-bunga"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return result, resp, err
}

// PenerimaanPengeluaran implements AsrRepository
func (a asrRepository) PenerimaanPengeluaran(req model.RequsestAsr, user model.ReqValidasiUsername) (model.ResPenerimaanPengeluaran, *resty.Response, error) {
	var result model.ResPenerimaanPengeluaran
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"tipe_user":  req.Company,
		"id_holding": req.IdHolding,
		"id_kkks":    req.IdKkks,
		"id_wk":      req.IdWk,
		"username":   user.Username,
	}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DANA_ASR_SERVICE"), "/api/penerimaan-pengeluaran"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return result, resp, err
}

// GrafikGiroDepo implements AsrRepository
func (a asrRepository) GrafikGiroDepo(req model.RequsestAsr, user model.ReqValidasiUsername) (model.DataGrafikGiroDepo, *resty.Response, error) {
	var result model.DataGrafikGiroDepo
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"tipe_user":  req.Company,
		"id_holding": req.IdHolding,
		"id_kkks":    req.IdKkks,
		"id_wk":      req.IdWk,
		"username":   user.Username,
	}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DANA_ASR_SERVICE"), "/api/giro-depo"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return result, resp, err
}

// TotalDanaAsr implements AsrRepository
func (a asrRepository) TotalDanaAsr(req model.RequsestAsr, user model.ReqValidasiUsername) (model.GrafikTotalDana, *resty.Response, error) {
	var result model.GrafikTotalDana
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"tipe_user":  req.Company,
		"id_holding": req.IdHolding,
		"id_kkks":    req.IdKkks,
		"id_wk":      req.IdWk,
		"username":   user.Username,
	}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DANA_ASR_SERVICE"), "/api/grafik-total"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return result, resp, err
}

// CardDashboard implements AsrRepository
func (a asrRepository) CardDashboard(req model.RequsestAsr, user model.ReqValidasiUsername) (model.CardDashboard, *resty.Response, error) {
	var result model.CardDashboard
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"tipe_user":  req.Company,
		"id_holding": req.IdHolding,
		"id_kkks":    req.IdKkks,
		"id_wk":      req.IdWk,
		"username":   user.Username,
	}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DANA_ASR_SERVICE"), "/api/data-card"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return result, resp, err
}

type AsrRepository interface {
	CardDashboard(model.RequsestAsr, model.ReqValidasiUsername) (model.CardDashboard, *resty.Response, error)
	TotalDanaAsr(model.RequsestAsr, model.ReqValidasiUsername) (model.GrafikTotalDana, *resty.Response, error)
	GrafikGiroDepo(model.RequsestAsr, model.ReqValidasiUsername) (model.DataGrafikGiroDepo, *resty.Response, error)
	PenerimaanPengeluaran(model.RequsestAsr, model.ReqValidasiUsername) (model.ResPenerimaanPengeluaran, *resty.Response, error)
	TrenBungaNettoPph(model.RequsestAsr, model.ReqValidasiUsername) (model.ResTrenBunga, *resty.Response, error)
}

func NewAsrRepository() AsrRepository {
	return asrRepository{
		client: resty.New().SetTimeout(600 * time.Second).SetRetryCount(3),
		//client: resty.New(),
	}
}
