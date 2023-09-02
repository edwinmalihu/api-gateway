package model

type RequestHolding struct {
	HoldingName string `json:"holding_name" binding:"required" validate:"patternAlphaUnderscore"`
}

type GetHolding struct {
	ID string `json:"id_holding" form:"id_holding" validate:"patternAlphaUnderscore"`
}

type UpdateHolding struct {
	ID          string `json:"id_holding" form:"id_holding" validate:"patternAlphaUnderscore"`
	HoldingName string `json:"holding_name" form:"holding_name" binding:"required" validate:"patternAlphaUnderscore"`
}

type ResListHoldng struct {
	No          int    `json:"nomor"`
	ID          int    `json:"id_holding"`
	HoldingName string `json:"holding_name"`
}
type DetailHolding struct {
	ID          string `json:"id_holding"`
	HoldingName string `json:"holding_name"`
}

type ResponAdd struct {
	Messege string      `json:"messege"`
	Respond interface{} `json:"respond"`
}

type Total struct {
	Total int `json:"total"`
}
