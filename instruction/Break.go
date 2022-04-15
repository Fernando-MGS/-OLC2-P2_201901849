package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"
)

type Break struct {
	Retorno interfaces.Expresion
	Tipo    bool
	line    string
	col     string
}

func NewBreak(retorno interfaces.Expresion, tipo bool, line, col int) Break {
	brk := Break{retorno, tipo, strconv.Itoa(line), strconv.Itoa(col)}
	return brk
}

func (b Break) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var ret interfaces.Value
	ret.Type = interfaces.NULL
	ambito := env.(environment.Environment).DevAmbito()
	if env.(environment.Environment).Control.Ciclo {
		if b.Tipo {

		} else {
			code := "//BREAK\n" + "goto " + env.(environment.Environment).Control.Salida + ";\n//FIN DE BREAK\n"
			gen.AddCodes(code, ambito)
		}
	} else {
		env.(environment.Environment).NewError("NO SE PUEDE USAR UN CONTINUE FUERA DE UN CICLO", b.line, b.col)
	}

	//fmt.Println("El retorno 2 es ")
	return b
}
