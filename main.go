package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type item struct {
	Words []string
	Bomba string
}

func main() {

	espiri := ""

	beyin := []item{
		{[]string{"ora", "lahmacun"}, "Lahmacunu ORA'da yemek lazım."},
		{[]string{"kırçiçeği", "çiçeği"}, "Gidelim de çiçeği kıralım."},
		{[]string{"dizimi", "dizi", "diz"}, "Eve gidip dizimi izleyeceğim."},
		{[]string{"kaptan", "burger"}, "Kaptan Burger yemem."},
		{[]string{"man", "adam"}, "Adam adama dalmış!"},
		{[]string{"kudret", "kudred"}, "Kudurmuş et."},
		{[]string{"erikson", "erriccson"}, "Bu erricson. Başka erik yok."},
		{[]string{"bitcoin", "coin"}, "Bit mi goin?"},
		{[]string{"git", "github", "biraz"}, "Git biraz github oku!"},
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rıza ^_^: Selam abi naber?: ")
	deger, _ := reader.ReadString('\n')

	words := strings.Split(deger, " ")

	for i := 0; i < len(beyin); i++ {
		list := beyin[i]

		for w := 0; w < len(words); w++ {

			word := words[w]

			for z := 0; z < len(list.Words); z++ {

				litem := list.Words[z]

				if strings.Contains(word, litem) {
					espiri = list.Bomba
					break
				}
			}

		}
	}

	if espiri == "" {
		espiri = "ee. şey. auahushsh güldüm."
	}

	fmt.Printf("Rıza ^_^: %s\n", espiri)

}
