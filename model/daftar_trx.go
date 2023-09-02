package model

type RequestDaftarTrx struct {
	Company     string `json:"tipe_user" form:"tipe_user" validate:"patternAlphaUnderscore"`
	IdHolding   string `json:"id_holding" form:"id_holding" validate:"patternAlphaUnderscore"`
	IdKkks      string `json:"id_kkks" form:"id_kkks" validate:"patternAlphaUnderscore"`
	IdWk        string `json:"id_wk" form:"id_wk" validate:"patternAlphaUnderscore"`
	Tipe        string `json:"tipe" form:"tipe" validate:"patternAlphaUnderscore"`
	AccountName string `json:"account_name" form:"account_name" validate:"patternAlphaUnderscore"`
	AccountNo   string `json:"account_no" form:"account_no" validate:"patternAlphaUnderscore"`
	StartDate   string `json:"start_date" form:"start_date" validate:"patternAlphaUnderscore"`
	EndDate     string `json:"end_date" form:"end_date" validate:"patternAlphaUnderscore"`
}

type ResDropAccountName struct {
	AccountNo   string `json:"account_no"`
	AccountName string `json:"account_name"`
}

type ResultListTrx struct {
	WkName      string `json:"wk_name"`
	Tipe        string `json:"tipe"`
	AccountNo   string `json:"account_no"`
	Bilyet      string `json:"bilyet"`
	AccountName string `json:"account_name"`
	TranDate    string `json:"trx_date"`
	Desc        string `json:"deskripsi"`
	DrCr        string `json:"dr_cr"`
	Penerimaan  int    `json:"penerimaan_asr"`
	Pengeluaran int    `json:"pengeluaran_asr"`
	Bunga       int    `json:"bunga"`
	Pajak       int    `json:"pajak"`
	Etc         int    `json:"lain_lain"`
	Nominal     int    `json:"nominal_saldo"`
}

type NominalTrx struct {
	Nominal int `json:"nominal"`
}

type ResponseSumaryList struct {
	SaldoAwal  int `json:"saldo_awal"`
	SaldoAkhir int `json:"saldo_akhir"`
	TotalIn    int `json:"total_penerimaan"`
	TotalOut   int `json:"total_pengeluaran"`
	TotalBunga int `json:"total_bunga"`
	TotalPajak int `json:"total_pajak"`
	TotalEtc   int `json:"total_etc"`
	NettoBunga int `json:"netto_bunga"`
}

type ResponseSummaryTrx struct {
	ListTable []ResultListTrx    `json:"list_tbl"`
	Summary   ResponseSumaryList `json:"summary"`
}
