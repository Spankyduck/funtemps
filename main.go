package main

import (
	"flag"
	"fmt"
	"strconv"

	"modul/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahrenheit float64
var celsius float64
var kelvin float64
var out string
var Funfacts string
var Svar float64
var funfactsunit string

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
	flag.Float64Var(&fahrenheit, "F", 0.0, "temperatur i grader fahrenheit")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&Funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&funfactsunit, "t", "C", "bruker setter input verdi, hvis ikke setter programmet til defualt verdi Celsius")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func GetFunFactsSun(about string) []float64 {
	if about == "Sun" {
		return []float64{5506.85, 15000000}
	}
	return nil
}

func Separate(num string) string {
	n := len(num)
	if n <= 3 {
		return num
	}
	return Separate(num[:n-3]) + " " + num[n-3:]
}

func StringToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func main() {

	flag.Parse()

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
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
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk

	Erlik := "="
	F := "°F"
	C := "°C"
	K := "°K"

	// FahrenheitToCelsius
	if out == "C" && isFlagPassed("F") {
		Svar = conv.FarhenheitToCelsius(fahrenheit)

		fmt.Printf("%.9g%s %s ", fahrenheit, F, Erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%s %s\n", Separate(strconv.Itoa(int(Svar))), C)
		} else {
			fmt.Printf("%s %s\n", Separate(strconv.FormatFloat(Svar, 'f', 2, 64)), C)
		}
	}

	// CelsiusToFahrenheit
	if out == "F" && isFlagPassed("C") {
		Svar = conv.CelsiusToFahrenheit(celsius)

		fmt.Printf("%.9g %s %s ", celsius, C, Erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%d %s\n", int(Svar), F)
		} else {
			fmt.Printf("%.2f %s\n", Svar, F)
		}
	}

	// CelsiusToKelvin
	if out == "K" && isFlagPassed("C") {
		Svar = conv.CelsiusToKelvin(celsius)

		fmt.Printf("%.9g %s %s ", celsius, C, Erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%d %s\n", int(Svar), K)
		} else {
			fmt.Printf("%.2f %s\n", Svar, K)
		}
	}

	// KelvinToCelsius
	if out == "C" && isFlagPassed("K") {
		Svar = conv.KelvinToCelsius(kelvin)

		fmt.Printf("%.9g %s %s ", kelvin, K, Erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%d %s\n", int(Svar), C)
		} else {
			fmt.Printf("%.2f %s\n", Svar, C)
		}
	}

	// FahrenheitToKelvin
	if out == "K" && isFlagPassed("F") {
		Svar = conv.FahrenheitToKelvin(fahrenheit)

		fmt.Printf("%.9g %s %s ", fahrenheit, F, Erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%d %s\n", int(Svar), K)
		} else {
			fmt.Printf("%.2f %s\n", Svar, K)
		}
	}

	// KelvinToFahrenheit
	if out == "F" && isFlagPassed("K") {
		Svar = conv.FarhenheitToCelsius(kelvin)

		fmt.Printf("%.9g %s %s ", kelvin, K, Erlik)
		if Svar == float64(int(Svar)) {
			fmt.Printf("%d %s\n", int(Svar), F)
		} else {
			fmt.Printf("%.2f %s\n", Svar, F)
		}
	}

}
