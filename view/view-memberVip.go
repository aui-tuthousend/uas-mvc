package view

import (
	"fmt"
	"uas-mvc/controller"
)

func InsertMemberVip() {
	var nama string
	var pass int
	fmt.Print("masukkan Nama Member VIP Baru : ")
	fmt.Scan(&nama)
	fmt.Print("masukkan Password : ")
	fmt.Scan(&pass)
	controller.ControllerInsertMemberVip(nama, pass)
}

func ViewMemberVip() {
	allMember := controller.ControllerFindAllMemberVip()
	for _, member := range allMember {
		fmt.Println("Nama Member VIP :", member.Nama)
		fmt.Println("Pass Member :", member.Pass)
	}
}
