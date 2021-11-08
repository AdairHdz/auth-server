package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/AdairHdz/auth-server/database"
	"github.com/AdairHdz/auth-server/entity"
	"github.com/AdairHdz/auth-server/helpers"
	"github.com/AdairHdz/auth-server/request"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
)

func GetToken() func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var loginInfo request.LoginInfo
		err := decoder.Decode(&loginInfo)

		if err != nil {			
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"error": "Bad input",
				"message": "Please make sure the data you've introduced has a valid format and try again",
			})
			return
		}

		var user entity.User
		err = database.Query(loginInfo.EmailAddress, &user)
		if err != nil {			
			rw.WriteHeader(http.StatusConflict)
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"error": "Email address not registered",
				"message": "The email address you provided does not match any of our records",
			})
			return
		}		

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
		if err != nil {			
			rw.WriteHeader(http.StatusForbidden)
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"error": "Password mismatch",
				"message": "The password you entered does not match with our database records. Please make sure the data is correct and try again",
			})
			return
		}

		if !user.Verified {			
			rw.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"error": "Non-verified account",
				"message": "It looks like you have not verified your account. Please check your email inbox",
			})
			return
		}

		token, err := helpers.SignString(user.ID, user.UserType, loginInfo.EmailAddress, time.Now().Add(15 * time.Minute))
		if err != nil {			
			rw.WriteHeader(http.StatusConflict)
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"error": "Internal error",
				"message": "There was an error while processing your request. Please try again later",
			})
			return
		}

		user.Token = token
		
		jwtTokenCookie := http.Cookie{
			Name: "jwt-token",
			Value: token,
			HttpOnly: true,
			Expires: time.Now().Add(time.Minute * 15),
		}		

		refreshToken, err := helpers.SignString(user.ID, user.UserType, loginInfo.EmailAddress, time.Now().Add(24 * time.Hour))
		if err != nil {			
			rw.WriteHeader(http.StatusConflict)
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"error": "Internal error",
				"message": "There was an error while processing your request. Please try again later",
			})
			return
		}

		refreshTokenCookie := http.Cookie {
			Name: "refresh-token",
			Value: refreshToken,
			HttpOnly: true,
			Expires: time.Now().Add(time.Hour * 24),
		}

		http.SetCookie(rw, &jwtTokenCookie)
		http.SetCookie(rw, &refreshTokenCookie)
		
		json.NewEncoder(rw).Encode(&user)
	}
}

func RefreshToken() func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		isValidToken := helpers.ValidateSignedString(token)
		
		if isValidToken {

			_, err := helpers.Get(token)

			if err == redis.Nil {
				
				newSignedString, err := helpers.SignString("userID", 0, "userEmailAddress", time.Now().Add(time.Minute * 15))

				if err != nil {
					panic("err")
				}

				json.NewEncoder(rw).Encode(map[string]interface{}{
					"token": newSignedString,
				})
				return
			}

			if err != nil {
				log.Println(err.Error())
				return
			}			
		}
	}
}

func SendTokenToBlackList() func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		err := helpers.Save(token, "")

		if err != nil {
			log.Println(err.Error())
		}
	}
}