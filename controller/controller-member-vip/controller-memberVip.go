package controller_member_vip

import (
	"fmt"
	"uas-mvc/entities"
	modelVip "uas-mvc/model/model-member-vip"
)

func ControllerFindAllMemberVip() []entities.MemberVip {
	return modelVip.ModelViewAllMemberVip()
}

func ControllerInsertMemberVip(namaVip string, passVip int) {
	dataVip := entities.MemberVip{
		Nama: namaVip,
		Pass: passVip,
	}
	fmt.Println("testing passing data controller ", dataVip)
	modelVip.ModelInsertMemberVip(dataVip)
}

func ControllerViewByIdMemberVip(nama string) *entities.MemberVip {
	current := modelVip.ModelViewByIdMemberVip(nama)
	if current == nil {
		return nil
	}
	return current
}
