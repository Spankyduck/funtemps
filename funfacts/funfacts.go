package funfacts

type Funfacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

var facts Funfacts

func GetFunFactsSun(about string) []string {
	if about == "Sun" {
		return []string{"Temperatur på ytre lag av solen 15 000 000", "Temperatur på ytre lag av solen"}
	}
	return facts.Sun
}

func GetFunFactsLuna(about string) []string {
	if about == "Luna" {
		return []string{"Temperaturen på Månens overflate om natten", "Temperatur på Månens overflate om dagen"}
	}
	return facts.Luna
}

func GetFunFactsTerra(about string) []string {
	if about == "Terra" {
		return []string{"Temperatur i Jordens indre kjerne", "Laveste temperatur målt på jordens overflate", "Høyeste temperatur målt på Jordens overflate"}
	}
	return facts.Terra
}

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra

  Sett inn alle Funfucts i en struktur
  type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
  }
*/
