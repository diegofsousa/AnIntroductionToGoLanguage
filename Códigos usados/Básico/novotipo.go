package main

import "fmt"

type TimesDoPiaui []string

func main() {

	times := make(TimesDoPiaui, 4)
	times[0] = "River"
	times[1] = "Parnahyba Sport Club"
	times[2] = "Sociedade Esportiva de Picos"
	times[3] = "Piauí Esporte Clube"

	times.TemPhb()

}

func (time TimesDoPiaui) TemPhb() {
	for _, value := range time {
		if value == "Parnahyba Sport Club" {
			fmt.Println("Tem Parnahyba Sport Club!")
		}
	}
}
