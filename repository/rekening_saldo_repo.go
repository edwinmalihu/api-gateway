package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type rekeningSalodRepository struct {
	client *resty.Client
}

// GetCsv implements RekeningSaldoRepository
func (r rekeningSalodRepository) GetCsv(req model.RequestRekeningSaldo) (model.ResponseUrl, *resty.Response, error) {
	var result model.ResponseUrl
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"id_holding": req.IdHolding,
			"id_kkks":    req.IdKkks,
			"id_wk":      req.IdWk,
			"tanggal":    req.Tanggal,
			"tipe":       req.Tipe,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_REKENING_SALDO_SERVICE"), "/api/csv"))

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

// GetExcel implements RekeningSaldoRepository
func (r rekeningSalodRepository) GetExcel(req model.RequestRekeningSaldo) (model.ResponseUrl, *resty.Response, error) {
	var result model.ResponseUrl
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"id_holding": req.IdHolding,
			"id_kkks":    req.IdKkks,
			"id_wk":      req.IdWk,
			"tanggal":    req.Tanggal,
			"tipe":       req.Tipe,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_REKENING_SALDO_SERVICE"), "/api/excel"))

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

// GetListTable implements RekeningSaldoRepository
func (r rekeningSalodRepository) GetListTable(req model.RequestRekeningSaldo, user model.ReqValidasiUsername) ([]model.ResultListTable, *resty.Response, error) {
	var result []model.ResultListTable
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"id_holding": req.IdHolding,
			"id_kkks":    req.IdKkks,
			"id_wk":      req.IdWk,
			"tanggal":    req.Tanggal,
			"tipe":       req.Tipe,
			"username":   user.Username,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_REKENING_SALDO_SERVICE"), "/api/list-table"))

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

// GetSummaryTotRek implements RekeningSaldoRepository
func (r rekeningSalodRepository) GetSummaryTotRek(req model.RequestRekeningSaldo, user model.ReqValidasiUsername) (model.ResCardRekeningSaldo, *resty.Response, error) {
	var result model.ResCardRekeningSaldo
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"id_holding": req.IdHolding,
			"id_kkks":    req.IdKkks,
			"id_wk":      req.IdWk,
			"tanggal":    req.Tanggal,
			"tipe":       req.Tipe,
			"username":   user.Username,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_REKENING_SALDO_SERVICE"), "/api/summary"))

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

// GetDropdown implements RekeningSaldoRepository
func (r rekeningSalodRepository) GetDropdown(req model.RequestRekeningSaldo, user model.ReqValidasiUsername) ([]model.ResDropdownWK, *resty.Response, error) {
	var result []model.ResDropdownWK
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"tipe_user":  req.Company,
			"id_holding": req.IdHolding,
			"id_kkks":    req.IdKkks,
			"id_wk":      req.IdWk,
			"username":   user.Username,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_REKENING_SALDO_SERVICE"), "/api/drop-wk"))
		//Get(fmt.Sprintf("%s%s", os.Getenv("HOST_REKENING_SALDO_SERVICE"), "/api/drop-wk"))

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

type RekeningSaldoRepository interface {
	GetDropdown(model.RequestRekeningSaldo, model.ReqValidasiUsername) ([]model.ResDropdownWK, *resty.Response, error)
	GetSummaryTotRek(model.RequestRekeningSaldo, model.ReqValidasiUsername) (model.ResCardRekeningSaldo, *resty.Response, error)
	GetListTable(model.RequestRekeningSaldo, model.ReqValidasiUsername) ([]model.ResultListTable, *resty.Response, error)
	GetExcel(model.RequestRekeningSaldo) (model.ResponseUrl, *resty.Response, error)
	GetCsv(model.RequestRekeningSaldo) (model.ResponseUrl, *resty.Response, error)
}

func NewRekeningSaldoRepository() RekeningSaldoRepository {
	return rekeningSalodRepository{
		client: resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}
