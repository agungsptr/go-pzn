package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  uint8
}

type UserSlice []User

func (v UserSlice) Len() int {
	return len(v)
}

func (v UserSlice) Less(i, j int) bool {
	return v[i].Name < v[j].Name
}

func (v UserSlice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func main() {
	users := []User{
		{
			Name: "Agung",
			Age:  23,
		},
		{
			Name: "Amni",
			Age:  24,
		},
		{
			Name: "Abbu",
			Age:  12,
		},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)

	// tmp := []string{"banana", "apple", "orange", "grape"}
	// sort.Strings(tmp)
	// fmt.Println(tmp)
}
