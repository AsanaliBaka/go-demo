package auth

import (
	conifgs "go/adv-demo/configs"
	"go/adv-demo/pkg/jwt"
	"go/adv-demo/pkg/req"
	"go/adv-demo/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*AuthService
	*conifgs.Config
	*jwt.JWT
}
type AuthHandler struct {
	*conifgs.Config
	*AuthService
	*jwt.JWT
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())

}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log, err := req.HandleBody[LoginRequest](&w, r)

		if err != nil {
			return
		}

		email, err := handler.AuthService.LoginService(log.Email, log.Password)

		if err != nil {
			res.JsonWriter(w, err.Error(), 404)
			return
		}

		mewtoken, err := jwt.NewJWT(handler.Auth.Secret).Create(jwt.JWTData{
			Email: email,
		})

		if err != nil {
			res.JsonWriter(w, err.Error(), 400)
		}
		data := LoginResponse{
			Token: mewtoken,
		}
		res.JsonWriter(w, data, 200)

	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)

		if err != nil {
			return
		}

		result, err := handler.AuthService.Register(body.Email, body.Password, body.Name)

		if err != nil {
			res.JsonWriter(w, err.Error(), 400)
		}
		res.JsonWriter(w, result, 200)
	}
}
