package main

import (	
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	NAME    string `json:"nombre"`
	SURNAME string `json:"apellidos"`
	AGE     int64  `json:"Edad"`
}

func main() {

	p := Person{"Person", "Surname", 20}
	argsWithProg := os.Args
	fmt.Println("size ", len(argsWithProg))
    fmt.Println("args ", argsWithProg)
	if len(argsWithProg) == 4 {
		p.NAME = argsWithProg[1]
		p.SURNAME = argsWithProg[2]
		intVar, err := strconv.ParseInt(argsWithProg[3], 0, 8)
		fmt.Println("value age error: ", err)
		p.AGE = intVar
		fmt.Println("With data")
	} else {
		fmt.Println("default")
	}
	fmt.Printf("%+v\n", p)
	r, _ := json.Marshal(p)
	fmt.Println(string(r))
}
