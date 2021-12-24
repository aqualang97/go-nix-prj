package main

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

const AccessSecret = "access_secret"
const RefreshSecret = "refresh_secret"

const AccessTokenLifetimeMinutes = 10
const RefreshTokenLifetimeMinutes = 60

func main() {
	//реализуем флоу логина юзера
	// юзер дает логин пароль
	// получаем ответ верный илинет юзер

	http.HandleFunc("/login", Login)     //умеем обрабатывать логин с помощью ф-ции логин
	http.HandleFunc("/profile", Profile) //умеем обрабатывать логин с помощью ф-ции логин
	http.HandleFunc("/refresh", Refresh)

	log.Fatal(http.ListenAndServe(":8080", nil)) //слушаем порт 8080 для входящих запросов
}
func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := new(LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil { //берем тенло запроса декодим и декодим в тело запроса
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := NewUserRepository().GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		tokenString, err := GenerateToken(user.ID, AccessTokenLifetimeMinutes, AccessSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		refreshString, err := GenerateToken(user.ID, RefreshTokenLifetimeMinutes, RefreshSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := LoginResponse{
			AccessToken:  tokenString,
			RefreshToken: refreshString,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func Profile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tokenString := GetTokenFromBearerString(r.Header.Get("Authorization"))
		claims, err := ValidateToken(tokenString, AccessSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		user, err := NewUserRepository().GetUserByID(claims.ID)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		resp := UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
func Refresh(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		resp := LoginResponse{
			AccessToken:  tokenString,
			RefreshToken: refreshString,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
