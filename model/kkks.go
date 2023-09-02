package model

import "time"

// request
type ReqAddK3s struct {
	KkksName string `json:"kkks_name"`
}

type GetDetailK3s struct {
	ID string `json:"id_kkks" form:"id_kkks" validate:"patternAlphaUnderscore"`
}

type ReqUpdateK3s struct {
	ID string `json:"id_kkks" form:"id_kkks"`
}

type ReqUpdate struct {
	ID       string `json:"id_kkks" binding:"required" validate:"patternAlphaUnderscore"`
	KkksName string `json:"kkks_name" binding:"required" validate:"patternAlphaUnderscore"`
}

type ReqDetailKkks struct {
	IdKkks      string           `json:"id_kkks" validate:"patternAlphaUnderscore"`
	KkksName    string           `json:"kkks_name" binding:"required" validate:"patternAlphaUnderscore"`
	ListAccount []ReqListAccount `json:"list_account"`
}

type ReqListAccount struct {
	IdAccount            string `json:"id_account"`
	AccountNo            string `json:"account_no" binding:"required" validate:"patternAlphaUnderscore" `
	Cabang               string `json:"cabang" binding:"required" validate:"patternAlphaUnderscore"`
	TanggalBukaRekening  string `json:"tgl_buka_rekening" binding:"required" validate:"patternAlphaUnderscore"`
	TanggalTutupRekening string `json:"tgl_tutup_rekening" validate:"patternAlphaUnderscore"`
	Bilyet               string `json:"bilyet" validate:"patternAlphaUnderscore" `
	AccountName          string `json:"account_name" binding:"required" validate:"patternAlphaUnderscore"`
	Tipe                 string `json:"tipe" validate:"patternAlphaUnderscore"`
	Ccy                  string `json:"ccy" validate:"patternAlphaUnderscore"`
	Rate                 string `json:"rate" binding:"required" validate:"patternAlphaUnderscore"`
	Tenor                uint   `json:"tenor" `
	IsActive             bool   `json:"is_active"`
}

// response

type ContainerResListKkks struct {
	Nomor       int    `json:"nomor"`
	IdKkks      int    `json:"id_kkks"`
	KkksName    string `json:"kkks_name"`
	WkName      string `json:"wk_name"`
	HoldingName string `json:"holding_name"`
}

type ResListKkks struct {
	Nomor       int    `json:"nomor"`
	IdKkks      string `json:"id_kkks"`
	KkksName    string `json:"kkks_name"`
	WkName      string `json:"wk_name"`
	HoldingName string `json:"holding_name"`
}

type ContainerDetailKkks struct {
	IdKkks               int       `json:"id_kkks"`
	KkksName             string    `json:"kkks_name"`
	IdAccount            int       `json:"id_account"`
	AccountNo            string    `json:"account_no"`
	AccountName          string    `json:"account_name"`
	Tipe                 string    `json:"tipe"`
	Ccy                  string    `json:"ccy"`
	Cabang               string    `json:"cabang"`
	Bilyet               string    `json:"bilyet"`
	TanggalBukaRekening  time.Time `json:"tgl_buka_rekening"`
	TanggalTutupRekening time.Time `json:"tgl_tutup_rekening" gorm:"tanggal_tutup_rekening"`
	IsActive             bool      `json:"is_active"`
}
type ResDetailKkks struct {
	IdKkks      string           `json:"id_kkks" gorm:"id_kkks"`
	KkksName    string           `json:"kkks_name"`
	ListAccount []ResListAccount `json:"list_account"`
}

type ResListAccount struct {
	Id                   string `json:"id_account"`
	AccountNo            string `json:"account_no"`
	AccountName          string `json:"account_name"`
	Tipe                 string `json:"tipe"`
	Ccy                  string `json:"ccy"`
	Cabang               string `json:"cabang"`
	Bilyet               string `json:"bilyet"`
	TanggalBukaRekening  string `json:"tgl_buka_rekening"`
	TanggalTutupRekening string `json:"tgl_tutup_rekening" gorm:"tanggal_tutup_rekening"`
	Rate                 string `json:"rate" `
	Tenor                uint   `json:"tenor" `
	IsActive             bool   `json:"is_active"`
}

type ResEditKkks struct {
	ID       string `json:"id_kkks"`
	KkksName string `json:"kkks_name"`
}

type ResAdd struct {
	Messege  string      `json:"messege"`
	Response interface{} `json:"response"`
}

type DataListKkks struct {
	Nomor    int    `json:"nomor"`
	IdKkks   string `json:"id_kkks"`
	KkksName string `json:"kkks_name"`
}
