package object

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
