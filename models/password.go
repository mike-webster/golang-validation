package models

// PasswordExample represents a users password
type PasswordExample struct {
	Username        string `binding:"required,gte=5,lte=30,alphanum"`
	OldPassword     string `binding:"required,gte=8,lte=30"`
	Password        string `binding:"required,gte=8,lte=30,nefield=OldPassword,excludes=password,excludesrune=^"`
	PasswordConfirm string `binding:"required,gte=8,lte=30,eqfield=Password,nefield=OldPassword"`
}