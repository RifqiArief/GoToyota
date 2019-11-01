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

func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return utils.Message(false, "Invalid email format"), false
	}
	if len(user.Password) < 2 {
		return utils.Message(false, "Password is required"), false
	}
	query := fmt.Sprintf(`
	select 
	coalesce(nullif(id_user,null),0) as id_user,
	coalesce(nullif(email,''),' ') as email ,
	coalesce(nullif(password,''),' ') as "password",
	coalesce(nullif(nama,''),' ') as nama,
	coalesce(nullif(telepon,''),' ') as telepon,
	coalesce(nullif(gambar,''),' ') as gambar,
	coalesce(nullif(alamat,''),' ') as alamat,
	coalesce(nullif(kota,''),' ') as kota,
	coalesce(nullif(created_at,null),'2000-01-01 00:00:00') as created_at,
	coalesce(nullif(updated_at,null),'2000-01-01 00:00:00') as updated_at,
	coalesce(nullif(delete_at,null),'2000-01-01 00:00:00') as delete_at
	from users where email = '%s'`, user.Email)
	utils.Logging.Println(query)
	var temp User
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

//add user
func (user *User) AddUser() (map[string]interface{}, bool) {
	res, ok := user.Validate()
	if !ok {
		utils.Logging.Println("model/userModel.go, line:52")
		utils.Logging.Println(res)
		return res, false
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.Message(false, err.Error()), false
	}

	createTime := time.Now()
	user.Password = string(hash)
	var idUser int
	query := `insert into users (
		created_at, email, password, nama, telepon, gambar,alamat,kota)
		values ($1,$2,$3,$4,$5,$6,$7,$8) returning id_user`
	err = db.QueryRow(query, createTime.Format("01-02-2006 15:04:05"), user.Email, user.Password, user.Nama, user.Telepon, user.Gambar, user.Alamat, user.Kota).Scan(&idUser)
	utils.Logging.Println(idUser)
	utils.Logging.Println(query)
	if err != nil {
		utils.Logging.
			Println("model/user.model.go, line:72")
		utils.Logging.Println(err.Error())
		return utils.Message(false, "Failed insert user"), false
	}

	data := &object.RegisterUser{
		IdUser:  idUser,
		Nama:    user.Nama,
		Email:   user.Email,
		Telepon: user.Telepon,
		Alamat:  user.Alamat,
		Kota:    user.Kota,
		Gambar:  user.Gambar,
	}

	response := utils.Message(true, "Success registration")
	response["response"] = data

	return response, true
}

//login
func LoginUser(email, password string) (map[string]interface{}, bool) {

	query := fmt.Sprintf(`
	select 
	coalesce(nullif(id_user,null),-1) as id_user,
	coalesce(nullif(email,''),' ') as email ,
	coalesce(nullif(password,''),' ') as "password",
	coalesce(nullif(nama,''),' ') as nama,
	coalesce(nullif(telepon,''),' ') as telepon,
	coalesce(nullif(gambar,''),' ') as gambar,
	coalesce(nullif(alamat,''),' ') as alamat,
	coalesce(nullif(kota,''),' ') as kota,
	coalesce(nullif(created_at,null),'2000-01-01 00:00:00') as created_at,
	coalesce(nullif(updated_at,null),'2000-01-01 00:00:00') as updated_at,
	coalesce(nullif(delete_at,null),'2000-01-01 00:00:00') as delete_at
	from users where email = '%s'`, email)

	var user User
	err := db.Get(&user, query)
	utils.Logging.Println(user.IdUser)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return utils.Message(false, "Email not registered"), false
		}
		return utils.Message(false, "model/userModel.go, line:118 "+err.Error()), false
	}

	//chek password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "Wrong password, please try again"), false
	}

	tk := &Token{UserId: user.IdUser}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return utils.Message(false, "Failed generate token string"), false
	}

	data := &object.LoginUser{
		IdUser:  user.IdUser,
		Nama:    user.Nama,
		Email:   user.Email,
		Telepon: user.Telepon,
		Alamat:  user.Alamat,
		Kota:    user.Kota,
		Gambar:  user.Gambar,
		Token:   tokenString,
	}

	response := utils.Message(true, "Success Login")
	response["response"] = data
	return response, true
}

