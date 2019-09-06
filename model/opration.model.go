package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/GoToyota/object"

	"github.com/GoToyota/utils"
)

func AddOpration(opration []*Oprasional) (map[string]interface{}, error) {
	utils.Logging.Println("add opration")
	utils.Logging.Println(opration)
	valueStrings := make([]string, 0, len(opration))
	valueArgs := make([]interface{}, 0, len(opration)*3)
	dt := time.Now()

	for i, post := range opration {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
		valueArgs = append(valueArgs, post.IdUser)
		valueArgs = append(valueArgs, dt.Format("01-02-2006 15:04:05"))
		valueArgs = append(valueArgs, post.Hari)
		valueArgs = append(valueArgs, post.Buka)
		valueArgs = append(valueArgs, post.Tutup)
	}
	stmt := fmt.Sprintf("insert into oprasional (id_bengkel, created_at,hari, buka, tutup) values %s", strings.Join(valueStrings, ","))
	utils.Logging.Println(stmt)
	_, err := db.Exec(stmt, valueArgs...)
	if err != nil {
		return nil, err
	}

	response := utils.Message(true, "Success insert oprational")

	return response, nil
}

func GetOpration(idBengkel int) (map[string]interface{}, []object.Oprasional) {
	query := fmt.Sprintf(`
	select 
	coalesce(nullif(hari,''),' ') as hari ,
	coalesce(nullif(buka,''),' ') as "buka",
	coalesce(nullif(tutup,''),' ') as "tutup"
	from oprasional where id_bengkel = '%v'`, idBengkel)

	utils.Logging.Println(query)

	var opr []object.Oprasional
	rows, err := db.Query(query)
	if err != nil {
		return utils.Message(false, "opration.model.go, line:54 "+err.Error()), nil
	}

	for rows.Next() {
		var o object.Oprasional
		err = rows.Scan(
			&o.Hari,
			&o.Buka,
			&o.Tutup,
		)
		if err != nil {
			return utils.Message(false, "opratin.model.go, line:68 "+err.Error()), nil
		}
		opr = append(opr, o)
	}

	return nil, opr
}
