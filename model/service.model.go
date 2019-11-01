package model

import (
	"fmt"
	"time"

	"github.com/GoToyota/object"
	"github.com/GoToyota/utils"
)

//Get jenis service
func GetJenisService() map[string]interface{} {
	query := fmt.Sprintf(`
	select 
	coalesce(nullif(id_jenis_service,null),0) as id_jenis_service,
	coalesce(nullif(nama_jenis_service,''),' ') as nama_jenis_service
	from jenis_service`)

	utils.Logging.Println(query)
	var jenisService []object.JenisService
	rows, err := db.Query(query)
	if err != nil {
		return utils.Message(false, "service.model.go, line:21 "+err.Error())
	}

	for rows.Next() {
		var s object.JenisService
		err = rows.Scan(
			&s.IdJenisService,
			&s.NamaJenisService,
		)
		if err != nil {
			return utils.Message(false, "service.model.go, line:31 "+err.Error())
		}
		jenisService = append(jenisService, s)
	}
	response := utils.Message(true, "Success")
	response["response"] = jenisService
	utils.Logging.Println(response)
	return response
}

//Add service
func (service *Service) AddService() map[string]interface{} {
	createTime := time.Now()
	var idService int
	query := `insert into service (
		created_at, id_jenis_service, id_bengkel, nama_service, harga)
		values ($1,$2,$3,$4,$5) returning id_service`
	err := db.QueryRow(query, createTime.Format("01-02-2006 15:04:05"), service.IdJenisService, service.IdBengkel, service.NamaService, service.Harga).Scan(&idService)
	utils.Logging.Println(query)
	utils.Logging.Println(idService)
	if err != nil {
		utils.Logging.
			Println("servivce.model.go, line:54")
		utils.Logging.Println(err.Error())
		return utils.Message(false, "Failed insert service")
	}

	response := utils.Message(true, "Success add menu service")

	return response
}

//Get service by id bengkel
func GetService(idBengkel int) map[string]interface{} {
	query := fmt.Sprintf(`
	select 
	coalesce(nullif(id_service,null),0) as id_service,
	coalesce(nullif(id_jenis_service,null),0) as id_jenis_service,
	coalesce(nullif(id_bengkel,null),0) as id_bengkel,
	coalesce(nullif(nama_service,''),' ') as nama_service,
	coalesce(nullif(harga,null),0) as harga
	from service where id_bengkel = %v`, idBengkel)

	utils.Logging.Println(query)
	var service []object.Service
	rows, err := db.Query(query)
	if err != nil {
		return utils.Message(false, "service.model.go, line:80 "+err.Error())
	}

	for rows.Next() {
		var s object.Service
		err = rows.Scan(
			&s.IdService,
			&s.IdJenisService,
			&s.IdBengkel,
			&s.NamaService,
			&s.Harga,
		)
		if err != nil {
			return utils.Message(false, "service.model.go, line:93 "+err.Error())
		}
		service = append(service, s)
	}
	response := utils.Message(true, "Success")
	response["response"] = service
	utils.Logging.Println(response)
	return response
}

func (service *Service) UpdateService() map[string]interface{} {

	return nil
}
