package main

import (
	"fmt"
	"math"
)

func main() {
	menu := true
	opc := 0

	for menu {
		fmt.Printf("\t\t Menu\n \t1.- Fibonacci\n\t2.- swap\n\t3. - getMax\n\t4. - Impares Generador\n\t5. - Salir\n")
		fmt.Println("opc: ")
		fmt.Scanln(&opc)
		switch opc {
		case 1:
			r := 0
			fmt.Println("Que numero quieres de Fibonacci (posicion): ")
			fmt.Scanln(&r)
			v :=fibonacci(r)
			fmt.Println(v)

		case 2:
			a, b := 0, 0
			fmt.Println("Valor de a: ")
			fmt.Scanln(&a)
			fmt.Println("Valor de b: ")
			fmt.Scanln(&b)
			swap(&a,&b)
			fmt.Println(a,b)

		case 3:
			var s [] int
			n :=0
			value := 0
			fmt.Println("Cuantos numeros ingresaras: ")
			fmt.Scanln(&n)
			for i:=0; i<n; i++ {
				fmt.Println("int ",i," : ")
				fmt.Scanln(&value)
				s = append(s,1)
			}
			x := getMax(s...)
			fmt.Println("Get Max : ", x)

		case 4:
			f := oddGenerator()
			for i:=0; i<10; i++ {
				fmt.Println(f())
			}
			break
		case 5:
			menu = false
		}
	}
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func getMax(args ...int)  float64 {
	max := math.Inf(-8)
	for _, value := range args {
		if max < float64(value) {
			max = float64(value)
		}
	}
	return max
}

func oddGenerator() func() uint {
	n := uint(0)
	return func() uint {
		odd := n * 2 + 1
		n++
		return odd
	}
}