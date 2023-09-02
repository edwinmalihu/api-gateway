package model

// request
type RequestAdd struct {
	Id                   string `json:"id_kkks"`
	AccountNo            string `json:"account_no" `
	Cabang               string `json:"cabang" `
	TanggalBukaRekening  string `json:"tgl_buka_rekening"`
	TanggalTutupRekening string `json:"tgl_tutup_rekening"`
	Bilyet               string `json:"bilyet" `
	AccountName          string `json:"account_name" `
	Tipe                 string `json:"tipe" `
	Ccy                  string `json:"ccy" `
	Rate                 string `json:"rate" `
	Tenor                uint   `json:"tenor" `
}

type ReqDetail struct {
	AccountNo string `json:"account_no" form:"account_no"`
}

type ReqUpdateAccount struct {
	AccountNo            string `json:"account_no"`
	Cabang               string `json:"cabang"`
	TanggalBukaRekening  string `json:"tgl_buka_rekening"`
	TanggalTutupRekening string `json:"tgl_tutup_rekening" `
	Bilyet               string `json:"bilyet" `
	AccountName          string `json:"account_name"`
	Tipe                 string `json:"tipe" `
	Ccy                  string `json:"ccy"  `
	Rate                 string `json:"rate" `
	Tenor                uint   `json:"tenor" `
	IsActive             bool   `json:"is_active" `
}

//response

type ResDetailAccount struct {
	ID                   string `json:"id_account"`
	AccountNo            string `json:"account_no"`
	Cabang               string `json:"cabang"`
	TanggalBukaRekening  string `json:"tgl_buka_rekening"`
	TanggalTutupRekening string `json:"tgl_tutup_rekening" `
	Bilyet               string `json:"bilyet" `
	AccountName          string `json:"account_name"`
	Tipe                 string `json:"tipe" `
	Ccy                  string `json:"ccy"  `
	Rate                 string `json:"rate" `
	Tenor                uint   `json:"tenor" `
	IsActive             bool   `json:"is_active" `
}
