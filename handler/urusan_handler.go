package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"microdata/kemendagri/sipd/service/boilerplate_go/controller"
	"strconv"
	"time"
)

type UrusanHandler struct {
	Controller *controller.UrusanController
	Validate   *validator.Validate
}

func NewUrusanHandler(r fiber.Router, vld *validator.Validate, controller *controller.UrusanController) {
	handler := &UrusanHandler{
		Controller: controller,
		Validate:   vld,
	}

	// strict route
	rStrict := r.Group("urusan")
	rStrict.Get("/", handler.Index)
}

// Index func for Menampilkan list urusan.
//
//	@Summary		Menampilkan List Master Data Urusan.
//	@Description	Menampilkan list master data urusan dengan id daerah yang sama dengan user yang mengaksesnya.
//	@Tags			Urusan
//	@Param			tahun		query	int		true	"Tahun yang ditampilkan"
//	@Param			kode_urusan	query	string	false	"Filter kode urusan (match)"
//	@Param			nama_urusan	query	string	false	"Filter nama urusan (like)"
//	@Param			page		query	int		false	"Halaman yang ditampilkan"
//	@Param			limit		query	int		false	"Jumlah data per halaman, maksimal 50 data"
//	@Produce		json
//	@success		200	{array}		model.UrusanModel	"Success"
//	@Failure		400	{object}	utils.RequestError	"Bad request"
//	@Failure		404	{object}	utils.RequestError	"Data not found"
//	@Failure		422	{object}	utils.RequestError	"Data validation failed"
//	@Failure		500	{object}	utils.RequestError	"Server error"
//	@Security		ApiKeyAuth
//	@Router			/strict/urusan [get]
func (h *UrusanHandler) Index(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return err
	}

	var limit int
	limit, err = strconv.Atoi(c.Query("limit", "20"))
	if err != nil {
		return err
	}

	tahun := c.QueryInt("tahun")
	if tahun <= 0 {
		tahun = time.Now().Year()
	}

	//limit max 50
	if limit > 50 {
		limit = 50
	}

	resp, totalCount, pageCount, err := h.Controller.Index(
		c.Locals("jwt").(*jwt.Token),
		page,
		limit,
		tahun,
		c.Query("kode_urusan", ""),
		c.Query("nama_urusan", ""),
	)
	if err != nil {
		return err
	}

	c.Append("x-pagination-total-count", strconv.Itoa(totalCount))
	c.Append("x-pagination-page-count", strconv.Itoa(pageCount))
	c.Append("x-pagination-page-size", strconv.Itoa(limit))
	if page > 1 {
		c.Append("x-pagination-previous-page", strconv.Itoa(page-1))
	}
	c.Append("x-pagination-current-page", strconv.Itoa(page))
	if page < pageCount {
		c.Append("x-pagination-next-page", strconv.Itoa(page+1))
	}

	return c.JSON(resp)
}
