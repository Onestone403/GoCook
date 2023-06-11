package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var secret = []byte("test")

func verifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(request.Header["Token"][0], claims, func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					writer.WriteHeader(http.StatusUnauthorized)
					_, err := writer.Write([]byte("You're Unauthorized!"))
					if err != nil {
						return nil, err
					}
				}
				return secret, nil

			})
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				log.Printf("Error parsing JWT: %v", err)
				_, err2 := writer.Write([]byte("You're Unauthorized due to error parsing the JWT"))
				if err2 != nil {
					return
				}
				return
			}
			if token.Valid {
				//Populate Context with relevant data
				ctx := context.WithValue(request.Context(), "userID", claims["userID"])
				ctx = context.WithValue(ctx, "path", strings.Split(strings.Trim(request.URL.Path, "/"), "/"))
				ctx = context.WithValue(ctx, "method", request.Method)
				endpointHandler(writer, request.WithContext(ctx))
			} else {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("You're Unauthorized due to invalid token"))
				if err != nil {
					return
				}
			}
		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			_, err := writer.Write([]byte("You're Unauthorized due to No token in the header"))
			if err != nil {
				return
			}
		}

	})
}
