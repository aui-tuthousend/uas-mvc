package model

import (
	"uas-mvc/database"
	"uas-mvc/entities"
)

func ModelInsertMember(container entities.Member) {
	newGerbong := entities.LinkedlistMember{}
	newGerbong.Data = container
	//fmt.Println("test passing data to model :", newGerbong.Data)
	temp := &database.DBMember
	if temp.Next == nil {
		temp.Next = &newGerbong
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newGerbong
	}
}

func ModelViewAllMember() []entities.Member {
	temp := database.DBMember.Next
	members := []entities.Member{}
	for temp != nil {
		members = append(members, temp.Data)
		temp = temp.Next
	}
	return members
}
