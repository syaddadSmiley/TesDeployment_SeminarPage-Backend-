package repository

import (
	"time"

	"github.com/google/uuid"
)

//ID entity ID
type ID = uuid.UUID

//NewID create a new entity ID
func NewID() ID {
	return ID(uuid.New())
}

//StringToID convert a string to an entity ID
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}

func ZeroID() string {
	return "00000000-0000-0000-0000-000000000000"
}

type User struct {
	Id       ID     `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Gambar   []byte `json:"gambar"`
}

type UserRequest struct {
	Id       ID     `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Gambar   string `json:"gambar"`
}

type UserResponse struct {
	Id       ID     `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Mail     string `json:"mail"`
	Password string `json:"-"`
	Role     string `json:"role"`
	Gambar   string `json:"gambar"`
}

type Task struct {
	Id        ID     `json:"id"`
	Jenis     string `json:"jenisProducts"`
	Gambar    string `json:"gambar"`
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
	Lokasi    string `json:"lokasi"`
	Harga     int    `json:"harga"`
	Waktu     string `json:"waktu"`
	Kapasitas int    `json:"kapasitas"`
}

type TaskResponse struct {
	Id        ID     `json:"id"`
	IdJenis   int    `json:"id_jenis"`
	Gambar    string `json:"gambar"`
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
	Lokasi    string `json:"lokasi"`
	Harga     int    `json:"harga"`
}

type JenisProducts struct {
	Id    int    `json:"id"`
	Jenis string `json:"jenis"`
}

type RegisterRequest struct {
	Id       ID     `json:"-"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Gambar   []byte `json:"gambar"`
}

type DeleteUserReqByUsername struct {
	Username string `json:"username"`
}

type Komentar struct {
	Id         ID        `json:"id"`
	Id_product ID        `db:"id_product"`
	Id_user    ID        `db:"id_user"`
	Username   string    `db:"username"`
	Content    string    `db:"content"`
	Like       int       `db:"like"`
	Dislike    int       `db:"dislike"`
	Rating     int       `db:"rating"`
	CreatedAt  time.Time `db:"created_at"`
}

type KomentarRequest struct {
	Id_product ID     `json:"id_product"`
	Username   string `json:"username"`
	Content    string `json:"content"`
}

type Wishlist struct {
	Id         ID   `json:"id"`
	Id_user    ID   `db:"id_user"`
	Id_product ID   `db:"id_product"`
	Status     bool `json:"status"`
}

type WishlistRequest struct {
	Id_user    ID `json:"id_user"`
	Id_product ID `json:"id_product"`
}

type WishlistResponse struct {
	Id       ID     `json:"id"`
	Username string `json:"username"`
	Product  string `json:"product"`
}

type NambahCoupon struct {
	Kode    string `json:"kode"`
	Diskon  int    `json:"diskon"`
	Minimal int    `json:"minimal"`
}

type UseCoupon struct {
	Id     ID     `json:"-"`
	IdUser ID     `json:"-"`
	Kode   string `json:"kode"`
	Status string `json:"status"`
}

type Notifikasi struct {
	Id    ID     `json:"id"`
	Pesan string `json:"pesan"`
}

type Basket struct {
	Id_user    ID `json:"id_user"`
	Id_product ID `json:"id_product"`
}

type BasketResponse struct {
	Id         ID `json:"id"`
	Id_user    ID `json:"username"`
	Id_product ID `json:"product"`
}

type CreateProductResponse struct {
	IdProduct ID     `json:"-"`
	IdJenis   int    `json:"id_jenis"`
	Gambar    string `json:"gambar"`
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
	Lokasi    string `json:"lokasi"`
	Harga     int    `json:"harga"`
	Waktu     string `json:"waktu"`
	Kapasitas int    `json:"kapasitas"`
}

type UpdateProductInput struct {
	IdProduct ID     `json:"id"`
	IdJenis   int    `json:"id_jenis"`
	Gambar    string `json:"gambar"`
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
	Lokasi    string `json:"lokasi"`
	Harga     int    `json:"harga"`
	Waktu     string `json:"waktu"`
	Kapasitas int    `json:"kapasitas"`
}
