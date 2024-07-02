package common

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendResponse(t *testing.T) {
	type args struct {
		w       *httptest.ResponseRecorder
		code    int
		message string
		data    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Ok response",
			args: args{
				w:       httptest.NewRecorder(),
				code:    http.StatusOK,
				message: "OK",
				data:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendResponse(tt.args.w, tt.args.code, tt.args.message, tt.args.data)
			assert.Equal(t, tt.args.w.Header().Get("Content-Type"), "application/json", "Content-Type should be application/json")
			assert.Equal(t, tt.args.w.Code, tt.args.code, "Response code should be equal to the input code")
			jsonBody, err := json.Marshal(Response{
				Code:    tt.args.code,
				Message: tt.args.message,
				Data:    tt.args.data,
			})
			if err != nil {
				t.Fatalf("Error marshalling response: %v", err)
			}

			assert.Equal(t, tt.args.w.Body.String(), string(jsonBody), "Response body should be equal to the input data")
		})
	}
}
