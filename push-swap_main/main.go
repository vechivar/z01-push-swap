package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	ps "pushswap"
	"sort"
	"strconv"
	"time"
)

var N int             // Nombre de nombres à traiter
var firstTierMax int  // valeur maximale du premier tiers (de 0 à firstTierMax)
var secondTierMax int // valeur maximale du deuxième tiers (de firstTierMax + 1 à secondTierMax)
var firstTierLen int  // taille du premier tiers
var secondTierLen int // taille du second tiers
var thirdTierLen int  // taille du derniers tiers

func main() {
	ps.OutputMode = true
	TreatInput()

	if IsSorted(ps.A) {
		return
	}

	if N == 6 {
		step1for6()
		step2for6()
		return
	}
	if N == 5 {
		step1for5()
		step2for5()
		return
	}

	ps.Count = 0
	firstTierMax = N/3 - 1
	firstTierLen = N / 3
	secondTierMax = (2*N)/3 - 1
	secondTierLen = (2*N)/3 - firstTierLen
	thirdTierLen = N - firstTierLen - secondTierLen

	Step1()
	Step2()
	Step3()
	Step4()

	// ps.PrintTabs()
	// ps.PrintCount()
	// AnalyseStat()
	// fmt.Println(ps.Count)
}

// Lit les nombres en entrée et effectue la conversion en nombres successifs.
func TreatInput() {
	input := ps.ReadNumbers()
	sortedInput := make([]int, len(input))
	copy(sortedInput, input)
	sort.Ints(sortedInput)
	for _, x := range input {
		for i, y := range sortedInput {
			if y == x {
				ps.AddInA(i)
			}
		}
	}
	N = len(ps.A)
}

// utilisé avec le script statCOunt pour estimer l'efficacité de l'algorithme.
func AnalyseStat() {
	f, _ := os.Open("statRes.txt")

	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)
	count := 0
	min := 800
	max := 0
	i := 0

	for fileScanner.Scan() {
		i++
		n, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			fmt.Println(err)
		} else {
			if n > 700 {
				fmt.Println(n)
				count++
			}
			if n <= min {
				min = n
			}
			if n >= max {
				max = n
			}
		}
	}
	fmt.Println("au-dessus de 700 : ")
	fmt.Println(count)
	fmt.Printf("total : %d min : %d max %d\n", i, min, max)
}

// pour tester avec des nombres aléatoires dans A
func GenerateRandomA(n int) {
	rand.Seed(time.Now().UnixMicro())
	ps.A = rand.Perm(n)
}

// Etape 1 de l'algorithme.
func Step1() {
	rotBneeded := 0 // rotations à effectuer dans B.

	for len(ps.A) != thirdTierLen {
		if ps.A[0] <= firstTierMax {
			// la valeur appartient au premier tiers. On push et on augmente les rotations nécessaires dans B
			ps.Pb()
			rotBneeded++
		} else if ps.A[0] <= secondTierMax {
			// la valeur appartient au second tiers. On effectue les rotations nécessaires avant de push.
			for rotBneeded > 0 {
				ps.Rb()
				rotBneeded--
			}
			ps.Pb()
		} else {
			// la valeur appartient au troisième tiers, on ne push pas et on optimise les rotations avec rr
			if rotBneeded > 0 {
				ps.Rr()
				rotBneeded--
			} else {
				ps.Ra()
			}
		}
	}
	// On effectue les dernières rotations si besoin.
	for rotBneeded > 0 {
		ps.Rb()
		rotBneeded--
	}
}

// Etape 2 de l'algorithme
func Step2() {
	min := secondTierMax + 1
	max := N - 1
	rot := 0
	needRotB := 0
	nextNeedRotB := 0

	for len(ps.A) > 0 {
		// On compte les optimisations possibles si on va chercher un max
		extras := 0
		for LookForMaxOptimizationStep2(max) {
			max--
			extras++
		}
		max += extras

		rotMin := RotNeeded(min, &ps.A)
		rotMax := RotNeeded(max, &ps.A)

		// choix du max ou du min
		if math.Abs(float64(rotMin)) > math.Abs(float64(rotMax))/float64(extras+1) {
			rot = rotMax
			max -= extras + 1
			nextNeedRotB += extras + 1
		} else {
			rot = rotMin
			min++
		}
		for i := 0; i < IntAbs(rot); i++ {
			// on rencontre un max supplémentaire, il faut l'envoyer.
			if ps.A[0] > max {
				for needRotB > 0 {
					// il reste des rotate à faire dans B avant de push
					ps.Rb()
					needRotB--
				}
				ps.Pb()
				if rot < 0 {
					// quand on parcourt la liste avec des rotate, un push enlève un rotate nécessaire
					// quand on la parcourt en reverse-rotate, ce n'est pas le cas
					ps.Rra()
				}

			} else if rot > 0 {
				if needRotB > 0 {
					// occasion de faire un double rotate
					ps.Rr()
					needRotB--
				} else {
					ps.Ra()
				}
			} else {
				ps.Rra()
			}
		}
		for needRotB > 0 {
			// On effectue les derniers rotate avant de push
			needRotB--
			ps.Rb()
		}
		ps.Pb()
		// On actualise les rotations nécessaires pour le prochain cycle
		needRotB = nextNeedRotB
		nextNeedRotB = 0
	}

	// on replace la valeur maximale en haut de B
	rot = RotNeeded(N-1, &ps.B)

	for i := 0; i < IntAbs(rot); i++ {
		ps.Rrb()
	}

	// Etape 2.5 de l'algorithme
	for i := 0; i < thirdTierLen; i++ {
		ps.Pa()
	}
}