//Get all user
func GetAllUser() map[string]interface{} {
	query := fmt.Sprintf(`
	select 
	coalesce(nullif(id_user,null),-1) as id_user,
	coalesce(nullif(nama,''),' ') as nama,
	coalesce(nullif(email,''),' ') as email ,
	coalesce(nullif(telepon,''),' ') as telepon,
	coalesce(nullif(alamat,''),' ') as alamat,
	coalesce(nullif(kota,''),' ') as kota,
	coalesce(nullif(gambar,''),' ') as gambar,
	coalesce(nullif(created_at,null),'2000-01-01 00:00:00') as created_at,
	coalesce(nullif(updated_at,null),'2000-01-01 00:00:00') as updated_at,
	coalesce(nullif(delete_at,null),'2000-01-01 00:00:00') as delete_at
	from users`)

	var user []object.GetAllUser
	rows, err := db.Query(query)
	if err != nil {
		return utils.Message(false, "user.model.go, line:169 "+err.Error())
	}

	for rows.Next() {
		var u object.GetAllUser
		err = rows.Scan(
			&u.IdUser,
			&u.Nama,
			&u.Email,
			&u.Telepon,
			&u.Alamat,
			&u.Kota,
			&u.Gambar,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.DeleteAt,
		)
		if err != nil {
			return utils.Message(false, "user.model.go, line:187 "+err.Error())
		}
		user = append(user, u)
	}
	response := utils.Message(true, "Success")
	response["response"] = user
	utils.Logging.Println(response)
	return response
}

//edit profil
func (user *User) UpdateUser() map[string]interface{} {
	var idUser int
	querycheck := fmt.Sprintf(`
	select 
	coalesce(nullif(id_user,null),0) as id_user 
	from users where id_user = %v`, user.IdUser)

	err := db.Get(&idUser, querycheck)
	utils.Logging.Println(querycheck)
	if err != nil && err != sql.ErrNoRows {
		return utils.Message(false, "user.model.go:209 "+err.Error())
	}

	if idUser == 0 {
		return utils.Message(false, "No STNK not registered")
	}

	query := fmt.Sprintf(`
	update users set 
	updated_at=$1, 
	nama= $2, 
	telepon=$3, 
	gambar=$4, 
	alamat=$5, 
	kota=$6 
	where id_user= $7 returning id_user`)

	createTime := time.Now()
	err = db.QueryRow(query, createTime.Format("01-02-2006 15:04:05"), user.Nama, user.Telepon, user.Gambar, user.Alamat, user.Kota, user.IdUser).Scan(&idUser)
	utils.Logging.Println(query)
	if err != nil {
		utils.Logging.Println("user.model.go, line:231")
		utils.Logging.Println(err.Error())
		return utils.Message(false, "Failed update kendaraan")
	}

	response := utils.Message(true, "Success edit profile")
	return response
}

//ganti password
func (data *ChangePassword) ChangePassword() map[string]interface{} {
	query := fmt.Sprintf(`
	select 
	coalesce(nullif(id_user,null),0) as id_user,
	coalesce(nullif(password,''),' ') as "password"
	from users where id_user = '%v'`, data.IdUser)

	var user ChangePassword
	err := db.Get(&user, query)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return utils.Message(false, "User not registered")
		}
		return utils.Message(false, "user.model.go, line:255 "+err.Error())
	}

	//chek password
	err = bcrypt.CompareHashAndPassword([]byte(user.OldPassword), []byte(data.OldPassword))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "Wrong password, please try again")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return utils.Message(false, err.Error())
	}

	query = fmt.Sprintf(`update users set "password"='%v' where id_user=%v`, string(hash), user.IdUser)

	_, err = db.Query(query)
	if err != nil {
		return utils.Message(false, "failed change password. Error : "+err.Error())
	}

	response := utils.Message(true, "Success change password")
	return response
}
