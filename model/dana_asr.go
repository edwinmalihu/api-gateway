package model

type RequsestAsr struct {
	Company   string `json:"tipe_user" form:"tipe_user" binding:"required" validate:"patternAlphaUnderscore"`
	IdHolding string `json:"id_holding" form:"id_holding" validate:"patternAlphaUnderscore"`
	IdKkks    string `json:"id_kkks" form:"id_kkks" validate:"patternAlphaUnderscore"`
	IdWk      string `json:"id_wk" form:"id_wk" validate:"patternAlphaUnderscore"`
}

type CardDashboard struct {
	Total    int `json:"total_saldo"`
	Giro     int `json:"total_giro"`
	Deposito int `json:"total_deposito"`
	KKKS     int `json:"total_kkks"`
}

type GrafikTotalDana struct {
	Databyyear  []interface{}            `json:"grafik_7_tahun"`
	Databymonth []GrafikTotalDanaByMonth `json:"grafik_month"`
}

type GrafikTotalDanaByMonth struct {
	Bulan   string `json:"bulan"`
	Nominal int    `json:"nominal"`
}

type RespNominalTahun struct {
	Tahun   []string `json:"tahun"`
	Nominal []int    `json:"nominal"`
}

type DataGrafikGiroDepo struct {
	Bulan    []string `json:"bulan"`
	Giro     []int    `json:"giro"`
	Deposito []int    `json:"deposito"`
}
type ResPenerimaanPengeluaran struct {
	Bulan       []string `json:"bulan"`
	Penerimaan  []int    `json:"penerimaan"`
	Pengeluaran []int    `json:"pengeluaran"`
}
type ResTrenBunga struct {
	Bulan []string `json:"bulan"`
	Bunga []int    `json:"bunga_netto"`
}

// Respond New

type RespondRelasiHolding struct {
	IdKkks int `json:"id_kkks"`
	IdWk   int `json:"id_wk"`
}
type TotalGrafik struct {
	Nominal int    `json:"nominal"`
	Tahun   string `json:"tahun"`
}
