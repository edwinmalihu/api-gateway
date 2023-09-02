package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type userKKPRepository struct {
	client *resty.Client
}

type UserKKPRepository interface {
	AddUserKKP(model.RequestAddUserKKP) (model.User, *resty.Response, error)
	GetUserKKPOnly() ([]model.UsersRoles, *resty.Response, error)
	DetailUserKKP(model.RequestGetDetail) (model.DetailUserKKP, *resty.Response, error)
	DetailsUserKKP(model.RequestGetDetail) (model.UserDetailKKP, *resty.Response, error)
	UpdateUserKKP(model.RequestAddUserKKP) (model.User, *resty.Response, error)
	ResetPassKKP(model.RequestGetDetail) (model.User, *resty.Response, error)
	ChangePasswordKKP(model.ChangePassword) (model.User, *resty.Response, error)
	DeleteUserKKP(model.RequestGetDetail) (string, *resty.Response, error)
	HardDeleteUserKKP(model.RequestGetDetail) (string, *resty.Response, error)
}

func NewUserKKPRepository() UserKKPRepository {
	return userKKPRepository{
		client: resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}

// AddUserKKP implements UserKKPRepository
func (u userKKPRepository) AddUserKKP(request model.RequestAddUserKKP) (model.User, *resty.Response, error) {
	var result model.User
	resp, err := u.client.R().
		SetBody(request).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/add-user-kkp"))
		//Post("http://localhost:8082/api/detail")

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

// DetailUserKKP implements UserKKPRepository
func (u userKKPRepository) DetailUserKKP(username model.RequestGetDetail) (model.DetailUserKKP, *resty.Response, error) {
	var user model.DetailUserKKP
	resp, err := u.client.R().SetQueryParams(map[string]string{
		"username": username.Username,
	}).SetResult(&user).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/detail-user-kkp"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	return user, resp, err
}

// GetUserKKPOnly implements UserKKPRepository
func (u userKKPRepository) GetUserKKPOnly() ([]model.UsersRoles, *resty.Response, error) {
	var result []model.UsersRoles
	resp, err := u.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/list-user-kkp"))

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

func (u userKKPRepository) UpdateUserKKP(request model.RequestAddUserKKP) (model.User, *resty.Response, error) {
	var result model.User
	resp, err := u.client.R().
		SetBody(request).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/update-user-kkp"))

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

func (u userKKPRepository) ResetPassKKP(request model.RequestGetDetail) (model.User, *resty.Response, error) {
	var user model.User
	resp, err := u.client.R().
		SetBody(request).
		SetResult(&user).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/reset-password"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return user, resp, err
}

func (u userKKPRepository) ChangePasswordKKP(change model.ChangePassword) (model.User, *resty.Response, error) {
	var user model.User
	resp, err := u.client.R().
		SetBody(change).
		SetResult(&user).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/change-password"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return user, resp, err
}

// Soft Delete
func (u userKKPRepository) DeleteUserKKP(username model.RequestGetDetail) (string, *resty.Response, error) {
	var user string
	resp, err := u.client.R().SetQueryParams(map[string]string{
		"username": username.Username,
	}).SetResult(&user).
		Delete(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/delete-user-kkp"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	return user, resp, err
}

// Hard Delete
func (u userKKPRepository) HardDeleteUserKKP(username model.RequestGetDetail) (string, *resty.Response, error) {
	var user string
	resp, err := u.client.R().SetQueryParams(map[string]string{
		"username": username.Username,
	}).SetResult(&user).
		Delete(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/hapus-user-kkp"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	return user, resp, err
}

// DetailUserKKP implements UserKKPRepository
func (u userKKPRepository) DetailsUserKKP(username model.RequestGetDetail) (model.UserDetailKKP, *resty.Response, error) {
	var user model.UserDetailKKP
	resp, err := u.client.R().SetQueryParams(map[string]string{
		"username": username.Username,
	}).SetResult(&user).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/details-users-kkp"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	return user, resp, err
}
