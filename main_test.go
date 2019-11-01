package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/airdb/bbs-api/web"
	"github.com/stretchr/testify/assert"
)

func performRequest(method, path string) (*httptest.ResponseRecorder, error) {
	router := web.NewRouter()
	req, err := http.NewRequest(method, path, nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w, err
}

//You could write the init logic like reset database code here
var testCases = []struct {
	TestID         int
	Uri            string
	Method         string
	bodyData       string
	ExpectedCode   int
	responseRegexg string
	msg            string
}{
	//Testing will run one by one, so you can combine it to a user story till another init().
	//And you can modified the header or body in the func(req *http.Request) {}

	//---------------------   Testing case register   ---------------------
	{
		1,
		"/v1/notice/list",
		"GET",
		`{"user":{"username": "wangzitian0","email": "wzt@gg.cn","password": "jakejxke"}}`,
		http.StatusOK,
		`{"user":{"username":"wangzitian0","email":"wzt@gg.cn","bio":"","image":null,"token":"([a-zA-Z0-9-_.]{115})"}}`,
		"valid data and should return StatusCreated",
	},
	{
		2,
		"/v1/article/list",
		"GET",
		`{"user":{"username": "wangzitian0","email": "wzt@gg.cn","password": "jakejxke"}}`,
		http.StatusOK,
		`{"user":{"username":"wangzitian0","email":"wzt@gg.cn","bio":"","image":null,"token":"([a-zA-Z0-9-_.]{115})"}}`,
		"valid data and should return StatusCreated",
	},
	{
		3,
		"/v1/square/list",
		"GET",
		`{"user":{"username": "wangzitian0","email": "wzt@gg.cn","password": "jakejxke"}}`,
		http.StatusOK,
		`{"user":{"username":"wangzitian0","email":"wzt@gg.cn","bio":"","image":null,"token":"([a-zA-Z0-9-_.]{115})"}}`,
		"valid data and should return StatusCreated",
	},
	{
		4,
		"/v1/article/list",
		"GET",
		`{"user":{"username": "wangzitian0","email": "wzt@gg.cn","password": "jakejxke"}}`,
		http.StatusOK,
		`{"user":{"username":"wangzitian0","email":"wzt@gg.cn","bio":"","image":null,"token":"([a-zA-Z0-9-_.]{115})"}}`,
		"valid data and should return StatusCreated",
	},
}

// https://github.com/gothinkster/golang-gin-realworld-example-app/blob/master/users/unit_test.go
func Test_API(t *testing.T) {
	asserts := assert.New(t)
	for _, testCase := range testCases {
		t.Log("test case:", testCase.TestID, testCase.Uri)
		uri := "/apis/bbs" + testCase.Uri

		resp, err := performRequest(testCase.Method, uri)
		asserts.NoError(err)

		assert.Equal(t, testCase.ExpectedCode, resp.Code)
		// assert.Regexp(testCase.responseRegexg, resp.Body.String(), )
	}
}
