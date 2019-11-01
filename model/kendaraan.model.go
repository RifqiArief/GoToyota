package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/GoToyota/object"
	"github.com/GoToyota/utils"
)

//add mobil
func (kendaraan *Kendaraan) AddKendaraan() (map[string]interface{}, bool) {

	var NoStnk int
	querycheck := fmt.Sprintf(`
	select 
	coalesce(nullif(no_stnk,null),0) as no_stnk 
	from kendaraan where no_stnk = %v`, kendaraan.NoStnk)

	err := db.Get(&NoStnk, querycheck)
	utils.Logging.Println(querycheck)
	if err != nil && err != sql.ErrNoRows {
		return utils.Message(false, "location.model.go:64 "+err.Error()), false
	}

	utils.Logging.Println(NoStnk)

	if NoStnk != 0 {
		return utils.Message(false, "No STNK already use"), false
	}

	createTime := time.Now()

	query := `insert into kendaraan (
		created_at, no_stnk, id_user, merk, type, jenis,model, tahun_pembuatan, isi_silinder, 
		no_rangka, no_mesin, warna, bahan_bakar, warna_tnkb, tahun_registrasi, no_bpkb, kode_lokasi)
		values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17) returning no_stnk`
	err = db.QueryRow(query, createTime.Format("01-02-2006 15:04:05"), kendaraan.NoStnk, kendaraan.IdUser, kendaraan.Merk, kendaraan.Type,
		kendaraan.Jenis, kendaraan.Model, kendaraan.TahunPembuatan, kendaraan.IsiSilinder, kendaraan.NoRangka, kendaraan.NoMesin, kendaraan.Warna,
		kendaraan.BahanBakar, kendaraan.WarnaTnkb, kendaraan.TahunRegistrasi, kendaraan.NoBpkb, kendaraan.KodeLokasi).Scan(&NoStnk)
	utils.Logging.Println(createTime.Format("01-02-2006 15:04:05"))
	utils.Logging.Println(query)
	if err != nil {
		utils.Logging.
			Println("model/user.model.go, line:72")
		utils.Logging.Println(err.Error())
		return utils.Message(false, "Failed insert user"), false
	}

	response := utils.Message(true, "Success add car")
	return response, true
}

//get kendaraan by id
func GetKendaraanId(IdKendaraan int) map[string]interface{} {
	query := fmt.Sprintf(`
	select 
	coalesce(nullif(no_stnk,null),0) as no_stnk,
	coalesce(nullif(id_user,null),0) as id_user,
	coalesce(nullif(merk,''),' ') as merk,
	coalesce(nullif(type,''),' ') as type,
	coalesce(nullif(jenis,''),' ') as jenis,
	coalesce(nullif(model,''),' ') as model,
	coalesce(nullif(tahun_pembuatan,null),0) as tahun_pembuatan,
	coalesce(nullif(isi_silinder,''),' ') as isi_silinder,
	coalesce(nullif(no_rangka,''),' ') as no_rangka,
	coalesce(nullif(no_mesin,''),' ') as no_mesin,
	coalesce(nullif(warna,''),' ') as warna,
	coalesce(nullif(bahan_bakar,''),' ') as bahan_bakar,
	coalesce(nullif(warna_tnkb,''),' ') as warna_tnkb,
	coalesce(nullif(tahun_registrasi,null),0) as tahun_registrasi,
	coalesce(nullif(no_bpkb,''),' ') as no_bpkb,
	coalesce(nullif(kode_lokasi,''),' ') as kode_lokasi
	from kendaraan where no_stnk = %v`, IdKendaraan)

	utils.Logging.Println(query)
	var kendaraan []object.Kendaraan
	rows, err := db.Query(query)
	if err != nil {
		return utils.Message(false, "kendaraan.model.go, line:81 "+err.Error())
	}

	for rows.Next() {
		var k object.Kendaraan
		err = rows.Scan(
			&k.NoStnk,
			&k.IdUser,
			&k.Merk,
			&k.Type,
			&k.Jenis,
			&k.Model,
			&k.TahunPembuatan,
			&k.IsiSilinder,
			&k.NoRangka,
			&k.NoMesin,
			&k.Warna,
			&k.BahanBakar,
			&k.WarnaTnkb,
			&k.TahunRegistrasi,
			&k.NoBpkb,
			&k.KodeLokasi,
		)
		if err != nil {
			return utils.Message(false, "kendaraan.model.go, line:105 "+err.Error())
		}
		kendaraan = append(kendaraan, k)
	}
	response := utils.Message(true, "Success")
	response["response"] = kendaraan
	utils.Logging.Println(response)
	return response
}

func (kendaraan *Kendaraan) UpdateKendaraan() map[string]interface{} {
	var NoStnk int
	querycheck := fmt.Sprintf(`
	select 
	coalesce(nullif(no_stnk,null),0) as no_stnk 
	from kendaraan where no_stnk = %v`, kendaraan.NoStnk)

	err := db.Get(&NoStnk, querycheck)
	utils.Logging.Println(querycheck)
	if err != nil && err != sql.ErrNoRows {
		return utils.Message(false, "kendaraan.model.go:125 "+err.Error())
	}

	utils.Logging.Println(NoStnk)

	if NoStnk == 0 {
		return utils.Message(false, "No STNK not registered")
	}

	query := fmt.Sprintf(`
	UPDATE kendaraan SET 
	updated_at=$1, 
	merk= $2, 
	type=$3, 
	jenis=$4, 
	model=$5, 
	tahun_pembuatan=$6, 
	isi_silinder=$7, 
	no_rangka=$8, 
	no_mesin=$9, 
	warna=$10, 
	bahan_bakar=$11, 
	warna_tnkb=$12, 
	tahun_registrasi=$13, 
	no_bpkb=$14, 
	kode_lokasi=$15
	WHERE no_stnk= $16 returning no_stnk`)

	createTime := time.Now()
	err = db.QueryRow(query, createTime.Format("01-02-2006 15:04:05"), kendaraan.Merk, kendaraan.Type, kendaraan.Jenis, kendaraan.Model,
		kendaraan.TahunPembuatan, kendaraan.IsiSilinder, kendaraan.NoRangka, kendaraan.NoMesin, kendaraan.Warna, kendaraan.BahanBakar, kendaraan.WarnaTnkb,
		kendaraan.TahunRegistrasi, kendaraan.NoBpkb, kendaraan.KodeLokasi, kendaraan.NoStnk).Scan(&NoStnk)
	utils.Logging.Println(createTime.Format("01-02-2006 15:04:05"))
	utils.Logging.Println(query)
	if err != nil {
		utils.Logging.
			Println("kendaraan.model.go, line:161")
		utils.Logging.Println(err.Error())
		return utils.Message(false, "Failed update kendaraan")
	}

	response := utils.Message(true, "Success update kendaraan")
	return response
}
