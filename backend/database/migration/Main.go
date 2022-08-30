package main

import (
	"database/sql"

	con "github.com/syaddadSmiley/SeminarPage/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := con.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	asd, err := CreateTable(db)
	if err != nil {
		panic(err)
	}
	println(asd)

	ad, err := CreateTableProducts(db)
	if err != nil {
		panic(err)
	}
	println(ad)

	ap, err := CreateTableJenisProduct(db)
	if err != nil {
		panic(err)
	}
	println(ap)

	ak, err := CreateTableKomentar(db)
	if err != nil {
		panic(err)
	}
	println(ak)

	az, err := wishlist_table(db)
	if err != nil {
		panic(err)
	}
	println(az)

	az, err = coupon_table(db)
	if err != nil {
		panic(err)
	}
	println(az)

	az, err = user_coupons_table(db)
	if err != nil {
		panic(err)
	}
	println(az)

	az, err = notifikasi_table(db)
	if err != nil {
		panic(err)
	}
	println(az)

	az, err = basket_table(db)
	if err != nil {
		panic(err)
	}
	println(az)
}

func CreateTable(db *sql.DB) (string, error) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS user (
		id string not null primary key UNIQUE,
		Nama varchar(255) not null,
		Username varchar(255) not null UNIQUE,
		mail varchar(255) not null UNIQUE,
		Password varchar(255) not null,
		role varchar(255) not null,
		gambar LONGBLOB 
	);

	`)
	if err != nil {
		return "Failed Make Table", err
	}
	return "Succes Make Table", nil

}

func CreateTableProducts(db *sql.DB) (string, error) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS products (
		id string not null primary key,
		id_jenis integer not null,
		gambar longblob,
		judul integer not null,
		deskripsi varchar(255) not null,
		lokasi varchar(255) not null DEFAULT 'Online',
		harga integer not null,
		waktu TIMESTAMP not null,
		kapasitas integer not null,

		FOREIGN KEY (id_jenis) REFERENCES JenisProducts(id)
	);

	`)
	if err != nil {
		return "Failed Make Table", err
	}
	return "Succes Make Table", nil
}

func CreateTableJenisProduct(db *sql.DB) (string, error) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS JenisProducts (
		id integer not null primary key AUTOINCREMENT,
		jenis varchar(255) not null
	);

	`)
	if err != nil {
		return "Failed Make Table", err
	}

	if err != nil {
		return "Failed Make Table", err
	}
	return "Succes Make Table", nil
}

func CreateTableKomentar(db *sql.DB) (string, error) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS komentar (
		id string not null primary key,
		id_product string not null ,
		id_user string not null,
		content varchar(255) not null,
		like integer not null,
		dislike integer not null,
		rating integer not null,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (id_product) REFERENCES products(id),
		FOREIGN KEY (id_user) REFERENCES user(id)
	);

	`)
	if err != nil {
		return "Failed Make Table", err
	}
	return "Succes Make Table", nil
}

func wishlist_table(db *sql.DB) (string, error) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS wishlist_item (
		id string not null primary key,
		id_user string not null,
		id_product string not null,
		FOREIGN KEY (id_user) REFERENCES user(id),
		FOREIGN KEY (id_product) REFERENCES products(id)
	);

	`)
	if err != nil {
		return "Failed Make Table", err
	}
	return "Succes Make Table", nil
}

func coupon_table(db *sql.DB) (string, error) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS coupons (
		kode string not null primary key,
		diskon integer not null,
		minimal integer not null
	);

	`)
	if err != nil {
		return "Failed Make Table", err
	}
	return "Succes Make Table", nil
}

func user_coupons_table(db *sql.DB) (string, error) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS user_coupons (
		id string NOT NULL PRIMARY KEY,
		id_user string NOT NULL,
		kode_coupon string NOT NULL,
		status string NOT NULL
	);

	`)
	if err != nil {
		return "Failed Make Table", err
	}
	return "Succes Make Table", nil
}

func notifikasi_table(db *sql.DB) (string, error) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS Notifikasi (
		id string NOT NULL PRIMARY KEY,
		Pesan string NOT NULL
	);

	`)
	if err != nil {
		return "Failed Make Table", err
	}
	return "Succes Make Table", nil
}

func basket_table(db *sql.DB) (string, error) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS basket (
		id string NOT NULL PRIMARY KEY,
		id_user string NOT NULL,
		id_product string NOT NULL,
		FOREIGN KEY (id_user) REFERENCES user(id),
		FOREIGN KEY (id_product) REFERENCES products(id)
	);

	`)
	if err != nil {
		return "Failed Make Table", err
	}
	return "Succes Make Table", nil
}
