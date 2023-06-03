package entity

type LoginReq struct {
	Email    string `json:"email" validate:"required" example:"admin@email.com"`
	Password string `json:"password" validate:"required" example:"1veG+hbzZNStOA6rUSvwVw=="`
}

type RefreshReq struct {
	Refresh string `json:"refresh" validate:"required" example:"c13035843647f520402b5b3d96fa302c9c23817d746af08dc9cab2447993cfb6.1687316436"`
}