// Etape 3 de l'algorithme
func Step3() {
	// fonctionnement similaire à l'étape 2
	min := firstTierMax + 1
	max := secondTierMax
	rot := 0
	needRotA := 0
	nextNeedRotA := 0

	for len(ps.B) > firstTierLen {
		// on trie les nombres dans l'autre sens (minimum en haut)
		// les optimisations se font donc en cherchant des minimums additionnels
		extras := 0
		for LookForMinOptimizationStep34(min) {
			extras++
			min++
		}
		min -= extras

		rotMin := RotNeeded(min, &ps.B)
		rotMax := RotNeeded(max, &ps.B)

		if math.Abs(float64(rotMin))/float64(extras+1) > math.Abs(float64(rotMax)) {
			rot = rotMax
			max--
		} else {
			rot = rotMin
			min += extras + 1
			nextNeedRotA += extras + 1
		}
		for i := 0; i < IntAbs(rot); i++ {
			// attention à ne pas envoyer une valeur du premier tiers
			if ps.B[0] < min && ps.B[0] > firstTierMax {
				for needRotA > 0 {
					ps.Ra()
					needRotA--
				}
				ps.Pa()
				if rot < 0 {
					ps.Rrb()
				}
			} else if rot > 0 {
				if needRotA > 0 {
					ps.Rr()
					needRotA--
				} else {
					ps.Rb()
				}
			} else {
				ps.Rrb()
			}
		}
		for needRotA > 0 {
			needRotA--
			ps.Ra()
		}
		ps.Pa()
		needRotA = nextNeedRotA
		nextNeedRotA = 0
	}

	rot = RotNeeded(firstTierMax+1, &ps.A)

	for i := 0; i < IntAbs(rot); i++ {
		if rot > 0 {
			ps.Ra()
		} else {
			ps.Rra()
		}
	}
}

// Etape 4 de l'algorithme
func Step4() {
	// fonctionnement similaire à l'étape 3
	min := 0
	max := firstTierMax
	rot := 0
	needRotA := 0
	nextNeedRotA := 0

	for len(ps.B) > 1 {
		extras := 0
		for LookForMinOptimizationStep34(min) {
			extras++
			min++
		}
		min -= extras

		rotMin := RotNeeded(min, &ps.B)
		rotMax := RotNeeded(max, &ps.B)

		if math.Abs(float64(rotMin))/float64(extras+1) > math.Abs(float64(rotMax)) {
			rot = rotMax
			max--
		} else {
			rot = rotMin
			min += extras + 1
			nextNeedRotA += extras + 1
		}
		for i := 0; i < IntAbs(rot); i++ {
			if ps.B[0] < min {
				for needRotA > 0 {
					ps.Ra()
					needRotA--
				}
				ps.Pa()
				if rot < 0 {
					ps.Rrb()
				}
			} else if rot > 0 {
				if needRotA > 0 {
					ps.Rr()
					needRotA--
				} else {
					ps.Rb()
				}
			} else {
				ps.Rrb()
			}
		}
		for needRotA > 0 {
			needRotA--
			ps.Ra()
		}
		ps.Pa()
		needRotA = nextNeedRotA
		nextNeedRotA = 0
	}
	// on effectue les dernières rotations nécessaires avant de push la dernière valeur.
	for needRotA > 0 {
		needRotA--
		ps.Ra()
	}
	ps.Pa()

	// On replace 0 en haut de la pile
	rot = RotNeeded(0, &ps.A)

	for i := 0; i < IntAbs(rot); i++ {
		if rot > 0 {
			ps.Ra()
		} else {
			ps.Rra()
		}
	}
}

// Trouve la position de x dans tab
func FindPos(x int, tab *[]int) int {
	for i, y := range *tab {
		if y == x {
			return i
		}
	}
	fmt.Println("problem")
	ps.PrintTabs()
	return -1
}

// Valeur absolue sur un entier
func IntAbs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

// Rotations nécessaires pour atteindre val dans tab
// Positif pour des rotate, négatif pour des reverse-rotate
func RotNeeded(val int, tab *[]int) int {
	pos := FindPos(val, tab)

	if pos < len(*tab)/2 {
		return pos
	} else {
		return -1 * (len(*tab) - pos)
	}
}

