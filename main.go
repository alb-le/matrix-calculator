package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Função para ler a matriz do usuário
func readMatrix() [][]int {
	var rows, cols int

	// Lendo o número de linhas e colunas da matriz
	fmt.Print("Informe o número de linhas da matriz: ")
	fmt.Scanln(&rows)
	fmt.Print("Informe o número de colunas da matriz: ")
	fmt.Scanln(&cols)

	// Criando uma matriz vazia
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}

	// Lendo os elementos da matriz
	fmt.Println("Informe os elementos da matriz (separados por espaço ou nova linha):")
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < rows; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			elements := strings.Fields(line)
			for j := 0; j < cols; j++ {
				if j < len(elements) {
					element, _ := strconv.Atoi(elements[j])
					matrix[i][j] = element
				}
			}
		}
	}

	return matrix
}

// Função para exibir a matriz
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, element := range row {
			fmt.Printf("%d ", element)
		}
		fmt.Println()
	}
}
func main() {
	// Lendo a matriz como entrada do usuário
	matrix := readMatrix()

	// Exibindo a matriz na saída
	fmt.Println("Matriz informada:")
	printMatrix(matrix)
}
