package ctxWithValue

import (
	"context"
	"fmt"
	"net/http"
)

type guidKey int

const (
	_ guidKey = iota
	key
)

func contextWithGUID(ctx context.Context, guid string) context.Context {
	return context.WithValue(ctx, key, guid)
}

func guidFromContext(ctx context.Context) (string, bool) {
	g, ok := ctx.Value(key).(string)
	return g, ok
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		if guid := req.Header.Get("X_GUID"); guid != "" {
			ctx = contextWithGUID(ctx, guid)
		} else {
			ctx = contextWithGUID(ctx, uuid.New().String())
		}
		req = req.WithContext(ctx)
		h.ServeHTTP(rw, req)
	})
}

type Logger struct{}

func (Logger) Log(ctx context.Context, message string) {
	if guid, ok := guidFromContext(ctx); ok {
		message = fmt.Sprintf("GUID: %s - %s", guid, message)
	}
	// do logging
	fmt.Println(message)
}

func Request(req *http.Request) *http.Request {
	ctx := req.Context()
	if guid, ok := guidFromContext(ctx); ok {
		req.Header.Add("X_GUID", guid)
	}
	return req
}

// generating business logic independent to all tracking info.

type LoggerIF interface {
	Log(context.Context, string)
}

type RequestDecorator func(*http.Request) *http.Request

type BusinessLogic struct {
	RequestDecorator RequestDecorator
	Logger           Logger
	Remote           string
}

func (bl BusinessLogic) bussinessLogic(ctx context.Context, user string, data string) (string, error) {
	bl.Logger.Log(ctx, "starting businessLogic for "+user+" with "+data)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, bl.Remote+"?query="+data, nil)
	if err != nil {
		bl.Logger.Log(ctx, "error building remote request:"+err.Error())
		return "", err
	}
	req = bl.RequestDecorator(req)
	resp, err := http.DefaultClient.Do(req)

	// processing futhermore
	return resp.Status, nil
}

/*
And in main.go,

bl := BusinessLogic {
	RequestDecorator: ctxWithValue.Request,
	Logger: ctxWithValue.Logger{},
	Remote: "http://www.example.com/query"
}
*/
