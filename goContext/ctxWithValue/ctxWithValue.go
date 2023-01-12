package ctxWithValue

import (
	"context"
	"net/http"
)

// not exposed constance
type userKey int

const (
	_ userKey = iota
	key
)

func ContextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, key, user)
}

func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(key).(string)
	return user, ok
}

// extractUser extracts user ID from cookie.
func extractUser(req *http.Request) (string, error) {
	userCookie, err := req.Cookie("user")
	if err != nil {
		return "", err
	}
	return userCookie.Value, nil
}

// Middleware1 extracts user id from cookie and save it to the context.
func Middleware1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		user, err := extractUser(req)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := req.Context()
		ctx = ContextWithUser(ctx, user)
		req = req.WithContext(ctx)
		h.ServeHTTP(rw, req)
	})
}

type Logic interface {
	businessLogic(ctx context.Context, user string, data string) (string, error)
}
type Controller struct {
	Logic
}

type tmp1 struct{}

var identity tmp1

func (i tmp1) UserFromContext(ctx context.Context) (string, bool)

func (c Controller) HandleRequest(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	user, ok := identity.UserFromContext(ctx)
	if !ok {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := req.URL.Query().Get("data")
	result, err := c.Logic.businessLogic(ctx, user, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte(result))
}
