package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type userSkkRepository struct {
	client *resty.Client
}

// SendOtpForgotHard implements UserSkkRepository
func (u userSkkRepository) SendOtpForgotHard(req model.CheckUser) (model.ResOtpForgotSkk, *resty.Response, error) {
	var result model.ResOtpForgotSkk
	resp, err := u.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/send-otp-forgot"))

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

// SendOtpForgotSoa implements UserSkkRepository
func (u userSkkRepository) SendOtpForgotSoa(req model.CheckUser) (model.ResOtpForgotSkk, *resty.Response, error) {
	var result model.ResOtpForgotSkk
	resp, err := u.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/soa-otp-forgot"))

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

// ValidasiOtpForgotHard implements UserSkkRepository
func (u userSkkRepository) ValidasiOtpForgotHard(req model.PostUsername) (model.ResLoginSKk, *resty.Response, error) {
	var result model.ResLoginSKk
	resp, err := u.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/validasi-otp"))

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

// ValidasiOtpForgotSoa implements UserSkkRepository
func (u userSkkRepository) ValidasiOtpForgotSoa(req model.PostUsername) (model.ResLoginSKk, *resty.Response, error) {
	var result model.ResLoginSKk
	resp, err := u.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/validasi-otp-soa"))

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

// NewLoginUserSkk implements UserSkkRepository
func (u userSkkRepository) NewLoginUserSkk(req model.LoginSkk) (model.ResLoginSKk, *resty.Response, error) {
	var result model.ResLoginSKk
	resp, err := u.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/login"))

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

// NewSendOtp implements UserSkkRepository
func (u userSkkRepository) NewSendOtp(req model.PostUsername) (model.ResponseAddUserSkk, *resty.Response, error) {
	var result model.ResponseAddUserSkk
	resp, err := u.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/otp"))

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

// UpdateProfil implements UserSkkRepository
func (u userSkkRepository) UpdateProfil(req model.UpdateProfil) (model.ResponseUpdateProfil, *resty.Response, error) {
	var result model.ResponseUpdateProfil
	resp, err := u.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/update-profil"))
		//Post("http://localhost:8082/api/update-profil")

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

// SendOtp implements UserSkkRepository
func (u userSkkRepository) SendOtp(req model.PostUsername) (model.ResponseAddUserSkk, *resty.Response, error) {
	var result model.ResponseAddUserSkk
	resp, err := u.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/send-otp"))

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

// UpdatePasswordSkk implements UserSkkRepository
func (u userSkkRepository) UpdatePasswordSkk(req model.UpdatePassword) (model.ResChangePassword, *resty.Response, error) {
	var result model.ResChangePassword
	resp, err := u.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/update-pass"))

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

// ChangePasswordUserSkk implements UserSkkRepository
func (u userSkkRepository) ChangePasswordUserSkk(req model.ChangePassword) (model.ResChangePassword, *resty.Response, error) {
	var result model.ResChangePassword
	resp, err := u.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/change-pass"))

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

// LoginUserSkk implements UserSkkRepository
func (u userSkkRepository) LoginUserSkk(req model.LoginSkk) (model.ResponseAddUserSkk, *resty.Response, error) {
	var result model.ResponseAddUserSkk
	resp, err := u.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/signin"))

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

// UpdateUserSkk implements UserSkkRepository
func (u userSkkRepository) UpdateUserSkk(req model.RequestUpdateUserSkk) (model.RequestUpdateUserSkk, *resty.Response, error) {
	var result model.RequestUpdateUserSkk
	resp, err := u.client.R().
		SetBody(req).
		SetResult(&result).
		Put(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/update"))
		//Put("http://localhost:8082/api/update")

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

// DetailUserSkk implements UserSkkRepository
func (u userSkkRepository) DetailUserSkk(req model.RequestDetailUserSkkConten) (model.UserSkkDetail, *resty.Response, error) {
	var result model.UserSkkDetail
	resp, err := u.client.R().SetQueryParams(map[string]string{
		"username":   req.Username,
		"user_login": req.Userlogin,
	}).SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/detail"))
		//Get("http://localhost:8082/api/detail")

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

// DetailUserSkk implements UserSkkRepository
func (u userSkkRepository) DetailuserskkProfil(req model.RequestDetailUserSkk) (model.UserSkkDetail, *resty.Response, error) {
	var result model.UserSkkDetail
	resp, err := u.client.R().SetQueryParams(map[string]string{
		"username": req.Username,
	}).SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/detail"))
		//Get("http://localhost:8082/api/detail")

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

// ListUserSkk implements UserSkkRepository
func (u userSkkRepository) ListUserSkk(req model.RequestParamUser) ([]model.ListUserSkk, *resty.Response, error) {
	var result []model.ListUserSkk
	resp, err := u.client.R().SetQueryParams(map[string]string{
		"id_holding": req.IdHolding,
		"id_kkks":    req.IdKkks,
		"username":   req.Username,
	}).
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/list"))
		//Get("http://localhost:8082/api/list")

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

// AddUser implements UserSkkRepository
func (u userSkkRepository) AddUserSkk(req model.RequestAddUserSkk) (model.ResponseAddUserSkk, *resty.Response, error) {
	var result model.ResponseAddUserSkk
	resp, err := u.client.R().
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("SERVICE_HOST_USER"), "/api/add"))

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

// TypeUserRole implements PolicySkkRepository
func (u userSkkRepository) TypeUserRole() (model.ResResult, *resty.Response, error) {
	var result model.ResResult
	resp, err := u.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_SERVICE_POLICY"), "/api/role-group"))

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

// UserRepository : represent the user's repository contract
type UserSkkRepository interface {
	AddUserSkk(model.RequestAddUserSkk) (model.ResponseAddUserSkk, *resty.Response, error)
	ListUserSkk(model.RequestParamUser) ([]model.ListUserSkk, *resty.Response, error)
	DetailUserSkk(model.RequestDetailUserSkkConten) (model.UserSkkDetail, *resty.Response, error)
	DetailuserskkProfil(model.RequestDetailUserSkk) (model.UserSkkDetail, *resty.Response, error)
	UpdateUserSkk(model.RequestUpdateUserSkk) (model.RequestUpdateUserSkk, *resty.Response, error)
	LoginUserSkk(model.LoginSkk) (model.ResponseAddUserSkk, *resty.Response, error)
	TypeUserRole() (model.ResResult, *resty.Response, error)
	ChangePasswordUserSkk(model.ChangePassword) (model.ResChangePassword, *resty.Response, error)
	UpdatePasswordSkk(model.UpdatePassword) (model.ResChangePassword, *resty.Response, error)
	SendOtp(model.PostUsername) (model.ResponseAddUserSkk, *resty.Response, error)
	UpdateProfil(model.UpdateProfil) (model.ResponseUpdateProfil, *resty.Response, error)

	// new Login
	NewLoginUserSkk(model.LoginSkk) (model.ResLoginSKk, *resty.Response, error)
	NewSendOtp(model.PostUsername) (model.ResponseAddUserSkk, *resty.Response, error)

	// new forgot
	SendOtpForgotHard(model.CheckUser) (model.ResOtpForgotSkk, *resty.Response, error)
	ValidasiOtpForgotHard(model.PostUsername) (model.ResLoginSKk, *resty.Response, error)
	SendOtpForgotSoa(model.CheckUser) (model.ResOtpForgotSkk, *resty.Response, error)
	ValidasiOtpForgotSoa(model.PostUsername) (model.ResLoginSKk, *resty.Response, error)
}

// NewUserRepository -> returns new user repository
func NewUserSkkRepository() UserSkkRepository {
	return userSkkRepository{
		client: resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}
