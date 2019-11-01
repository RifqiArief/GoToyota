package model

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/GoToyota/object"
	"github.com/GoToyota/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (bengkel *Bengkel) ValidateBengkel() (map[string]interface{}, bool) {
	if !strings.Contains(bengkel.Email, "@") {
		return utils.Message(false, "Invalid email format"), false
	}
	if len(bengkel.Password) < 2 {
		return utils.Message(false, "Password is required"), false
	}
	query := fmt.Sprintf(`
	select 
	coalesce(nullif(id_bengkel,null),0) as id_bengkel,
	coalesce(nullif(email,''),' ') as email ,
	coalesce(nullif(password,''),' ') as "password",
	coalesce(nullif(nama,''),' ') as nama,
	coalesce(nullif(telepon,''),' ') as telepon,
	coalesce(nullif(gambar,''),' ') as gambar,
	coalesce(nullif(created_at,null),'2000-01-01 00:00:00') as created_at,
	coalesce(nullif(updated_at,null),'2000-01-01 00:00:00') as updated_at,
	coalesce(nullif(delete_at,null),'2000-01-01 00:00:00') as delete_at
	from bengkel where email = '%s'`, bengkel.Email)
	utils.Logging.Println(query)
	var temp Bengkel
	err := db.Get(&temp, query)
	if err != nil && err != sql.ErrNoRows {
		return utils.Message(false, err.Error()), false
	}
	utils.Logging.Printf("%+v\n", temp)
	if temp.Email != "" {
		return utils.Message(false, "Email already use"), false
	}
	return utils.Message(false, "Required passed"), true
}

//add bengkel
func (bengkel *Bengkel) AddBengkel() (map[string]interface{}, *object.RegisterBengkel) {
	res, ok := bengkel.ValidateBengkel()
	if !ok {
		utils.Logging.Println("bengkel.model.go, line:50")
		utils.Logging.Println(res)
		return res, nil
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(bengkel.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.Message(false, err.Error()), nil
	}
	createTime := time.Now()
	bengkel.Password = string(hash)
	var idBengkel int
	query := `insert into bengkel(
		created_at, email, password, nama, telepon, gambar)
		values ($1,$2,$3,$4,$5,$6) returning id_bengkel`
	err = db.QueryRow(query, createTime.Format("01-02-2006 15:04:05"), bengkel.Email, bengkel.Password, bengkel.Nama, bengkel.Telepon, bengkel.Gambar).Scan(&idBengkel)
	utils.Logging.Println(idBengkel)
	utils.Logging.Println(query)
	if err != nil {
		utils.Logging.Println("bengkel.model.go, line:68")
		utils.Logging.Println(err.Error())
		return utils.Message(false, "Failed insert bengekel"), nil
	}

	data := &object.RegisterBengkel{
		IdBengkel: idBengkel,
		Nama:      bengkel.Nama,
		Email:     bengkel.Email,
		Telepon:   bengkel.Telepon,
		Gambar:    bengkel.Gambar,
	}

	response := utils.Message(true, "Success registration")

	return response, data
}

//login bengkel
func LoginBengkel(email, password string) (map[string]interface{}, *Bengkel) {

	query := fmt.Sprintf(`
	select 
	coalesce(nullif(id_bengkel,null),-1) as id_bengkel,
	coalesce(nullif(email,''),' ') as email ,
	coalesce(nullif(password,''),' ') as "password",
	coalesce(nullif(nama,''),' ') as nama,
	coalesce(nullif(telepon,''),' ') as telepon,
	coalesce(nullif(gambar,''),' ') as gambar,
	coalesce(nullif(id_role,null),0) as id_role,
	coalesce(nullif(created_at,null),'2000-01-01 00:00:00') as created_at,
	coalesce(nullif(updated_at,null),'2000-01-01 00:00:00') as updated_at,
	coalesce(nullif(delete_at,null),'2000-01-01 00:00:00') as delete_at
	from bengkel where email = '%s'`, email)

	var bengkel Bengkel
	err := db.Get(&bengkel, query)
	utils.Logging.Println(query)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return utils.Message(false, "Email not registered"), nil
		}
		return utils.Message(false, "bengkel.model.go, line:111 "+err.Error()), nil
	}

	//chek password
	err = bcrypt.CompareHashAndPassword([]byte(bengkel.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "Wrong password, please try again"), nil
	}

	tk := &Token{UserId: bengkel.IdBengkel}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return utils.Message(false, "Failed generate token string"), nil
	}

	data := &Bengkel{
		IdBengkel: bengkel.IdBengkel,
		Nama:      bengkel.Nama,
		Email:     bengkel.Email,
		Telepon:   bengkel.Telepon,
		Gambar:    bengkel.Gambar,
		Token:     tokenString,
		IdRole:    bengkel.IdRole,
	}
	return nil, data
}

