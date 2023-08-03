package dojo

import(
	"strings"
	// "sort"
	// "fmt"
)

func voxtron(frase string) []string{
	fraseSplitada := strings.Split(frase, " ")
	m := make(map[string]int)

	for _, palavra := range fraseSplitada{
		m[palavra] += 1
		// for i := 0; i < len(fraseSplitada); i++{
			// if palavra == fraseSplitada[i]{
			// }
		// }
	}
	// for key, value := range m{
	// 	fmt.Println("key:", key, "value: ", value)
	// }

	keys := make(map[int]string)

	for key, value := range m{
		keys[value] = key
	}
	var array []string
	for i := len(keys); i > len(keys) - 3; i-- {
		array = append(array, keys[i])
	}
	//sort.Ints(keys)
	// for key, value := range keys{
	// 	fmt.Println("key:", key, "value: ", value)
	// }

	// v1 := m["aab"]

	// fmt.Println("valor v1: %i", v1)
	// sort.Strings(fraseSplitada)

	return array
}
// [aab,
// for 
// 	m[word] += 1