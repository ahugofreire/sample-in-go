package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	GetDataInSlice()
	CloneSlice()
	CompactSlice()
	CompareSlices()
	ContainsSlice()
	DeleteSlice()
	EqualSlice()
	IndexSlice()
	insertSlice()
	// BinarySearch()
}

func GetDataInSlice() {
	contact := []string{"Amanda", "Julia", "Monica", "Marcia", "Beatriz"}
	position, isFound := slices.BinarySearch(contact, "Ana")
	fmt.Println(isFound, position)
}

// func BinarySearch() {
// 	type Contact struct {
// 		Name  string
// 		Phone string
// 	}

// 	contacts := []Contact{
// 		{
// 			"Ana",
// 			"113334-1122",
// 		},
// 		{
// 			"Bruno",
// 			"453366-2500",
// 		},
// 		{
// 			"Monica",
// 			"115555-3333",
// 		},
// 	}

// 	// Para utilizar a busca binaria o slice deve estar ordenado ou via da erro
// 	bruno := Contact{"Bruno", ""}
// 	position, isFound := slices.BinarySearchFunc(contacts, bruno, func(c1, c2, Contact) int {
// 		return cmp.Compare(c1.Name, c2.Name)
// 	})

// 	fmt.Println(position, isFound)
// }

func CloneSlice() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("s", &s[0])

	//usando o clone os slices se trona totalmente indepedentes
	b := slices.Clone(s) // != b := s
	fmt.Println("b", b)
}

func CompactSlice() {
	//Remove valores iguais em sequencia
	a := []int{1, 1, 1, 2, 2, 2, 3, 3, 4, 4, 4, 5, 1, 3}
	fmt.Println(a)
	b := slices.Compact(a)
	fmt.Println(b)

	//Compactando  em squecia
	c := []string{"alice", "Alice", "Bruna", "bruna", "Lena"}
	fmt.Println(c)
	d := slices.CompactFunc(c, func(s1, s2 string) bool {
		return strings.ToLower(s1) == strings.ToLower(s2)
	})
	fmt.Println(d)
}

func CompareSlices() {
	a := []int{1, 2, 3}
	// Se os slices for iguais retorna 0
	fmt.Println(slices.Compare(a, []int{1, 2, 3}))
	// Se o for maior que o comparado retorna 1
	fmt.Println(slices.Compare(a, []int{1, 2, 2}))
	// Se o for menor que o comparado retorna -1
	fmt.Println(slices.Compare(a, []int{1, 2, 4}))

}

func ContainsSlice() {
	a := []int{1, 2, 3, 4, 5}
	// Retorna true se o valor existe no slice
	fmt.Println(slices.Contains(a, 3))
	// Retorna false se não existe o valor no slice
	fmt.Println(slices.Contains(a, 9))
}

func DeleteSlice() {
	i := []int{1, 2, 3, 4, 5}
	fmt.Println(i)
	//passa o range de indices
	a := slices.Delete(i, 2, 5)
	fmt.Println(a)
}

func EqualSlice() {
	a := []int{1, 2, 3}
	fmt.Println(slices.Equal(a, []int{1, 2, 3})) //true
	fmt.Println(slices.Equal(a, []int{1, 5, 9})) //false
}

func IndexSlice() {
	v := []int64{1, 2, 3, 4, 5}
	fmt.Println(slices.Index(v, 5)) //retorna o index do valor
	fmt.Println(slices.Index(v, 9)) //retorna -1 se o valor não existir no slice

}

func insertSlice() {
	a := []int{1, 2, 4}
	//recebe o slice, o index que deseja adicioar e o valor
	a = slices.Insert(a, 2, 3)
	fmt.Println(a)
}
