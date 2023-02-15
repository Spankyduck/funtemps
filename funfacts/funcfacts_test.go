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
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string
		want  []string
	}

	// Her må du legge inn korrekte testverdier
	//tests := []test{
	//  {input: , want: },
	//}
	tests := []test{
		{input: "sun", want: []string{"Temperatur i Solens kjerne", "Temperatur på ytre lag av solen"}},
		{input: "luna", want: []string{"Temperatur i Solens kjerne", "Temperatur på ytre lag av solen"}},
		{input: "terra", want: []string{"Temperatur i Solens kjerne", "Temperatur på ytre lag av solen"}},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
