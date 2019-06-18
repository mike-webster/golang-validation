package main

import "time"

type CarExample struct {
	Make  string `binding:"required,lte=20,gte=3"`
	Model string `binding:"required,lte=15,gte=2"`
}

type AlbumExample struct {
	Artist []string `biding:"required,gte=1,lte=5,dive,gte=2,lte=50"`
	Name string `binding:"required,gte=2,lte=50,alpha"`
}

type PasswordExample struct {
	Username string `binding:"required,gte=5,lte=30,alphanum"`
        OldPassword string `binding:"required,gte=8,lte=30"`
	Password string `binding:"required,gte=8,lte=30,nefield=OldPassword,excludes=password,excludesrune=^"`
	PasswordConfirm string `binding:"required,gte=8,lte=30,eqfield=Password,nefield=OldPassword"`
}

type LeadSourceExample struct {
	VisitorID string `binding:"required,uuid4"`
	Source string `binding:"required,eq=google|yahoo|other"`
}

type SignupExample struct {
	Username string `binding:"required,gte=5,lte=30,alphanum"`
	Email string `binding:"required,email,max=100"`
}

type StudioSessionExample struct {
	BandName string `binding:"required,max=30,alphanum"`
	BandMembers int `binding:"required,numeric,max=8"`
	StartTime time.Time `binding:"required"`
	EndTime time.Time `binding:"required,gtfield=StartTime"`
}

type PartnershipRequestExample struct {
	CompanyName string `binding:"required,max=50,alphanum"`
	Website string `binding:"required,url"`
	Referrer string `binding:"uri"`
}

type PostCoordinatesExample struct {
	UserID int `binding:"required,int"`
	Lat string `binding:"required,latitude"`
	Long string `binding:"required,longitude"`
}

type UploadCsvsExample struct {
	Content [][]string `binding:"required,max=5,dive,gte=3,max=50,dive,required,gte=5,max=1000,alpha"`
}
