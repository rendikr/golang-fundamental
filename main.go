package main

import (
	"fmt"
)

func main() {
	luas, keliling := calculate(4, 5)
	fmt.Println(luas)
	fmt.Println(keliling)
}

func calculate(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	return
}
