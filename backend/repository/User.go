package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) LoginUser(username string) (*UserRequest, error) {
	sqlStatement := `SELECT * FROM user WHERE Username = ?;`

	rows, err := u.db.Query(sqlStatement, username)
	if err != nil {
		return nil, err
	}

	var user UserRequest
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Nama, &user.Username, &user.Mail, &user.Password, &user.Role, &user.Gambar)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (u *UserRepo) RegisterUser(regis RegisterRequest) (*User, error) {
	stmt := `INSERT INTO user (id, username, password, mail, nama ,role, gambar) 
	VALUES (?, ?, ?, ?, ?, ?, ?);`

	rows, err := u.db.Query(stmt, regis.Id, regis.Username, regis.Password, regis.Mail, regis.Nama, regis.Role, regis.Gambar)
	if err != nil {
		return nil, err
	}

	var user User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Nama, &user.Username, &user.Mail, &user.Password, &user.Role, &user.Gambar)
		if err != nil {
			return nil, err
		}
	}

	user.Id = regis.Id
	user.Nama = regis.Nama
	user.Username = regis.Username
	user.Mail = regis.Mail
	user.Password = regis.Password
	user.Role = regis.Role
	user.Gambar = regis.Gambar

	return &user, nil
}

func (u *UserRepo) UpdateUserByID(update UserRequest) (int64, error) {
	var sqlStatement string
	if update.Gambar == "" {
		sqlStatement = `UPDATE user SET nama = ?, username = ? WHERE id = ?;`
	} else {
		sqlStatement = `UPDATE user SET nama = ?, username = ?, gambar = ? WHERE id = ?;`
	}
	if update.Gambar != "" {
		byteGambar := []byte(update.Gambar)
		result, err := u.db.Exec(sqlStatement, update.Nama, update.Username, byteGambar, update.Id)
		if err != nil {
			return 0, err
		}
		return result.RowsAffected()
	} else {
		result, err := u.db.Exec(sqlStatement, update.Nama, update.Username, update.Id)
		if err != nil {
			return 0, err
		}
		return result.RowsAffected()
	}

}

