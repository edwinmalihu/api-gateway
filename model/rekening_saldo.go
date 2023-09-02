package model

type RequestRekeningSaldo struct {
	Company   string `json:"tipe_user" form:"tipe_user" validate:"patternAlphaUnderscore"`
	IdHolding string `json:"id_holding" form:"id_holding" validate:"patternAlphaUnderscore"`
	IdKkks    string `json:"id_kkks" form:"id_kkks" validate:"patternAlphaUnderscore"`
	IdWk      string `json:"id_wk" form:"id_wk" validate:"patternAlphaUnderscore"`
	Tipe      string `json:"tipe" form:"tipe" validate:"patternAlphaUnderscore"`
	Tanggal   string `json:"tanggal" form:"tanggal" validate:"patternAlphaUnderscore"`
}

type ResCardRekeningSaldo struct {
	Summary       int `json:"summary"`
	TotalRekening int `json:"total_rekening"`
}

type ResDropdownWK struct {
	IdWk   int    `json:"id_wk"`
	WkName string `json:"wk_name"`
}

type ResultListTable struct {
	WkName        string `json:"wk_name"`
	KkksName      string `json:"kkks_name"`
	Tipe          string `json:"tipe"`
	AccountNo     string `json:"account_no"`
	Bilyet        string `json:"bilyet"`
	AccountName   string `json:"account_name"`
	TglBukaRek    string `json:"tgl_buka_rek"`
	TglJatuhTempo string `json:"tgl_jatuh_temp"`
	StatusBunga   string `json:"status_bunga"`
	Cabang        string `json:"cabang"`
	Rate          string `json:"rate"`
	SukuBunga     string `json:"suku_bunga"`
	Nominal       int    `json:"nominal_saldo"`
}

type ResponseUrl struct {
	Url string `json:"url"`
}
