package main

import (
	translate "translateStrings/translateStrings"
	factory "translateStrings/translateStrings/factory"
	"fmt"
)

func main() {
	factory := NewFactoryEncoder()

	translate := translate.NewTranslateString(factory)

	textoBinario := translate.Translate("TRADUCIR","TEXT","BINARY")

	fmt.Printf(textoBinario)
}
