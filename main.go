package main

import (
	"fmt"
	"math"
)

type converter struct{}

type (
	min float64

	sec float64

	cm float64

	feet      float64
	millisec  float64
	farenheit float64
	celsius   float64
	rad       float64
	deg       float64
	kg        float64
	pounds    float64
	naira     float64
	dollar    float64
)

func (a converter) minToSec(m min) sec {

	return sec(m * 60)
}

func (a converter) secTomin(s sec) min {
	return min(s / 60)
}

func (a converter) cmToFeet(c cm) feet {
	return feet(c / 30.48)
}

func (a converter) feetToCent(ft feet) cm {
	return cm(ft * 30.48)
}

func (a converter) celsiusToFareheit(cel celsius) farenheit {
	return farenheit((9 * cel / 5) + (32))
}

func (a converter) farenheitTocelsius(F farenheit) celsius {
	return celsius((F - 32) * 5 / 9)
}

func (a converter) radToDeg(r rad) deg {
	return deg(r * 180 / math.Pi)
}

func (a converter) degToRad(d deg) rad {
	return rad(d * math.Pi / 180)
}

func (a converter) secTomillisec(s sec) millisec {
	return millisec(s * 1000)
}

func (a converter) millisecToSec(ms millisec) sec {
	return sec(ms / 1000)
}

func (a converter) kgToPound(k kg) pounds {
	return pounds(k * 2.205)
}

func (a converter) poundToKg(p pounds) kg {
	return kg(p / 2.205)
}

func (a converter) nairaToDollar(N naira) dollar {
	return dollar(N * 385)
}

func (a converter) dollarToNaira(d dollar) naira {
	return naira(d / 385)
}
func main() {
	cvr := converter{}
	fmt.Println(cvr.minToSec(2))
	fmt.Println(cvr.secTomin(1200))
	fmt.Println(cvr.cmToFeet(3048))
	fmt.Println(cvr.feetToCent(10))
	fmt.Println(cvr.radToDeg(6.284))
	fmt.Println(cvr.degToRad(1800))
	fmt.Println(cvr.millisecToSec(3000))
	fmt.Println(cvr.secTomillisec(4))
	fmt.Println(cvr.poundToKg(2000))
	fmt.Println(cvr.kgToPound(1000))
	fmt.Println(cvr.celsiusToFareheit(100))
	fmt.Println(cvr.farenheitTocelsius(212))
	fmt.Println(cvr.nairaToDollar(360000000))
	fmt.Println(cvr.dollarToNaira(5300))
}
