package main

import "fmt"

type Figuras2D interface {
	Area() float64
}

type Cuadrado struct {
	Base float64
}
type Rectangulo struct {
	Base   float64
	Altura float64
}

func (c Cuadrado) Area() float64 {
	return c.Base * c.Base
}

func (r Rectangulo) Area() float64 {
	return r.Base * r.Base
}
func Calcular(f Figuras2D) {
	fmt.Printf("Area %2f \n", f.Area())
}
func main() {
	myCuadrado := Cuadrado{Base: 2}
	myRectangulo := Rectangulo{Base: 2.111, Altura: 4.434}
	Calcular(myCuadrado)
	Calcular(myRectangulo)

	//Lista de interfaces
	myInterface := []interface{}{"hola", 12, 14.9}
	fmt.Println(myInterface...)
}
