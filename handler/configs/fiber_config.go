package configs

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"microdata/kemendagri/sipd/service/boilerplate_go/utils"
	"os"
	"strconv"
	"time"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		AppName:     os.Getenv("APP_NAME"),
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		//Prefork:       true,
		//CaseSensitive: true,
		//StrictRouting: true,
		ServerHeader: os.Getenv("SERVER_NAME"),
		ErrorHandler: fiberErrorHandler,
	}
}

func fiberErrorHandler(ctx *fiber.Ctx, err error) error {
	var re utils.RequestError
	var vlde validator.ValidationErrors

	switch {
	case errors.As(err, &re):
		_ = ctx.Status(re.Code).JSON(re)
	case errors.As(err, &vlde):
		var eReq = new(utils.RequestError)
		eReq.Code = fiber.StatusInternalServerError

		for _, err := range vlde {
			errObj := utils.DataValidationError{Field: err.Field()}

			// penyesuaian pesan error berdasarkan jenis validasinya
			switch err.Tag() {
			case "gte":
				errObj.Message = fmt.Sprintf("%s harus lebih besar atau sama dengan %s", err.Field(), err.Param())
			case "gt":
				errObj.Message = fmt.Sprintf("%s harus lebih besar dari %s", err.Field(), err.Param())
			case "e164":
				errObj.Message = "Invalid phone number format (E.164)"
			case "alphanumspace":
				errObj.Message = "Only alphanumeric and space allowed"
			case "alphanumslashasterisk":
				errObj.Message = "Only alphanumeric, slash and asterisk allowed"
			default:
				errObj.Message = err.Tag()
			}

			eReq.Fields = append(eReq.Fields, errObj)
		}
		_ = ctx.Status(fiber.StatusUnprocessableEntity).JSON(eReq)
	default:
		_ = ctx.Status(fiber.StatusInternalServerError).JSON(utils.GlobalError{Message: err.Error()})
	}

	// Return from handler
	return nil
}
