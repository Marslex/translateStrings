package translateStrings

import (
	"translateStrings/translateStrings/factory"
	"errors"
)

type translateString struct {
	factory factory.Factory
}

//contructor
//parametros:
//	factory: factory usado para obtener el encoder a usar para traducir el texto desde el formato
//			origen hacia el formato destino
func NewTranslateString(factory factory.Factory) *translateString {
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
	var encoder map[string]string
	var outText string
	var err error

	switch fmtOrigin {
	case TEXT: 
		if(fmtDestiny == BINARY){
			encoder = *translate.factory.GetEncoder("TEXT_BINARY")
		}
		//aquí debe recorrer el string de entrada y usar el encoder para traducirlo
		outText = encoder[inText]
	default: 
		err = errors.New("ocurrio un error, formato invalido")	
	}
	return outText,err
}

