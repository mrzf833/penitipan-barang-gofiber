package middleware

import (
	"gofiber-penitipan-barang/application/config"
	"gofiber-penitipan-barang/application/database"
	"gofiber-penitipan-barang/application/model"
	"gofiber-penitipan-barang/application/response"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var AuthenticationGetUser response.UserAuthentication

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		regexToken, _ := regexp.Compile(`^Bearer\s(\S+)$`)

		tokenBearer := c.Get("Authorization")

		tokens := regexToken.FindStringSubmatch(tokenBearer)
		if len(tokens) > 1 {
			token := tokens[1]

			claims := &config.Claims{}

			_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
				return config.JwtKey, nil
			})
			if err != nil {
				panic(fiber.NewError(500, err.Error()))
			}

			var user model.User
			database.DB.First(&user, "username", claims.Username)
			if user.Id == 0 {
				panic(fiber.NewError(500, "data error in auth.go line 40"))
			}

			AuthenticationGetUser = response.UserAuthentication{
				Id:       user.Id,
				Name:     user.Name,
				Username: user.Username,
				Role:     user.Role,
			}
			return c.Next()
		} else {
			return fiber.ErrUnauthorized
		}
	}
}

// func Protected() fiber.Handler {
// 	return jwtware.New(jwtware.Config{
// 		SigningKey:   []byte(config.Config("SECRET")),
// 		ErrorHandler: jwtError,
// 	})
// }

// func jwtError(c *fiber.Ctx, err error) error {
// 	if err.Error() == "Missing or malformed JWT" {
// 		return c.Status(fiber.StatusBadRequest).
// 			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
// 	}
// 	return c.Status(fiber.StatusUnauthorized).
// 		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
// }
