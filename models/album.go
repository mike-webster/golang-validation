package models

// AlbumExample represents an album
type AlbumExample struct {
	Artist []string `binding:"required,gte=1,lte=5,dive,gte=2,lte=50"`
	Name   string   `binding:"required,gte=2,lte=50"`
}