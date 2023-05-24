package view

import (
	"fmt"
	"uas-mvc/controller"
)

func InsertMember() {
	var nama string
	var pass int
	fmt.Print("masukkan Nama Member Baru : ")
	fmt.Scan(&nama)
	fmt.Print("masukkan Password : ")
	fmt.Scan(&pass)
	controller.ControllerInsertMember(nama, pass)
}

func ViewMember() {
	allMember := controller.ControllerFindAllMember()
	for _, member := range allMember {
		fmt.Println("Nama Member :", member.Nama)
		fmt.Println("Pass Member :", member.Pass)
	}
}