func (u *UserRepo) DeleteUserByID(id ID) (int64, error) {
	sqlStatement := `DELETE FROM user WHERE id = ?;`

	result, err := u.db.Exec(sqlStatement, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (u *UserRepo) CheckAccount(username, mail string) (*User, error) {
	sqlStatement := `SELECT * FROM user WHERE username = ? or mail = ?;`

	rows, err := u.db.Query(sqlStatement, username, mail)
	if err != nil {
		return nil, err
	}

	var user User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Nama, &user.Username, &user.Mail, &user.Password, &user.Role, &user.Gambar)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (u *UserRepo) GetProfile(id ID) (*UserResponse, error) {
	sqlStatement := `SELECT * FROM user WHERE id = ?;`

	rows, err := u.db.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}

	var user UserResponse
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Nama, &user.Username, &user.Mail, &user.Password, &user.Role, &user.Gambar)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (u *UserRepo) Allproduk(limit int, offset int) ([]Task, error) {
	sqlStmt := `
	SELECT
		g.Id AS id,
		u.jenis AS jenis,
		g.gambar AS gambar,
		g.judul AS judul,
		g.deskripsi AS deskripsi,
		g.lokasi AS lokasi,
		g.harga AS harga
	FROM
		products AS g
	JOIN JenisProducts AS u ON (u.id = g.id_jenis)
	LIMIT ?
	OFFSET ?
	`

	rows, err := u.db.Query(sqlStmt, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	barang := []Task{}
	for rows.Next() {
		var barangs Task
		err := rows.Scan(
			&barangs.Id,
			&barangs.Jenis,
			&barangs.Gambar,
			&barangs.Judul,
			&barangs.Deskripsi,
			&barangs.Lokasi,
			&barangs.Harga,
		)
		if err != nil {
			return nil, err
		}
		barang = append(barang, barangs)
	}
	return barang, nil
}

func (u *UserRepo) GetRowProducts() (int, error) {
	sqlStmt := `SELECT COUNT(*) FROM products`
	var total int
	err := u.db.QueryRow(sqlStmt).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

func (u *UserRepo) DeleteUser(username string) error {
	_, err := u.db.Exec("DELETE FROM user WHERE username = ?", username)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) SortingAsc() ([]Task, error) {
	sqlStmt := `
	SELECT
		g.Id AS id,
		u.jenis AS jenis,
		g.gambar AS gambar,
		g.judul AS judul,
		g.deskripsi AS deskripsi,
		g.lokasi AS lokasi,
		g.harga AS harga
	FROM
		products AS g
	JOIN JenisProducts AS u ON (u.id = g.id_jenis)
	ORDER BY g.harga ASC
	`

	rows, err := u.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	barang := []Task{}
	for rows.Next() {
		var barangs Task
		err := rows.Scan(
			&barangs.Id,
			&barangs.Jenis,
			&barangs.Gambar,
			&barangs.Judul,
			&barangs.Deskripsi,
			&barangs.Lokasi,
			&barangs.Harga,
		)
		if err != nil {
			return nil, err
		}
		barang = append(barang, barangs)
	}
	return barang, nil
}

func (u *UserRepo) SortingDsc() ([]Task, error) {
	sqlStmt := `
	SELECT
		g.Id AS id,
		u.jenis AS jenis,
		g.gambar AS gambar,
		g.judul AS judul,
		g.deskripsi AS deskripsi,
		g.lokasi AS lokasi,
		g.harga AS harga
	FROM
		products AS g
	JOIN JenisProducts AS u ON (u.id = g.id_jenis)
	ORDER BY g.harga DESC
	`

	rows, err := u.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	barang := []Task{}
	for rows.Next() {
		var barangs Task
		err := rows.Scan(
			&barangs.Id,
			&barangs.Jenis,
			&barangs.Gambar,
			&barangs.Judul,
			&barangs.Deskripsi,
			&barangs.Lokasi,
			&barangs.Harga,
		)
		if err != nil {
			return nil, err
		}
		barang = append(barang, barangs)
	}
	return barang, nil
}

func (u *UserRepo) CreateKomentar(Id ID, IdProduct ID, IdUser ID, Content string, Rating int) (int64, error) {
	sqlStatement := `INSERT INTO komentar (id, id_product, id_user, content, like, dislike, rating, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	stmt, err := u.db.Prepare(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(Id, IdProduct, IdUser, Content, 0, 0, Rating, time.Now().UTC())
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	return result.LastInsertId()

}

func (u *UserRepo) DeleteKomentar(id ID) error {
	_, err := u.db.Exec("DELETE FROM komentar WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) CheckKomentar(id ID) (*Komentar, error) {
	rows, err := u.db.Query("SELECT * FROM komentar WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	var komentar Komentar
	for rows.Next() {
		err := rows.Scan(
			&komentar.Id,
			&komentar.Id_product,
			&komentar.Id_user,
			&komentar.Content,
			&komentar.Like,
			&komentar.Dislike,
			&komentar.Rating,
			&komentar.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
	}
	return &komentar, nil
}

func (u *UserRepo) GetProductByID(id ID) ([]Task, error) {
	sqlStmt := `SELECT
		products.Id,
		JenisProducts.Jenis AS JenisProducts,
		products.gambar,
		products.Judul,
		products.deskripsi,
		products.lokasi,
		products.harga
	FROM products
	INNER JOIN JenisProducts
	ON products.Id_jenis = JenisProducts.Id
	WHERE products.Id = ?`

	row := u.db.QueryRow(sqlStmt, id)
	if row == nil {
		return nil, errors.New("Product not found")
	}
	var barang Task
	err := row.Scan(&barang.Id, &barang.Jenis, &barang.Gambar, &barang.Judul, &barang.Deskripsi, &barang.Lokasi, &barang.Harga)
	if err != nil {
		return nil, err
	}
	return []Task{barang}, nil
}

func (u *UserRepo) FetchWishlist() ([]Wishlist, error) {
	sqlStmt := `
	SELECT w.Id_user AS id_user,w.Id_product AS id_product FROM wishlist AS w INNER JOIN TaskResponse p ON w.Id_product = p.Id`

	rows, err := u.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	Wish := []Wishlist{}
	for rows.Next() {
		var Wishs Wishlist
		err := rows.Scan(
			&Wishs.Id_user,
			&Wishs.Id_product,
		)
		if err != nil {
			return nil, err
		}
		Wish = append(Wish, Wishs)
	}
	return Wish, nil
}

func (u *UserRepo) InsertWishlist(Id ID, Id_user ID, Id_product ID) (string, error) {
	sqlStatement := `INSERT INTO wishlist_item (id, id_user, id_product) VALUES (?, ?, ?) RETURNING id;`

	LastInsertId := ""
	err := u.db.QueryRow(sqlStatement, Id, Id_user, Id_product).Scan(&LastInsertId)
	if err != nil {
		panic(err)
	}

	return LastInsertId, nil
}

func (u *UserRepo) DeleteWishlistByID(Id ID) (int64, error) {
	data, err := u.db.Exec("DELETE FROM wishlist_item WHERE id = ?", Id)
	if err != nil {
		return 0, err
	}
	return data.LastInsertId()
}

func (u *UserRepo) CheckWishlist(Id_user ID, Id_product ID) (*WishlistResponse, error) {
	sqlStatement := `
	SELECT 
		t.id AS id,
		u.username AS username, 
		p.judul AS judul
	FROM 
		wishlist_item as t
	INNER JOIN 
		user as u ON (u.id = t.id_user)
	INNER JOIN 
		products as p ON (p.id = t.id_product)
	WHERE 
		t.id_user = ? AND t.id_product = ?`

	rows, err := u.db.Query(sqlStatement, Id_user, Id_product)
	if err != nil {
		return nil, err
	}

	var wish WishlistResponse
	for rows.Next() {
		err := rows.Scan(
			&wish.Id,
			&wish.Username,
			&wish.Product,
		)
		if err != nil {
			return nil, err
		}
	}
	return &wish, nil
}

func (u *UserRepo) SearchTask(search string) ([]Task, error) {
	sqlStatement := `SELECT 
			t.id,
			p.jenis,
			t.gambar,
			t.judul,
			t.deskripsi,
			t.lokasi,
			t.harga
			FROM products AS t 
			INNER JOIN JenisProducts AS p 
			ON t.id_jenis = p.id 
			WHERE t.judul LIKE ?`

	searchX := "%" + search + "%"
	rows, err := u.db.Query(sqlStatement, searchX)
	if rows == nil {
		return nil, errors.New("Product not found")
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := []Task{}
	products := Task{}

	for rows.Next() {
		err = rows.Scan(&products.Id, &products.Jenis, &products.Gambar, &products.Judul, &products.Deskripsi, &products.Lokasi, &products.Harga)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, products)

	}
	if len(tasks) == 0 {
		return nil, errors.New("Product not found")
	}
	fmt.Println(tasks)
	return tasks, nil
}

func (u *UserRepo) GetTask() ([]Task, error) {
	rows, err := u.db.Query(`
	SELECT
		t.id,
		p.jenis,
		t.gambar,
		t.judul,
		t.deskripsi,
		t.lokasi,
		t.harga,
		t.waktu,
		t.kapasitas
	FROM products AS t 
	INNER JOIN JenisProducts AS p
	ON t.id_jenis = p.id`)
	if err != nil {
		return []Task{}, err
	}

	defer rows.Close()

	result := []Task{}
	for rows.Next() {
		products := Task{}
		err = rows.Scan(&products.Id, &products.Jenis, &products.Gambar, &products.Judul, &products.Deskripsi, &products.Lokasi, &products.Harga, &products.Waktu, &products.Kapasitas)
		if err != nil {
			return []Task{}, err

		}
		result = append(result, products)
	}
	fmt.Println(result[0].Waktu)
	return result, nil
}

func (u *UserRepo) UpdatePassword(Password string) (int64, error) {
	sqlStatement := `UPDATE user SET password = ?;`

	stmt, err := u.db.Prepare(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(Password)
	if err != nil {
		panic(err)
	}

	return result.RowsAffected()
}

func (u *UserRepo) DeleteAkun(Id int) (int64, error) {
	data, err := u.db.Exec("DELETE FROM user WHERE id = ?", Id)
	if err != nil {
		return 0, err
	}
	return data.LastInsertId()
}

func (u *UserRepo) FilterByGame(id int) ([]Task, error) {
	sqlStatement := `SELECT 
			t.id,
			p.jenis,
			t.gambar,
			t.judul,
			t.deskripsi,
			t.lokasi,
			t.harga
			FROM products AS t 
			INNER JOIN JenisProducts AS p 
			ON t.id_jenis = p.id 
			WHERE t.id_jenis = ?`

	rows, err := u.db.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for rows.Next() {
		products := Task{}
		err = rows.Scan(&products.Id, &products.Jenis, &products.Gambar, &products.Judul, &products.Deskripsi, &products.Lokasi, &products.Harga)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, products)
	}
	return tasks, nil
}

func (u *UserRepo) UseCoupon(kupon UseCoupon) (*UseCoupon, error) {
	sqlStatement := `INSERT INTO user_coupons (id, id_user, kode_coupon, status) VALUES (?, ?, ?, ?);`

	rows, err := u.db.Query(sqlStatement, kupon.Id, kupon.IdUser, kupon.Kode, kupon.Status)
	if err != nil {
		return nil, err
	}

	var data UseCoupon
	for rows.Next() {
		err = rows.Scan(&data.Id, &data.IdUser, &data.Kode, &data.Status)
		if err != nil {
			return nil, err
		}
	}

	return &data, nil
}

func (u *UserRepo) ValidateCoupon(kupon UseCoupon) (*UseCoupon, error) {
	sqlStatement := `SELECT * FROM user_coupons WHERE kode_coupon = ? AND id_user = ?`

	rows, err := u.db.Query(sqlStatement, kupon.Kode, kupon.IdUser)
	if err != nil {
		return nil, err
	}

	var data UseCoupon
	for rows.Next() {
		err = rows.Scan(&data.Id, &data.IdUser, &data.Kode, &data.Status)
		if err != nil {
			return nil, err
		}
	}

	return &data, nil
}

func (u *UserRepo) CheckCoupon(Kode string) (NambahCoupon, error) {
	sqlStatement := `SELECT * FROM coupons WHERE kode = ?;`

	rows, err := u.db.Query(sqlStatement, Kode)
	if err != nil {
		return NambahCoupon{}, err
	}

	var coupon NambahCoupon
	for rows.Next() {
		err = rows.Scan(&coupon.Kode, &coupon.Diskon, &coupon.Minimal)
		if err != nil {
			return coupon, err
		}
	}

	return coupon, nil

}

func (u *UserRepo) GetNotification() ([]Notifikasi, error) {
	sqlStatement := `SELECT * FROM notifications`

	rows, err := u.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	notification := []Notifikasi{}
	for rows.Next() {
		var notifications Notifikasi
		err := rows.Scan(
			&notifications.Id,
			&notifications.Pesan,
		)
		if err != nil {
			return nil, err
		}
		notification = append(notification, notifications)
	}
	return notification, nil
}

func (u *UserRepo) AddBasket(id ID, Id_user ID, Id_product ID) (Basket, error) {
	sqlStatement := `INSERT INTO Basket (id, id_user, id_product) VALUES (?, ?, ?);`

	stmt, err := u.db.Prepare(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	row, err := stmt.Query(id, Id_user, Id_product)
	if err != nil {
		panic(err)
	}

	var basket Basket

	for row.Next() {
		err = row.Scan(&basket.Id_user, &basket.Id_product)
		if err != nil {
			panic(err)
		}
	}

	return basket, nil
}

func (u *UserRepo) CheckBasket(Id_user ID, Id_product ID) (BasketResponse, error) {
	sqlStatement := `SELECT * FROM Basket WHERE id_user = ? AND id_product = ?`
	rows, err := u.db.Query(sqlStatement, Id_user, Id_product)
	if err != nil {
		return BasketResponse{}, err
	}

	var basket BasketResponse
	for rows.Next() {
		err = rows.Scan(&basket.Id, &basket.Id_user, &basket.Id_product)
		if err != nil {
			return basket, err
		}
	}

	return basket, nil
}

func (u *UserRepo) GetWishlist(id_user ID, Id_product ID) ([]Wishlist, error) {
	sqlStatement := `SELECT * FROM wishlist_item WHERE id_user = ? and id_product = ?`
	rows, err := u.db.Query(sqlStatement, id_user, Id_product)
	if err != nil {
		return nil, err
	}

	var wishlist []Wishlist
	for rows.Next() {
		var wishlists Wishlist
		err = rows.Scan(&wishlists.Id, &wishlists.Id_user, &wishlists.Id_product)
		if err != nil {
			return nil, err
		}
		wishlists.Status = true
		wishlist = append(wishlist, wishlists)
	}

	return wishlist, nil
}

func (u *UserRepo) GetAllWishlist(id_user ID) ([]Task, error) {
	sqlStatement := `SELECT
		t.id,
		p.jenis,
		t.gambar,
		t.judul,
		t.deskripsi,
		t.lokasi,
		t.harga,
		t.waktu,
		t.kapasitas
	FROM products AS t 
	INNER JOIN JenisProducts AS p
	ON t.id_jenis = p.id
	WHERE t.id IN (SELECT id_product FROM wishlist_item WHERE wishlist_item.id_user = ?)`

	fmt.Println(id_user)
	rows, err := u.db.Query(sqlStatement, id_user)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Wishlists := []Task{}
	products := Task{}
	for rows.Next() {
		err = rows.Scan(&products.Id, &products.Jenis, &products.Gambar, &products.Judul, &products.Deskripsi, &products.Lokasi, &products.Harga, &products.Waktu, &products.Kapasitas)
		if err != nil {
			return []Task{}, err
		}

		Wishlists = append(Wishlists, products)
	}
	return Wishlists, nil
}
