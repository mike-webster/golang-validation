package controllers

import (
	"testing"

	"github.com/mike-webster/golang-validation/models"
)

func TestPostCar(t *testing.T) {
	t.Run("CarTests", func(t *testing.T) {
		tests := []testCase{
			testCase{
				Name:        "no-make-or-model-provided",
				Path:        "/car",
				ExpCode:     400,
				ExpFields:   []string{"Make", "Model"},
				ExpMessages: []string{"Make is required", "Model is required"},
			},
			testCase{
				Name:        "no-make-provided",
				Path:        "/car",
				ExpCode:     400,
				Body:        models.CarExample{Model: "test model"},
				ExpFields:   []string{"Make"},
				ExpMessages: []string{"Make is required"},
			},
			testCase{
				Name:        "no-model-provided",
				Path:        "/car",
				ExpCode:     400,
				Body:        models.CarExample{Make: "test make"},
				ExpFields:   []string{"Model"},
				ExpMessages: []string{"Model is required"},
			},
			testCase{
				Name:        "make-too-short",
				Path:        "/car",
				ExpCode:     400,
				Body:        models.CarExample{Make: "aa", Model: "test model"},
				ExpFields:   []string{"Make"},
				ExpMessages: []string{"Make must contain at least 3 characters"},
			},
			testCase{
				Name:        "make-too-long",
				Path:        "/car",
				ExpCode:     400,
				Body:        models.CarExample{Make: "aaaaaaaaaaaaaaaaaaaaa", Model: "test model"},
				ExpFields:   []string{"Make"},
				ExpMessages: []string{"Make must contain no more than 20 characters"},
			},
			testCase{
				Name:        "model-too-short",
				Path:        "/car",
				ExpCode:     400,
				Body:        models.CarExample{Make: "test make", Model: "q"},
				ExpFields:   []string{"Model"},
				ExpMessages: []string{"Model must contain at least 2 characters"},
			},
			testCase{
				Name:        "model-too-long",
				Path:        "/car",
				ExpCode:     400,
				Body:        models.CarExample{Make: "test make", Model: "aaaaaaaaaaaaaaaa"},
				ExpFields:   []string{"Model"},
				ExpMessages: []string{"Model must contain no more than 15 characters"},
			},
		}
		runTests(t, tests, GetRouter())
	})
}
