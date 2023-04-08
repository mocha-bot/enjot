package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mocha-bot/enjot/pkg/token"
)

func Authorization(tokenKeySecret []byte) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("Authorization")

			authorizations := strings.Split(authorizationHeader, " ")

			tokenString := authorizations[0]

			if len(authorizations) > 1 {
				tokenString = authorizations[1]
			}

			mapClaims := make(jwt.MapClaims)

			valid, err := token.Verify(mapClaims, tokenKeySecret, tokenString)

			if err == jwt.ErrTokenExpired {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)

				json.NewEncoder(w).Encode(map[string]string{
					"message": err.Error(),
				})

				return
			}

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)

				json.NewEncoder(w).Encode(map[string]string{
					"message": err.Error(),
				})

				return
			}

			if !valid {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)

				json.NewEncoder(w).Encode(map[string]string{
					"message": token.ErrTokenNotValid.Error(),
				})

				return
			}

			ctx := r.Context()
			ctx = parseClaimsToCtx(ctx, mapClaims)

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func parseClaimsToCtx(ctx context.Context, mapClaims jwt.MapClaims) context.Context {
	// parsing {key:value} from `mapClaims`
	// and set it to the context, so we can use it later in our handler, service and repository
	for k, v := range mapClaims {
		ctx = context.WithValue(ctx, k, v)
	}

	return ctx
}
