package routes

import (
	"net/http"	
	authController "github.com/AdairHdz/auth-server/controllers"
	"github.com/rs/cors"
)

func init() {	
	mux := http.NewServeMux()

	mux.HandleFunc("/login", authController.GetToken())	
	mux.HandleFunc("/refresh", authController.RefreshToken())	
	mux.HandleFunc("/logout", authController.SendTokenToBlackList())	
	
	handler := cors.AllowAll().Handler(mux)

	

	http.ListenAndServe("0.0.0.0:50000", handler)
}