package model

import (
	"time"
)

type RequestAddUserSkk struct {
	Userlogin       string `json:"user_login" `
	Username        string `json:"username" `
	Name            string `json:"name" validate:"patternAlphaUnderscore"`
	NoTlpn          string `json:"no_tlpn" validate:"patternAlphaUnderscore"`
	Email           string `json:"email" binding:"required,email"`
	Company         string `json:"tipe_user" validate:"patternAlphaUnderscore"`
	Tenant          string `json:"tenant" validate:"patternAlphaUnderscore"`
	Password        string `json:"password"`
	Role            string `json:"role"`
	IdHolding       int    `json:"id_holding"`
	IdKkks          int    `json:"id_kkks"`
	IdWk            []int  `json:"id_wk"`
	AccountIsActive bool   `json:"status"`
}

type ResponseAddUserSkk struct {
	Name                string      `json:"name" gorm:"type:varchar(255)"`
	Email               string      `json:"email" gorm:"type:varchar(255)"`
	NoTlpn              string      `json:"no_tlpn" gorm:"type:varchar(255)"`
	Username            string      `json:"username" gorm:"type:varchar(255);unique;not null"`
	Company             string      `json:"tipe_user" gorm:"type:varchar(255)"`
	Role                interface{} `json:"roles"`
	Tenant              string      `json:"tenant" gorm:"type:varchar(255);not null"`
	Token               string      `json:"token"`
	Password            string      `json:"password" gorm:"type:varchar(255);not null"`
	LastChangePassword  time.Time   `json:"last_change_password" gorm:"autoCreateTime"`
	LastLoginAttempt    time.Time   `json:"last_login_attempt" gorm:"timestamptz;DEFAULT:null"`
	LoginAttemptsFailed int         `json:"login_attempts_failed" gorm:"DEFAULT:0"`
	AccountLocked       bool        `json:"account_locked" gorm:"DEFAULT:False"`
	AccountDeleted      bool        `json:"account_deleted" gorm:"DEFAULT:False"`
	AccountReseted      bool        `json:"account_reseted" gorm:"DEFAULT:False"`
	AccountDeleteTime   time.Time   `json:"account_delete_time" gorm:"type:timestamptz;DEFAULT:null"`
	AccountIsActive     bool        `json:"account_is_active" gorm:"DEFAULT:True"`
	IdHolding           int         `json:"id_holding" `
	IdKkks              int         `json:"id_kkks"`
	IdWk                interface{} `json:"id_wk"`
	WkName              interface{} `json:"wk_name"`
	Idle                int         `json:"idle_time"`
}

type ListUserSkk struct {
	Username        string `json:"username"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Company         string `json:"tipe_user"`
	AccountIsActive bool   `json:"account_is_active"`
}

type RequestDetailUserSkk struct {
	Userlogin string `json:"user_login" form:"user_login" validate:"patternAlphaUnderscore`
	Username  string `json:"username" form:"username"`
}
type RequestDetailUserSkkConten struct {
	Userlogin string `json:"user_login" form:"user_login" validate:"patternAlphaUnderscore`
	Username  string `json:"username" form:"username" binding:"required" validate:"patternAlphaUnderscore`
}

