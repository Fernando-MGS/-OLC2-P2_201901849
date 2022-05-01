package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"
)

type LenT struct {
	Acceso interfaces.Expresion
	Line   string
	Col    string
}

func NewLenT(expr interfaces.Expresion, line, col int) LenT {
	return LenT{expr, strconv.Itoa(line), strconv.Itoa(col)}
}

func (l LenT) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	ambito := env.(environment.Environment).DevAmbito()
	array := l.Acceso.Ejecutar(env, gen)
	if array.Type != interfaces.NULL {
		if array.Type == interfaces.ARRAY || array.Type == interfaces.VECTOR {
			retorno.Type = interfaces.INTEGER
			size := gen.NewTemp()
			index := gen.NewTemp()
			gen.AddCodes(index+"=P+1;\n", ambito)
			gen.AddCodes("STACK[(int)"+index+"]="+array.Value+";", ambito)
			gen.AddFuncExtra("LENVECTOR")
			gen.AddCodes("proc_LenVector();", ambito)
			gen.AddCodes(size+"=STACK[(int)P];", ambito)
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
