package main

import "fmt"

func main() {
	mapaDoPiaui := map[string]string{}
	mapaDoPiaui["Parnaíba"] = "Curso de Go"
	mapaDoPiaui["Picos"] = "FixCode"
	mapaDoPiaui["Oeiras"] = "WebDesign"

	fmt.Println(mapaDoPiaui)

	for key, value := range mapaDoPiaui {
		fmt.Println("Em", key, "tem", value)
	}

}
