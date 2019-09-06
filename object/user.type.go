package object

type RegisterUser struct {
	IdUser  int    `json:"id_user"`
	Nama    string `json:"nama"`
	Email   string `json:"email"`
	Telepon string `json:"telepon"`
	Alamat  string `json:"alamat"`
	Kota    string `json:"kota"`
	Gambar  string `json:"image"`
}

type LoginUser struct {
	IdUser  int    `json:"id_user"`
	Nama    string `json:"nama"`
	Email   string `json:"email"`
	Telepon string `json:"telepon"`
	Alamat  string `json:"alamat"`
	Kota    string `json:"kota"`
	Gambar  string `json:"image"`
	Token   string `json:"token"`
}

type GetAllUser struct {
	IdUser    int    `json:"id_user"`
	Nama      string `json:"nama"`
	Email     string `json:"email"`
	Telepon   string `json:"telepon"`
	Alamat    string `json:"alamat"`
	Kota      string `json:"kota"`
	Gambar    string `json:"image"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeleteAt  string `json:"delete_at"`
}
