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
	testHeaders := map[string]string{"Content-Type": "application/json"}
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

		for _, ct := range tests {
			t.Run(ct.Name, func(t *testing.T) {
				bytes, _ := json.Marshal(ct.Body)
				req := performRequest(r, "POST", ct.Path, &bytes, testHeaders)

				assertCodeAndMessages(t, ct, req)
			})
		}
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
			Name: "artist-entry-too-long",
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
				Name: "name-not-provided",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Name"},
				ExpMessages: []string{"Name is required"},
				Body: AlbumExample{
					Artist: []string{"blink 182"},
				},
			},
			testCase{
				Name: "name-too-short",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Name"},
				ExpMessages: []string{"Name must contain at least 2 characters"},
				Body: AlbumExample{
					Artist: []string{"blink 182"},
					Name: "a",
				},
			},
			testCase{
				Name: "name-too-long",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Name"},
				ExpMessages: []string{"Name must contain no more than 50 characters"},
				Body: AlbumExample{
					Artist: []string{"blink 182"},
					Name: "asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasd",
				},
			},
		}

		for _, at := range tests {
			t.Run(at.Name, func(t *testing.T) {
				bytes, _ := json.Marshal(at.Body)
				req := performRequest(r, "POST", at.Path, &bytes, testHeaders)

				assertCodeAndMessages(t, at, req)
			})
		}
	})

	t.Run("PasswordTests", func(t *testing.T) {
		// type PasswordExample struct {
		// 	Username string `binding:"required,gte=5,lte=30,alphanum"`
		// 		OldPassword string `binding:"required,gte=8,lte=30"`
		// 	Password string `binding:"required,gte=8,lte=30,nefield=OldPassword,excludes=password,excludesrune=^"`
		// 	PasswordConfirm string `binding:"required,gte=8,lte=30,eqfield=Password,nefield=OldPassword"`
		// }
	})

	t.Run("LeadSourceTests", func(t *testing.T) {
		// type LeadSourceExample struct {
		// 	VisitorID string `binding:"required,uuid4"`
		// 	Source string `binding:"required,eq=google|yahoo|other"`
		// }
	})

	t.Run("SignupTests", func(t *testing.T) {
		// type SignupExample struct {
		// 	Username string `binding:"required,gte=5,lte=30,alphanum"`
		// 	Email string `binding:"required,email,max=100"`
		// }
	})

	t.Run("StudioSessionTests", func(t *testing.T) {
		// type StudioSessionExample struct {
		// 	BandName string `binding:"required,max=30,alphanum"`
		// 	BandMembers int `binding:"required,numeric,max=8"`
		// 	StartTime time.Time `binding:"required"`
		// 	EndTime time.Time `binding:"required,gtfield=StartTime"`
		// }
	})

	t.Run("PartnershipRequestTests", func(t *testing.T) {
		// type PartnershipRequestExample struct {
		// 	CompanyName string `binding:"required,max=50,alphanum"`
		// 	Website string `binding:"required,url"`
		// 	Referrer string `binding:"uri"`
		// }
	})

	t.Run("PostCoordinatesTests", func(t *testing.T) {
		// type PostCoordinatesExample struct {
		// 	UserID int `binding:"required,int"`
		// 	Lat string `binding:"required,latitude"`
		// 	Long string `binding:"required,longitude"`
		// }
	})

	t.Run("UploadCSVsTests", func(t *testing.T) {
		// type UploadCsvsExample struct {
		// 	Content [][]string `binding:"required,max=5,dive,gte=3,max=50,dive,required,gte=5,max=1000,alpha"`
		// }
	})
}

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
