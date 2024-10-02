package main

import "fmt"

func main() {
	var name, nim, class string
	fmt.Scan(&name, &nim, &class)
	fmt.Print("Perkenalkan saya adalah ", name, " salah satu mahasiswa Prodi S1-IF dari kelas ", class, " dengan NIM ", nim)
}
