package auth

import (
	"encoding/json"
	"log"
	"time"
	"transport_accessibility_bus/configuration"
	"transport_accessibility_bus/datastore"
	"transport_accessibility_bus/env"

	jwt "github.com/LdDl/fiber-jwt/v2"
	"github.com/gofiber/fiber/v2"
)

// login Аутентификационные данные
// swagger:model
type login struct {
	// Логин пользователя
	Username string `form:"username" json:"username" binding:"required"`
	// Пароль пользователя
	Password string `form:"password" json:"password" binding:"required"`
}

// SuccessAuthentication Успешная аутентификация
// swagger:model
type SuccessAuthentication struct {
	// Код ответа сервера
	Code int `json:"code" example:"200"`
	// Срок истечения действия токена
	Expire time.Time `json:"exp" example:"2019-01-23T15:33:21+03:00"`
	// Строка токена
	Token string `json:"token" example:"uIjoibW9ydGFsIeyJhbGciOiJIUzI1"`
}

// InitAuth Middware для аутентификации
// @Summary Аутентификация
// @Tags Аутентификация
// @Produce json
// @Param username_password body auth.login true "Введите Ваш логин и пароль от учётной записи"
// @Success 200 {object} auth.SuccessAuthentication
// @Failure 401 {object} codes.Error401
// @Failure 403 {object} codes.Error403
// @Failure 500 {object} codes.Error500
// @Failure 503 {object} codes.Error503
// @Router /tula/doauth [POST]
func InitAuth(dbConn *datastore.DataBase, cfg *configuration.JWTKeys) *jwt.FiberJWTMiddleware {
	var identityKey = "login"
	authMiddleware, err := jwt.New(&jwt.FiberJWTMiddleware{
		Realm:            "Moscow",
		Key:              []byte("transport-accessibility=2021-01-01"),
		Timeout:          time.Hour * 7 * 24,
		MaxRefresh:       time.Hour * 7 * 24,
		IdentityKey:      identityKey,
		SigningAlgorithm: "RS512",
		PubKeyFile:       cfg.PubFile,
		PrivKeyFile:      cfg.KeyFile,
		PayloadFunc: func(userId interface{}) jwt.MapClaims {
			user, err := dbConn.GetAData(userId.(string))
			if err != nil {
				return nil
			}
			return jwt.MapClaims{
				"login":      userId.(string),
				"role":       user.RoleID,
				"guid":       user.GUID,
				"changepass": user.ChangePassword,
			}
		},
		IdentityHandler: func(c *fiber.Ctx) interface{} {
			claims := jwt.ExtractClaims(c)
			return &datastore.AData{
				Login:          claims["login"].(string),
				RoleID:         claims["role"].(string),
				GUID:           claims["guid"].(string),
				ChangePassword: claims["changepass"].(bool),
			}
		},
		Authenticator: func(ctx *fiber.Ctx) (interface{}, error) {
			loginVals := login{}
			bodyBytes := ctx.Context().PostBody()
			if err := json.Unmarshal(bodyBytes, &loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			user, err := dbConn.GetADataPwd(userID, password)
			if err != nil {
				return userID, jwt.ErrFailedAuthentication
			}
			if user.Access == "Authentication" || user.Access == "Authorization" {
				return userID, nil
			}
			return userID, jwt.ErrFailedAuthentication
		},
		Authorizator: func(userId interface{}, ctx *fiber.Ctx) bool {
			user, err := dbConn.GetAData(userId.(*datastore.AData).Login)
			if err != nil {
				return false
			}
			if user.Access == "Authorization" {
				return true
			}
			return false
		},
		Unauthorized: func(ctx *fiber.Ctx, code int, message string) error {
			if message == jwt.ErrFailedAuthentication.Error() {
				return ctx.Status(401).JSON(fiber.Map{"Error": string(ctx.Context().URI().Path()) + ";Unauthorized"})
			}
			return ctx.Status(403).JSON(fiber.Map{"Error": string(ctx.Context().URI().Path()) + message})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Println("Can not init JWT auth", err)
	}
	return authMiddleware
}

// MyRole Роль текущего пользователя
// swagger:model
type MyRole struct {
	// Роль пользователя: admin / user
	Role string `json:"role" example:"operator"`
}

// GetMyRole Роль текущего пользователя
// @Summary Роль текущего пользователя
// @Tags Аутентификация
// @Security AuthJWT
// @Produce json
// @Success 200 {object} auth.MyRole
// @Failure 401 {object} codes.Error401
// @Failure 403 {object} codes.Error403
// @Failure 500 {object} codes.Error500
// @Failure 503 {object} codes.Error503
// @Router /tula/whoami [GET]
func GetMyRole(envdb *env.Env) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		dbConn := datastore.DataBase{DB: envdb.DBs["its_db_name"]}
		userID := GetUserGUID(ctx)
		model, err := dbConn.GetRole(userID)
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"Error": string(ctx.Context().RequestURI()) + ";Can't select role for current user",
			})
		}

		if model.PolicyRole == nil {
			return ctx.Status(500).JSON(fiber.Map{
				"Error": string(ctx.Context().RequestURI()) + ";nil role",
			})
		}

		ans := MyRole{
			Role: model.PolicyRole.RoleName,
		}
		return ctx.Status(200).JSON(ans)
	}
}

// GetUserGUID Получение GUID'а пользователя
func GetUserGUID(ctx *fiber.Ctx) string {
	claims := jwt.ExtractClaims(ctx)
	if _, ok := claims["guid"]; !ok {
		return ""
	}
	return claims["guid"].(string)
}
