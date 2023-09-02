package model

type SignatureRequest struct {
	Host    string `json:"host" form:"host"`
	Path    string `json:"path" form:"path"`
	Payload string `json:"payload" form:"payload"`
}

type ReqValidasiUsername struct {
	Username string `json:"username" form:"username" validate:"patternAlphaUnderscore`
}
type ReqValidasiUserJson struct {
	Username string `json:"username"`
}
