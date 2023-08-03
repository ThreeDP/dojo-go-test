package dojo

import (
	// "fmt"
	"testing"

)

func TestVoxtron(test *testing.T){
	
	// test.Run("testamos passar 'ola tchau tudo' e esperamos o retorno de uma array de string", func(test *testing.T) {
	// 	frase := "tchau ola tudo"
	// 	result := voxtron(frase)

	// 	esperado := []string{"ola", "tchau", "tudo"}
	// 	for i, palavra := range result{
	// 		if palavra != esperado[i]{
	// 			test.Fatalf("Deu ruim, valor esperado : (%s) -> valor retornado: (%s)", esperado[i], palavra)
	// 		}
	// 		// fmt.Println(palavra)
	// 	}
	// })

	// test.Run("testando ordernar o serie de strings", func(test *testing.T) {
	// 	frase := "xyz abc aab"
	// 	result := voxtron(frase)

	// 	esperado := []string{"aab", "abc", "xyz"}
	// 	for i, palavra := range result{
	// 		if palavra != esperado[i]{
	// 			test.Fatalf("Deu ruim, valor esperado : (%s) -> valor retornado: (%s)", esperado[i], palavra)
	// 		}
	// 		// fmt.Println(palavra)
	// 	}
	// })

	test.Run("testando ordernar o serie de strings", func(test *testing.T) {
		frase := "xyz abc aab yyt aab hhh abc abc hhh hhh hhh"
		result := voxtron(frase)

		esperado := []string{"hhh", "abc", "aab"}
		for i, palavra := range result{
			if palavra != esperado[i]{
				test.Fatalf("Deu ruim, valor esperado : (%s) -> valor retornado: (%s)", esperado[i], palavra)
			}
			// fmt.Println(palavra)
		}
	})

}