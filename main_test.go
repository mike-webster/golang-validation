package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bmizerany/assert"
	gin "github.com/gin-gonic/gin"
)

type testCase struct {
	Name        string
	Path        string
	ExpCode     int
	ExpFields   []string
	ExpMessages []string
	Body        interface{} // I made this an interface so that it could be used by all test cases
}

func TestMain(t *testing.T) {
	r := getRouter()

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
				Body:        CarExample{Model: "test model"},
				ExpFields:   []string{"Make"},
				ExpMessages: []string{"Make is required"},
			},
			testCase{
				Name:        "no-model-provided",
				Path:        "/car",
				ExpCode:     400,
				Body:        CarExample{Make: "test make"},
				ExpFields:   []string{"Model"},
				ExpMessages: []string{"Model is required"},
			},
			testCase{
				Name:        "make-too-short",
				Path:        "/car",
				ExpCode:     400,
				Body:        CarExample{Make: "aa", Model: "test model"},
				ExpFields:   []string{"Make"},
				ExpMessages: []string{"Make must contain at least 3 characters"},
			},
			testCase{
				Name:        "make-too-long",
				Path:        "/car",
				ExpCode:     400,
				Body:        CarExample{Make: "aaaaaaaaaaaaaaaaaaaaa", Model: "test model"},
				ExpFields:   []string{"Make"},
				ExpMessages: []string{"Make must contain no more than 20 characters"},
			},
			testCase{
				Name:        "model-too-short",
				Path:        "/car",
				ExpCode:     400,
				Body:        CarExample{Make: "test make", Model: "q"},
				ExpFields:   []string{"Model"},
				ExpMessages: []string{"Model must contain at least 2 characters"},
			},
			testCase{
				Name:        "model-too-long",
				Path:        "/car",
				ExpCode:     400,
				Body:        CarExample{Make: "test make", Model: "aaaaaaaaaaaaaaaa"},
				ExpFields:   []string{"Model"},
				ExpMessages: []string{"Model must contain no more than 15 characters"},
			},
		}
		runTests(t, tests, r)
	})

	t.Run("AlbumTests", func(t *testing.T) {
		tests := []testCase{
			testCase{
				Name:        "artists-not-provided",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Artist"},
				ExpMessages: []string{"Artist is required"},
				Body: AlbumExample{
					Name: "dude ranch",
				},
			},
			testCase{
				Name:        "artists-empty",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Artist"},
				ExpMessages: []string{"Artist must contain at least 1 entry"},
				Body: AlbumExample{
					Artist: []string{},
					Name:   "dude ranch",
				},
			},
			testCase{
				Name:        "artists-too-large",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Artist"},
				ExpMessages: []string{"Artist must contain no more than 5 entries"},
				Body: AlbumExample{
					Artist: []string{"fda", "fda", "fdas", "fdas", "fdas", "fda"},
					Name:   "dude ranch",
				},
			},
			testCase{
				Name:        "artist-entry-too-short",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Artist[0]"},
				ExpMessages: []string{"Artist [ 0 ] must contain at least 2 characters"},
				Body: AlbumExample{
					Artist: []string{"f"},
					Name:   "dude ranch",
				},
			},
			testCase{
				Name:        "artist-entry-too-long",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Artist[0]"},
				ExpMessages: []string{"Artist [ 0 ] must contain no more than 50 characters"},
				Body: AlbumExample{
					Artist: []string{"asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasd"},
					Name:   "dude ranch",
				},
			},
			testCase{
				Name:        "name-not-provided",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Name"},
				ExpMessages: []string{"Name is required"},
				Body: AlbumExample{
					Artist: []string{"blink 182"},
				},
			},
			testCase{
				Name:        "name-too-short",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Name"},
				ExpMessages: []string{"Name must contain at least 2 characters"},
				Body: AlbumExample{
					Artist: []string{"blink 182"},
					Name:   "a",
				},
			},
			testCase{
				Name:        "name-too-long",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Name"},
				ExpMessages: []string{"Name must contain no more than 50 characters"},
				Body: AlbumExample{
					Artist: []string{"blink 182"},
					Name:   "asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasd",
				},
			},
		}
		runTests(t, tests, r)
	})

	t.Run("PasswordTests", func(t *testing.T) {
		tests := []testCase{
			testCase{
				Name:        "username-not-provided",
				Path:        "/password",
				ExpCode:     400,
				ExpFields:   []string{"Username"},
				ExpMessages: []string{"Username is required"},
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
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
				Body: PasswordExample{
					Username:        "fdsafdfdsfds",
					Password:        "oldtestpass",
					PasswordConfirm: "oldtestpass",
					OldPassword:     "oldtestpass",
				},
			},
		}
		runTests(t, tests, r)
	})

	t.Run("UploadCSVsTests", func(t *testing.T) {
		tests := []testCase{}
		runTests(t, tests, r)
	})
}

// performRequest performs the request ;)
func performRequest(r http.Handler,
	method string,
	path string,
	body *[]byte,
	headers map[string]string) *httptest.ResponseRecorder {

	var req *http.Request
	if body != nil {
		reader := strings.NewReader(string(*body))
		log.Println("Sending test body")
		req, _ = http.NewRequest(method, path, reader)
	} else {
		log.Println("Sending empty test body")
		req, _ = http.NewRequest(method, path, nil)
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// runTests will take a slice of test cases and a gin router and perform
// all of the assertions for the tests.
func runTests(t *testing.T, cases []testCase, r *gin.Engine) {
	testHeaders := map[string]string{"Content-Type": "application/json"}
	for _, iCase := range cases {
		t.Run(iCase.Name, func(t *testing.T) {
			bytes, _ := json.Marshal(iCase.Body)
			req := performRequest(r, "POST", iCase.Path, &bytes, testHeaders)

			assertCodeAndMessages(t, iCase, req)
		})
	}
}

// assertCodeAndMessages will check the code and messages in the given test case
// to ensure the response values were what we were expecting.
//
// there's one flaw to the way I built this - we're only checking to make sure
// that the expected fields and messages are contained in the response. This means
// that there could be additional messages that we weren't intending... but it's
// just a testing shortcut :)
func assertCodeAndMessages(t *testing.T, tc testCase, req *httptest.ResponseRecorder) {
	t.Run("ExpectedCode", func(t *testing.T) {
		assert.Equal(t, tc.ExpCode, req.Code, req.Body)
	})

	errs := map[string]string{}
	_ = json.Unmarshal([]byte(req.Body.String()), &errs)
	t.Run("ExpectedFields", func(t *testing.T) {
		for _, f := range tc.ExpFields {
			found := false
			for k := range errs {
				if k == f {
					found = true
				}
			}
			assert.Equal(t, true, found, fmt.Sprint("Field: ", f), fmt.Sprint(" -- Body: ", string(req.Body.String())))
		}
	})
	t.Run("ExpectedMessages", func(t *testing.T) {
		for _, m := range tc.ExpMessages {
			found := false
			for _, v := range errs {
				if v == m {
					found = true
				}
			}
			assert.Equal(t, true, found, fmt.Sprint("Messages: ", errs), fmt.Sprint("\nExpected: ", m), fmt.Sprint(" -- Body: ", string(req.Body.String())))
		}
	})
}
