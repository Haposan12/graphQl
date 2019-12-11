package gql

import (
	"github.com/graphql-go/graphql"
	programProvider "graphQL/repository"
)

// constanta untuk sort
const (
	Sort1 = "created asc"
	Sort2 = "created desc"
	Sort3 = "program_name asc"
	Sort4 = "program_name desc"
)

func SpecialProgramResolver(p graphql.ResolveParams)(interface{}, error)  {
	page := p.Args["page"].(int)
	size := p.Args["size"].(int)
	sort := p.Args["sort"].(int)
	categoryID := p.Args["category"].(int)
	id := p.Args["id"].(int)

	var sortType string

	switch sort {
	case 0:
		sortType = "null"
	case 1:
		sortType = Sort1
	case 2:
		sortType = Sort2
	case 3:
		sortType = Sort3
	case 4:
		sortType = Sort4
	}

	//cek jika id nya gak 0
	if id != 0{
		result, err := programProvider.GetSpecialProgramById(id)
		//fmt.Println(result)
		if err != nil{
			return result, err
		}
		return result, nil
	}

	//hasil jika tidak ada id
	programs, err := programProvider.GetSpecialProgram(page, size, categoryID, sortType)
	if err != nil{
		return programs, err
	}
	return programs, nil
}