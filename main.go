package main

import (
	"fmt"
)

func main() {
	tow := NewTower(22)
	fmt.Println(tow)

	GoHanoiTower(&tow)
	fmt.Println(tow)
}

type Tower struct {
	start []int
	end   []int
	tmp   []int
}

//NewTower принимает количество элементов башни, создает структуру.
//Пирамида на позиции start
func NewTower(n int) Tower {
	arr := make([]int, 0, n)
	for i := n; i > 0; i-- {
		arr = append(arr, i)
	}
	return Tower{arr, make([]int, 0, n), make([]int, 0, n)}
}

//GoHanoiTower запускает метод Hanoi, возвращает измененную структуру.
func GoHanoiTower(t *Tower) Tower {
	t.Hanoi(&t.start, &t.end, &t.tmp, len(t.start))
	return *t
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
