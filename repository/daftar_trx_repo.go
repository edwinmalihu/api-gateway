package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type daftratrxRepository struct {
	client *resty.Client
}

// GetDropdownWk implements DaftarTrxRepository
func (d daftratrxRepository) GetDropdownWk(req model.RequestRekeningSaldo, user model.ReqValidasiUsername) ([]model.ResDropdownWK, *resty.Response, error) {
	var result []model.ResDropdownWK
	resp, err := d.client.R().
		SetQueryParams(map[string]string{
			"tipe_user":  req.Company,
			"id_holding": req.IdHolding,
			"id_kkks":    req.IdKkks,
			"id_wk":      req.IdWk,
			"username":   user.Username,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_REKENING_SALDO_SERVICE"), "/api/drop-wk"))

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

// ExportCsv implements DaftarTrxRepository
func (d daftratrxRepository) ExportCsv(req model.RequestDaftarTrx) (model.ResponseUrl, *resty.Response, error) {
	var result model.ResponseUrl
	resp, err := d.client.R().
		SetQueryParams(map[string]string{
			"id_holding":   req.IdHolding,
			"id_kkks":      req.IdKkks,
			"id_wk":        req.IdWk,
			"tipe":         req.Tipe,
			"account_name": req.AccountName,
			"account_no":   req.AccountNo,
			"start_date":   req.StartDate,
			"end_date":     req.EndDate,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DAFTAR_TRX"), "/api/csv"))

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

// ExportExcel implements DaftarTrxRepository
func (d daftratrxRepository) ExportExcel(req model.RequestDaftarTrx) (model.ResponseUrl, *resty.Response, error) {
	var result model.ResponseUrl
	resp, err := d.client.R().
		SetQueryParams(map[string]string{
			"id_holding":   req.IdHolding,
			"id_kkks":      req.IdKkks,
			"id_wk":        req.IdWk,
			"tipe":         req.Tipe,
			"account_name": req.AccountName,
			"account_no":   req.AccountNo,
			"start_date":   req.StartDate,
			"end_date":     req.EndDate,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DAFTAR_TRX"), "/api/excel"))

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

// CardBunga implements DaftarTrxRepository
func (d daftratrxRepository) CardBunga(req model.RequestDaftarTrx) (model.NominalTrx, *resty.Response, error) {
	var result model.NominalTrx
	resp, err := d.client.R().
		SetQueryParams(map[string]string{
			"id_holding":   req.IdHolding,
			"id_kkks":      req.IdKkks,
			"id_wk":        req.IdWk,
			"tipe":         req.Tipe,
			"account_name": req.AccountName,
			"account_no":   req.AccountNo,
			"start_date":   req.StartDate,
			"end_date":     req.EndDate,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DAFTAR_TRX"), "/api/card-bunga"))

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

// CardLainLain implements DaftarTrxRepository
func (d daftratrxRepository) CardLainLain(req model.RequestDaftarTrx) (model.NominalTrx, *resty.Response, error) {
	var result model.NominalTrx
	resp, err := d.client.R().
		SetQueryParams(map[string]string{
			"id_holding":   req.IdHolding,
			"id_kkks":      req.IdKkks,
			"id_wk":        req.IdWk,
			"tipe":         req.Tipe,
			"account_name": req.AccountName,
			"account_no":   req.AccountNo,
			"start_date":   req.StartDate,
			"end_date":     req.EndDate,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DAFTAR_TRX"), "/api/card-lain"))

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

// CardPenerimaan implements DaftarTrxRepository
func (d daftratrxRepository) CardPenerimaan(req model.RequestDaftarTrx) (model.NominalTrx, *resty.Response, error) {
	var result model.NominalTrx
	resp, err := d.client.R().
		SetQueryParams(map[string]string{
			"id_holding":   req.IdHolding,
			"id_kkks":      req.IdKkks,
			"id_wk":        req.IdWk,
			"tipe":         req.Tipe,
			"account_name": req.AccountName,
			"account_no":   req.AccountNo,
			"start_date":   req.StartDate,
			"end_date":     req.EndDate,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DAFTAR_TRX"), "/api/card-penerimaan"))

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

// CardPengeluaran implements DaftarTrxRepository
func (d daftratrxRepository) CardPengeluaran(req model.RequestDaftarTrx) (model.NominalTrx, *resty.Response, error) {
	var result model.NominalTrx
	resp, err := d.client.R().
		SetQueryParams(map[string]string{
			"id_holding":   req.IdHolding,
			"id_kkks":      req.IdKkks,
			"id_wk":        req.IdWk,
			"tipe":         req.Tipe,
			"account_name": req.AccountName,
			"account_no":   req.AccountNo,
			"start_date":   req.StartDate,
			"end_date":     req.EndDate,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DAFTAR_TRX"), "/api/card-pengeluaran"))

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

// CardSaldo implements DaftarTrxRepository
func (d daftratrxRepository) CardSaldo(req model.RequestDaftarTrx) (model.NominalTrx, *resty.Response, error) {
	var result model.NominalTrx
	resp, err := d.client.R().
		SetQueryParams(map[string]string{
			"id_holding":   req.IdHolding,
			"id_kkks":      req.IdKkks,
			"id_wk":        req.IdWk,
			"tipe":         req.Tipe,
			"account_name": req.AccountName,
			"account_no":   req.AccountNo,
			"start_date":   req.StartDate,
			"end_date":     req.EndDate,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DAFTAR_TRX"), "/api/card-saldo"))

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

// ListTable implements DaftarTrxRepository
func (d daftratrxRepository) ListTable(req model.RequestDaftarTrx) (model.ResponseSummaryTrx, *resty.Response, error) {
	var result model.ResponseSummaryTrx
	resp, err := d.client.R().
		SetQueryParams(map[string]string{
			"id_holding":   req.IdHolding,
			"id_kkks":      req.IdKkks,
			"id_wk":        req.IdWk,
			"tipe":         req.Tipe,
			"account_name": req.AccountName,
			"account_no":   req.AccountNo,
			"start_date":   req.StartDate,
			"end_date":     req.EndDate,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DAFTAR_TRX"), "/api/table-trx"))

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

// GetDropdownAccount implements DaftarTrxRepository
func (d daftratrxRepository) GetDropdownAccount(req model.RequestDaftarTrx) ([]model.ResDropAccountName, *resty.Response, error) {
	var result []model.ResDropAccountName
	resp, err := d.client.R().
		SetQueryParams(map[string]string{
			"tipe_user":  req.Company,
			"id_holding": req.IdHolding,
			"id_kkks":    req.IdKkks,
			"id_wk":      req.IdWk,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_DAFTAR_TRX"), "/api/drop-account"))

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

type DaftarTrxRepository interface {
	GetDropdownAccount(model.RequestDaftarTrx) ([]model.ResDropAccountName, *resty.Response, error)
	ListTable(model.RequestDaftarTrx) (model.ResponseSummaryTrx, *resty.Response, error)
	CardSaldo(model.RequestDaftarTrx) (model.NominalTrx, *resty.Response, error)
	CardBunga(model.RequestDaftarTrx) (model.NominalTrx, *resty.Response, error)
	CardPenerimaan(model.RequestDaftarTrx) (model.NominalTrx, *resty.Response, error)
	CardPengeluaran(model.RequestDaftarTrx) (model.NominalTrx, *resty.Response, error)
	CardLainLain(model.RequestDaftarTrx) (model.NominalTrx, *resty.Response, error)
	ExportExcel(model.RequestDaftarTrx) (model.ResponseUrl, *resty.Response, error)
	ExportCsv(model.RequestDaftarTrx) (model.ResponseUrl, *resty.Response, error)
	GetDropdownWk(model.RequestRekeningSaldo, model.ReqValidasiUsername) ([]model.ResDropdownWK, *resty.Response, error)
}

func NewDaftarTrxRepository() DaftarTrxRepository {
	return daftratrxRepository{
		client: resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}
