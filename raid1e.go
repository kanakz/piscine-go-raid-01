package piscine

import "github.com/01-edu/z01"

func Raid1e(x, y int) {
	if x < 0 || y < 0 {

	} else if x >= 1 || y >= 1 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if (i == 1 && j == 1) || (j == x && i == y) {
					z01.PrintRune('A')
				} else if (i == 1 && j == x) || (i == y && j == 1) {
					z01.PrintRune('C')
				} else if j == 1 || j == x || i == 1 || i == y {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}

// func main() {
// 	z01.PrintRune('\n')
// 	Raid1e(2, 1)
// 	z01.PrintRune('\n')
// 	Raid1e(1, 2)
// 	z01.PrintRune('\n')
// 	Raid1e(2, 2)
// 	z01.PrintRune('\n')
// 	Raid1e(3, 3)
// 	z01.PrintRune('\n')
// 	Raid1e(-1, 3)
// 	z01.PrintRune('\n')
// 	Raid1e(1, 4)
// 	z01.PrintRune('\n')
// 	Raid1e(4, 1)
// 	z01.PrintRune('\n')
// 	Raid1e(5, 3)
// 	z01.PrintRune('\n')
// 	Raid1e(2, 19)
// 	z01.PrintRune('\n')
// 	Raid1e(20, 2)
// 	z01.PrintRune('\n')
// }
