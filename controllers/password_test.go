package controllers

import (
	"testing"

	"github.com/mike-webster/golang-validation/models"
)

func TestPostPassword(t *testing.T) {
	t.Run("PasswordTests", func(t *testing.T) {
		tests := []testCase{
			testCase{
				Name:        "username-not-provided",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"Username"},
				ExpMessages: []string{"Username is required"},
				Body: models.PasswordExample{
					Password:        "testpass",
					PasswordConfirm: "testpass",
					OldPassword:     "oldtestpass",
				},
			},
			testCase{
				Name:        "username-too-short",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"Username"},
				ExpMessages: []string{"Username must contain at least 5 characters"},
				Body: models.PasswordExample{
					Username:        "a",
					Password:        "testpass",
					PasswordConfirm: "testpass",
					OldPassword:     "oldtestpass",
				},
			},
			testCase{
				Name:        "username-too-long",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"Username"},
				ExpMessages: []string{"Username must contain no more than 30 characters"},
				Body: models.PasswordExample{
					Username:        "fdsafdsafdsafdsafdsafdsafdasfds",
					Password:        "testpass",
					PasswordConfirm: "testpass",
					OldPassword:     "oldtestpass",
				},
			},
			testCase{
				Name:        "username-not-alphanum",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"Username"},
				ExpMessages: []string{"Username must be alphanumeric"},
				Body: models.PasswordExample{
					Username:        "fdsafd fds",
					Password:        "testpass",
					PasswordConfirm: "testpass",
					OldPassword:     "oldtestpass",
				},
			},
			testCase{
				Name:        "old-password-not-provided",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"OldPassword"},
				ExpMessages: []string{"Old password is required"},
				Body: models.PasswordExample{
					Username:        "fdsafdfdsfds",
					Password:        "testpass",
					PasswordConfirm: "testpass",
				},
			},
			testCase{
				Name:        "old-password-too-short",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"OldPassword"},
				ExpMessages: []string{"Old password must contain at least 8 characters"},
				Body: models.PasswordExample{
					Username:        "fdsafdfdsfds",
					Password:        "testpass",
					PasswordConfirm: "testpass",
					OldPassword:     "a",
				},
			},
			testCase{
				Name:        "old-password-too-long",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"OldPassword"},
				ExpMessages: []string{"Old password must contain no more than 30 characters"},
				Body: models.PasswordExample{
					Username:        "fdsafdfdsfds",
					Password:        "testpass",
					PasswordConfirm: "testpass",
					OldPassword:     "fdsafdsafdasfdasfdasfdasfdsafds",
				},
			},
			testCase{
				Name:        "password-not-provided",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"Password"},
				ExpMessages: []string{"Password is required"},
				Body: models.PasswordExample{
					Username:        "fdsafdfdsfds",
					PasswordConfirm: "testpass",
					OldPassword:     "oldtestpass",
				},
			},
			testCase{
				Name:        "password-too-short",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"Password"},
				ExpMessages: []string{"Password must contain at least 8 characters"},
				Body: models.PasswordExample{
					Username:        "fdsafdfdsfds",
					Password:        "test",
					PasswordConfirm: "testpass",
					OldPassword:     "oldtestpass",
				},
			},
			testCase{
				Name:        "password-too-long",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"Password"},
				ExpMessages: []string{"Password must contain no more than 30 characters"},
				Body: models.PasswordExample{
					Username:        "fdsafdfdsfds",
					Password:        "fdsafdsafdasfdasfdasfdasfdsafds",
					PasswordConfirm: "testpass",
					OldPassword:     "oldtestpass",
				},
			},
			testCase{
				Name:        "password-not-equal-old-password",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"Password"},
				ExpMessages: []string{"Password must not be the same as Old password"},
				Body: models.PasswordExample{
					Username:        "fdsafdfdsfds",
					Password:        "oldtestpass",
					PasswordConfirm: "oldtestpass",
					OldPassword:     "oldtestpass",
				},
			},
			testCase{
				Name:        "password-not-equal-'password'",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"Password"},
				ExpMessages: []string{"Password must not be 'password'"},
				Body: models.PasswordExample{
					Username:        "fdsafdfdsfds",
					Password:        "password",
					PasswordConfirm: "password",
					OldPassword:     "oldtestpass",
				},
			},
			testCase{
				Name:        "password-doesnt-contain-'^'",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"Password"},
				ExpMessages: []string{"Password must not contain '^'"},
				Body: models.PasswordExample{
					Username:        "fdsafdfdsfds",
					Password:        "passw^rd",
					PasswordConfirm: "password",
					OldPassword:     "oldtestpass",
				},
			},
			testCase{
				Name:        "password-confirm-not-provided",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"PasswordConfirm"},
				ExpMessages: []string{"Password confirm is required"},
				Body: models.PasswordExample{
					Username:    "fdsafdfdsfds",
					Password:    "testpass",
					OldPassword: "oldtestpass",
				},
			},
			testCase{
				Name:        "password-confirm-too-short",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"PasswordConfirm"},
				ExpMessages: []string{"Password confirm must contain at least 8 characters"},
				Body: models.PasswordExample{
					Username:        "fdsafdfdsfds",
					Password:        "testpass",
					PasswordConfirm: "t",
					OldPassword:     "oldtestpass",
				},
			},
			testCase{
				Name:        "password-confirm-too-long",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"PasswordConfirm"},
				ExpMessages: []string{"Password confirm must contain no more than 30 characters"},
				Body: models.PasswordExample{
					Username:        "fdsafdfdsfds",
					Password:        "testpass",
					PasswordConfirm: "fdasfdsafdsafdsafdasfdsafdsafds",
					OldPassword:     "oldtestpass",
				},
			},
			testCase{
				Name:        "password-confirm-equal-password",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"PasswordConfirm"},
				ExpMessages: []string{"Password confirm must match Password"},
				Body: models.PasswordExample{
					Username:        "fdsafdfdsfds",
					Password:        "testpass",
					PasswordConfirm: "oldtestpass",
					OldPassword:     "oldtestpass",
				},
			},
			testCase{
				Name:        "password-confirm-not-equal-old-password",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"PasswordConfirm"},
				ExpMessages: []string{"Password confirm must not be the same as Old password"},
				Body: models.PasswordExample{
					Username:        "fdsafdfdsfds",
					Password:        "oldtestpass",
					PasswordConfirm: "oldtestpass",
					OldPassword:     "oldtestpass",
				},
			},
		}
		runTests(t, tests, GetRouter())
	})
}
