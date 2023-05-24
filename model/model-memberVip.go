package model

import (
	"fmt"
	"uas-mvc/database"
	"uas-mvc/entities"
)

func ModelInsertMemberVip(container entities.MemberVip) {
	newGerbongVip := entities.LinkedlistMemberVip{}
	newGerbongVip.DataVip = container
	fmt.Println("test passing data to model :", newGerbongVip.DataVip)
	temp := &database.DBMemberVip
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
	temp := database.DBMemberVip.Next
	membersVip := []entities.MemberVip{}
	for temp != nil {
		membersVip = append(membersVip, temp.DataVip)
		temp = temp.Next
	}
	return membersVip
}