//get bengkel by kota
func GetBengkelKota(kota string) map[string]interface{} {
	query := fmt.Sprintf(`
	select 
	coalesce(nullif(bengkel.id_bengkel,null),-1) as id_bengkel,
	coalesce(nullif(nama,''),' ') as nama,
	coalesce(nullif(email,''),' ') as email ,
	coalesce(nullif(telepon,''),' ') as telepon,
	coalesce(nullif(alamat,''),' ') as alamat,
	coalesce(nullif(kota,''),' ') as kota,
	coalesce(nullif(provinsi,''),' ') as provinsi,
	coalesce(nullif(gambar,''),' ') as gambar,
	coalesce(nullif(longitude,null),0) as longitude,
	coalesce(nullif(latitude,null),0) as latitude
	from bengkel left join lokasi on bengkel.id_bengkel = lokasi.id_bengkel
	where lokasi.kota = '%s'`, strings.ToLower(kota))

	utils.Logging.Println(query)
	var bengkel []object.GetAllBengkel
	rows, err := db.Query(query)
	if err != nil {
		return utils.Message(false, "bengkel.model.go, line:158 "+err.Error())
	}

	for rows.Next() {
		var b object.GetAllBengkel
		err = rows.Scan(
			&b.IdBengkel,
			&b.Nama,
			&b.Email,
			&b.Telepon,
			&b.Alamat,
			&b.Kota,
			&b.Provinsi,
			&b.Gambar,
			&b.Longitude,
			&b.Latitude,
		)
		if err != nil {
			return utils.Message(false, "bengkel.model.go, line:176 "+err.Error())
		}
		bengkel = append(bengkel, b)
	}
	response := utils.Message(true, "Success")
	response["response"] = bengkel
	utils.Logging.Println(response)
	return response
}

//get bengkel by provinsi
func GetBengkelProvinsi(provinsi string) map[string]interface{} {
	query := fmt.Sprintf(`
	select 
	coalesce(nullif(bengkel.id_bengkel,null),-1) as id_bengkel,
	coalesce(nullif(nama,''),' ') as nama,
	coalesce(nullif(email,''),' ') as email ,
	coalesce(nullif(telepon,''),' ') as telepon,
	coalesce(nullif(alamat,''),' ') as alamat,
	coalesce(nullif(kota,''),' ') as kota,
	coalesce(nullif(provinsi,''),' ') as provinsi,
	coalesce(nullif(gambar,''),' ') as gambar,
	coalesce(nullif(longitude,null),0) as longitude,
	coalesce(nullif(latitude,null),0) as latitude
	from bengkel left join lokasi on bengkel.id_bengkel = lokasi.id_bengkel
	where lokasi.provinsi = '%s'`, strings.ToLower(provinsi))

	utils.Logging.Println(query)
	var bengkel []object.GetAllBengkel
	rows, err := db.Query(query)
	if err != nil {
		return utils.Message(false, "bengkel.model.go, line:208 "+err.Error())
	}

	for rows.Next() {
		var b object.GetAllBengkel
		err = rows.Scan(
			&b.IdBengkel,
			&b.Nama,
			&b.Email,
			&b.Telepon,
			&b.Alamat,
			&b.Kota,
			&b.Provinsi,
			&b.Gambar,
			&b.Longitude,
			&b.Latitude,
		)
		if err != nil {
			return utils.Message(false, "bengkel.model.go, line:226 "+err.Error())
		}
		bengkel = append(bengkel, b)
	}
	response := utils.Message(true, "Success")
	response["response"] = bengkel
	utils.Logging.Println(response)
	return response
}

