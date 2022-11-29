package middlewares

import (
	"deliveroo/db"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

func AuthenUserLoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, resquest *http.Request) {
		username := resquest.Header.Get("Username")
		// salt := resquest.Header.Get("Salt")
		authenInfo := resquest.Header.Get("Password")

		log.Info("Middleware: Authen user login")
		//get user hash password from db
		passwordHash, err := db.GetUserHashPassword(username)
		if err != nil {
			log.Error(err)
			response.WriteHeader(http.StatusInternalServerError)
			return
		}
		if passwordHash == "" {
			log.Error("can not get password of username :", username)
			response.WriteHeader(http.StatusBadRequest)
			return
		}
		// mac := hmac.New(sha256.New, []byte(salt))
		// mac.Write([]byte(passwordHash))
		// hmacAuthen := hex.EncodeToString(mac.Sum(nil))
		if strings.Compare(authenInfo, passwordHash) != 0 {
			log.Info("Authen failed , authen info not valid ")
			response.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(response, "Username or password is incorrect")
			return
		}

		log.Info("Middleware: Authen pass, serve next")

		next.ServeHTTP(response, resquest)
	})
}

//AuthenTokenDev check access token of dev's api
func AuthenTokenDev(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ah := r.Header.Get("accessToken"); ah == "" {
			log.Error("Missing token authen info ")
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			token, err := jwt.Parse(ah, func(t *jwt.Token) (interface{}, error) {

				signkey := "weriwoxcr342f234"
				// if err != nil {
				// 	log.Error("get signkey in JWT authen error ")
				// 	return nil, err
				// }
				return []byte(signkey), nil
			})

			if err == nil && token.Valid {
				userInfo, ok := token.Claims.(jwt.MapClaims)
				if !ok {
					log.Error("cannot parse token claim to map claim ")
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				userName, ok := userInfo["iss"].(string)
				if !ok {
					log.Error("cannot parse contact claim ")
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				log.Info(" get request from user :", userName)
				// var DevAuthInfo models.DevAuthInfo
				// DevAuthInfo.User = userName
				// ctx := context.WithValue(r.Context(), models.ContextKeyDevAuthInfo, DevAuthInfo)
				// r = r.WithContext(ctx)
			} else {
				log.Info("Token invalid ")
				if err != nil {
					log.Error(err)
				}
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
