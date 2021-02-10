package main

import (
	translate "translateStrings/translateStrings"
	factory "translateStrings/translateStrings/factory"
	"fmt"
)

func main() {
	factory := factory.NewFactoryEncoder()

	translate := translate.NewTranslateString(factory)

	textoBinario,_ := translate.Translate("A","TEXT","BINARY")

	fmt.Printf(textoBinario)
}
