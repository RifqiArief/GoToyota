package model

import "github.com/dgrijalva/jwt-go"

type Token struct {
	UserId int
	jwt.StandardClaims
}

type User struct {
	IdUser    int    `json:"id_user" db:"id_user"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	DeleteAt  string `json:"delete_at" db:"delete_at"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Nama      string `json:"nama" db:"nama"`
	Telepon   string `json:"telepon" db:"telepon"`
	Gambar    string `json:"gambar" db:"gambar"`
	Alamat    string `json:"alamat" db:"alamat"`
	Kota      string `json:"kota" db:"kota"`
}

type Bengkel struct {
	IdBengkel int    `json:"id_user" db:"id_bengkel"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	DeleteAt  string `json:"delete_at" db:"delete_at"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Nama      string `json:"nama" db:"nama"`
	Telepon   string `json:"telepon" db:"telepon"`
	Gambar    string `json:"gambar" db:"gambar"`
	Token     string `json:"token"`
}

type Lokasi struct {
	IdLokasi  int     `json:"id_lokasi" db:"id_lokasi"`
	IdBengkel int     `json:"id_user" db:"id_bengkel"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	UpdatedAt string  `json:"updated_at" db:"updated_at"`
	DeleteAt  string  `json:"delete_at" db:"delete_at"`
	Alamat    string  `json:"alamat" db:"alamat"`
	Kota      string  `json:"kota" db:"kota"`
	Provinsi  string  `json:"provinsi" db:"provinsi"`
	Longitude float64 `json:"longitude" db:"longitude"`
	Latitude  float64 `json:"latitude" db:"latitude"`
}

type Oprasional struct {
	IdOprasional int    `json:"id_oprasional" db:"id_oprasional"`
	IdUser       int    `json:"id_user" db:"id_user"`
	CreatedAt    string `json:"create_at" db:"created_at"`
	UpdatedAt    string `json:"update_at" db:"updated_at"`
	DeleteAt     string `json:"delete_at" db:"delete_at"`
	Hari         string `json:"hari" db:"hari"`
	Buka         string `json:"buka" db:"buka"`
	Tutup        string `json:"tutup" db:"tutup"`
}
