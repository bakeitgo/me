package main

func MultiplicationTable(size int) (matrice [][]int) {
	for i := 1; i <= size; i++ {
		var part []int
		for j := i; j <= size*i; j += i {
			part = append(part, j)
		}
		matrice = append(matrice, part)
	}
	return
}

func main() {

}
