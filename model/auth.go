package model

type RegisterParams struct {
	Email     string `json:"email" xml:"email" form:"email" validate:"required,email"`
	FirstName string `json:"first_name" xml:"first_name" form:"first_name"`
	LastName  string `json:"last_name" xml:"last_name" form:"last_name"`
	Password  string `json:"password" xml:"password" form:"password" validate:"required"`
}

type LoginParams struct {
	Email    string `json:"email" xml:"email" form:"email"`
	Password string `json:"password xml:"password" form:"password"`
}
