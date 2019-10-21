package student

import "github.com/01-edu/z01"

func Raid1b(x, y int) {
	if x < 0 || y < 0 {

	} else if x >= 1 || y >= 1 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if x == 1 || y == 1 {
					if i == 1 && j == 1 {
						z01.PrintRune('/')
					} else if (i == y && j == 1) || (j == x && i == 1) {
						z01.PrintRune('\\')
					} else if (i == y && y == 1) || (j == x && x == 1) {
						z01.PrintRune('*')
					} else {
						z01.PrintRune(' ')
					}
				} else if (i == 1 && j == 1) || (i == y && j == x) {
					z01.PrintRune('/')
				} else if (i == y && j == 1) || (j == x && i == 1) {
					z01.PrintRune('\\')
				} else if j == 1 || j == x || i == 1 || i == y {
					z01.PrintRune('*')
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
// 	Raid1b(2, 1)
// 	z01.PrintRune('\n')
// 	Raid1b(1, 2)
// 	z01.PrintRune('\n')
// 	Raid1b(2, 2)
// 	z01.PrintRune('\n')
// 	Raid1b(3, 3)
// 	z01.PrintRune('\n')
// 	Raid1b(-1, 3)
// 	z01.PrintRune('\n')
// 	Raid1b(1, 4)
// 	z01.PrintRune('\n')
// 	Raid1b(4, 1)
// 	z01.PrintRune('\n')
// 	Raid1b(5, 3)
// 	z01.PrintRune('\n')
// 	Raid1b(2, 19)
// 	z01.PrintRune('\n')
// 	Raid1b(20, 2)
// 	z01.PrintRune('\n')
// 	Raid1b(-1, 6)
// 	z01.PrintRune('\n')
// }
