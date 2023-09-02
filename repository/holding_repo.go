package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type holdingRepository struct {
	client *resty.Client
}

// TotalHolding implements HolodngRepository
func (a holdingRepository) TotalHolding() (model.Total, *resty.Response, error) {
	var result model.Total
	resp, err := a.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_HOLDING_SERVICE"), "/api/total"))
		// Get("http://localhost:8080/api/total")

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
func (a holdingRepository) UpdateHolding(req model.UpdateHolding) (model.DetailHolding, *resty.Response, error) {
	var result model.DetailHolding
	resp, err := a.client.R().
		SetBody(req).
		SetResult(&result).
		Put(fmt.Sprintf("%s%s", os.Getenv("HOST_HOLDING_SERVICE"), "/api/update"))
		//Put("http://localhost:8080/api/update")

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
func (a holdingRepository) DetailHolding(req model.GetHolding) (model.DetailHolding, *resty.Response, error) {
	var result model.DetailHolding
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"id_holding": req.ID,
	}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_HOLDING_SERVICE"), "/api/detail"))
		//Get("http://localhost:8080/api/detail")

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

// ListHolding implements HolodngRepository
func (a holdingRepository) ListHolding() ([]model.ResListHoldng, *resty.Response, error) {
	var result []model.ResListHoldng
	resp, err := a.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_HOLDING_SERVICE"), "/api/list"))
		//Get("http://localhost:8080/api/list")

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
func (a holdingRepository) AddHolding(request model.RequestHolding) (model.ResponAdd, *resty.Response, error) {
	var result model.ResponAdd
	resp, err := a.client.R().
		SetBody(request).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_HOLDING_SERVICE"), "/api/add"))
		//Post("http://localhost:8080/api/add")

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

type HolodngRepository interface {
	AddHolding(model.RequestHolding) (model.ResponAdd, *resty.Response, error)
	ListHolding() ([]model.ResListHoldng, *resty.Response, error)
	DetailHolding(model.GetHolding) (model.DetailHolding, *resty.Response, error)
	UpdateHolding(model.UpdateHolding) (model.DetailHolding, *resty.Response, error)
	TotalHolding() (model.Total, *resty.Response, error)
}

func NewHoldingRepository() HolodngRepository {
	return holdingRepository{
		client: resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}
