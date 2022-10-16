package service

import (
	"errors"
	"gofiber-penitipan-barang/application/config"
	"gofiber-penitipan-barang/application/database"
	"gofiber-penitipan-barang/application/middleware"
	"gofiber-penitipan-barang/application/model"
	"gofiber-penitipan-barang/application/request"
	"gofiber-penitipan-barang/application/response"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	DB       *gorm.DB
	validate *validator.Validate
}

func NewAuthService(DB *gorm.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		DB:       DB,
		validate: validate,
	}
}

func (service *AuthServiceImpl) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (service *AuthServiceImpl) getUserByUsername(username string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Username: username}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (service *AuthServiceImpl) Login(c *fiber.Ctx) response.LoginResponse {
	// input := new(LoginInput)
	// var ud UserData
	var userData model.User
	loginRequest := request.LoginRequest{}

	err := c.BodyParser(&loginRequest)

	if err != nil {
		panic(c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err}))
	}

	user, err := service.getUserByUsername(loginRequest.Username)
	if err != nil {
		panic(c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err}))
	}
	if user != nil {
		userData = model.User{
			Id:       user.Id,
			Name:     user.Name,
			Username: user.Username,
			Password: user.Password,
			Role:     user.Role,
		}
	}

	if !service.CheckPasswordHash(loginRequest.Password, user.Password) {
		panic(c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil}))
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["user_id"] = userData.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		panic(c.SendStatus(fiber.StatusInternalServerError))
	}

	loginResponse := response.LoginResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
		Role:     user.Role,
		Token:    tokenString,
	}
	return loginResponse
}

func (service *AuthServiceImpl) User(c *fiber.Ctx) response.LoginResponse {
	user := middleware.AuthenticationGetUser

	userResponse := response.LoginResponse{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Role:     user.Role,
	}

	return userResponse
}
