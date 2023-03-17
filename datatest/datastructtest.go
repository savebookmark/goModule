// package main
package datatest

import (
	"fmt"
	"log"

	"github.com/Workiva/go-datastructures/set"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/francoispqt/gojay"
	boom "github.com/tylertreat/BoomFilters"
)

type user struct {
	id    int
	name  string
	email string
}

// implement gojay.UnmarshalerJSONObject
func (u *user) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "id":
		return dec.Int(&u.id)
	case "name":
		return dec.String(&u.name)
	case "email":
		return dec.String(&u.email)
	}
	return nil
}
func (u *user) NKeys() int {
	return 3
}

func main2() {
	// gods
	list := arraylist.New()
	list.Add("list", "list2", "list3", "list4")
	// workiva
	set := set.New()
	set.Add(1, 2, 3)

	fmt.Println(list)
	fmt.Println(set)
	fmt.Println("/////////////////////////////////////////")

	sbf := boom.NewDefaultStableBloomFilter(10000, 0.01)
	fmt.Println("stable point", sbf.StablePoint())

	sbf.Add([]byte(`a`))
	if sbf.Test([]byte(`a`)) {
		fmt.Println("contains a")
	}

	if !sbf.TestAndAdd([]byte(`b`)) {
		fmt.Println("doesn't contain b")
	}

	if sbf.Test([]byte(`b`)) {
		fmt.Println("now it contains b!")
	}
	// Restore to initial state.
	sbf.Reset()
	fmt.Println("//////////////////////////////////////////////")

	u := &user{}
	d := []byte(`{"id":1,"name":"gojay","email":"gojay@email.com"}`)
	err := gojay.UnmarshalJSONObject(d, u)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}
