package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"pushswap"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) > 3 {
		pushswap.Error()
	}

	if len(os.Args) == 3 {
		// Permet de générer des nombres aléatoires dans un fichier
		if os.Args[1] == "--generaterandom" {
			n, err := strconv.Atoi(os.Args[2])
			if err != nil {
				pushswap.Error()
			} else {
				rand.Seed(time.Now().UnixMicro())
				x := rand.Perm(n)
				f, _ := os.Create("random-" + strconv.Itoa(n))
				for i := 0; i < n-1; i++ {
					f.WriteString(strconv.Itoa(x[i]) + " ")
				}
				f.WriteString(strconv.Itoa(x[n-1]))
				f.Close()
				os.Exit(0)
			}
		} else {
			pushswap.Error()
		}
	}

	pushswap.A = pushswap.ReadNumbers()
	ExecuteInstructions()
	CheckResult()
}

// Affichage de l'échec.
func Fail() {
	fmt.Println("KO")
	os.Exit(0)
}

// Lit les instructions sur l'entrée standard et effectue les opérations correspondantes.
func ExecuteInstructions() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		x := scanner.Text()
		switch x {
		case "pa":
			pushswap.Pa()
		case "pb":
			pushswap.Pb()
		case "sa":
			pushswap.Sa()
		case "sb":
			pushswap.Sb()
		case "ss":
			pushswap.Ss()
		case "ra":
			pushswap.Ra()
		case "rb":
			pushswap.Rb()
		case "rr":
			pushswap.Rr()
		case "rra":
			pushswap.Rra()
		case "rrb":
			pushswap.Rrb()
		case "rrr":
			pushswap.Rrr()
		case "":
		default:
			pushswap.Error()
		}
	}
}

// Vérifie que le résultat est conforme
func CheckResult() {
	if len(pushswap.B) != 0 {
		Fail()
	}

	for i := 0; i < len(pushswap.A)-1; i++ {
		if pushswap.A[i] >= pushswap.A[i+1] {
			Fail()
		}
	}

	fmt.Println("OK")
}
