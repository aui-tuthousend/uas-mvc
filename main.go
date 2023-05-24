package main

import (
	"fmt"
	"uas-mvc/view"
)

func main() {

	for {
		var menu int
		view.MenuUtama()
		fmt.Scanln(&menu)
		switch menu {
		case 1:
			var pilih int
			view.MenuReguler()
			fmt.Scanln(&pilih)
		case 2:
			var pilih int
			view.MenuVip()
			fmt.Scanln(&pilih)
		}
	}
	//view.InsertMember()
	//view.InsertMemberVip()
	//view.ViewMember()
	//view.ViewMemberVip()
}
