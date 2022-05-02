package main

import (
	"fmt"
	"si2/src/si2/algorithm"
)

func runBC(alg algorithm.Algorithm, varH, valH int) {
	fmt.Printf("\n(%d, %d)\n", varH, valH)
	for i := 0; i < 10; i++ {
		fmt.Println(alg.Backtracking(varH, valH))
	}
}

func runFC(alg algorithm.Algorithm, varH, valH int) {
	fmt.Printf("\n(%d, %d)\n", varH, valH)
	for i := 0; i < 10; i++ {
		fmt.Println(alg.Backtracking(varH, valH))
	}
}

func runFullBinary(name, path string) {
	var alg = algorithm.Algorithm{}
	alg.ImportBinary(path)
	fmt.Println(name)
	fmt.Println("BC")
	runBC(alg, 0, 0)
	runBC(alg, 0, 1)
	runBC(alg, 1, 1)
	runBC(alg, 1, 0)
	fmt.Println("\nFC")
	runFC(alg, 0, 0)
	runFC(alg, 0, 1)
	runFC(alg, 1, 1)
	runFC(alg, 1, 0)
}

func runFullFuto(name, path string) {
	var alg = algorithm.Algorithm{}
	alg.ImportFuto(path)
	fmt.Println()
	fmt.Println(name)
	fmt.Println("BC")
	runBC(alg, 0, 0)
	runBC(alg, 0, 1)
	runBC(alg, 1, 1)
	runBC(alg, 1, 0)
	fmt.Println("\nFC")
	runFC(alg, 0, 0)
	runFC(alg, 0, 1)
	runFC(alg, 1, 1)
	runFC(alg, 1, 0)
}

func main() {
	// bPath := "D:\\go\\data\\binary_6x6"
	runFullBinary("binary 6x6", "D:\\go\\data\\binary_6x6")
	runFullBinary("binary 8x8", "D:\\go\\data\\binary_8x8")
	runFullBinary("binary 10x10", "D:\\go\\data\\binary_10x10")

	fmt.Println()
	runFullFuto("futo 4x4", "D:\\go\\data\\futoshiki_4x4")
	runFullFuto("futo 5x5", "D:\\go\\data\\futoshiki_5x5")
	runFullFuto("futo 6x6", "D:\\go\\data\\futoshiki_6x6")
	// var binaryAlg = algorithm.Algorithm{}
	// binaryAlg.ImportBinary(bPath)
	// fmt.Println("binary 6x6")
	// fmt.Println("backtracking")
	// fmt.Println("(0, 0)")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(binaryAlg.Backtracking(0, 0))
	// }
	// fmt.Println("\n(0, 1)")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(binaryAlg.Backtracking(0, 1))
	// }
	// fmt.Println("\n(1, 1)")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(binaryAlg.Backtracking(1, 1))
	// }
	// fmt.Println("\n(1, 0)")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(binaryAlg.Backtracking(1, 0))
	// }
	// fmt.Println(binaryAlg.Backtracking(0, 1))
	// fmt.Println(binaryAlg.Backtracking(1, 1))
	// fmt.Println(binaryAlg.Backtracking(1, 0))
	// fmt.Println(binaryAlg.ForwardChecking(0, 0))
	// fmt.Println(binaryAlg.ForwardChecking(0, 1))
	// fmt.Println(binaryAlg.ForwardChecking(1, 1))
	// fmt.Println(binaryAlg.ForwardChecking(1, 0))

	// fPath := "D:\\go\\data\\futoshiki_5x5"
	// var futoAlg = algorithm.Algorithm{}
	// fmt.Println("\n\nfutoshiki")
	// futoAlg.ImportFuto(fPath)
	// fmt.Println(futoAlg.Backtracking(0, 0))
	// fmt.Println(futoAlg.Backtracking(0, 1))
	// fmt.Println(futoAlg.Backtracking(1, 1))
	// fmt.Println(futoAlg.Backtracking(1, 0))
	// fmt.Println(futoAlg.ForwardChecking(0, 0))
	// fmt.Println(futoAlg.ForwardChecking(0, 1))
	// fmt.Println(futoAlg.ForwardChecking(1, 1))
	// fmt.Println(futoAlg.ForwardChecking(1, 0))
}