// Recherche d'optimisations avec des maximums additionnels
func LookForMaxOptimizationStep2(max int) bool {
	rots := RotNeeded(max, &ps.A)

	if rots > 0 {
		for i := 0; i < FindPos(max, &ps.A); i++ {
			if ps.A[i] == max-1 {
				return true
			}
		}
	}
	if rots < 0 {
		for i := len(ps.A) - 1; i >= FindPos(max, &ps.A); i-- {
			if ps.A[i] == max-1 {
				return true
			}
		}
	}
	return false
}

// Recherche d'optimisations avec des minimums additionnels
func LookForMinOptimizationStep34(min int) bool {
	rots := RotNeeded(min, &ps.B)

	if rots > 0 {
		for i := 0; i < FindPos(min, &ps.B); i++ {
			if ps.B[i] == min+1 {
				return true
			}
		}
	}
	if rots < 0 {
		for i := len(ps.B) - 1; i >= FindPos(min, &ps.B); i-- {
			if ps.B[i] == min+1 {
				return true
			}
		}
	}
	return false
}

func step1for6() { //envoyer la moitié dans B
	m := 0
	minv := Min(ps.A)
	for s := 0; m < 3; s++ {
		if ps.A[0] == minv || ps.A[0] == minv+1 || ps.A[0] == minv+2 {
			ps.Pb()
			m++
		} else if (ps.A[1] == minv || ps.A[1] == minv+1 || ps.A[1] == minv+2) && (m == 2) {
			ps.Sa()
		} else {
			ps.Ra()
		}
	}
}
func step2for6() { // trier les deux piles et renvoyer B dans A
	optiA := 0
	optiB := 0
	//trier la pile A
	if ps.A[0] > ps.A[1] && ps.A[0] > ps.A[2] && ps.A[1] < ps.A[2] {
		ps.Ra()
	}
	if ps.A[0] > ps.A[1] && ps.A[1] > ps.A[2] {
		ps.Ra()
		optiA = 1
	}
	if ps.A[0] > ps.A[1] && ps.A[1] < ps.A[2] {
		optiA = 1
	}
	if ps.A[2] < ps.A[0] {
		ps.Rra()
	}
	if ps.A[2] < ps.A[1] {
		ps.Rra()
		optiA = 1
	}
	//trier la pile B
	if ps.B[2] > ps.B[1] && ps.B[2] > ps.B[0] && ps.B[1] < ps.B[0] {
		ps.Rrb()
	}
	if ps.B[0] < ps.B[2] && ps.B[2] < ps.B[1] {
		ps.Rb()
	}
	if ps.B[0] > ps.B[1] && ps.B[1] < ps.B[2] {
		ps.Rrb()
		optiB = 1
	}
	if ps.B[2] > ps.B[1] && ps.B[1] > ps.B[0] {
		ps.Rb()
		optiB = 1
	}
	if ps.B[0] > ps.B[2] && ps.B[0] < ps.B[1] {
		optiB = 1
	}
	//optimiser swap
	if optiA == 0 && optiB == 1 {
		ps.Sb()
	} else if optiA == 1 && optiB == 0 {
		ps.Sa()
	} else if optiA == 1 && optiB == 1 {
		ps.Ss()
	}
	//envoyer B sur A
	ps.Pa()
	ps.Pa()
	ps.Pa()
}
func step1for5() { // envoyer les 2  chiffres  les plus faible de la liste A sur la liste B
	m := 0
	minv := Min(ps.A)
	for s := 0; m < 2; s++ {
		if ps.A[0] == minv || ps.A[0] == minv+1 {
			ps.Pb()
			m++
		} else if (ps.A[1] == minv || ps.A[1] == minv+1) && (m == 1) {
			ps.Sa()
		} else {
			ps.Ra()
		}
	}
}
func step2for5() { //trier les deux listes et ensuite envoyer B sur A
	w := ps.Sa
	if ps.B[0] < ps.B[1] {
		w = ps.Ss
	}
	if ps.A[0] > ps.A[1] && ps.A[0] > ps.A[2] && ps.A[1] < ps.A[2] {
		ps.Ra()
		if ps.B[0] < ps.B[1] {
			ps.Sb()
		}
	}
	if ps.A[0] > ps.A[1] && ps.A[1] > ps.A[2] {
		ps.Ra()
		w()
	}
	if ps.A[0] > ps.A[1] {
		w()
	}
	if ps.A[2] < ps.A[0] {
		ps.Rra()
		if ps.B[0] < ps.B[1] {
			ps.Sb()
		}
	}
	if ps.A[2] < ps.A[1] {
		ps.Rra()
		w()
	}
	if ps.B[0] < ps.B[1] {
		ps.Sb()
	}
	//envoyer B sur A
	ps.Pa()
	ps.Pa()
}

func Min(values []int) int {
	min := values[0] //assign the first element equal to min
	for _, number := range values {
		if number < min {
			min = number
		}
	}
	return min
}

func IsSorted(tab []int) bool {
	for i := 0; i < len(tab)-1; i++ {
		if tab[i] >= tab[i+1] {
			return false
		}
	}
	return true
}
