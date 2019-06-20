package models

// LeadSourceExample represents a potential lead
type LeadSourceExample struct {
	VisitorID string `binding:"required,uuid4"`
	Source    string `binding:"required,eq=google|eq=yahoo|eq=other"`
}
