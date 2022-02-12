package middlewares

import (
	"errors"
	"ms-fetch/domain/constants"
	"ms-fetch/domain/constants/messages"
	"ms-fetch/libraries"
	"ms-fetch/server/http/handlers"
	"ms-fetch/usecase"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/gofiber/fiber/v2"
)

type JWTMiddleware struct {
	*usecase.Contract
}

func NewJWTMiddleware(ucContract *usecase.Contract) JWTMiddleware {
	return JWTMiddleware{ucContract}
}

func (middleware JWTMiddleware) New(ctx *fiber.Ctx) error {

	// validate
	if err := middleware.validate(ctx); err != nil {
		return handlers.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnauthorized)
	}

	return ctx.Next()
}

func (middleware JWTMiddleware) validate(ctx *fiber.Ctx) (err error) {

	// check header
	header := ctx.Get("Authorization")
	if !strings.Contains(header, "Bearer") {
		return errors.New(messages.TokenIsNotProvidedMessage)
	}

	// check token is valid
	token := strings.Replace(header, "Bearer ", "", -1)
	claims, IsValid := libraries.NewJWTLibrary(middleware.JwtSecret).ValidateToken(token)
	if !IsValid {
		return errors.New(messages.TokenIsNotValidMessage)
	}

	// check live time
	if expInt, ok := claims[constants.JWTPayloadTokenLiveTime].(float64); ok {
		now := time.Now().Unix()
		if now > int64(expInt) {
			return errors.New(messages.TokenIsExpiredMessage)
		}
	} else {
		return errors.New(messages.InterfaceConversionErrorMessage)
	}
	// insert payload
	if err = middleware.insertPayload(claims); err != nil {
		return err
	}

	return err
}

func (middleware JWTMiddleware) insertPayload(claims jwt.MapClaims) (err error) {

	// username
	if payloadUsername, ok := claims[constants.JWTPayloadUsername].(string); ok {
		middleware.Contract.Auth.UserName = payloadUsername
	} else {
		return errors.New(messages.InterfaceConversionErrorMessage)
	}

	// phone
	if payloadPhone, ok := claims[constants.JWTPayloadPhone].(string); ok {
		middleware.Contract.Auth.UserPhone = payloadPhone
	} else {
		return errors.New(messages.InterfaceConversionErrorMessage)
	}

	// role
	if payloadRole, ok := claims[constants.JWTPayloadRole].(string); ok {
		middleware.Contract.Auth.UserRole = payloadRole
	} else {
		return errors.New(messages.InterfaceConversionErrorMessage)
	}

	return err
}

func (middleware JWTMiddleware) RoleBiasaOnly(ctx *fiber.Ctx) error {

	// validate
	if err := middleware.validate(ctx); err != nil {
		return handlers.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnauthorized)
	}

	// user must be biasa
	if middleware.Contract.Auth.UserRole != constants.RoleBiasa {
		return handlers.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, messages.AccountIsNotBiasaRole, nil, http.StatusForbidden)
	}

	return ctx.Next()
}

func (middleware JWTMiddleware) RoleAdminOnly(ctx *fiber.Ctx) error {

	// validate
	if err := middleware.validate(ctx); err != nil {
		return handlers.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnauthorized)
	}

	// user must be admin
	if middleware.Contract.Auth.UserRole != constants.RoleAdmin {
		return handlers.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, messages.AccountIsNotAdminRole, nil, http.StatusForbidden)
	}

	return ctx.Next()
}

func (middleware JWTMiddleware) RoleBebas(ctx *fiber.Ctx) error {

	// validate
	if err := middleware.validate(ctx); err != nil {
		return handlers.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnauthorized)
	}

	return ctx.Next()
}
