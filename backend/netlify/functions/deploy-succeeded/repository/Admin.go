package repository

import (
	"database/sql"
	"fmt"
)

type AdminRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *AdminRepo {
	return &AdminRepo{db: db}
}

func (a *AdminRepo) GetTask() ([]Task, error) {
	rows, err := a.db.Query(`
	SELECT
		products.Id,
		JenisProducts.Jenis AS JenisProducts,
		products.gambar,
		products.Judul,
		products.deskripsi,
		products.lokasi,
		products.harga,
		products.waktu,
		products.kapasitas
	FROM products
	INNER JOIN JenisProducts
	ON products.Id_jenis = JenisProducts.Id`)
	if err != nil {
		return []Task{}, err
	}

	defer rows.Close()

	result := []Task{}
	for rows.Next() {
		admin := Task{}
		err = rows.Scan(&admin.Id, &admin.Jenis, &admin.Gambar, &admin.Judul, &admin.Deskripsi, &admin.Lokasi, &admin.Harga, &admin.Waktu, &admin.Kapasitas)
		if err != nil {
			return []Task{}, err
		}
		result = append(result, admin)
	}

	return result, nil
}

func (a *AdminRepo) PutTask(products CreateProductResponse) (int64, error) {
	if products.Gambar == "" {
		products.Gambar = "data:image/png;base64,"
		sqlStatement := `INSERT INTO products (id, Id_jenis, gambar,  Judul, Deskripsi, lokasi, harga, waktu, kapasitas) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
		res, err := a.db.Exec(sqlStatement, products.IdProduct, products.IdJenis, products.Gambar, products.Judul, products.Deskripsi, products.Lokasi, products.Harga, products.Waktu, products.Kapasitas)
		if err != nil {
			return 0, err
		}
		return res.RowsAffected()
	} else {
		sqlStatement := `INSERT INTO products (id, id_jenis, gambar, judul, deskripsi, lokasi, harga, waktu, kapasitas) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`

		stmt, err := a.db.Prepare(sqlStatement)
		if err != nil {
			panic(err)
		}

		defer stmt.Close()
		gambarByte := []byte(products.Gambar)
		result, err := stmt.Exec(products.IdProduct, products.IdJenis, gambarByte, products.Judul, products.Deskripsi, products.Lokasi, products.Harga, products.Waktu, products.Kapasitas)
		if err != nil {
			panic(err)
		}

		return result.LastInsertId()
	}
}
func (a *AdminRepo) UpdateTask(product UpdateProductInput) (int64, error) {
	if product.Gambar == "" {
		sqlStatement := `UPDATE products SET id_jenis = ?, judul = ?, deskripsi = ?, lokasi = ?, harga = ?, waktu = ?, kapasitas = ? WHERE id = ?;`

		stmt, err := a.db.Prepare(sqlStatement)
		if err != nil {
			return 0, err
		}

		defer stmt.Close()

		res, err := stmt.Exec(product.IdJenis, product.Judul, product.Deskripsi, product.Lokasi, product.Harga, product.Waktu, product.Kapasitas, product.IdProduct)
		if err != nil {
			return 0, err
		}

		return res.RowsAffected()
	} else {
		sqlStatement := `UPDATE products SET id_jenis = ?, gambar = ?, judul = ?, deskripsi = ?, lokasi = ?, harga = ?, waktu = ?, kapasitas = ? WHERE id = ?;`
		stmt, err := a.db.Prepare(sqlStatement)
		if err != nil {
			return 0, err
		}
		gambarByte := []byte(product.Gambar)

		res, err := stmt.Exec(product.IdJenis, gambarByte, product.Judul, product.Deskripsi, product.Lokasi, product.Harga, product.Waktu, product.Kapasitas, product.IdProduct)
		if err != nil {
			return 0, err
		}
		defer stmt.Close()
		return res.RowsAffected()
	}
}
func (a *AdminRepo) DeleteTask(id string) (int64, error) {
	sqlStatement := `DELETE FROM products WHERE Id = ?;`

	result, err := a.db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	return result.RowsAffected()
}

func (a *AdminRepo) GetBarang() ([]JenisProducts, error) {
	rows, err := a.db.Query(`SELECT * FROM JenisProducts;`)
	if err != nil {
		return []JenisProducts{}, err
	}

	defer rows.Close()

	result := []JenisProducts{}
	for rows.Next() {
		nama := JenisProducts{}
		err = rows.Scan(&nama.Id, &nama.Jenis)
		if err != nil {
			return []JenisProducts{}, err
		}
		result = append(result, nama)
	}

	return result, nil
}

func (a *AdminRepo) DeleteCategory(Id int) (int64, error) {
	sqlStatement := `DELETE FROM JenisProducts WHERE Id = ?;`

	result, err := a.db.Exec(sqlStatement, Id)
	if err != nil {
		panic(err)
	}

	return result.RowsAffected()
}

func (a *AdminRepo) CreateCategory(Jenis string) (int64, error) {
	sqlStatement := `INSERT INTO JenisProducts (jenis) VALUES (?);`

	stmt, err := a.db.Prepare(sqlStatement)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(Jenis)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (a *AdminRepo) UpdateCategory(Id int, Jenis string) (int64, error) {
	sqlStatement := `UPDATE JenisProducts SET jenis = ? WHERE id = ?;`

	stmt, err := a.db.Prepare(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(Jenis, Id)
	if err != nil {
		panic(err)
	}

	return result.RowsAffected()
}

func (a *AdminRepo) GetCategory() ([]JenisProducts, error) {
	rows, err := a.db.Query(`SELECT * FROM JenisProducts;`)
	if err != nil {
		return []JenisProducts{}, err
	}

	defer rows.Close()

	result := []JenisProducts{}
	for rows.Next() {
		nama := JenisProducts{}
		err = rows.Scan(&nama.Id, &nama.Jenis)
		if err != nil {
			return []JenisProducts{}, err
		}
		result = append(result, nama)
	}

	return result, nil
}

func (a *AdminRepo) GetJenisByID(id int) (JenisProducts, error) {
	row := a.db.QueryRow(`SELECT * FROM JenisProducts WHERE id = ?`, id)
	fmt.Println(row)
	jenis := JenisProducts{}
	err := row.Scan(&jenis.Id, &jenis.Jenis)
	if err != nil {
		fmt.Println(err)
		return jenis, err
	}
	return jenis, nil
}

func (a *AdminRepo) CheckJenis(jenis string) (JenisProducts, error) {
	sqlStatement := `SELECT * FROM JenisProducts WHERE jenis = ?;`

	rows, err := a.db.Query(sqlStatement, jenis)
	if err != nil {
		return JenisProducts{}, err
	}

	var admin JenisProducts
	for rows.Next() {
		err = rows.Scan(&admin.Id, &admin.Jenis)
		if err != nil {
			return admin, err
		}
	}

	return admin, nil
}

func (a *AdminRepo) AddCoupon(Kode string, Diskon int, Minimal int) (NambahCoupon, error) {
	sqlStatement := `INSERT INTO coupons (kode, diskon, minimal) VALUES (?, ?, ?);`

	stmt, err := a.db.Prepare(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	row, err := stmt.Query(Kode, Diskon, Minimal)
	if err != nil {
		panic(err)
	}

	var coupon NambahCoupon

	for row.Next() {
		err = row.Scan(&coupon.Kode, &coupon.Diskon, &coupon.Minimal)
		if err != nil {
			panic(err)
		}
	}

	return coupon, nil
}

func (a *AdminRepo) CheckCoupon(Kode string) (NambahCoupon, error) {
	sqlStatement := `SELECT * FROM coupons WHERE kode = ?;`

	rows, err := a.db.Query(sqlStatement, Kode)
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

func (a *AdminRepo) Notification(Notifikasi string) (int64, error) {
	sqlStatement := `INSERT INTO notifikasi (notifikasi) VALUES (?);`

	stmt, err := a.db.Prepare(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(Notifikasi)
	if err != nil {
		panic(err)
	}

	return result.LastInsertId()

}
