package repository

import (
	"auth-services/model"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type userRepository struct {
	client *resty.Client
}

// UserRepository : represent the user's repository contract
type UserRepository interface {
	AddUserAdmin(model.Register) (model.User, *resty.Response, error)
	AddUser(model.Register) (model.User, *resty.Response, error)
	AddRoleToUser(model.RequestAddRoleToUser) (model.User, *resty.Response, error)
	DeleteRoleFromUser(model.RequestAddRoleToUser) (model.User, *resty.Response, error)
	DeleteRolesFromUser(model.RequestUpdateUserAndRoles) (model.User, *resty.Response, error)
	SignIn(model.Login) (model.User, *resty.Response, error)
	GetUser(model.RequestGetDetail) (model.User, *resty.Response, error)
	GetAllUser() ([]model.User, *resty.Response, error)
	UpdateUser(model.RequestUpdateUserAndRoles) (model.User, *resty.Response, error)
	DeleteUser(model.RequestGetDetail) (model.User, *resty.Response, error)
	ChangePassword(model.ChangePassword) (model.User, *resty.Response, error)
	ResetPassword(model.RequestGetDetail) (model.User, *resty.Response, error)
	DetailUserPln(model.RequestGetDetail) (model.DetailUserPln, *resty.Response, error)
}

// NewUserRepository -> returns new user repository
func NewUserRepository() UserRepository {
	return userRepository{
		client: resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}

func (u userRepository) GetUser(username model.RequestGetDetail) (model.User, *resty.Response, error) {
	var result model.User
	resp, err := u.client.R().SetQueryParams(map[string]string{
		"username": username.Username,
	}).SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/detail"))

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

func (u userRepository) DetailUserPln(username model.RequestGetDetail) (model.DetailUserPln, *resty.Response, error) {
	var result model.DetailUserPln
	resp, err := u.client.R().SetQueryParams(map[string]string{
		"username": username.Username,
	}).SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/detail-user-pln"))

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

func (u userRepository) GetAllUser() ([]model.User, *resty.Response, error) {
	var result []model.User
	resp, err := u.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/all-users"))

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

func (u userRepository) AddUserAdmin(register model.Register) (model.User, *resty.Response, error) {
	var result model.User
	resp, err := u.client.R().
		SetBody(register).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/admin"))

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

func (u userRepository) AddUser(register model.Register) (model.User, *resty.Response, error) {
	var result model.User
	resp, err := u.client.R().
		SetBody(register).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/register"))

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

func (u userRepository) AddRoleToUser(role model.RequestAddRoleToUser) (model.User, *resty.Response, error) {
	var result model.User
	resp, err := u.client.R().
		SetBody(role).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/add-role-to-user"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	fmt.Println("TEST REQUEST ADD ROLE TO USER")

	return result, resp, err
}

func (u userRepository) DeleteRoleFromUser(role model.RequestAddRoleToUser) (model.User, *resty.Response, error) {
	var result model.User
	resp, err := u.client.R().
		SetBody(role).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/delete-role-from-user"))

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	fmt.Println("TEST REQUEST ADD ROLE TO USER")

	return result, resp, err
}

func (u userRepository) DeleteRolesFromUser(role model.RequestUpdateUserAndRoles) (model.User, *resty.Response, error) {
	var result model.User
	resp, err := u.client.R().
		SetBody(role).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/delete-allrole-from-user"))

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

func (u userRepository) SignIn(login model.Login) (model.User, *resty.Response, error) {
	var result model.User
	resp, err := u.client.R().
		SetBody(login).
		SetResult(&result).
		Post(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/signin"))

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

func (u userRepository) UpdateUser(request model.RequestUpdateUserAndRoles) (model.User, *resty.Response, error) {
	var result model.User
	resp, err := u.client.R().
		SetBody(request).
		SetResult(&result).
		Put(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/update-user"))

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

func (u userRepository) DeleteUser(username model.RequestGetDetail) (model.User, *resty.Response, error) {
	var user model.User
	resp, err := u.client.R().SetQueryParams(map[string]string{
		"username": username.Username,
	}).SetResult(&user).
		Delete(fmt.Sprintf("%s%s", os.Getenv("HOST_USER_SERVICE"), "/api/delete-user"))

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

func (u userRepository) ChangePassword(change model.ChangePassword) (model.User, *resty.Response, error) {
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

func (u userRepository) ResetPassword(request model.RequestGetDetail) (model.User, *resty.Response, error) {
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
