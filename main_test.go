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

func TestMain(t *testing.T) {
	testHeaders := map[string]string{"Content-Type": "application/json"}
	r := getRouter()

	t.Run("CarTests", func(t *testing.T) {
		type cases struct {
			Name        string
			Path        string
			ExpCode     int
			ExpFields   []string
			ExpMessages []string
			Body        CarExample
		}

		tests := []cases{
			cases{
				Name:        "no-make-or-model-provided",
				Path:        "/car",
				ExpCode:     400,
				ExpFields:   []string{"Make", "Model"},
				ExpMessages: []string{"Make is required", "Model is required"},
			},
			cases{
				Name:        "no-make-provided",
				Path:        "/car",
				ExpCode:     400,
				Body:        CarExample{Model: "test model"},
				ExpFields:   []string{"Make"},
				ExpMessages: []string{"Make is required"},
			},
			cases{
				Name:        "no-model-provided",
				Path:        "/car",
				ExpCode:     400,
				Body:        CarExample{Make: "test make"},
				ExpFields:   []string{"Model"},
				ExpMessages: []string{"Model is required"},
			},
			cases{
				Name:        "make-too-short",
				Path:        "/car",
				ExpCode:     400,
				Body:        CarExample{Make: "aa", Model: "test model"},
				ExpFields:   []string{"Make"},
				ExpMessages: []string{"Make must be at least 3 characters"},
			},
			cases{
				Name:        "make-too-long",
				Path:        "/car",
				ExpCode:     400,
				Body:        CarExample{Make: "aaaaaaaaaaaaaaaaaaaaa", Model: "test model"},
				ExpFields:   []string{"Make"},
				ExpMessages: []string{"Make must be less than or equal to 20 characters"},
			},
			cases{
				Name:        "model-too-short",
				Path:        "/car",
				ExpCode:     400,
				Body:        CarExample{Make: "test make", Model: "q"},
				ExpFields:   []string{"Model"},
				ExpMessages: []string{"Model must be at least 2 characters"},
			},
			cases{
				Name:        "model-too-long",
				Path:        "/car",
				ExpCode:     400,
				Body:        CarExample{Make: "test make", Model: "aaaaaaaaaaaaaaaa"},
				ExpFields:   []string{"Model"},
				ExpMessages: []string{"Model must be less than or equal to 15 characters"},
			},
		}

		for _, ct := range tests {
			t.Run(ct.Name, func(t *testing.T) {
				bytes, _ := json.Marshal(ct.Body)
				req := performRequest(r, "POST", "/car", &bytes, testHeaders)
				t.Run("ExpectedCode", func(t *testing.T) {
					assert.Equal(t, ct.ExpCode, req.Code, req.Body)
				})

				errs := map[string]string{}
				_ = json.Unmarshal([]byte(req.Body.String()), &errs)
				t.Run("ExpectedFields", func(t *testing.T) {
					for _, f := range ct.ExpFields {
						found := false
						for k := range errs {
							if k == f {
								found = true
							}
						}
						assert.Equal(t, true, found, fmt.Sprint("Field: ", f))
					}
				})
				t.Run("ExpectedMessages", func(t *testing.T) {
					for _, m := range ct.ExpMessages {
						found := false
						for _, v := range errs {
							if v == m {
								found = true
							}
						}
						assert.Equal(t, true, found, fmt.Sprint("Messages: ", errs), fmt.Sprint("\nExpected: ", m), fmt.Sprint("\nBody: ", string(bytes)))
					}
				})
			})
		}
	})

	t.Run("AlbumTests", func(t *testing.T) {
		// type AlbumExample struct {
		// 	Artist []string `biding:"required,gte=1,lte=5,dive,gte=2,lte=50"`
		// 	Name string `binding:"required,gte=2,lte=50,alpha"`
		// }
		type cases struct {
			Name    string
			Path    string
			ExpCode int
			Body    AlbumExample
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
