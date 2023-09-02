package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type rateRepository struct {
	client *resty.Client
}

// DetailSofr implements RateRepository.
func (r rateRepository) DetailSofr(req model.GetIdRate) (model.Libor, *resty.Response, error) {
	var result model.Libor
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"id": req.Id,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_RATE"), "/api/detail-sofr"))

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

// ListSofr implements RateRepository.
func (r rateRepository) ListSofr() ([]model.ResponseLibor, *resty.Response, error) {
	var result []model.ResponseLibor
	resp, err := r.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_RATE"), "/api/list-sofr"))

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

// UpdateSofr implements RateRepository.
func (r rateRepository) UpdateSofr(req model.RequestLibor) (model.Libor, *resty.Response, error) {
	var result model.Libor
	resp, err := r.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_RATE"), "/api/update-sofr"))

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

// UpdateLibor implements RateRepository
func (r rateRepository) UpdateLibor(req model.RequestLibor) (model.Libor, *resty.Response, error) {
	var result model.Libor
	resp, err := r.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_RATE"), "/api/update-libor"))

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

// DetailLibor implements RateRepository
func (r rateRepository) DetailLibor(req model.GetIdRate) (model.Libor, *resty.Response, error) {
	var result model.Libor
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"id": req.Id,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_RATE"), "/api/detail-libor"))

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

// ListLibor implements RateRepository
func (r rateRepository) ListLibor() ([]model.ResponseLibor, *resty.Response, error) {
	var result []model.ResponseLibor
	resp, err := r.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_RATE"), "/api/list-libor"))

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

// UpdateLps implements RateRepository
func (r rateRepository) UpdateLps(req model.RequestLps) (model.Lps, *resty.Response, error) {
	var result model.Lps
	resp, err := r.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_RATE"), "/api/update-lps"))

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

// DetailLps implements RateRepository
func (r rateRepository) DetailLps(req model.GetIdRate) (model.DetailLps, *resty.Response, error) {
	var result model.DetailLps
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"id": req.Id,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_RATE"), "/api/detail-lps"))

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

// ListLps implements RateRepository
func (r rateRepository) ListLps() ([]model.ResponseLps, *resty.Response, error) {
	var result []model.ResponseLps
	resp, err := r.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_RATE"), "/api/list-lps"))

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

func NewRateRepository() RateRepository {
	return rateRepository{
		client: resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}

type RateRepository interface {
	ListLps() ([]model.ResponseLps, *resty.Response, error)
	DetailLps(model.GetIdRate) (model.DetailLps, *resty.Response, error)
	UpdateLps(model.RequestLps) (model.Lps, *resty.Response, error)
	//libor
	ListLibor() ([]model.ResponseLibor, *resty.Response, error)
	DetailLibor(model.GetIdRate) (model.Libor, *resty.Response, error)
	UpdateLibor(model.RequestLibor) (model.Libor, *resty.Response, error)
	//Sofr
	ListSofr() ([]model.ResponseLibor, *resty.Response, error)
	DetailSofr(model.GetIdRate) (model.Libor, *resty.Response, error)
	UpdateSofr(model.RequestLibor) (model.Libor, *resty.Response, error)
}
