package ctxPattern

import (
	"context"
	"fmt"
	"net/http"
)

// Common pattern
func CommonMW(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		// wrapping various jobs with contetx
		req = req.WithContext(ctx)
		handler.ServeHTTP(rw, req)
	})
}

// Calling handler with context
func Handler(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	err := req.ParseForm()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	data := req.FormValue("data")
	result, err := logic(ctx, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte(result))
}

func logic(ctx context.Context, data string) (string, error) {
	fmt.Println(data)
	return data, nil
}

// Setting context at the request going out from my own HTTP Service
type ServiceCaller struct {
	client *http.Client
}

func (sc ServiceCaller) callAnotherService(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "http://example.com?data="+data, nil)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)
	resp, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpectd status code %d", resp.StatusCode)
	}
	//id ,err := processResponse(resp.body)
	id := ""
	return id, err
}
