package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	persegi := Persegi{5}
	persegiPanjang := PersegiPanjang{5, 4}

	luasPersegi := persegi.HitungLuas()
	luasPersegiPanjang := persegiPanjang.HitungLuas()

	fmt.Println("Luas Persegi :", luasPersegi)
	fmt.Println("Luas Persegi Panjang :", luasPersegiPanjang)

	fmt.Println("===")

	persegi = Persegi{Sisi: 4}
	luas := SeberapaLuas(persegi)
	fmt.Println("Luas Persegi :", luas)

	persegiPanjang = PersegiPanjang{Panjang: 6, Lebar: 5}
	luas = SeberapaLuas(persegiPanjang)
	fmt.Println("Luas Persegi Panjang :", luas)

}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