type DetailUser struct {
	Username        string `json:"username"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	NoTlpn          string `json:"no_tlpn"`
	AccountIsActive bool   `json:"account_is_active"`
	Company         string `json:"tipe_user"`
	Role            string `json:"role"`
	IdHolding       int    `json:"id_holding"`
	IdKkks          int    `json:"id_kkks"`
	IdWk            string `json:"id_wk"`
}

type RespondDetailUser struct {
	Username        string        `json:"username"`
	Name            string        `json:"name"`
	Email           string        `json:"email"`
	NoTlpn          string        `json:"no_tlpn"`
	AccountIsActive bool          `json:"account_is_active"`
	Company         string        `json:"tipe_user"`
	UserBni         []UserBni     `json:"user_bni"`
	UserSkk         []UserSkk     `json:"user_skk"`
	UserHolding     []UserHolding `json:"user_holding"`
	UserKkks        []UserKkks    `json:"user_kkks"`
}

// user bni
type UserBni struct {
	Role string `json:"role"`
}

// USER SKK
type TampSKK struct {
	HoldingName string `json:"holding_name"`
	KkksName    string `json:"kkks_name"`
	WkName      string `json:"wk_name"`
}

type UserSkk struct {
	SkkData []TampSKK `json:"list_skk"`
	Role    string    `json:"role"`
}

// USER kkkS

type TampKkks struct {
	IdKkks      int    `json:"id_kkks"`
	KkksName    string `json:"kkks_name"`
	HoldingName string `json:"holding_name"`
	WkName      string `json:"wk_name"`
}

type UserKkks struct {
	KkksData []TampKkks `json:"list_kkks"`
	Role     string     `json:"role"`
}

// holding
type UserHolding struct {
	TampArrHolding []TampArrHolding `json:"list_holding"`
	TampWk         []TampWk         `json:"unholding"`
	Role           string           `json:"role"`
}

type TampHolding struct {
	IdHolding   int    `json:"id_holding"`
	HoldingName string `json:"holding_name"`
	KkksName    string `json:"kkks_name"`
	WkName      string `json:"wk_name"`
}

type TampArrHolding struct {
	IdHolding   int      `json:"id_holding"`
	HoldingName string   `json:"holding_name"`
	KkksName    []string `json:"kkks_name"`
	WkName      []string `json:"wk_name"`
}

//unholding

type UserWk struct {
	TampWk []TampWk `json:"list_unholding"`
}

type TampWk struct {
	IdWk        int    `json:"id_wk"`
	WkName      string `json:"wk_name"`
	HoldingName string `json:"holding_name"`
	KkksName    string `json:"kkks_name"`
}

type RequestUpdateUserSkk struct {
	Username        string `json:"username" validate:"patternAlphaUnderscore"`
	AccountIsActive bool   `json:"status"`
	Userlogin       string `json:"user_login" `
}

type RequestGetRoleSkk struct {
	Username string `json:"username" form:"username"`
}

type UserSkkLogin struct {
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

type LoginSkk struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Detail
type UserSkkDetail struct {
	Username        string `json:"username"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	NoTlpn          string `json:"no_tlpn"`
	Company         string `json:"tipe_user"`
	Role            string `json:"role"`
	AccountIsActive bool   `json:"account_is_active"`
}

type ChangePassword struct {
	Key             string `json:"key"`
	Username        string `json:"username" binding:"required" validate:"patternAlphaUnderscore"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ResChangePassword struct {
	Username           string    `json:"username"`
	Password           string    `json:"password"`
	LastChangePassword time.Time `json:"last_change_password"`
}

type UpdatePassword struct {
	Username        string `json:"username" validate:"patternAlphaUnderscore"`
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type RespondSendOtp struct {
	Username string `json:"username"`
	NoTlpn   string `json:"no_tlpn"`
	Msg      string `json:"messege"`
}

type UpdateProfil struct {
	Username        string `json:"username" validate:"patternAlphaUnderscore"`
	Name            string `json:"name" validate:"patternAlphaUnderscore"`
	Email           string `json:"email" binding:"required,email"`
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ResponseUpdateProfil struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type ResLoginSKk struct {
	Messege string `json:"msg"`
	NoTlpn  string `json:"no_tlpn"`
	Key     string `json:"key"`
}

type PostUsername struct {
	Username string `json:"username" binding:"required" validate:"patternAlphaUnderscore"`
	Otp      string `json:"otp" binding:"required" validate:"patternAlphaUnderscore"`
	NoTlpn   string `json:"no_tlpn" binding:"required" validate:"patternAlphaUnderscore"`
}

type CheckUser struct {
	Username string `json:"username"`
}

type ResOtpForgotSkk struct {
	Username string `json:"username" binding:"required" validate:"patternAlphaUnderscore"`
	NoTlpn   string `json:"no_tlpn"`
}

type RequestParamUser struct {
	IdHolding string `json:"id_holding" form:"id_holding" validate:"patternAlphaUnderscore"`
	IdKkks    string `json:"id_kkks" form:"id_kkks" validate:"patternAlphaUnderscore"`
	Username  string `json:"username" form:"username"`
}
