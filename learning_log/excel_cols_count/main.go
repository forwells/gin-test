package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Excel列数计算:")

	fmt.Println("A=", 'A')
	fmt.Println("Z=", 'Z')
	fmt.Println("Z-A=", 'Z'-'A')
	fmt.Printf("A=%d Z=%d的\n", 'A', 'Z')
	base := 'Z' - 'A' + 1
	baseInt := int(base)
	fmt.Println("base:", base)
	var total int
	total += 'D' - 'A' + 1
	total += baseInt * ('F' - 'A' + 1)
	total += baseInt * baseInt * ('X' - 'A' + 1)
	fmt.Println("DFX值=", total)

	// AFZ的值计算
	fmt.Println("AFZ的值计算:", 1*math.Pow(26, 2)+6*math.Pow(26, 1)+26*math.Pow(26, 0))

}
