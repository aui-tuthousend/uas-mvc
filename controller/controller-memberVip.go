package controller

import (
	"fmt"
	"uas-mvc/entities"
	"uas-mvc/model"
)

func ControllerFindAllMemberVip() []entities.MemberVip {
	return model.ModelViewAllMemberVip()
}

func ControllerInsertMemberVip(namaVip string, passVip int) {
	dataVip := entities.MemberVip{
		Nama: namaVip,
		Pass: passVip,
	}
	fmt.Println("testing passing data controller ", dataVip)
	model.ModelInsertMemberVip(dataVip)
}
