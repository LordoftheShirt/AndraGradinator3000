package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var buffer string
	var a, b, c float64

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("format ax² + bx + c = 0, ange variablerna. Kom ihåg använda punkt istället för kommatecken vid användning av decimaler.")
	fmt.Printf("a:")
	buffer, _ = reader.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	a, _ = strconv.ParseFloat(buffer, 64)

	fmt.Printf("b:")
	buffer, _ = reader.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	b, _ = strconv.ParseFloat(buffer, 64)

	fmt.Printf("c:")
	buffer, _ = reader.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	c, _ = strconv.ParseFloat(buffer, 64)
	fmt.Println(a, b, c)

	// ax² + bx + c = 0
	//a := 2.
	//b := 2.
	//c := 20.
	// Huvud variabler

	fmt.Printf("\n Din ekvation: \n\n %vx²", a)
	if b < 0 {
		fmt.Printf(" %vx", b)
	} else {
		fmt.Printf(" + %vx", b)
	}

	if c < 0 {
		fmt.Printf(" %v = 0", c)
	} else {
		fmt.Printf(" + %v = 0", c)
	}

	if a < 0 {
		a = a * -1
		b = b * -1
		c = c * -1
	}

	b = b / a
	c = c / a

	fmt.Printf("\n\n x²")
	if b < 0 {
		fmt.Printf(" %vx", b)
	} else {
		fmt.Printf(" + %vx", b)
	}

	if c < 0 {
		fmt.Printf(" %v = 0", c)
	} else {
		fmt.Printf(" + %v = 0", c)
	}
	b = (b * -1) / 2
	d := (b*b + c*-1)
	fmt.Printf("\n\n x = %v ± √%v", b, d)
	if d >= 0 {
		e := math.Sqrt(d)
		fmt.Printf("\n\n √%v = %v", d, e)
		x1 := b + e
		x2 := b - e
		fmt.Printf("\n\n x₁ = %v \n x₂ = %v", x1, x2)
	} else {
		d = d * (-1)
		f := math.Sqrt(d)
		fmt.Printf("\n\n x = %v ± %vi", b, f)
		fmt.Printf("\n\n x₁ = %v + %vi \n x₂ = %v - %vi", b, f, b, f)
	}
}
