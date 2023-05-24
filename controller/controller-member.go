package controller

import (
	"fmt"
	"uas-mvc/entities"
	"uas-mvc/model"
)

func ControllerFindAllMember() []entities.Member {
	return model.ModelViewAllMember()
}

func ControllerInsertMember(nama string, pass int) {
	data := entities.Member{
		Nama: nama,
		Pass: pass,
	}
	fmt.Println("testing passing data controller ", data)
	model.ModelInsertMember(data)
}
