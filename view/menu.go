package view

import "fmt"

func MenuUtama() {
	fmt.Print("\033[H\033[2J")

	fmt.Println("\n------------------------------------------------")
	fmt.Println("          Selamat Datang di Warnet PMA")
	fmt.Println("------------------------------------------------")
	fmt.Println("Data yang tersedia: ")
	fmt.Println("[1]. Member Reguler")
	fmt.Println("[2]. Member VIP")
	fmt.Print("Masukkan data yang akan diolah: ")
}

func MenuReguler() {
	fmt.Println("\nMenu Member Reguler")
	fmt.Println("Pilihan Menu:")
	fmt.Println("[1]. Register Member")
	fmt.Println("[2]. Login")
	fmt.Println("[3]. View")
	fmt.Println("[4]. View By ID")
	fmt.Println("[5]. Menu Utama")
	fmt.Print("Masukkan Pilihan: ")
}

func MenuVip() {
	fmt.Println("\nMenu Member VIP")
	fmt.Println("Pilihan Menu:")
	fmt.Println("[1]. Register Member")
	fmt.Println("[2]. Login")
	fmt.Println("[3]. View")
	fmt.Println("[4]. View By ID")
	fmt.Println("[5]. Menu Utama")
	fmt.Print("Masukkan Pilihan: ")
}
