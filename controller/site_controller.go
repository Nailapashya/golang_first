package controller

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type SiteController struct {
	contextTimeout time.Duration
	pgxConn        *pgxpool.Pool
}

func NewSiteController(conn *pgxpool.Pool, timeout time.Duration) (controller *SiteController) {
	controller = &SiteController{
		pgxConn:        conn,
		contextTimeout: timeout,
	}

	return
}

func (c *SiteController) Index() string {
	return "site"
}
