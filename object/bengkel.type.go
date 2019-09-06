package object

type Lokasi struct {
	Alamat    string  `json:"alamat"`
	Kota      string  `json:"kota"`
	Provinsi  string  `json:"provinsi"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Oprasional struct {
	Hari  string `json:"hari"`
	Buka  string `json:"buka"`
	Tutup string `json:"tutup"`
}

type RegisterBengkel struct {
	IdBengkel int    `json:"id_bengkel"`
	Nama      string `json:"nama"`
	Email     string `json:"email"`
	Telepon   string `json:"telepon"`
	Gambar    string `json:"image"`
	Lokasi    Lokasi `json:"lokasi"`
}

type LoginBengkel struct {
	IdBengkel  int          `json:"id_bengkel"`
	Nama       string       `json:"nama"`
	Email      string       `json:"email"`
	Telepon    string       `json:"telepon"`
	Gambar     string       `json:"image"`
	Token      string       `json:"token"`
	Lokasi     Lokasi       `json:"lokasi"`
	Oprasional []Oprasional `json:"oprasional"`
}

type GetAllBengkel struct {
	IdBengkel int     `json:"id_bengkel"`
	Nama      string  `json:"nama"`
	Email     string  `json:"email"`
	Telepon   string  `json:"telepon"`
	Alamat    string  `json:"alamat"`
	Kota      string  `json:"kota"`
	Provinsi  string  `json:"provinsi"`
	Gambar    string  `json:"image"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
