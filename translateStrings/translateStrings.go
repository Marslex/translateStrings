package translateStrings

import (
	"translateStrings/translateStrings/factory"
	"fmt"
)

type translateString struct {
	factory *factory.Factory
}

func NewTranslateString(factory *factory.Factory) *translateString {
	return &translateString{
		factory: factory,
	}
}

//función de entrada al componente de traducción
//parametros:
//	inText: texto a traducir
//	fmtOrigin: formato origen
//	fmtDestiny: formato destino
func (translate *translateString) Translate(inText string,fmtOrigin string, fmtDestiny string) (string,error) {
	var encoder *map[string]string
	var outText string
	var err error

	switch fmtOrigin {
	case TEXT: 
		if(fmtDestiny == BINARY){
			encoder = factory.GetEncoder("TEXT_BINARY")[]
		}
	default: 
		err = errors.New("ocurrio un error, formato invalido")	
	}
	return outText,err
}

