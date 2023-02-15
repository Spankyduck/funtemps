package funfacts

type Funfacts struct {
  Sun []string
  Luna []string
  Terra []string
}

var facts Funfacts

func GetFunFacts(about string) []string {
  if about == "sun"  {
    return []string{"Temperatur i Solens kjerne", "Temperatur på ytre lag av Solen"}
  }
  return facts.Luna
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
