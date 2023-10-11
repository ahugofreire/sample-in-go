package apislices

import (
	"cmp"
	"fmt"
	"slices"
)

func GetDataInSlice() {
	contact := []string{"Amanda", "Julia", "Monica", "Marcia", "Beatriz"}
	position, isFound := slices.BinarySearch(contact, "Ana")
	fmt.Println(isFound, position)
}

func BinarySearch() {
	type Contact struct {
		Name  string
		Phone string
	}

	contacts := []Contact{
		{
			"Ana",
			"113334-1122",
		},
		{
			"Bruno",
			"453366-2500",
		},
		{
			"Monica",
			"115555-3333",
		},
	}

	// Para utilizar a busca binaria o slice deve estar ordenado ou via da erro
	bruno := Contact{"Bruno", ""}
	position, isFound := slices.BinarySearchFunc(contacts, bruno, func(c1, c2, Contact) int {
		return cmp.Compare(c1.Name, c2.Name)
	})

	fmt.Println(position, isFound)
}