//get bengkel by Id
func GetBengkelId(IdBengkel int) map[string]interface{} {
	query := fmt.Sprintf(`
	select 
	coalesce(nullif(bengkel.id_bengkel,null),-1) as id_bengkel,
	coalesce(nullif(nama,''),' ') as nama,
	coalesce(nullif(email,''),' ') as email ,
	coalesce(nullif(telepon,''),' ') as telepon,
	coalesce(nullif(alamat,''),' ') as alamat,
	coalesce(nullif(kota,''),' ') as kota,
	coalesce(nullif(provinsi,''),' ') as provinsi,
	coalesce(nullif(gambar,''),' ') as gambar,
	coalesce(nullif(longitude,null),0) as longitude,
	coalesce(nullif(latitude,null),0) as latitude
	from bengkel left join lokasi on bengkel.id_bengkel = lokasi.id_bengkel
	where bengkel.id_bengkel = %v`, IdBengkel)

	utils.Logging.Println(query)
	var bengkel []object.GetAllBengkel
	rows, err := db.Query(query)
	if err != nil {
		return utils.Message(false, "bengkel.model.go, line:258 "+err.Error())
	}

	for rows.Next() {
		var b object.GetAllBengkel
		err = rows.Scan(
			&b.IdBengkel,
			&b.Nama,
			&b.Email,
			&b.Telepon,
			&b.Alamat,
			&b.Kota,
			&b.Provinsi,
			&b.Gambar,
			&b.Longitude,
			&b.Latitude,
		)
		if err != nil {
			return utils.Message(false, "bengkel.model.go, line:276 "+err.Error())
		}
		bengkel = append(bengkel, b)
	}
	response := utils.Message(true, "Success")
	response["response"] = bengkel
	utils.Logging.Println(response)
	return response
}

//get all bengkel
func GetAllBengkel() map[string]interface{} {
	query := fmt.Sprintf(`
	select 
	coalesce(nullif(bengkel.id_bengkel,null),-1) as id_bengkel,
	coalesce(nullif(nama,''),' ') as nama,
	coalesce(nullif(email,''),' ') as email ,
	coalesce(nullif(telepon,''),' ') as telepon,
	coalesce(nullif(alamat,''),' ') as alamat,
	coalesce(nullif(kota,''),' ') as kota,
	coalesce(nullif(provinsi,''),' ') as provinsi,
	coalesce(nullif(gambar,''),' ') as gambar,
	coalesce(nullif(longitude,null),0) as longitude,
	coalesce(nullif(latitude,null),0) as latitude
	from bengkel left join lokasi on bengkel.id_bengkel = lokasi.id_bengkel`)

	utils.Logging.Println(query)
	var bengkel []object.GetAllBengkel
	rows, err := db.Query(query)
	if err != nil {
		return utils.Message(false, "bengkel.model.go, line:305 "+err.Error())
	}

	for rows.Next() {
		var b object.GetAllBengkel
		err = rows.Scan(
			&b.IdBengkel,
			&b.Nama,
			&b.Email,
			&b.Telepon,
			&b.Alamat,
			&b.Kota,
			&b.Provinsi,
			&b.Gambar,
			&b.Longitude,
			&b.Latitude,
		)
		if err != nil {
			return utils.Message(false, "bengkel.model.go, line:323 "+err.Error())
		}
		bengkel = append(bengkel, b)
	}
	response := utils.Message(true, "Success")
	response["response"] = bengkel
	utils.Logging.Println(response)
	return response
}

//Edit bengkel
func (bengkel *Bengkel) EditBengkel() map[string]interface{} {
	var idBengkel int
	createTime := time.Now()
	query := `update bengkel set 
	updated_at=$1, 
	nama= $2, 
	telepon=$3, 
	gambar=$4, 
	id_role=$5
	where id_bengkel= $6 returning id_bengkel`
	err := db.QueryRow(query, createTime.Format("01-02-2006 15:04:05"), bengkel.Nama, bengkel.Telepon, bengkel.Gambar, bengkel.IdRole, bengkel.IdBengkel).Scan(&idBengkel)
	utils.Logging.Println(idBengkel)
	utils.Logging.Println(query)

	if err != nil {
		utils.Logging.Println("bengkel.model.go, line:350")
		utils.Logging.Println(err.Error())
		return utils.Message(false, "Failed edit bengekel, "+err.Error())
	}

	response := utils.Message(true, "Success edit bengkel")
	return response
}

//Edit bengkel location
func (loc *Lokasi) EditLocation() map[string]interface{} {
	var idBengkel int
	createTime := time.Now()
	query := `update lokasi set 
	updated_at=$1, 
	alamat= $2, 
	kota=$3, 
	provinsi=$4, 
	longitude=$5,
	latitude=$6
	where id_bengkel=$7  returning id_bengkel`
	err := db.QueryRow(query, createTime.Format("01-02-2006 15:04:05"), loc.Alamat, loc.Kota, loc.Provinsi, loc.Longitude, loc.Latitude, loc.IdBengkel).Scan(&idBengkel)
	utils.Logging.Println(query)

	if err != nil {
		utils.Logging.Println("bengkel.model.go, line:375")
		utils.Logging.Println(err.Error())
		return utils.Message(false, "Failed edit location, "+err.Error())
	}

	response := utils.Message(true, "Success edit bengkel")
	return response
}

//delete bengkel
