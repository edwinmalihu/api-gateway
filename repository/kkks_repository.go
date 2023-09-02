package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type kkksRepository struct {
	client *resty.Client
}

// UpdateAccountKkks implements KkksRepository
func (a kkksRepository) UpdateAccountKkks(request model.ReqDetailKkks) (model.ResAdd, *resty.Response, error) {
	var result model.ResAdd
	resp, err := a.client.R().
		SetBody(&request).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_KKKS_SERVICE"), "/api/edit"))

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

// ListKkks implements KkksRepository
func (a kkksRepository) ListKkks() ([]model.DataListKkks, *resty.Response, error) {
	var result []model.DataListKkks
	resp, err := a.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_KKKS_SERVICE"), "/api/list"))

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

// EditKkks implements KkksRepository
func (a kkksRepository) EditKkks(req model.GetDetailK3s) (model.ResEditKkks, *resty.Response, error) {
	var result model.ResEditKkks
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"id_kkks": req.ID,
	}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_KKKS_SERVICE"), "/api/edit"))

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

// TotalHolding implements HolodngRepository
func (a kkksRepository) TotalKkks() (model.Total, *resty.Response, error) {
	var result model.Total
	resp, err := a.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_KKKS_SERVICE"), "/api/total"))

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

// UpdateHolding implements HolodngRepository
func (a kkksRepository) UpdateKkks(req model.ReqUpdate) (model.ResAdd, *resty.Response, error) {
	var result model.ResAdd
	resp, err := a.client.R().
		SetBody(req).
		SetResult(&result).
		Put(fmt.Sprintf("%s%s", os.Getenv("HOST_KKKS_SERVICE"), "/api/update"))

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

// DetailHolding implements HolodngRepository
func (a kkksRepository) DetailKkks(req model.GetDetailK3s) (model.ReqDetailKkks, *resty.Response, error) {
	var result model.ReqDetailKkks
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"id_kkks": req.ID,
	}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_KKKS_SERVICE"), "/api/detail"))

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

// Add Holding
func (a kkksRepository) AddKkks(request model.ReqDetailKkks) (model.ResAdd, *resty.Response, error) {
	var result model.ResAdd
	resp, err := a.client.R().
		SetBody(&request).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_KKKS_SERVICE"), "/api/add"))

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

type KkksRepository interface {
	AddKkks(model.ReqDetailKkks) (model.ResAdd, *resty.Response, error)
	DetailKkks(model.GetDetailK3s) (model.ReqDetailKkks, *resty.Response, error)
	EditKkks(model.GetDetailK3s) (model.ResEditKkks, *resty.Response, error)
	UpdateKkks(model.ReqUpdate) (model.ResAdd, *resty.Response, error)
	TotalKkks() (model.Total, *resty.Response, error)
	ListKkks() ([]model.DataListKkks, *resty.Response, error)
	UpdateAccountKkks(model.ReqDetailKkks) (model.ResAdd, *resty.Response, error)
}

func NewKkksRepository() KkksRepository {
	return kkksRepository{
		client: resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}
