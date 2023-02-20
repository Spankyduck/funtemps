package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
// Her må du legge inn korrekte testverdier
//tests := []test{
//  {input: , want: },
//}
func TestGetFunFactsSun(t *testing.T) {
	type test struct {
		input string
		want  []string
	}

	tests := []test{
		{input: "Sun", want: []string{"Temperatur i Solens kjerne", "Temperatur på ytre lag av solen 15 000 000"}},
	}
	for _, tc := range tests {
		got := GetFunFactsSun(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestGetFunFactsLuna(t *testing.T) {
	type test struct {
		input string
		want  []string
	}

	tests := []test{
		{input: "Luna", want: []string{"Temperatur i Solens kjerne", "Temperatur på ytre lag av solen"}},
	}
	for _, tc := range tests {
		got := GetFunFactsLuna(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestGetFunFactsTerra(t *testing.T) {
	type test struct {
		input string
		want  []string
	}

	tests := []test{
		{input: "Terra", want: []string{"Temperatur i Solens kjerne", "Temperatur på ytre lag av solen"}},
	}
	for _, tc := range tests {
		got := GetFunFactsTerra(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
