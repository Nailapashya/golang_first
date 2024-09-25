package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"microdata/kemendagri/sipd/service/boilerplate_go/controller"
)

type SiteHandler struct {
	Controller *controller.SiteController
	Validate   *validator.Validate
}

func NewSiteHandler(app *fiber.App, vld *validator.Validate, controller *controller.SiteController) {
	handler := &SiteHandler{
		Controller: controller,
		Validate:   vld,
	}

	rSite := app.Group("/site")
	rSite.Get("/", handler.Index)
}

// Index func for get site page.
//
//	@Summary		site page
//	@Description	get site page.
//	@Tags			Site
//	@Accept			json
//	@Produce		json
//	@success		200	{object}	string						"Success"
//	@Failure		400	{object}	utils.RequestError			"Bad request"
//	@Failure		401	{object}	utils.RequestError			"Unauthorized"
//	@Failure		404	{object}	utils.RequestError			"Not found"
//	@Failure		422	{array}		utils.DataValidationError	"Data validation failed"
//	@Failure		500	{object}	utils.RequestError			"Server error"
//	@Router			/site [get]
func (h *SiteHandler) Index(c *fiber.Ctx) error {
	resp := h.Controller.Index()
	return c.JSON(resp)
}
