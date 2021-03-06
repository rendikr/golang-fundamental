package main

import "fmt"

type Luas interface {
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

	fmt.Println(luasPersegi)
	fmt.Println(luasPersegiPanjang)
}
