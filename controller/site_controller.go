package controller

import (
	"database/sql" // Misalkan Anda menggunakan database SQL
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// SiteController struct
type SiteController struct {
	db             *sql.DB // Tambahkan field untuk koneksi DB
	contextTimeout time.Duration
}

// NewSiteController creates a new instance of SiteController
func NewSiteController(db *sql.DB, timeout time.Duration) *SiteController {
	return &SiteController{
		db:             db,
		contextTimeout: timeout,
	}
}

func (c *SiteController) Index(ctx *gin.Context) {
	birthYear := 1997
	umur := calculateAge(birthYear)

	// Query untuk mengambil data pengguna dari database
	var name, email, addressHome, addressNow, phone, gitlab string
	err := c.db.QueryRow("SELECT name, email, address_home, address_now, phone, gitlab FROM users WHERE id = $1", 1).Scan(&name, &email, &addressHome, &addressNow, &phone, &gitlab) // Misalkan ID pengguna yang ingin diambil adalah 1
	if err != nil {
		log.Println("Error querying database:", err)
		ctx.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	data := map[string]interface{}{
		"1. Nama":           name,
		"2. Umur":           umur,
		"3. Email":          email,
		"4. Alamat Rumah":   addressHome,
		"5. Alamat Sekarang": addressNow,
		"6. Nomor Telepon":  phone,
		"7. Gitlab":         gitlab,
	}

	ctx.JSON(200, data)
}
