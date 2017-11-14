package main

import (
	"fmt"
)

func main() {
	timesDoPiaui := map[string]string{
		"RIV": "River",
		"PAR": "Parnahyba Sport Club",
		"SEP": "Sociedade Esportiva de Picos",
	}

	timesDoPiaui["PEC"] = "Piaui Esporte Clube"

	fmt.Println(timesDoPiaui)
}
