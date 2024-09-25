package controller

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"math"
	"microdata/kemendagri/sipd/service/boilerplate_go/model"
	"microdata/kemendagri/sipd/service/boilerplate_go/utils"
	"net/http"
	"time"
)

type UrusanController struct {
	contextTimeout time.Duration
	pgxConn        *pgxpool.Pool
}

func NewUrusanController(conn *pgxpool.Pool, timeout time.Duration) (controller *UrusanController) {
	controller = &UrusanController{
		pgxConn:        conn,
		contextTimeout: timeout,
	}

	return
}

func (c *UrusanController) Index(
	user *jwt.Token,
	page,
	limit,
	tahun int,
	kodeUrusan,
	namaUrusan string,
) (r []model.UrusanModel, totalCount, pageCount int, err error) {
	var q string
	r = make([]model.UrusanModel, 0)

	claims := user.Claims.(jwt.MapClaims)
	// sub := claims["sub"].(string)
	// tahun := int64(claims["tahun"].(float64))
	idDaerah := int64(claims["id_daerah"].(float64))

	pageCount = 1
	offset := limit * (page - 1)

	var filterCond string

	if kodeUrusan != "" {
		filterCond += " AND kode_urusan = " + kodeUrusan
	}
	if namaUrusan != "" {
		filterCond += " AND nama_urusan ilike %" + namaUrusan + "%"
	}

	// hitung total data
	q = `SELECT count(*) FROM public.r_urusan u
		WHERE
		    tahun=$1 AND id_daerah=$2` + filterCond
	err = c.pgxConn.QueryRow(context.Background(), q, tahun, idDaerah).Scan(&totalCount)
	if err != nil {
		err = utils.RequestError{
			Code:    http.StatusInternalServerError,
			Message: "gagal menghitung total data - " + err.Error(),
		}
		return
	}

	// ambil data
	q = `SELECT
			id_urusan,
			tahun,
			id_daerah,
			kode_urusan,
			nama_urusan,
			id_unik,
			is_locked
		FROM public.r_urusan u
		WHERE
			tahun = $1
			AND id_daerah = $2
			AND is_locked = 0` + filterCond
	q += `ORDER BY id_urusan LIMIT $3 OFFSET $4`

	var rows pgx.Rows
	rows, err = c.pgxConn.Query(
		context.Background(),
		q,
		tahun,
		idDaerah,
		limit,
		offset,
	)
	if err != nil {
		err = utils.RequestError{
			Code:    http.StatusInternalServerError,
			Message: "gagal mengambil data - " + err.Error(),
		}
		return
	}
	for rows.Next() {
		m := model.UrusanModel{}
		err = rows.Scan(
			&m.IdUrusan,
			&m.Tahun,
			&m.IdDaerah,
			&m.KodeUrusan,
			&m.NamaUrusan,
			&m.IdUnik,
			&m.IsLocked,
		)

		r = append(r, m)
	}
	if rows.Err() != nil {
		err = utils.RequestError{
			Code:    http.StatusInternalServerError,
			Message: "gagal mengambil data (rows) - " + rows.Err().Error(),
		}
		return
	}

	defer rows.Close()

	if totalCount > 0 && totalCount > limit {
		pageCount = int(math.Ceil(float64(totalCount) / float64(limit)))
	}

	return
}
