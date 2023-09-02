package model

type ReqRoleSkk struct {
	Role     string   `json:"role" bind:"required"`
	Tenant   string   `json:"tenant"`
	IsActive bool     `json:"status"`
	Object   []Object `json:"object"`
}

// List SKK
type RoleSkk struct {
	Roles []DataRoleSKk `json:"roles" `
}

type DataRoleSKk struct {
	Role   string `json:"role" `
	Name   string `json:"name"`
	Status bool   `json:"is_active"`
}

type StatusPrivalegeSkk struct {
	Object string `json:"object"`
	Status bool   `json:"is_active"`
}

type ResultListSkk struct {
	Role        string        `json:"role" `
	ListRoleSkk []ListRoleSkk `json:"akses"`
}

type ListRoleSkk struct {
	Name   string `json:"name"`
	Status bool   `json:"is_active"`
}

// RESPOND ROLE GROUP
type RoleBni struct {
	BniRole     Group `json:"user_bni"`
	SkkRole     Group `json:"user_skk"`
	HoldingRole Group `json:"user_holding"`
	KkksRole    Group `json:"user_kkks"`
}
type UserHoldingBNI struct {
	BniRole Group `json:"user_bni"`
}

type Group struct {
	Role []string `json:""`
}

type ResultLah struct {
	BniRole Group `json:"user_bni"`
	UserHolding
}

// holding
type WkRName struct {
	WkName string `json:"wk_name"`
}

type KkksRName struct {
	KkksName string `json:"kkks_name"`
}

type ResponseRelasi struct {
	IdHolding   int    `json:"id_holding"`
	HoldingName string `json:"holding_name"`
	KkksName    string `json:"kkks_name"`
	WkName      string `json:"wk_name"`
}

type ResponseRelasiContainer struct {
	IdHolding   int    `json:"id_holding"`
	HoldingName string `json:"holding_name"`
	KkksName    string `json:"kkks_name"`
	WkName      string `json:"wk_name"`
}

type HoldingResponse struct {
	IdHolding   int      `json:"id_holding"`
	HoldingName string   `json:"holding_name"`
	KkksName    []string `json:"kkks_name"`
	WkName      []string `json:"wk_name"`
}

type Result struct {
	HoldingData   []HoldingResponse `json:"list_holding"`
	UnholdingData []WKResponse      `json:"unholding"`
	Role          []string          `json:"role"`
}

type ResResult struct {
	BniUser     []BniResult  `json:"user_bni"`
	SkkUser     []SkkResult  `json:"user_skk"`
	HoldingUser []Result     `json:"user_holding"`
	KkksUser    []KkksResult `json:"user_kkks"`
}

// kkks
type ResponseRelasiKkks struct {
	IdKkks      int    `json:"id_kkks"`
	KkksName    string `json:"kkks_name"`
	HoldingName string `json:"holding_name"`
	WkName      string `json:"wk_name"`
}

type KkksResponse struct {
	IdKkks      int    `json:"id_kkks"`
	KkksName    string `json:"kkks_name"`
	HoldingName string `json:"holding_name"`
	WkName      string `json:"wk_name"`
}

type KkksResult struct {
	KkksData []KkksResponse `json:"list_kkks"`
	Role     []string       `json:"role"`
}

type SkkResponse struct {
	HoldingName string `json:"holding_name"`
	KkksName    string `json:"kkks_name"`
	WkName      string `json:"wk_name"`
}

type SkkResult struct {
	SkkData []SkkResponse `json:"list_skk"`
	Role    []string      `json:"role"`
}

// BNI
type BniResult struct {
	Role []string `json:"role"`
}

// wk & unholding
type WKResponse struct {
	IdWk     int    `json:"id_wk"`
	WkName   string `json:"wk_name"`
	KkksName string `json:"kkks_name"`
}

// penambahan
type RequestData struct {
	Username string   `json:"username" validate:"patternAlphaUnderscore`
	Role     string   `json:"role" binding:"required" validate:"patternAlphaUnderscore"`
	Label    string   `json:"label" binding:"required" validate:"patternAlphaUnderscore"`
	Tenant   string   `json:"tenant" binding:"required" validate:"patternAlphaUnderscore" `
	IsActive bool     `json:"status"`
	Object   []Object `json:"object" `
}

type DetailtGetRole struct {
	Role string `form:"role" json:"role" validate:"patternAlphaUnderscore"`
}

type GetRoleByUsername struct {
	Username string `json:"username" `
	Data     []Data `json:"data" `
}

type Data struct {
	Role      string      `json:"role" `
	Label     string      `json:"label" `
	Tenant    string      `json:"tenant" `
	Privilege []Privilege `json:"privileges" `
}

type Privilege struct {
	Object     string   `json:"object"`
	Permission []string `json:"permission" `
}

type Object struct {
	Name  string  `json:"name" validate:"patternAlphaUnderscore"`
	Child []Child `json:"child" `
}

type Child struct {
	Name      string   `json:"name" validate:"patternAlphaUnderscore"`
	Privilege []string `json:"privilege" validate:"patternAlphaUnderscore" `
}
