package member_vip_view

import (
	"fmt"
	controllerVip "uas-mvc/controller/controller-member-vip"
)

func InsertMemberVip() {
	var namaVip string
	var pass int
	fmt.Print("masukkan Nama Member VIP Baru : ")
	fmt.Scan(&namaVip)
	fmt.Print("masukkan Password : ")
	fmt.Scan(&pass)
	controllerVip.ControllerInsertMemberVip(namaVip, pass)
}

func ViewMemberVip() {
	allMemberVip := controllerVip.ControllerFindAllMemberVip()
	for _, member := range allMemberVip {
		fmt.Println("Nama Member VIP :", member.Nama)
		fmt.Println("Pass Member :", member.Pass)
	}
}

func ViewByIdMemberVip() {
	var namaVip string
	fmt.Print("Masukan Nama Member: ")
	fmt.Scan(&namaVip)
	currentVip := controllerVip.ControllerViewByIdMemberVip(namaVip)
	if currentVip != nil {
		fmt.Println(currentVip.Nama)
		fmt.Println(currentVip.Pass)
	} else {
		fmt.Println("Nama member tidak di temukan")
	}
}
