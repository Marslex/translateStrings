package translateStrings

import (
	"reflect"
	"testing"
)

func Test_translateString_Translate(t *testing.T) {
	type args struct {
		inText     String
		fmtOrigin  String
		fmtDestiny String
	}
	factory := NewFactoryEncoder()
	translate := translate.NewTranslateString(factory)

	testArgs := args{
		inText: 	"A"
		fmtOrigin:  "TEXTO"
		fmtDestiny: "BINARY"
	}

	tests := []struct {
		name      string
		translate *translateString
		args      args
		want      String
		wantErr   bool
	}{
		name      "test de translate"
		translate translate
		args      testArgs
		want      "01000001"
		wantErr   false
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.translate.Translate(tt.args.inText, tt.args.fmtOrigin, tt.args.fmtDestiny)
			if (err != nil) != tt.wantErr {
				t.Errorf("translateString.Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("translateString.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
