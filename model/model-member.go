package model

import (
	DB "uas-mvc/database"
	"uas-mvc/entities"
)

func ModelInsertMember(container entities.Member) bool {
	newGerbong := entities.LinkedlistMember{}
	newGerbong.Data = container
	//fmt.Println("test passing data to model :", newGerbong.Data)
	temp := &DB.DBMember
	if temp.Next == nil {
		temp.Next = &newGerbong
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newGerbong
	}
	return true
}

func ModelViewAllMember() []entities.Member {
	temp := DB.DBMember.Next
	members := []entities.Member{}
	for temp != nil {
		members = append(members, temp.Data)
		temp = temp.Next
	}
	return members
}

func ModelViewByIdMember(nama string) *entities.Member {
	temp := &DB.DBMember
	for temp != nil {
		if temp.Data.Nama == nama {
			return &temp.Data
		}
		temp = temp.Next
	}
	return nil
}
