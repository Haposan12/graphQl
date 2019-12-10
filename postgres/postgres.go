package postgres

import (
	"database/sql"
	"fmt"
	"graphQL/models"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func New(conn string) (*DB, error)  {
	db, err := sql.Open("postgres", conn)
	if err != nil{
		return nil, err
	}

	err = db.Ping()
	if err != nil{
		return nil, err
	}

	return &DB{db}, nil
}

func ConnString(host string, port int, user string, password string, dbName string) string  {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
}


func (d *DB) GetSpecialProgramById(id int)(data []models.Program)   {
	query := `SELECT * FROM program WHERE id=$1`

	stmt, err := d.Prepare(query)
	if err != nil{
		log.Println(err)
	}
	rows, err := stmt.Query(id)
	defer rows.Close()

	if err != nil{
		fmt.Println(err)
	}
	var program models.Program
	for rows.Next(){
		err = rows.Scan(
			&program.Id,
			&program.Created,
			&program.CreatedBy,
			&program.Modified,
			&program.ModifiedBy,
			&program.Active,
			&program.IsDeleted,
			&program.Deleted,
			&program.Deleted_by,
			&program.ProgramName,
			&program.ProgramImage,
			&program.ProgramStartDate,
			&program.ProgramEndDate,
			&program.ProgramDescription,
			&program.Card,
			&program.OutletID,
			&program.MerchantId,
			&program.CategoryId,
			&program.Benefit,
			&program.TermsAndCondition,
			&program.Tier,
			&program.RedeemRules,
			&program.RewardTarget,
			&program.QRCodeId)

		data = append(data, program)
	}
	return
}

func (d *DB) GetSpecialProgram(page, size, categoryId int, sortType string)(specialPrograms []models.Program)   {
	query := `SELECT * FROM program LIMIT $1 OFFSET $2`

	if sortType != "null"{
		query = `SELECT * FROM program ORDER BY `+sortType+` LIMIT $1 OFFSET $2`
	}
	stmt, err := d.Prepare(query)
	if err != nil{
		log.Println(err)
	}
	rows, err := stmt.Query(size, page)
	defer rows.Close()

	if err != nil{
		fmt.Println(err)
	}
	var program models.Program
	for rows.Next(){
		err = rows.Scan(
			&program.Id,
			&program.Created,
			&program.CreatedBy,
			&program.Modified,
			&program.ModifiedBy,
			&program.Active,
			&program.IsDeleted,
			&program.Deleted,
			&program.Deleted_by,
			&program.ProgramName,
			&program.ProgramImage,
			&program.ProgramStartDate,
			&program.ProgramEndDate,
			&program.ProgramDescription,
			&program.Card,
			&program.OutletID,
			&program.MerchantId,
			&program.CategoryId,
			&program.Benefit,
			&program.TermsAndCondition,
			&program.Tier,
			&program.RedeemRules,
			&program.RewardTarget,
			&program.QRCodeId)

		// jika category id ga 0
		if categoryId != 0 && program.CategoryId == categoryId{
			specialPrograms = append(specialPrograms, program)
		}else if categoryId == 0{
			specialPrograms = append(specialPrograms, program)
		}

	}
	return
}

func (d *DB) GetMerchantName(data []models.Program)([]models.Program)  {
	for i:=0;i<len(data) ;i++  {
		//get merchant name
		name := d.GetMerchantNameFromDb(data[i].MerchantId)
		data[i].MerchantName = name
	}

	return data
}

func (d *DB) GetMerchantNameFromDb(id int)(name string)  {
	query := `SELECT merchant_name FROM merchant where id=$1`

	 _ = d.QueryRow(query, id).Scan(&name)
	 return
}