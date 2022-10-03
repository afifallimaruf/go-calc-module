package go_calc_module

import (
	"fmt"
)

type Calc struct {
	Number1 int
	Number2 int
}

type Operations interface {
	Tambah()
	Kurang()
	Kali()
	Bagi()
}

func (c Calc) Tambah() {
	result := c.Number1 + c.Number2
	fmt.Printf("%d + %d = %d\n", c.Number1, c.Number2, result)
}

func (c Calc) Kurang() {
	result := c.Number1 - c.Number2
	fmt.Printf("%d - %d = %d\n", c.Number1, c.Number2, result)
}

func (c Calc) Kali() {
	result := c.Number1 * c.Number2
	fmt.Printf("%d * %d = %d\n", c.Number1, c.Number2, result)
}

func (c Calc) Bagi() {
	var number1 float64 = float64(c.Number1)
	var number2 float64 = float64(c.Number2)
	result := number1 / number2
	fmt.Printf("%.1f / %.1f = %.1f\n", number1, number2, result)
}
