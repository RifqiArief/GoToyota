package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/GoToyota/object"

	"github.com/GoToyota/utils"
)

func (loc *Lokasi) AddLocation(idBengkel int) (map[string]interface{}, *object.Lokasi) {
	query := `
	insert into lokasi (
		id_bengkel,
		created_at,
		alamat, 
		kota,
		provinsi, 
		longitude,
		latitude
	) values ($1,$2,$3,$4,$5,$6,$7)`
	utils.Logging.Println(query)
	createTime := time.Now()
	_, err := db.Exec(query, idBengkel, createTime.Format("01-02-2006 15:04:05"), loc.Alamat, loc.Kota, loc.Provinsi, loc.Longitude, loc.Latitude)
	if err != nil {
		utils.Logging.Println("model/userLocation.go, line:28")
		utils.Logging.Println(err.Error())
		return utils.Message(false, err.Error()), nil
	}

	data := &object.Lokasi{
		Alamat:    loc.Alamat,
		Kota:      loc.Kota,
		Provinsi:  loc.Provinsi,
		Longitude: loc.Longitude,
		Latitude:  loc.Latitude,
	}

	return utils.Message(true, "Success Registration"), data
}

func GetLocation(id_bengkel int) (map[string]interface{}, object.Lokasi) {
	query := fmt.Sprintf(`
	select 
	coalesce(nullif(id_lokasi,null),0) as "id_lokasi",
	coalesce(nullif(id_bengkel,null),0) as "id_bengkel",
	coalesce(nullif(alamat,''),' ') as "alamat",
	coalesce(nullif(kota,''),' ') as "kota",
	coalesce(nullif(provinsi,''),' ') as "provinsi",
	coalesce(nullif(longitude,null),0) as "longitude",
	coalesce(nullif(latitude,null),0) as "latitude",
	coalesce(nullif(created_at,null),'2000-01-01 00:00:00') as created_at,
	coalesce(nullif(updated_at,null),'2000-01-01 00:00:00') as updated_at,
	coalesce(nullif(delete_at,null),'2000-01-01 00:00:00') as delete_at
	from lokasi where 
	id_bengkel = %v`, id_bengkel)

	utils.Logging.Println(query)
	var loc Lokasi
	err := db.Get(&loc, query)
	if err != nil && err != sql.ErrNoRows {
		return utils.Message(false, "location.model.go:64"+err.Error()), object.Lokasi{}
	}

	data := object.Lokasi{
		Alamat:    loc.Alamat,
		Kota:      loc.Kota,
		Provinsi:  loc.Provinsi,
		Longitude: loc.Longitude,
		Latitude:  loc.Latitude,
	}

	return nil, data
}

//edit location
