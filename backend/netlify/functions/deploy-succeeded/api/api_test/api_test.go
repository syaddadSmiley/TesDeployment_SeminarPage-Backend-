package api_test

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	api "github.com/syaddadSmiley/SeminarPage/api"
	repository "github.com/syaddadSmiley/SeminarPage/repository"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	_ "github.com/mattn/go-sqlite3"
)

type TokenStr struct {
	Token string `json:"token"`
}
type loginRespon struct {
	Code    int      `json:"code"`
	Data    TokenStr `json:"data"`
	Message string   `json:"message"`
}

var cookie *http.Cookie
var mainApi *api.API

var DB *sql.DB

var responseLogin loginRespon

var _ = Describe("api", func() {

	BeforeEach(func() {
		db, err := sql.Open("sqlite3", "../../database/final_project2.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS user (
				Id integer not null primary key AUTOINCREMENT,
				Nama varchar(255) not null,
				Username varchar(255) not null,
				mail varchar(255) not null,
				Password varchar(255) not null,
				role varchar(255) not null
			);

			CREATE TABLE IF NOT EXISTS products (
				Id integer not null primary key AUTOINCREMENT,
				Judul varchar(255) not null,
				Tanggal varchar(255) not null,
				Id_Penulis integer not null,
				Deskripsi varchar(255) not null,
				FOREIGN KEY (id_penulis) REFERENCES penulis(id)
			);

			CREATE TABLE IF NOT EXISTS penulis (
				Id integer not null primary key AUTOINCREMENT,
				nama varchar(255) not null
			);

			INSERT INTO user (Nama, Username, mail, Password, role) 
			VALUES ("nanda","nanda","nanda@gmail.com","$2a$10$S4ENBIoIA7BGPMQVE3h.H.W9attf8bAR3uq94rY5Ynw3v4o1Ch.Vm","user");
		`)

		if err != nil {
			panic(err)
		}

		userRepo := repository.NewUserRepo(db)
		adminRepo := repository.NewTaskRepo(db)
		mainApi = api.NewAPI(*userRepo, *adminRepo)
	})

	AfterEach(func() {
		db, err := sql.Open("sqlite3", "../../database/final_project2.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
		DROP TABLE IF EXISTS user;
		DROP TABLE IF EXISTS products;
		DROP TABLE IF EXISTS penulis;`)

		if err != nil {
			panic(err)
		}
	})

	When("username and password are correct", func() {
		It("should return message login success", func() {
			bodyReader := strings.NewReader(`{"username": "nanda", "password": "nanda"}`)

			w := httptest.NewRecorder()
			r, err := http.NewRequest("POST", "/Login", bodyReader)
			if err != nil {
				log.Fatal(err)
			}
			mainApi.Handler().ServeHTTP(w, r)

			err = json.Unmarshal([]byte(w.Body.String()), &responseLogin)
			Expect(responseLogin.Message).To(Equal("login success"))
			cookie = &http.Cookie{Name: "token", Value: responseLogin.Data.Token}
		})
	})

	When("username is correct but password is incorrect", func() {
		It("should return 401", func() {
			bodyReader := strings.NewReader(`{"username": "nanda", "password": "nanda123"}`)

			w := httptest.NewRecorder()
			r, err := http.NewRequest("POST", "/Login", bodyReader)
			if err != nil {
				log.Fatal(err)
			}
			mainApi.Handler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusUnauthorized))
		})
	})

	When("username is incorrect but password is correct", func() {
		It("should return 401", func() {
			bodyReader := strings.NewReader(`{"username": "nanda123", "password": "nanda"}`)

			w := httptest.NewRecorder()
			r, err := http.NewRequest("POST", "/Login", bodyReader)
			if err != nil {
				log.Fatal(err)
			}
			mainApi.Handler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusUnauthorized))
		})
	})

	When("username is incorrect but password is incorrect", func() {
		It("should return 401", func() {
			bodyReader := strings.NewReader(`{"username": "nanda123", "password": "nanda123"}`)

			w := httptest.NewRecorder()
			r, err := http.NewRequest("POST", "/Login", bodyReader)
			if err != nil {
				log.Fatal(err)
			}
			mainApi.Handler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusUnauthorized))
		})
	})

	When("Register Test", func() {
		It("should return message register success", func() {
			bodyReader := strings.NewReader(`{"nama": "user_test_1233","username": "user_test_1232", "password": "user_test1232", "Mail" : "user_test_123z2@email.com"}`)
			r, err := http.NewRequest("POST", "/RegisterAdmin", bodyReader)
			w := httptest.NewRecorder()
			if err != nil {
				log.Fatal(err)
			}

			mainApi.Handler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))

			// delete user
			bodyReader = strings.NewReader(`{"username": "user_test_1232"}`)
			r, err = http.NewRequest("DELETE", "/DeleteUser", bodyReader)
			w = httptest.NewRecorder()
			if err != nil {
				log.Fatal(err)
			}
			mainApi.Handler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})

	When("Register admin Test", func() {
		It("should return message register admin success", func() {
			bodyReader := strings.NewReader(`{"nama": "user_test_1233","username": "user_test_1232", "password": "user_test1232", "Mail" : "user_test_123z2@email.com"}`)
			r, err := http.NewRequest("POST", "/RegisterAdmin", bodyReader)
			w := httptest.NewRecorder()
			if err != nil {
				log.Fatal(err)
			}

			mainApi.Handler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))

			// delete user
			bodyReader = strings.NewReader(`{"username": "user_test_1232"}`)
			r, err = http.NewRequest("DELETE", "/DeleteUser", bodyReader)
			w = httptest.NewRecorder()
			if err != nil {
				log.Fatal(err)
			}
			mainApi.Handler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})

	When("logout Test", func() {
		It("should return message logout success", func() {
			db, err := sql.Open("sqlite3", "./final_project2.db")
			if err != nil {
				panic(err)
			}
			defer db.Close()
			userRepo := repository.NewUserRepo(db)
			userAdmin := repository.NewTaskRepo(db)

			route := api.NewAPI(*userRepo, *userAdmin).Handler()
			r, err := http.NewRequest("POST", "/Logout", nil)
			w := httptest.NewRecorder()
			if err != nil {
				log.Fatal(err)
			}
			r.AddCookie(cookie)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})

})
