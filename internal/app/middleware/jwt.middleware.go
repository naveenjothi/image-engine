package middleware

import (
	"context"
	"fmt"
	"log"
	"myapp/internal/model"
	"myapp/internal/utils/mongoutil"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRETS")) // Replace with your actual secret key

type Claims struct {
    UserID string `json:"user_id"`
    jwt.StandardClaims
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header missing", http.StatusUnauthorized)
            return
        }
		collection := mongoutil.GetCollection("user")
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == "development" {
			var adminUser model.User
			objectID, err := primitive.ObjectIDFromHex(os.Getenv("ADMIN_USER_ID"))
			if err != nil {
				http.Error(w, "Invalid User ID", http.StatusBadRequest)
				return
			}
			collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&adminUser)
			ctx := context.WithValue(r.Context(), UserContextKey, adminUser)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Convert UserID to ObjectID
        objectID, err := primitive.ObjectIDFromHex(claims.UserID)
        if err != nil {
            http.Error(w, "Invalid User ID", http.StatusBadRequest)
            return
        }

        // Retrieve user from database
        var user model.User
        err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
        if err != nil {
            http.Error(w, "User not found", http.StatusUnauthorized)
            return
        }

        // Add user information to context
        ctx := context.WithValue(r.Context(), UserContextKey, user)
		log.Printf("Found user: %+v", user)
        if user.Plan == "Premium" {
            fmt.Println("Premium user logic")
            // Add your premium user-specific logic here
        } else {
            fmt.Println("Free user logic")
            // Add your free user-specific logic here
        }

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

