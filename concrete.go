package main

import (
	"fmt"
	"math"
)

func calcβ(fcuk float64) (β float64) {
	if fcuk <= 50 {
		β = 0.8
	} else {
		β = fcuk*(-0.002) + 0.9
	}
	return β
}

func calcα(fcuk float64) (α float64) {
	α = calcβ(fcuk) + 0.2
	return α
}

func RoundWithPrecision(x float64, precision int) float64 {
	var p = math.Pow10(precision)
	if precision == 0 {
		return math.Round(x)
	}
	return math.Round(x*p) / p
}

func calcxi(N, α, fc, b, h0, γ float64) (xi float64) {
	xi = RoundWithPrecision(γ*N*1000/(α*fc*b*h0), 3)
	return xi
}

func calcx(xi, h0 float64) (x float64) {
	x = xi * h0
	return x
}

func calcAs(M, N, e, fc, fy, h, h0, as, b, x, α, γ float64, flag bool) (As float64) {
	ei := RoundWithPrecision(M*1000/N, 3) + math.Max(h/30, 20)
	if flag {
		e = ei - h/2 + as
		As = γ * N * 1000 * e / (fy * (h0 - as))
	} else {
		e = ei + h/2 - as
		fmt.Println(h0)
		fmt.Println(x)
		fmt.Println(fy * (h0 - as))
		As = (γ*N*1000*e - α*fc*b*x*(h0-x/2)) / (fy * (h0 - as))
	}
	return As
}

func main() {
	fmt.Print("输入fcuk:")
	var fcuk float64
	var α, β float64
	var xi float64
	var M, N, fc, b, h0, h, ρ float64
	var x float64
	var as float64
	var ξ float64
	var e, fy float64
	var As float64
	var flag bool
	var γ float64
	//calc fcuk
	fmt.Scanln(&fcuk)
	if fcuk == 0 {
		fmt.Print("输入α:")
		fmt.Scanln(&α)
		fmt.Print("输入β:")
		fmt.Scanln(&β)
	} else {
		β = calcβ(fcuk)
		α = calcα(fcuk)
	}
	fmt.Print("输入γ0:")
	fmt.Scanln(&γ)

	//calc x
	fmt.Print("输入N(单位为kN·M）:")
	fmt.Scanln(&N)
	fmt.Print("输入fc:")
	fmt.Scanln(&fc)
	fmt.Print("输入b:")
	fmt.Scanln(&b)
	fmt.Print("输入as:")
	fmt.Scanln(&as)
	fmt.Print("输入h:")
	fmt.Scanln(&h)
	h0 = h - as

	//calc As
	fmt.Print("输入M:")
	fmt.Scanln(&M)
	fmt.Print("输入fy:")
	fmt.Scanln(&fy)
	fmt.Print("输入ξb:")
	fmt.Scanln(&ξ)
	fmt.Print("输入ρ:")
	fmt.Scanln(&ρ)
	for {
		xi = calcxi(N, α, fc, b, h0, γ)
		x = calcx(xi, h0)
		if x < 2*as {
			flag = true
			As = calcAs(M, N, e, fc, fy, h, h0, as, b, x, α, γ, flag)
		} else if x < xi*b*h0 {
			//fmt.Println("小于")
			flag = false
			As = calcAs(M, N, e, fc, fy, h, h0, as, b, x, α, γ, flag)
		} else {
			fmt.Println("大于")
			flag = false
			fmt.Println("重新设计截面")
			fmt.Print("重新输入b:")
			fmt.Scanln(&b)
			fmt.Print("重新输入h:")
			fmt.Scanln(&h)
			continue
		}
		if 2*As/(b*h) >= ρ {
			fmt.Printf("As=%f", As)
			break
		} else {
			fmt.Println("重新设计截面")
			fmt.Print("重新输入b:")
			fmt.Scanln(&b)
			fmt.Print("重新输入h:")
			fmt.Scanln(&h)
		}
	}
}
