package models

// Proptip: If you can't figure out why your validation is working... check that you spelled
// "binding" correctly ;)

// CarExample represents a car
type CarExample struct {
	Make  string `binding:"required,lte=20,gte=3"`
	Model string `binding:"required,lte=15,gte=2"`
}
