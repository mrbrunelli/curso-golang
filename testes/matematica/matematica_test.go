package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

// gerar arquivo .out
// go test -coverprofile=resultado.out

// analisar resultado.out
// go tools cover -func=resultado.out

// gerar pagina html com coverage
// go tool cover -html=resultado.out
