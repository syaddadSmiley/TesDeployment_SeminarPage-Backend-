package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	repoz "github.com/syaddadSmiley/SeminarPage/repository"

	"github.com/gin-gonic/gin"
)

func (api *API) CreateProduct(c *gin.Context) {
	api.alloworigin(c)
	if c.Request.Method == "OPTIONS" {
		fmt.Println(c.Request.Method)
		c.Writer.WriteHeader(http.StatusOK)
		return
	}

	if c.Request.Method != "POST" {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Method Not Allowed",
		})
		return
	}
	var products repoz.CreateProductResponse
	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if products.IdProduct.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Jenis invalid input"})
		return
	} else if products.Judul == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Judul tidak boleh kosong"})
		return
	} else if products.Deskripsi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Deskripsi tidak boleh kosong"})
		return
	} else if products.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Lokasi tidak boleh kosong"})
		return
	} else if products.Harga < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Harga invalid input"})
		return
	} else if products.IdJenis < 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Jenis tidak boleh kosong"})
		return
	}

	_, err := api.adminRepo.GetJenisByID(products.IdJenis)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "jenis tidak ditemukan",
		})
		return
	}

	products.IdProduct = repoz.NewID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	_, err = api.adminRepo.PutTask(products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "produk berhasil ditambahkan",
	})
}

func (api *API) UpdateProduct(c *gin.Context) {
	api.alloworigin(c)
	if c.Request.Method == "OPTIONS" {
		fmt.Println(c.Request.Method)
		c.Writer.WriteHeader(http.StatusOK)
		return
	}

	if c.Request.Method != "PUT" {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Method Not Allowed",
		})
		return
	}
	var products repoz.UpdateProductInput
	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if products.IdProduct.String() == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Id tidak boleh kosong"})
		return
	} else if products.Judul == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Judul tidak boleh kosong"})
		return
	} else if products.Deskripsi == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "deskripsi tidak boleh kosong"})
		return
	} else if products.Lokasi == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "lokasi tidak boleh kosong"})
		return
	} else if products.Harga == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "harga tidak boleh kosong"})
		return
	} else if products.IdJenis < 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "jenis tidak boleh kosong"})
		return
	} else if products.Waktu == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "waktu tidak boleh kosong"})
		return
	} else if products.Kapasitas < 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "kapasitas tidak boleh kosong"})
		return
	}

	data, err := api.adminRepo.GetJenisByID(products.IdJenis)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "jenis tidak ditemukan",
		})
		return
	}

	if data.Id == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "jenis tidak ditemukan",
		})
		return
	}

	_, err = api.adminRepo.UpdateTask(products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "produk berhasil diubah",
	})
}

func (api *API) DeleteProduct(c *gin.Context) {
	api.alloworigin(c)
	id := c.Query("id")
	fmt.Println("ID", id)
	_, err := api.adminRepo.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Barang berhasil dihapus",
	})
}

//////////////////////CRUD CATEGORY//////////////////////////////
func (api *API) CreateCategory(c *gin.Context) {
	api.alloworigin(c)
	if c.Request.Method == "OPTIONS" {
		fmt.Println(c.Request.Method)
		c.Writer.WriteHeader(http.StatusOK)
		return
	}

	if c.Request.Method != "POST" {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Method Not Allowed",
		})
		return
	}
	var barang NambahKategori
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(barang)
	if barang.Jenis == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jenis barang tidak boleh kosong"})
		return
	}

	check, err := api.adminRepo.CheckJenis(barang.Jenis)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	if check.Id != 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "category sudah ada",
		})
		return
	}
	_, err = api.adminRepo.CreateCategory(barang.Jenis)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Barang berhasil ditambahkan",
	})
}

func (api *API) UpdateCategory(c *gin.Context) {
	api.alloworigin(c)
	if c.Request.Method == "OPTIONS" {
		fmt.Println(c.Request.Method)
		c.Writer.WriteHeader(http.StatusOK)
		return
	}

	if c.Request.Method != "PUT" {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Method Not Allowed",
		})
		return
	}
	var barang UpdateKategori
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(barang)
	if barang.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id barang tidak boleh kosong"})
		return
	} else if barang.Jenis == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jenis barang tidak boleh kosong"})
		return
	}
	_, err := api.adminRepo.UpdateCategory(barang.Id, barang.Jenis)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Barang berhasil dirubah",
	})
}

func (api *API) DeleteCategory(c *gin.Context) {
	api.alloworigin(c)
	id, _ := strconv.Atoi(c.Query("id"))
	log.Println(id)
	_, err := api.adminRepo.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Barang berhasil dihapus",
	})
}

func (api *API) GetCategory(c *gin.Context) {
	api.alloworigin(c)
	if c.Request.Method == "OPTIONS" {
		fmt.Println(c.Request.Method)
		c.Writer.WriteHeader(http.StatusOK)
		return
	}

	if c.Request.Method != "GET" {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Method Not Allowed",
		})
		return
	}
	var barang repoz.JenisProducts
	fmt.Println(barang)
	data, err := api.adminRepo.GetCategory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Berhasil",
		"data":    data,
	})
}

func (api *API) AddCoupon(c *gin.Context) {
	api.alloworigin(c)
	if c.Request.Method == "OPTIONS" {
		fmt.Println(c.Request.Method)
		c.Writer.WriteHeader(http.StatusOK)
		return
	}

	if c.Request.Method != "POST" {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Method Not Allowed",
		})
		return
	}

	var barang repoz.NambahCoupon
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(barang)
	if barang.Kode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kode kupon tidak boleh kosong"})
		return
	} else if barang.Diskon == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "diskon tidak boleh 0"})
		return
	} else if barang.Minimal < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "minimal tidak boleh 0"})
		return
	}

	check, err := api.adminRepo.CheckCoupon(barang.Kode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	if check.Kode != "" {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "kode kupon sudah ada",
		})
		return
	}
	_, err = api.adminRepo.AddCoupon(barang.Kode, barang.Diskon, barang.Minimal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Coupon berhasil ditambahkan",
	})
}

func (api *API) Notification(c *gin.Context) {
	api.alloworigin(c)
	if c.Request.Method == "OPTIONS" {
		fmt.Println(c.Request.Method)
		c.Writer.WriteHeader(http.StatusOK)
		return
	}

	if c.Request.Method != "POST" {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Method Not Allowed",
		})
		return
	}

	var Notification repoz.Notifikasi
	if err := c.ShouldBindJSON(&Notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(Notification)
	if Notification.Pesan == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pesan tidak boleh kosong"})
		return
	}
	_, err := api.adminRepo.Notification(Notification.Pesan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Notifikasi berhasil dikirim",
	})
}
