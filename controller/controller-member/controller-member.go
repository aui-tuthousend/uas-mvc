package controller_member

import (
	"uas-mvc/entities"
	model "uas-mvc/model/model-member"
)

func ControllerFindAllMember() []entities.Member {
	return model.ModelViewAllMember()
}

func ControllerInsertMember(nama string, pass int) {
	data := entities.Member{
		Nama: nama,
		Pass: pass,
	}
	//fmt.Println("testing passing data controller ", data)
	model.ModelInsertMember(data)
}

func ControllerViewByIdMember(nama string) *entities.Member {
	current := model.ModelViewByIdMember(nama)
	if current == nil {
		return nil
	}
	return current
}
