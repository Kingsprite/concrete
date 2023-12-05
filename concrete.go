package main

import (
	"fmt"
)

func calcbeta(fcuk float64) (beta float64) {
	if fcuk <= 50 {
		beta = 0.8
	} else {
		beta = fcuk*(-0.002) + 0.9
	}
	return beta
}

func calcalpha(fcuk float64) (alpha float64) {
	alpha = calcbeta(fcuk) + 0.2
	return alpha
}

func main() {
	fmt.Println("输入fcuk")
	var fcuk float64
	var alpha, beta float64
	//calc fcuk
	fmt.Scanln(&fcuk)
	if fcuk == 0 {
		fmt.Println("输入alpha")
		fmt.Scanln(&alpha)
		fmt.Println("输入beta")
		fmt.Scanln(&beta)
	} else {
		beta = calcbeta(fcuk)
		alpha = calcalpha(fcuk)
	}
	fmt.Println()

	//calc x

	//calc as
	fmt.Println("输入as")

}
