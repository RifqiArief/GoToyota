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
	IdBengkel int    `json:"id_bengkel" db:"id_bengkel"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	DeleteAt  string `json:"delete_at" db:"delete_at"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Nama      string `json:"nama" db:"nama"`
	Telepon   string `json:"telepon" db:"telepon"`
	Gambar    string `json:"gambar" db:"gambar"`
	Token     string `json:"token"`
	IdRole    int    `json:"id_role" db:"id_role"`
}

type Lokasi struct {
	IdLokasi  int     `json:"id_lokasi" db:"id_lokasi"`
	IdBengkel int     `json:"id_bengkel" db:"id_bengkel"`
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

type Kendaraan struct {
	NoStnk          int    `json:"no_stnk" db:"no_stnk"`
	IdUser          int    `json:"id_user" db:"id_user"`
	CreatedAt       string `json:"created_at" db:"created_at"`
	UpdatedAt       string `json:"updated_at" db:"updated_at"`
	DeleteAt        string `json:"delete_at" db:"delete_at"`
	Merk            string `json:"merk" db:"merk"`
	Type            string `json:"type" db:"type"`
	Jenis           string `json:"jenis" db:"jenis"`
	Model           string `json:"model" db:"model"`
	TahunPembuatan  int    `json:"tahun_pembuatan" db:"tahun_pembuatan"`
	IsiSilinder     string `json:"isi_silinder" db:"isi_silinder"`
	NoRangka        string `json:"no_rangka" db:"no_rangka"`
	NoMesin         string `json:"no_mesin" db:"no_mesin"`
	Warna           string `json:"warna" db:"warna"`
	BahanBakar      string `json:"bahan_bakar" db:"bahan_bakar"`
	WarnaTnkb       string `json:"warna_tnkb" db:"warna_tnkb"`
	TahunRegistrasi int    `json:"tahun_registrasi" db:"tahun_registrasi"`
	NoBpkb          string `json:"no_bpkb" db:"no_bpkb"`
	KodeLokasi      string `json:"kode_lokasi" db:"kode_lokasi"`
}

type JenisService struct {
	IdJenisService   int    `json:"id_jenis_service" db:"id_jenis_service"`
	NamaJenisService string `json:"nama_jenis_service" db:"nama_jenis_service"`
}

type Service struct {
	IdService      int    `json:"id_service" db:"id_service"`
	IdJenisService int    `json:"id_jenis_service" db:"id_jenis_service"`
	IdBengkel      int    `json:"id_bengkel" db:"id_bengkel"`
	NamaService    string `json:"nama_service" db:"nama_service"`
	Harga          int    `json:"harga" db:"harga"`
}

type ChangePassword struct {
	IdUser      int    `json:"id_user" db:"id_user"`
	OldPassword string `json:"old_password" db:"password"`
	NewPassword string `json:"new_password"`
}
