package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"
)

type Len struct {
	Acceso interfaces.Expresion
	Line   string
	Col    string
}

func NewLen(expr interfaces.Expresion, line, col int) Len {
	return Len{expr, strconv.Itoa(line), strconv.Itoa(col)}
}

func (l Len) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	ambito := env.(environment.Environment).DevAmbito()
	array := l.Acceso.Ejecutar(env, gen)
	if array.Type != interfaces.NULL {
		if array.Type == interfaces.ARRAY || array.Type == interfaces.VECTOR {
			retorno.Type = interfaces.INTEGER
			size := gen.NewTemp()
			gen.AddCodes(size+"=HEAP[(int)"+array.Value+"];", ambito)
			retorno.Value = size
			retorno.IsTemp = true
		} else {
			env.(environment.Environment).NewError("SE ESPERABA UN VECTOR O UN ARRAY", l.Line, l.Col)
			retorno.Type = interfaces.NULL
		}
	} else {
		retorno.Type = interfaces.NULL
	}
	return retorno
}
