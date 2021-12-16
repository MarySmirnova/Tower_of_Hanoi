package main

import (
	"fmt"
	"sort"
)

func main() {
	tow := NewTower([]int{1, 5, 3, 6, 2, 4, 8, 7})
	fmt.Println(tow)
	Hanoi(0, &tow, startArr)

	fmt.Println(tow)
}

type Tower struct {
	start []int
	end   []int
	tmp   []int
}

func NewTower(arr []int) Tower {
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	return Tower{arr, make([]int, 0, len(arr)), make([]int, 0, len(arr))}
}

type ArrayType int

const (
	startArr ArrayType = 0
	endArr   ArrayType = 1
	tmpArr   ArrayType = 2
)

func (t *Tower) DefineArr(arrtype ArrayType) *[]int {
	var arr *[]int
	switch arrtype {
	case startArr:
		arr = &t.start
	case endArr:
		arr = &t.end
	case tmpArr:
		arr = &t.tmp
	}
	return arr
}

func (t *Tower) RearrElement(arrtype ArrayType, index int) {
	arr := t.DefineArr(arrtype)

	for i := 0; i < 3; i++ {
		if i == int(arrtype) {
			continue
		}
		switch {
		case i == 0 && (len(t.start) == 0 || t.start[len(t.start)-1] > (*arr)[index]) && index%2 != 0:
			t.start = append(t.start, (*arr)[index])
			*arr = (*arr)[:len(*arr)-1]

		case i == 1 && (len(t.end) == 0 || t.end[len(t.end)-1] > (*arr)[index]) && index%2 == 0:
			t.end = append(t.end, (*arr)[index])
			*arr = (*arr)[:len(*arr)-1]

		case i == 2 && (len(t.tmp) == 0 || t.tmp[len(t.tmp)-1] > (*arr)[index]) && index%2 != 0:
			t.tmp = append(t.tmp, (*arr)[index])
			*arr = (*arr)[:len(*arr)-1]
		}
	}

}

func Hanoi(index int, t *Tower, arrtype ArrayType) {
	arr := t.DefineArr(arrtype)

	if index == len(*arr)-1 {
		fmt.Println("end ", t.start[index], index)
		t.RearrElement(arrtype, len(t.start)-1)
		return
	}
	fmt.Println("enter ", t.start[index])
	Hanoi(index+1, t, arrtype)
	fmt.Println("exit ", t.start[index])
}
