package model

// Request
type ReqAddWk struct {
	Username  string `json:"username" validate:"patternAlphaUnderscore"`
	WkName    string `json:"wk_name" validate:"patternAlphaUnderscore"`
	IdHolding string `json:"id_holding" validate:"patternAlphaUnderscore"`
	IdKkks    string `json:"id_kkks" validate:"patternAlphaUnderscore"`
}

type ReqGetWk struct {
	IdWk string `json:"id_wk" form:"id_wk" validate:"patternAlphaUnderscore"`
}

type ReqEditWK struct {
	Username  string `json:"username" validate:"patternAlphaUnderscore"`
	IdWk      string `json:"id_wk" form:"id_wk" validate:"patternAlphaUnderscore"`
	WkName    string `json:"wk_name" form:"wk_name" validate:"patternAlphaUnderscore"`
	IdHolding string `json:"id_holding" form:"id_holding" validate:"patternAlphaUnderscore"`
	IdKkks    string `json:"id_kkks" form:"id_kkks" validate:"patternAlphaUnderscore"`
}

// Response
type ResSuccess struct {
	Messege  string      `json:"messege"`
	Response interface{} `json:"response"`
}

type ResRelasiWk struct {
	IdWk        string `json:"id_wk"`
	WkName      string `json:"wk_name"`
	IdHolding   string `json:"id_holding"`
	HoldingName string `json:"holding_name"`
	IdKkks      string `json:"id_kkks"`
	KkksName    string `json:"kkks_name"`
}

type ResListWK struct {
	Nomor       int    `json:"nomor"`
	IdWk        string `json:"id_wk"`
	WkName      string `json:"wk_name"`
	IdHolding   string `json:"id_holding"`
	HoldingName string `json:"holding_name"`
	IdKkks      string `json:"id_kkks"`
	KkksName    string `json:"kkks_name"`
}

type ResListAccountWkDetail struct {
	AccountNo            string `json:"account_no"`
	AccountName          string `json:"account_name"`
	Tipe                 string `json:"tipe"`
	Ccy                  string `json:"ccy"`
	Cabang               string `json:"cabang"`
	Bilyet               string `json:"bilyet"`
	TanggalBukaRekening  string `json:"tgl_buka_rekening"`
	TanggalTutupRekening string `json:"tgl_tutup_rekening"`
	IsActive             bool   `json:"is_active"`
}

type ResDetailWk struct {
	IdWk        string                   `json:"id_wk"`
	WKName      string                   `json:"wk_name"`
	KkksName    string                   `json:"kkks_name"`
	ListAccount []ResListAccountWkDetail `json:"list_account"`
}

type ListDropdownHolidng struct {
	ID          string `json:"id_holding"`
	HoldingName string `json:"holding_name"`
}

type ListDropdownK3S struct {
	ID       string `json:"id_kkks"`
	KkksName string `json:"kkks_name"`
}

type ReqParam struct {
	IdHolding string `json:"id_holding" form:"id_holding" validate:"patternAlphaUnderscore`
	IdKkks    string `json:"id_kkks" form:"id_kkks" validate:"patternAlphaUnderscore`
	IdWk      string `json:"id_wk" form:"id_wk" validate:"patternAlphaUnderscore`
}
