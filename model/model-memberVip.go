package model

import (
	"fmt"
	DB "uas-mvc/database"
	"uas-mvc/entities"
)

func ModelInsertMemberVip(container entities.MemberVip) {
	newGerbongVip := entities.LinkedlistMemberVip{}
	newGerbongVip.DataVip = container
	fmt.Println("test passing data to model :", newGerbongVip.DataVip)
	temp := &DB.DBMemberVip
	if temp.Next == nil {
		temp.Next = &newGerbongVip
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newGerbongVip
	}
}

func ModelViewAllMemberVip() []entities.MemberVip {
	temp := DB.DBMemberVip.Next
	membersVip := []entities.MemberVip{}
	for temp != nil {
		membersVip = append(membersVip, temp.DataVip)
		temp = temp.Next
	}
	return membersVip
}

func ModelViewByIdMemberVip(nama string) *entities.MemberVip {
	temp := &DB.DBMemberVip
	for temp != nil {
		if temp.DataVip.Nama == nama {
			return &temp.DataVip
		}
		temp = temp.Next
	}
	return nil
}
