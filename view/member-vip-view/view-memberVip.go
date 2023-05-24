package member_vip_view

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

func ViewByIdMemberVip() {
	var nama string
	fmt.Print("Masukan Nama Member: ")
	fmt.Scan(&nama)
	current := controller.ControllerViewByIdMemberVip(nama)
	if current != nil {
		fmt.Println(current.Nama)
		fmt.Println(current.Pass)
	} else {
		fmt.Println("Nama member tidak di temukan")
	}
}
