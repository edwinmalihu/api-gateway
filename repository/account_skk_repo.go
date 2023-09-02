package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type accountSkkRepository struct {
	client *resty.Client
}

// UpdateAccount implements AccountSkkRepository
func (a accountSkkRepository) UpdateAccount(req model.ReqUpdateAccount) (model.ResAdd, *resty.Response, error) {
	var result model.ResAdd
	resp, err := a.client.R().
		SetBody(&req).
		SetResult(&result).
		Put(fmt.Sprintf("%s%s", os.Getenv("HOST_ACCOUNT_SERVICE"), "/api/update"))

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

// DetailAccount implements AccountSkkRepository
func (a accountSkkRepository) DetailAccount(req model.ReqDetail) (model.ResDetailAccount, *resty.Response, error) {
	var result model.ResDetailAccount
	resp, err := a.client.R().
		SetQueryParams(map[string]string{
			"account_no": req.AccountNo,
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_ACCOUNT_SERVICE"), "/api/detail"))

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

// AddAccount implements AccountSkkRepository
func (a accountSkkRepository) AddAccount(req model.RequestAdd) (model.ResAdd, *resty.Response, error) {
	var result model.ResAdd
	resp, err := a.client.R().
		SetBody(&req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_ACCOUNT_SERVICE"), "/api/add"))

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

// List implements AccountSkkRepository
func (a accountSkkRepository) List() ([]model.ResListAccount, *resty.Response, error) {
	var result []model.ResListAccount
	resp, err := a.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_ACCOUNT_SERVICE"), "/api/list"))

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

type AccountSkkRepository interface {
	List() ([]model.ResListAccount, *resty.Response, error)
	AddAccount(model.RequestAdd) (model.ResAdd, *resty.Response, error)
	DetailAccount(model.ReqDetail) (model.ResDetailAccount, *resty.Response, error)
	UpdateAccount(model.ReqUpdateAccount) (model.ResAdd, *resty.Response, error)
}

func NewAccountSkkRepository() AccountSkkRepository {
	return accountSkkRepository{
		client: resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}
