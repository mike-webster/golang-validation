package controllers

import (
	"testing"

	"github.com/mike-webster/golang-validation/models"
)

func TestPostLead(t *testing.T) {
	t.Run("LeadTests", func(t *testing.T) {
		tests := []testCase{
			testCase{
				Name:        "no-visitor-id-provided",
				Path:        "/lead",
				ExpCode:     400,
				ExpFields:   []string{"VisitorID"},
				ExpMessages: []string{"Visitor id is required"},
				Body: models.LeadSourceExample{
					Source: "google",
				},
			},
			testCase{
				Name:        "visitor-id-not-uuidv4",
				Path:        "/lead",
				ExpCode:     400,
				ExpFields:   []string{"VisitorID"},
				ExpMessages: []string{"Visitor id is not a valid uuidv4"},
				Body: models.LeadSourceExample{
					VisitorID: "not-valid-uuid",
					Source:    "google",
				},
			},
			testCase{
				Name:        "source-not-provided",
				Path:        "/lead",
				ExpCode:     400,
				ExpFields:   []string{"Source"},
				ExpMessages: []string{"Source is required"},
				Body: models.LeadSourceExample{
					VisitorID: "f6a91ca9-a517-458a-80f1-2e31b58f9cc2",
				},
			},
			testCase{
				Name:        "source-not-valid",
				Path:        "/lead",
				ExpCode:     400,
				ExpFields:   []string{"Source"},
				ExpMessages: []string{"Source is not valid"},
				Body: models.LeadSourceExample{
					VisitorID: "f6a91ca9-a517-458a-80f1-2e31b58f9cc2",
					Source:    "not-a-valid-source",
				},
			},
			testCase{
				Name:    "success",
				Path:    "/lead",
				ExpCode: 200,
				Body: models.LeadSourceExample{
					VisitorID: "f6a91ca9-a517-458a-80f1-2e31b58f9cc2",
					Source:    "google",
				},
			},
		}
		runTests(t, tests, GetRouter())
	})
}
