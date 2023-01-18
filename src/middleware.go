package main

import (
	sheeters "Joe/sheeter/pkg/general"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func validateJWT(next http.Handler) http.Handler{
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    cookie, err := r.Cookie("sheeter_token")
    if err != nil{
      if err == http.ErrNoCookie{
        http.Redirect(w, r, "/login", http.StatusUnauthorized)
        return
      }

      http.Redirect(w, r, "/login", http.StatusBadRequest)
      return
    }

    tknString := cookie.Value
    claims := &sheeters.Claims{}

    tkn, err := jwt.ParseWithClaims(tknString, claims, func(token *jwt.Token) (any, error){
      return []byte(JWTKEY), nil 
    })

    if err != nil || !tkn.Valid{
      http.Redirect(w, r, "/login", http.StatusBadRequest)
      return
    }

    next.ServeHTTP(w, r)
  })
}
