package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(parseAuthorizedUsersFromEnv())
}

func parseAuthorizedUsersFromEnv() gin.Accounts {
	authorizedUsers := gin.Accounts{}

	// ENV should be: "username1:password1,username2:password2"
	usersEnv := os.Getenv("AUTHORIZED_USERS")
	usersList := strings.Split(usersEnv, ",")
	for _, user := range usersList {
		parts := strings.Split(user, ":")
		if len(parts) == 2 {
			authorizedUsers[parts[0]] = parts[1]
		}
	}

	return authorizedUsers
}
