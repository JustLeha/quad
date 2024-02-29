package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for r := 0; r < y; r++ {
		for c := 0; c < x; c++ {
			if r == 0 && c == 0 {
				z01.PrintRune('A')
			} else if r == y-1 && c == x-1 {
				if y == 1 {
					z01.PrintRune('C')
				} else if x == 1 {
					z01.PrintRune(('C'))
				} else {
					z01.PrintRune('A')
				}
			} else if r == 0 && c == x-1 {
				z01.PrintRune('C')
			} else if r == y-1 && c == 0 {
				z01.PrintRune('C')
			} else if c == 0 || c == x-1 {
				z01.PrintRune('B')
			} else if !(r == 0 || r == y-1) {
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('B')
			}
		}
		z01.PrintRune('\n')
	}
}
