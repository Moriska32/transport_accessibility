package rest

import (
	"log"
	"time"
	"transport_accessibility_bus/datastore"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

// BasicAuthorizer обёртка над объектом casbin, Enforcer (для того, чтобы объявить собственные методы)
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// CheckPermission Проверка разрешений на использование сервисов
func (a *BasicAuthorizer) CheckPermission(ctx *fiber.Ctx, userID string) bool {
	if userID != "" {
		method := "allow"
		path := string(ctx.Context().URI().Path())
		authed, err := a.enforcer.Enforce(userID, path, method)
		if err != nil {
			return false
		}
		return authed
	}
	return false
}

// CasbinCheckerMiddleware Проверка пользователя для ролевой модели
func CasbinCheckerMiddleware(e *casbin.Enforcer) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ba := &BasicAuthorizer{enforcer: e}
		// userID := auth.GetUserGUID(ctx)
		if !ba.CheckPermission(ctx, "") {
			return ctx.Status(403).JSON(fiber.Map{
				"Error": "You have no permissions to this service. Contact your advisor.",
			})
		}
		return ctx.Next()
	}
}

// InitCasbinPoliciesServices Инициализация ролевой политики по СЕРВИСАМ
func InitCasbinPoliciesServices(e *casbin.Enforcer, dbConn *datastore.DataBase) error {
	// Предварительная очистка модели
	e.ClearPolicy()

	model, err := dbConn.GetAllPolicies()
	if err != nil {
		return err
	}

	for i := range model {
		if model[i].User == nil { // Общая политика 'p'
			_, err := e.AddPolicy(model[i].PolicyRole.RoleName, model[i].Service.MainPath+model[i].Service.Path, "*", model[i].Permission.PermName)
			if err != nil {
				return err
			}
		} else { // Привязка к группам 'g'
			if model[i].ServiceID != nil { // Политика для пользователя на конкретный сервис
				_, err := e.AddPermissionForUser(model[i].User.ID, model[i].Service.MainPath+model[i].Service.Path, "*", model[i].Permission.PermName)
				if err != nil {
					return err
				}
			} else {
				_, err := e.AddRoleForUser(model[i].User.ID, model[i].PolicyRole.RoleName)
				if err != nil {
					return err
				}
			}

		}
	}
	log.Println("Casbin model has been loaded")
	return nil
}

// UpdatePolicyListener Слушаем изменение политики в базе данных
func UpdatePolicyListener(dbConn *datastore.DataBase, e *casbin.Enforcer) {
	ln := dbConn.Listen(dbConn.Context(), "add_new_policy")
	ch := ln.Channel()
	for {
		select {
		case n := <-ch:
			_ = n.Payload
			err := InitCasbinPoliciesServices(e, dbConn)
			if err != nil {
				log.Printf("RBAC was not reloaded due the error: %s\n", err.Error())
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}
