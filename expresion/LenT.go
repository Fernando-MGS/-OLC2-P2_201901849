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
	//ambito := env.(environment.Environment).DevAmbito()
	name := env.(environment.Environment).Control.Id
	array := l.Acceso.Ejecutar(env, gen)
	if array.Type != interfaces.NULL {
		if array.Type == interfaces.ARRAY || array.Type == interfaces.VECTOR {
			retorno.Type = interfaces.INTEGER
			size := gen.NewTemp()
			index := gen.NewTemp()

			gen.NewComentario("CALCULANDO LONGITUD DE UN ARRAY O UN VECTOR", name, true, false, "")
			//gen.AddCodes(index+"=P+1;\n", ambito)
			gen.NewOperacion(index, "P", "+", "1", false, "", name, true, true, l.Line)
			//gen.AddCodes("STACK[(int)"+index+"]="+array.Value+";", ambito)
			gen.NewStack(index, array.Value, false, "", name, true, false, l.Line)
			gen.AddFuncExtra("LENVECTOR")
			//gen.AddCodes("proc_LenVector();", ambito)
			gen.NewLlamada("proc_LenVector", false, "", name, true, true, l.Line)
			//gen.AddCodes(size+"=STACK[(int)P];", ambito)
			gen.NewCallStack(size, "P", false, "", name, true, true, l.Line)
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
