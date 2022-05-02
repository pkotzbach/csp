package algorithm

import (
	"fmt"
	"sort"
	"time"

	"golang.org/x/exp/slices"
)

type Field struct {
	val         int
	domainAv    []int
	constant    bool
	next        *Field
	Constraints []func() bool
}

func (f Field) String() string {
	return fmt.Sprintf("%p", &f)
}

type inequality struct {
	lower, bigger int
}

type toFirst struct {
	time  time.Time
	nodes int
}

type Algorithm struct {
	fields           []*Field
	edgeSize         int
	Domain           []int
	FirstField       *Field
	resultsSize      int
	nodesVisited     int
	first            toFirst
	valHeuristicFunc func(int, *[]int) int
}

func (a *Algorithm) clear() {
	a.resultsSize = 0
	a.nodesVisited = 0
}

func (a *Algorithm) satisfyConstraints(field *Field, val int) bool {
	temp := field.val
	field.val = val
	defer func() { field.val = temp }()
	for _, constraint := range field.Constraints {
		satisfied := constraint()
		if !satisfied {
			return false
		}
	}
	return true
}

func (a *Algorithm) setNextFields() {
	for i := 0; i < len(a.fields)-1; i++ {
		a.fields[i].next = a.fields[i+1]
	}
}

func (a *Algorithm) ImportBinary(path string) {
	a.fields, a.edgeSize = importBinary(path)
	a.Domain = make([]int, 2)
	a.Domain[1] = 1
	a.setNextFields()
	BinaryAllConstraints(a)
	a.FirstField = a.fields[0]
}

func (a *Algorithm) ImportFuto(path string) {
	var ineq []inequality
	a.fields, ineq, a.edgeSize = importFuto(path)
	a.Domain = make([]int, a.edgeSize)
	for i := 0; i < a.edgeSize; i++ {
		a.Domain[i] = i + 1
	}
	a.setNextFields()
	FutoAllConstraints(a, ineq)
	a.FirstField = a.fields[0]
}

func (a *Algorithm) Print() {
	fmt.Println("--------------")
	field := a.FirstField
	for i := 0; i < a.edgeSize; i++ {
		for j := 0; j < a.edgeSize; j++ {
			fmt.Print(field.val, " ")
			field = field.next
		}
		fmt.Println()
	}
}

func (a *Algorithm) findEmptyLocation() int {

	for i := 0; i < len(a.fields); i++ {
		if a.fields[i].val == -1 {
			return i
		}
	}
	return -1
}

func (a *Algorithm) setVarHeuristics(id int) {
	// fmt.Print("var heuristic: ")
	switch id {
	case 0:
		// fmt.Println("no heuristic")
	case 1:
		// fmt.Println("most constraints first")
		sort.SliceStable(a.fields, func(i, j int) bool {
			return len(a.fields[i].Constraints) > len(a.fields[j].Constraints)
		})
	}
}

func (a *Algorithm) setValHeuristics(id int) {
	// fmt.Print("val heuristic: ")
	switch id {
	case 0:
		// fmt.Println("no heuristic")
		a.valHeuristicFunc = func(pos int, dom *[]int) int {
			for i, val := range *dom {
				if val != -1 {
					(*dom)[i] = -1
					return val
				}
			}
			panic("ERROR!")
		}
	case 1:
		// fmt.Println("least used value first")
		a.valHeuristicFunc = func(pos int, dom *[]int) int {
			ammount := make(map[int]int)
			for i := 0; i < pos; i++ {
				ammount[a.fields[i].val]++
			}

			min := -1
			ammount[min] = 10000

			for k, v := range ammount {
				if slices.Contains(*dom, k) && v < ammount[min] {
					min = k
				}
			}

			if min == -1 {
				for i, val := range *dom {
					if val != -1 {
						(*dom)[i] = -1
						return val
					}
				}
			}

			for i, val := range *dom {
				if val == min {
					(*dom)[i] = -1
					return val
				}
			}
			return -100
		}
	}
}

func (a *Algorithm) Backtracking(varHeuristic, valHeuristic int) (int, int, int64, int, int64) {
	a.clear()
	start := time.Now()
	// fmt.Println("\nbacktracking")
	a.setVarHeuristics(varHeuristic)
	a.setValHeuristics(valHeuristic)
	a.backtracking()
	return a.resultsSize, a.nodesVisited, time.Since(start).Microseconds(), a.first.nodes, a.first.time.Sub(start).Microseconds()
}

func (a *Algorithm) backtracking() {
	cords := a.findEmptyLocation()
	if cords == -1 {
		// a.Print()
		a.resultsSize++
		if a.resultsSize == 1 {
			a.first.nodes = a.nodesVisited
			a.first.time = time.Now()
		}
		return
	}
	a.nodesVisited++

	for _, num := range a.Domain {
		a.fields[cords].val = num
		if a.satisfyConstraints(a.fields[cords], num) {
			a.backtracking()
		}
		a.fields[cords].val = defaultVal
	}
}

func (a *Algorithm) checkForward(i int) bool {
	for j := i + 1; j < len(a.fields); j++ {
		field := a.fields[j]
		if !field.constant {
			wipeout := true
			for _, val := range a.Domain {
				if field.domainAv[val] == defaultVal {
					if a.satisfyConstraints(field, val) {
						wipeout = false
					} else {
						field.domainAv[val] = i
					}
				}
			}
			if wipeout {
				return false
			}
		}
	}
	return true
}

func (a *Algorithm) restore(i int) {
	for j := i + 1; j < len(a.fields); j++ {
		field := a.fields[j]
		if !field.constant {
			field.val = defaultVal
			for _, dom := range a.Domain {
				if field.domainAv[dom] == i {
					field.domainAv[dom] = defaultVal
				}
			}
		}
	}
}

func (a *Algorithm) ForwardChecking(varHeuristic, valHeuristic int) (int, int, int64, int, int64) {
	a.clear()
	start := time.Now()
	fmt.Println("\nforward checking")
	a.setVarHeuristics(varHeuristic)
	a.setValHeuristics(valHeuristic)
	a.forwardChecking(0)
	return a.resultsSize, a.nodesVisited, time.Since(start).Microseconds(), a.first.nodes, a.first.time.Sub(start).Microseconds()
}

func (a *Algorithm) forwardChecking(i int) {
	a.nodesVisited++
	if a.fields[i].constant {
		if i == len(a.fields)-1 {
			// a.Print()
			a.resultsSize++
			if a.resultsSize == 1 {
				a.first.nodes = a.nodesVisited
				a.first.time = time.Now()
			}
		} else {
			if a.checkForward(i) {
				a.forwardChecking(i + 1)
			}
			a.restore(i)
		}
	} else {
		usedDomains := make([]int, len(a.Domain))
		copy(usedDomains, a.Domain)
		for j := 0; j < len(a.Domain); j++ {
			val := a.valHeuristicFunc(i, &usedDomains)

			if a.fields[i].domainAv[val] == defaultVal {
				a.fields[i].val = val
				if i == len(a.fields)-1 {
					// a.Print()
					a.resultsSize++
					if a.resultsSize == 1 {
						a.first.nodes = a.nodesVisited
						a.first.time = time.Now()
					}
				} else {
					if a.checkForward(i) {
						a.forwardChecking(i + 1)
					}
					a.restore(i)
				}
			}
		}
	}
}
