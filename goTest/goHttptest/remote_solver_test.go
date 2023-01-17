package goHttptest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteSolver_Reslove(t *testing.T) {
	type info struct {
		expression string
		code       int
		body       string
	}
	var io info

	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			expression := req.URL.Query().Get("expresssion")
			if expression != io.expression {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte("invalid expression: " + io.expression))
				return
			}
			rw.WriteHeader(io.code)
			rw.Write([]byte(io.body))
		}))
	defer server.Close()

	rs := RemoteSolver{
		MathServerURL: server.URL,
		Client:        server.Client(),
	}

	data := []struct {
		name   string
		io     info
		result float64
	}{
		{"case1", info{"2 + 2 * 10", http.StatusOK, "22"}, 22},
		{"case2", info{"( 2 + 2 ) * 10", http.StatusOK, "40"}, 40},
		{"case3", info{"( 2 + 2 * 10 ", http.StatusBadRequest, ""}, 0},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			io = d.io
			result, err := rs.Resolve(context.Background(), d.io.expression)
			if result != d.result {
				t.Errorf("io `%f`, got `%f`", d.result, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("io error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
