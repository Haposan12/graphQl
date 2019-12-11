package repository

import (
	"fmt"
	"graphQL/models"
	"graphQL/config"
)
//get special program from db by id
func GetSpecialProgramById(id int)(data []models.Program, err error)  {
	var program models.Program
	query := `SELECT t1.*, t2.merchantname FROM program t1 JOIN merchant t2 ON t1.merchantid=t2.id WHERE t1.id=$1`

	rs, err := config.GlobalDbSql.Query(query, id)
	defer rs.Close()

	for rs.Next(){
		err = rs.Scan(
			&program.Id,
			&program.Created,
			&program.CreatedBy,
			&program.Modified,
			&program.ModifiedBy,
			&program.Active,
			&program.IsDeleted,
			&program.Deleted,
			&program.DeletedBy,
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
			&program.RedemRules,
			&program.RewardTarget,
			&program.QRCodeId,
			&program.MerchantName)
		fmt.Println(program.MerchantName)
		data = append(data, program)
	}

	return
}

//get special program without id
func GetSpecialProgram(page, size, categoryId int, sortType string) (data []models.Program, err error) {
	var program models.Program

	query := `SELECT t1.*, t2.merchant_name FROM program t1 JOIN merchant t2 ON t1.merchantid=t2.id LIMIT $1 OFFSET $2`

	if sortType != "null"{
		query = `SELECT t1.*, t2.merchant_name FROM program t1 JOIN merchant t2 ON t1.merchantid=t2.id ORDER BY `+sortType+` LIMIT $1 OFFSET $2`
	}

	rs, err := config.GlobalDbSql.Query(query, size, page)
	defer rs.Close()

	for rs.Next(){
		err = rs.Scan(
			&program.Id,
			&program.Created,
			&program.CreatedBy,
			&program.Modified,
			&program.ModifiedBy,
			&program.Active,
			&program.IsDeleted,
			&program.Deleted,
			&program.DeletedBy,
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
			&program.RedemRules,
			&program.RewardTarget,
			&program.QRCodeId,
			&program.MerchantName)

		// jika category id ga 0
		if categoryId != 0 && program.CategoryId == categoryId{
			data = append(data, program)
		}else if categoryId == 0{
			data = append(data, program)
		}
	}
	return
}