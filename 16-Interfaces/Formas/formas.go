package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}
type retangulo struct {
	largura float64
	altura  float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func imprimirArea(f forma) {
	fmt.Println(f.area())
}

func main() {

	r := retangulo{10, 15}
	imprimirArea(r)

	c := circulo{10}
	imprimirArea(c)

}
