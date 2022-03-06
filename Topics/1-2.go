package Topics

import "fmt"

// GRIDW 这个值表示移动范围的行宽度
const gGRIDW uint8 = 3

func ChessSolution() {
	type i struct {
		a uint8
		b uint8
	}
	area := gGRIDW * gGRIDW
	tmp := i{a: 4, b: 4}
	for tmp.a = 1; tmp.a <= area; tmp.a++ {
		for tmp.b = 1; tmp.b <= area; tmp.b++ {
			if tmp.a%gGRIDW != tmp.b%gGRIDW {
				fmt.Printf("A=%d, B=%d\n", tmp.a, tmp.b)
			}
		}
	}
}
