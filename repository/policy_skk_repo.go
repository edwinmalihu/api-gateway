package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type policySkkRepository struct {
	client *resty.Client
}

type PolicySkkRepository interface {
	AddRoleSkk(model.RequestData) (string, *resty.Response, error)
	UpdateRoleSkk(model.RequestData, string) (string, *resty.Response, error)
	DeleteRoleSkk(string) (string, *resty.Response, error)
	ListRoleSkk() ([]model.ResultListSkk, *resty.Response, error)
	DetailRoleSkk(model.DetailtGetRole) (model.RequestData, *resty.Response, error)
	GetRoleSkk(model.RequestGetRoleSkk) (model.GetRoleByUsername, *resty.Response, error)
}

func NewPolicySkkRepository() PolicySkkRepository {
	return policySkkRepository{
		client: resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}

// AddRoleName implements policySkkRepository
func (r policySkkRepository) AddRoleSkk(role model.RequestData) (string, *resty.Response, error) {
	var result string
	resp, err := r.client.R().
		SetBody(role).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_SERVICE_POLICY"), "/api/add"))

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

func (r policySkkRepository) UpdateRoleSkk(role model.RequestData, rolename string) (string, *resty.Response, error) {
	var result string
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"role": rolename,
		}).
		SetBody(role).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_SERVICE_POLICY"), "/api/update"))

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

func (r policySkkRepository) DeleteRoleSkk(rolename string) (string, *resty.Response, error) {
	var result string
	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"role": rolename,
		}).
		SetResult(&result).
		Delete(fmt.Sprintf("%s%s", os.Getenv("HOST_SERVICE_POLICY"), "/api/delete"))

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

// GetAllRoleName implements policySkkRepository
func (r policySkkRepository) ListRoleSkk() ([]model.ResultListSkk, *resty.Response, error) {
	var result []model.ResultListSkk
	resp, err := r.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_SERVICE_POLICY"), "/api/list"))

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

// GetDetailRoleName implements policySkkRepository
func (r policySkkRepository) DetailRoleSkk(req model.DetailtGetRole) (model.RequestData, *resty.Response, error) {
	var role model.RequestData
	resp, err := r.client.R().SetQueryParams(map[string]string{
		"role": req.Role,
	}).SetResult(&role).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_SERVICE_POLICY"), "/api/detail"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	return role, resp, err
}

// GetRoleSkk implements PolicySkkRepository
func (r policySkkRepository) GetRoleSkk(req model.RequestGetRoleSkk) (model.GetRoleByUsername, *resty.Response, error) {
	var role model.GetRoleByUsername
	resp, err := r.client.R().SetQueryParams(map[string]string{
		"username": req.Username,
	}).SetResult(&role).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_SERVICE_POLICY"), "/api/user-role"))
		//Get("http://localhost:8083/api/user-role")

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	return role, resp, err
}
