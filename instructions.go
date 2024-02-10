package pushswap

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var A []int
var B []int

var Count int // Décompte des opérations utilisées

var OutputMode bool // Permet d'écrire les instructions utilisées dans la sortie standard

func Pa() {
	first := B[0]
	B = append(B[:0], B[0+1:]...)
	res := make([]int, len(A)+1)
	copy(res[1:], A)
	res[0] = first
	A = res
	Count++
	if OutputMode {
		fmt.Println("pa")
		// PrintTabs()
	}
}

func Pb() {
	first := A[0]
	A = append(A[:0], A[0+1:]...)
	res := make([]int, len(B)+1)
	copy(res[1:], B)
	res[0] = first
	B = res
	Count++
	if OutputMode {
		fmt.Println("pb")
		// PrintTabs()
	}
}

func Sa() {
	A[0], A[1] = A[1], A[0]
	Count++
	if OutputMode {
		fmt.Println("sa")
		// PrintTabs()
	}
}

func Sb() {
	B[0], B[1] = B[1], B[0]
	Count++
	if OutputMode {
		fmt.Println("sb")
		// PrintTabs()
	}
}

func Ss() {
	A[0], A[1] = A[1], A[0]
	B[0], B[1] = B[1], B[0]
	Count++
	if OutputMode {
		fmt.Println("ss")
		// PrintTabs()
	}
}

func Ra() {
	first := A[0]
	A = append(A, first)
	A = A[1:]
	Count++
	if OutputMode {
		fmt.Println("ra")
		// PrintTabs()
	}
}

func Rb() {
	first := B[0]
	B = append(B, first)
	B = B[1:]
	Count++
	if OutputMode {
		fmt.Println("rb")
		// PrintTabs()
	}
}

func Rr() {
	firstb := B[0]
	B = append(B, firstb)
	B = B[1:]
	firsta := A[0]
	A = append(A, firsta)
	A = A[1:]
	Count++
	if OutputMode {
		fmt.Println("rr")
		// PrintTabs()
	}
}

func Rra() {
	last := A[len(A)-1]
	res := make([]int, len(A)+1)
	copy(res[1:], A)
	res[0] = last
	res = res[:len(res)-1]
	A = res
	Count++
	if OutputMode {
		fmt.Println("rra")
		// PrintTabs()
	}
}

func Rrb() {
	last := B[len(B)-1]
	res := make([]int, len(B)+1)
	copy(res[1:], B)
	res[0] = last
	res = res[:len(res)-1]
	B = res
	Count++
	if OutputMode {
		fmt.Println("rrb")
		// PrintTabs()
	}
}

func Rrr() {
	last := B[len(B)-1]
	res := make([]int, len(B)+1)
	copy(res[1:], B)
	res[0] = last
	res = res[:len(res)-1]
	A = res
	Count++
	if OutputMode {
		fmt.Println("rrr")
		// PrintTabs()
	}
}

func AddInA(x int) {
	A = append(A, x)
}

func Error() {
	fmt.Println("Error")
	os.Exit(0)
}

// lit les nombres en entrée
func ReadNumbers() []int {
	if len(os.Args) == 1 {
		os.Exit(0)
	}
	if len(os.Args) != 2 {
		Error()
	}

	var res []int

	numbers := strings.Split(os.Args[1], " ")
	for _, i := range numbers {
		x, err := strconv.Atoi(i)
		if err != nil {
			Error()
		}
		for _, n := range res {
			if n == x {
				Error()
			}
		}
		res = append(res, x)
	}

	return res
}

func PrintCount() {
	fmt.Printf("Count : %d\n", Count)
}

func PrintTabs() {
	fmt.Printf("%d || %d\n", A, B)
}
