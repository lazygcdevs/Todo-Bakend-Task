package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieName := os.Getenv("COOKIE_NAME")
		if cookieName == "" {
			cookieName = "todo_user_id"
		}

		userID, err := c.Cookie(cookieName)

		// If no cookie exists or cookie is invalid, generate a new user ID
		if err != nil || userID == "" {
			userID = uuid.New().String()

			// Set cookie with 24-hour expiration
			c.SetCookie(
				cookieName, // name
				userID,     // value
				24*60*60,   // max age in seconds (24 hours)
				"/",        // path
				"",         // domain
				false,      // secure (set to true in production with HTTPS)
				true,       // httpOnly
			)
		}

		// Add user ID to the context
		c.Set("user_id", userID)
		c.Next()
	}
}
