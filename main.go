package main

import (
	"fmt"
	view "uas-mvc/view"
	reguler "uas-mvc/view/member-view"
	vip "uas-mvc/view/member-vip-view"
)

func main() {

	var lanjut string
	var menu int

	for {
		view.MenuUtama()
		fmt.Scanln(&menu)
		//fmt.Println(menu)

		if menu == 1 {
			var pilih1 int
			view.MenuReguler()
			fmt.Scanln(&pilih1)
			//fmt.Println(pilih1)

			switch pilih1 {
			case 1:
				{
					reguler.InsertMember()
				}
			case 2:
				{
					fmt.Println("blom")
				}
			case 3:
				{
					reguler.ViewMember()
				}
			case 4:
				{
					reguler.ViewByIdMember()
				}
			case 5:
				continue
			default:
				{
					fmt.Println("menu tidak ada")
				}
			}
		} else if menu == 2 {
			var pilih2 int
			view.MenuVip()
			fmt.Scanln(&pilih2)
			//fmt.Println(pilih2)

			switch pilih2 {
			case 1:
				{
					vip.InsertMemberVip()
				}
			case 2:
				{
					fmt.Println("blom")
				}
			case 3:
				{
					vip.ViewMemberVip()
				}
			case 4:
				{
					vip.ViewByIdMemberVip()
				}
			case 5:
				continue
			default:
				{
					fmt.Println("menu tidak ada")
				}
			}
		} else {
			fmt.Println("menu tidak ada")
		}

		fmt.Print("\nApakah ingin lanjut? [y/t]: ")
		fmt.Scanln(&lanjut)
	}
}
