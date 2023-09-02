package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type wkRepository struct {
	client *resty.Client
}

// DeleteWK implements WkRepository
func (a wkRepository) DeleteWK(req model.ReqGetWk) (model.ResSuccess, *resty.Response, error) {
	var result model.ResSuccess
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"id_wk": req.IdWk,
	}).
		SetResult(&result).
		Delete(fmt.Sprintf("%s%s", os.Getenv("HOST_WK_SERVICE"), "/api/delete"))

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

// DrodownHolding implements WkRepository
func (a wkRepository) DrodownHolding() ([]model.ListDropdownHolidng, *resty.Response, error) {
	var result []model.ListDropdownHolidng
	resp, err := a.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_WK_SERVICE"), "/api/drop-holding"))

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

// DropdownKkks implements WkRepository
func (a wkRepository) DropdownKkks() ([]model.ListDropdownK3S, *resty.Response, error) {
	var result []model.ListDropdownK3S
	resp, err := a.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_WK_SERVICE"), "/api/drop-kkks"))

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

// AddWk implements WkRepository
func (a wkRepository) AddWk(request model.ReqAddWk) (model.ResSuccess, *resty.Response, error) {
	var result model.ResSuccess
	resp, err := a.client.R().
		SetBody(&request).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_WK_SERVICE"), "/api/add"))
		// Post("http://localhost:8087/api/add")

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

// DetailWk implements WkRepository
func (a wkRepository) DetailWk(req model.ReqGetWk, user model.ReqValidasiUsername) (model.ResDetailWk, *resty.Response, error) {
	var result model.ResDetailWk
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"id_wk":    req.IdWk,
		"username": user.Username,
	}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_WK_SERVICE"), "/api/detail"))
		//Get("http://localhost:8087/api/detail")

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

// EditKWk implements WkRepository
func (a wkRepository) EditKWk(request model.ReqEditWK) (model.ResSuccess, *resty.Response, error) {
	var result model.ResSuccess
	resp, err := a.client.R().
		SetBody(&request).
		SetResult(&result).
		Put(fmt.Sprintf("%s%s", os.Getenv("HOST_WK_SERVICE"), "/api/edit"))

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

// GetByWk implements WkRepository
func (a wkRepository) GetByWk(req model.ReqGetWk) (model.ResRelasiWk, *resty.Response, error) {
	var result model.ResRelasiWk
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"id_wk": req.IdWk,
	}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_WK_SERVICE"), "/api/update"))

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

// ListWk implements WkRepository
func (a wkRepository) ListWk(req model.ReqParam, user model.ReqValidasiUsername) ([]model.ResListWK, *resty.Response, error) {
	var result []model.ResListWK
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"id_holding": req.IdHolding,
		"id_kkks":    req.IdKkks,
		"id_wk":      req.IdWk,
		"username":   user.Username,
	}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_WK_SERVICE"), "/api/list"))

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

// TotalWk implements WkRepository
func (a wkRepository) TotalWk(req model.ReqParam, user model.ReqValidasiUsername) (model.Total, *resty.Response, error) {
	var result model.Total
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"id_holding": req.IdHolding,
		"id_kkks":    req.IdKkks,
		"id_wk":      req.IdWk,
		"username":   user.Username,
	}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_WK_SERVICE"), "/api/total"))

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

type WkRepository interface {
	AddWk(model.ReqAddWk) (model.ResSuccess, *resty.Response, error)
	DetailWk(model.ReqGetWk, model.ReqValidasiUsername) (model.ResDetailWk, *resty.Response, error)
	EditKWk(model.ReqEditWK) (model.ResSuccess, *resty.Response, error)
	GetByWk(model.ReqGetWk) (model.ResRelasiWk, *resty.Response, error)
	TotalWk(model.ReqParam, model.ReqValidasiUsername) (model.Total, *resty.Response, error)
	ListWk(model.ReqParam, model.ReqValidasiUsername) ([]model.ResListWK, *resty.Response, error)
	DrodownHolding() ([]model.ListDropdownHolidng, *resty.Response, error)
	DropdownKkks() ([]model.ListDropdownK3S, *resty.Response, error)
	DeleteWK(model.ReqGetWk) (model.ResSuccess, *resty.Response, error)
}

func NewWkRepository() WkRepository {
	return wkRepository{
		client: resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}
