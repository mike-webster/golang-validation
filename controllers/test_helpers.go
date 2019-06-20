package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
)

type testCase struct {
	Name        string
	Path        string
	ExpCode     int
	ExpFields   []string
	ExpMessages []string
	Body        interface{} // I made this an interface so that it could be used by all test cases
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
