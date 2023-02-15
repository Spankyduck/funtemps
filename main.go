package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"Funtemps-oppg/funtemps/conv"
	
)

// Definerer flag-variablene i hoved-"scope"
var fahr string
var fahrFloat float64

var kelvin string
var kelvinFloat float64

var celsius string
var celsiusFloat float64

var scale string
var out string
var funfacts string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.StringVar(&fahr, "F", "0.0", "temperatur i grader fahrenheit")
	flag.StringVar(&celsius, "C", "0.0", "Temperatur i grader celsius")
	flag.StringVar(&kelvin, "K", "0.0", "Temperatur i Kelvin")

	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises
	flag.StringVar(&scale, "t", "", "Funfacts tempreaturskala i C, F, K.")

}

func main() {

	flag.Parse()
    
	celsiusFloat := toFloat(celsius)
	fahrFloat := toFloat(fahr)
	kelvinFloat := toFloat(kelvin)

	if isFlagPassed("funfacts") || isFlagPassed("t") {
	    if isFlagPassed("funfacts") {
	     fmt.Println("funfacts can only be used with t")
		 os.Exit(1)
		} else if isFlagPassed("t") {
	     fmt.Println("t and funfacts are true")
		 os.Exit(1)
		}
    } else if isFlagPassed("out") {
	     fmt.Printf("Can not convert from %s to %s.\n", out, out)
		 os.Exit(1)
    } else {
	switch {
	case out == "C" && isFlagPassed("F"): //F til C
	         fmt.Println("%vF er %vC", fahrFloat, formatOutput(conv.FarhenheitToCelsius(fahrFloat)))
	case out == "C" && isFlagPassed("F"): //F til C
	         fmt.Println("%vF er %vK", fahrFloat, formatOutput(conv.FarhenheitToKelvin(fahrFloat)))
	case out == "C" && isFlagPassed("F"): //F til C
	         fmt.Println("%vC er %vF", celsiusFloat, formatOutput(conv.CelsiusToFahrenheit(fahrFloat)))
	case out == "C" && isFlagPassed("F"): //F til C
	         fmt.Println("%vC er %vK", celsiusFloat, formatOutput(conv.CelsiusToKelvin(fahrFloat)))
	case out == "C" && isFlagPassed("F"): //F til C
	         fmt.Println("%vK er %vC", kelvinFloat, formatOutput(conv.KelivnToCelsius(fahrFloat)))
	case out == "C" && isFlagPassed("F"): //F til C
	         fmt.Println("%vK er %vF", kelvinFloat, formatOutput(conv.KelvinToFarhenheit(fahrFloat)))
	}
  } 
} 

	/*   Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	

	// Her er noen eksempler du kan bruke i den manuelle testingen
	fmt.Println(fahr, out, funfacts)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	Eksempel på enkel logikk
	if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		fmt.Println("0°F er -17.78°C")
	}

	

Funksjonen sjekker om flagget er spesifisert på kommandolinje
*/
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func toFloat(str string) float64 {
    value, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Printf("Error converting string to float: %v\n", err)
        os.Exit(1)
    }
    return value
}

func formatOutput(value float64) string {
    return strconv.FormatFloat(value, 'f', 2, 64)
}

