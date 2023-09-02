package model

type GetIdRate struct {
	Id string `json:"id" form:"id" binding:"required" validate:"patternAlphaUnderscore"`
}

type RequestLps struct {
	Id            int    `json:"id"`
	Periode       string `json:"periode" binding:"required"`
	BankUmumIdr   string `json:"bank_umum_idr" binding:"required" `
	BankUmumValas string `json:"bank_umum_valas" binding:"required" `
	Bpr           string `json:"bpr" binding:"required" `
}
type RequestLibor struct {
	Id      int    `json:"id" binding:"required"`
	Periode string `json:"periode" binding:"required"`
	Rate    string `json:"rate" binding:"required"`
}

type Lps struct {
	Id            int    `json:"id"`
	Periode       string `json:"periode"`
	BankUmumIdr   string `json:"bank_umum_idr" `
	BankUmumValas string `json:"bank_umum_valas"`
	Bpr           string `json:"bpr"  `
}
type Libor struct {
	Id      int    `json:"id"`
	Periode string `json:"periode" `
	Rate    string `json:"rate" `
}

type ResponseLps struct {
	Id            int    `json:"id"`
	Periode       string `json:"periode"`
	BankUmumIdr   string `json:"bank_umum_idr"`
	BankUmumValas string `json:"bank_umum_valas"`
	Bpr           string `json:"bpr" `
}

type ResponseLibor struct {
	Id      int    `json:"id"`
	Periode string `json:"periode"`
	Rate    string `json:"rate"`
}

type DetailLps struct {
	Id            int    `json:"id"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	BankUmumIdr   string `json:"bank_umum_idr"`
	BankUmumValas string `json:"bank_umum_valas"`
	Bpr           string `json:"bpr" `
}
