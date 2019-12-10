package models

import "time"

type Program struct {
	Id         			int
	Created    			time.Time
	CreatedBy  			string
	Modified   			time.Time
	ModifiedBy 			string
	Active     			bool
	IsDeleted  			bool
	Deleted    			*time.Time
	Deleted_by 			string
	ProgramName 		string
	ProgramImage 		string
	ProgramStartDate 	time.Time
	ProgramEndDate 		time.Time
	ProgramDescription 	string
	Card 				string
	OutletID			int
	MerchantId			int
	MerchantName 		string
	CategoryId			int
	Benefit				string
	TermsAndCondition	string
	Tier 				string
	RedeemRules			string
	RewardTarget		float64
	QRCodeId			string
}
