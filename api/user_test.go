package api

import (
	"bytes"
	"encoding/json"
	"github.com/AveryKing/tasks/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUserAPI(t *testing.T) {
	username := util.RandomString(6)
	password := util.RandomString(6)

	testCases := []struct {
		name          string
		body          gin.H
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username": username,
				"password": password,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, 200, recorder.Code)
			},
		},
		{
			name: "DuplicateUsername",
			body: gin.H{
				"username": username,
				"password": password,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, 403, recorder.Code)
			},
		},
		{
			name: "NonAlphaUsername",
			body: gin.H{
				"username": "!!$(#)@!)_!",
				"password": password,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, 400, recorder.Code)
			},
		},
		{
			name: "TooShortPassword",
			body: gin.H{
				"username": util.RandomString(6),
				"password": "12345",
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, 400, recorder.Code)
			},
		},
		{
			name: "EmptyBody",
			body: gin.H{},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, 400, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			router := newTestServer(t).setupRouter()
			recorder := httptest.NewRecorder()
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			req, _ := http.NewRequest("POST", "/users", bytes.NewReader(data))
			router.ServeHTTP(recorder, req)
			tc.checkResponse(recorder)
		})
	}
}
