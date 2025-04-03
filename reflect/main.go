package main

import (
	"fmt"
	"reflect"
)

type People struct {
	id   string
	name string
}

func CheckStruct(i any) {
	t := reflect.TypeOf(i) // ✅ Gunakan i, bukan z

	if t.Kind() == reflect.Struct { // ✅ Gunakan reflect.Struct langsung
		fmt.Println("Struct Name:", t.Name())
	} else {
		fmt.Println("Kind:", t.Kind()) // ✅ Gunakan t.Kind() bukan t.kind
	}
}

func main() {
	x := 30
	var checkReflect = reflect.ValueOf(x)
	fmt.Println(checkReflect)
	fmt.Println("Tipe dari x:", checkReflect.Type())

	z := People{
		id:   "13",
		name: "Jhon",
	}

	checkStructValue := reflect.ValueOf(z) // ✅ Hindari nama variabel yang sama dengan fungsi
	fmt.Println("Struct Data:", z)
	fmt.Println("Struct Type:", checkStructValue.Type())

	// Memanggil fungsi dengan benar
	CheckStruct(z) // ✅ Perbaikan pemanggilan fungsi
}
