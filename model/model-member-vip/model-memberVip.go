package model_member_vip

import (
	DB "uas-mvc/database"
	"uas-mvc/entities"
)

func ModelInsertMemberVip(container entities.MemberVip) {
	newGerbongVip := entities.LinkedlistMemberVip{}
	newGerbongVip.DataVip = container
	//fmt.Println("test passing data to model :", newGerbongVip.DataVip)
	tempvip := &DB.DBMemberVip
	if tempvip.Next == nil {
		tempvip.Next = &newGerbongVip
	} else {
		for tempvip.Next != nil {
			tempvip = tempvip.Next
		}
		tempvip.Next = &newGerbongVip
	}
}

func ModelViewAllMemberVip() []entities.MemberVip {
	tempvip := DB.DBMemberVip.Next
	membersVip := []entities.MemberVip{}
	for tempvip != nil {
		membersVip = append(membersVip, tempvip.DataVip)
		tempvip = tempvip.Next
	}
	return membersVip
}

func ModelViewByIdMemberVip(nama string) *entities.MemberVip {
	tempvip := &DB.DBMemberVip
	for tempvip != nil {
		if tempvip.DataVip.Nama == nama {
			return &tempvip.DataVip
		}
		tempvip = tempvip.Next
	}
	return nil
}
