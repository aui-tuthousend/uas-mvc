package member_view

import (
	"fmt"
	controller "uas-mvc/controller"
)

func InsertMember() {
	var nama string
	var pass int
	fmt.Print("masukkan Nama Member Baru : ")
	fmt.Scan(&nama)
	fmt.Print("masukkan Password : ")
	fmt.Scan(&pass)
	controller.ControllerInsertMember(nama, pass)
	fmt.Println("Member Berhasil Dibuat!")
}

func ViewMember() {
	allMember := controller.ControllerFindAllMember()
	for _, member := range allMember {
		fmt.Println("Nama Member :", member.Nama)
		fmt.Println("Pass Member :", member.Pass)
	}
}

func ViewByIdMember() {
	var nama string
	fmt.Print("Masukan Nama Member: ")
	fmt.Scan(&nama)
	current := controller.ControllerViewByIdMember(nama)
	if current != nil {
		fmt.Println(current.Nama)
		fmt.Println(current.Pass)
	} else {
		fmt.Println("Nama member tidak di temukan")
	}
}
