package main

import (
	"fmt"
	"sort"
)

func main() {
	tow, n := NewTower([]int{1, 5, 3, 6, 2, 4, 8, 7, 9, 10, 12, 11, 13, 1, 5})
	fmt.Println(tow)

	tow.Hanoi(&tow.start, &tow.end, &tow.tmp, n)
	fmt.Println(tow)
}

type Tower struct {
	start []int
	end   []int
	tmp   []int
}

//NewTower создает элемент структуры, массив сортируется по убыванию
//и находится на позиции start
func NewTower(arr []int) (Tower, int) {
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	return Tower{arr, make([]int, 0, len(arr)), make([]int, 0, len(arr))}, len(arr)
}

//Hanoi перекладывает элементы в позицию end в изначальном порядке
//по правилам "Ханойской башни"
func (t *Tower) Hanoi(src, dest, tmp *[]int, n int) {
	if n == 0 {
		transposit(src, dest)
		return
	}

	t.Hanoi(src, tmp, dest, n-1)
	if len(*src) != 0 {
		transposit(src, dest)
	}
	t.Hanoi(tmp, dest, src, n-1)
}

//transposit переставляет элементы межу двумя массивами
func transposit(source, dest *[]int) {
	*dest = append((*dest), (*source)[len(*source)-1])
	*source = (*source)[:len(*source)-1]
}
