package model

import (
	"time"

	"gorm.io/gorm"
)

type Register struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Company  string   `json:"company"`
	Role     []string `json:"role"`
	Tenant   string   `json:"tenant"`
	Password string   `json:"password"`
}

type User struct {
	gorm.Model
	Name                string      `json:"name"`
	Email               string      `json:"email"`
	Username            string      `json:"username"`
	Company             string      `json:"company"`
	Roles               interface{} `json:"roles"`
	Tenant              string      `json:"tenant"`
	Token               string      `json:"token"`
	LastChangePassword  time.Time   `json:"last_change_password"`
	LastLoginAttempt    time.Time   `json:"last_login_attempt"`
	LoginAttemptsFailed int         `json:"login_attempts_failed"`
	AccountLocked       bool        `json:"account_locked"`
	AccountDeleted      bool        `json:"account_deleted"`
	AccountReseted      bool        `json:"account_reseted" gorm:"DEFAULT:False"`
	AccountDeleteTime   time.Time   `json:"account_delete_time"`
}

type UserDetailPln struct {
	gorm.Model
	Name                string      `json:"name"`
	Email               string      `json:"email"`
	Username            string      `json:"username"`
	Company             string      `json:"company"`
	Roles               interface{} `json:"roles"`
	Tenant              string      `json:"tenant"`
	Token               string      `json:"token"`
	LastChangePassword  time.Time   `json:"last_change_password"`
	LastLoginAttempt    time.Time   `json:"last_login_attempt"`
	LoginAttemptsFailed int         `json:"login_attempts_failed"`
	AccountLocked       bool        `json:"account_locked"`
	AccountDeleted      bool        `json:"account_deleted"`
	AccountReseted      bool        `json:"account_reseted" gorm:"DEFAULT:False"`
	AccountDeleteTime   time.Time   `json:"account_delete_time"`
	UserId              string      `json:"user_id"`
	CorporateId         string      `json:"corporate_id"`
	ApiKey              string      `json:"apikey"`
}

type RequestUpdateUserAndRoles struct {
	Username string   `json:"username"`
	Company  string   `json:"company"`
	Role     []string `json:"role"`
	Tenant   string   `json:"tenant"`
}

type DetailUserPln struct {
	Username    string `json:"username"`
	UserId      string `json:"user_id"`
	CorporateId string `json:"corporate_id"`
	ApiKey      string `json:"api_key"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestAddRoleToUser struct {
	Username string   `json:"username"`
	Tenant   string   `json:"tenant"`
	Role     []string `json:"role"`
}

type RequestGetDetail struct {
	Username string `form:"username"`
}

// new request and respond user kkp

type RequestAddUserKKP struct {
	Company        string   `json:"tipe_user"`
	KodeSatker     string   `json:"kode_user"`
	Name           string   `json:"name"`
	Username       string   `json:"username"`
	Tenant         string   `json:"tenant"`
	Role           []string `json:"role"`
	AccountDeleted bool     `json:"status"`
}

type UsersRoles struct {
	Name     string `json:"name" gorm:"name"`
	Username string `json:"username" gorm:"username" `
	Company  string `json:"company"`
	Role     string `json:"role" gorm:"role" `
	Status   string `json:"status" `
}

type DetailUserKKP struct {
	Company        string `json:"company"`
	KodeSatker     string `json:"kode_user"`
	Name           string `json:"name" gorm:"name"`
	Username       string `json:"username" gorm:"username" `
	Role           string `json:"role" gorm:"role" `
	AccountDeleted bool   `json:"status"`
}

type UserDetailKKP struct {
	Name       string `json:"name" gorm:"name"`
	Username   string `json:"username" gorm:"username" `
	Company    string `json:"company"`
	KodeSatker string `json:"kode_satker"`
	Role       string `json:"role" gorm:"role" `
}
type DetailUsersKKP struct {
	Name       string `json:"name" gorm:"name"`
	Username   string `json:"username" gorm:"username" `
	Company    string `json:"company"`
	Role       string `json:"role" gorm:"role" `
	BankName   string `json:"bank_name"`
	BankCode   string `json:"bank_code"`
	KppnName   string `json:"kppn_name"`
	KppnCode   string `json:"kppn_code"`
	KodeSatker string `json:"kode_satker"`
	SatkerName string `json:"satker_name"`
}
