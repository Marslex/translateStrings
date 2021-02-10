package factory

type Factory interface{
	GetEncoder(encoder string) *map[string]string
	setTBinary(encoder string)
}

type factoryEncoder struct {
	tBinary *map[string]string
	tMorse *map[string]string
	bText *map[string]string
	mText *map[string]string
}

const (
	textToBinary = "TEXT_BINARY"
	textToMorse  = "TEXT_MORSE"
	binaryToText = "BINARY_TEXT"
	morseToText  = "MORSE_TEXT"
)

func NewFactoryEncoder() *factoryEncoder {
	return &factoryEncoder{
	}
}

func (factory factoryEncoder) GetEncoder(encode string) *map[string]string {

	var m *map[string]string
	switch encode {
	case "TEXT_BINARY": 
		if factory.tBinary==nil {
			factory.setTBinary()
		}
		m = factory.tBinary
	}

	return m
}

func (factory factoryEncoder) setTBinary(){
	factory.tBinary = &map[string]string{
		"A": "01000001",
		"B": "01000010",
		"C": "01000011",
		"D": "01000100",
		"E": "01000101",
		"F": "01000110",
		"G": "01000111",
		"H": "01001000",
		"I": "01001001",
		"J": "01001010",
		"K": "01001011",
		"L": "01001100",
		"M": "01001101",
		"N": "01001110",
		"O": "01001111",
		"P": "01010000",
		"Q": "01010001",
		"R": "01010010",
		"S": "01010011",
		"T": "01010100",
		"U": "01010101",
		"V": "01010110",
		"W": "01010111",
		"X": "01011000",
		"Y": "01011001",
		"Z": "01011010",
	}
}
