package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/mocha-bot/enjot/pkg/logger"
)

func Logger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			wrapRW := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			defer func() {
				scheme := "http"
				if r.TLS != nil {
					scheme = "https"
				}

				endpoint := fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

				requestBody, _ := ioutil.ReadAll(r.Body)

				if wrapRW.Status() != http.StatusOK {
					logger.Error().
						Str("context", "http.call").
						Str("method", r.Method).
						Str("endpoint", endpoint).
						Int("status", wrapRW.Status()).
						Str("data", string(requestBody)).
						Send()
				}
			}()

			next.ServeHTTP(wrapRW, r)
		}

		return http.HandlerFunc(fn)
	}
}
