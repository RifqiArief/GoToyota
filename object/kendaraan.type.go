package object

type Kendaraan struct {
	NoStnk          int    `json:"no_stnk" db:"no_stnk"`
	IdUser          int    `json:"id_user" db:"id_user"`
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
