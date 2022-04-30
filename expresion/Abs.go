package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"
)

type Abs struct {
	Arg  interfaces.Expresion
	Line string
	Col  string
}

func NewAbsoluto(arg interfaces.Expresion, line, col int) Abs {
	return Abs{arg, strconv.Itoa(line), strconv.Itoa(col)}
}

func (a Abs) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	argumento := a.Arg.Ejecutar(env, gen)
	ambito := env.(environment.Environment).DevAmbito()
	if argumento.Type == interfaces.INTEGER || argumento.Type == interfaces.FLOAT {
		resultado := gen.NewTemp()
		positivo := gen.NewLabel()
		salida := gen.NewLabel()
		gen.AddCodes("if ("+argumento.Value+">=0) goto "+positivo+";", ambito)
		gen.AddCodes(resultado+"="+argumento.Value+"*-1;", ambito)
		gen.AddCodes("goto "+salida+";", ambito)
		gen.AddCodes(positivo+":", ambito)
		gen.AddCodes(resultado+"="+argumento.Value+";", ambito)
		gen.AddCodes(salida+":", ambito)
		retorno.Value = resultado
		retorno.Type = argumento.Type
		retorno.IsTemp = true
	} else {
		env.(environment.Environment).NewError("SE ESPERABA UNA EXPRESION NUMERICA", a.Line, a.Col)
	}
	return retorno
}
