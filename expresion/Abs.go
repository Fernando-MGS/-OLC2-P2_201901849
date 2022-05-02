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
	name := env.(environment.Environment).Control.Id
	argumento := a.Arg.Ejecutar(env, gen)
	//ambito := env.(environment.Environment).DevAmbito()
	if argumento.Type == interfaces.INTEGER || argumento.Type == interfaces.FLOAT {
		resultado := gen.NewTemp()
		positivo := gen.NewLabel()
		negativo := gen.NewTemp()
		salida := gen.NewLabel()
		//gen.AddCodes("if ("+argumento.Value+">=0) goto "+positivo+";", ambito)
		gen.NewIf(argumento.Value, ">=", "0", positivo, false, "", name, true, true, a.Line)
		gen.NewOperacion(negativo, "0", "-", "1", false, "", name, true, true, a.Line)
		//gen.AddCodes(resultado+"="+argumento.Value+"*-1;", ambito)
		gen.NewOperacion(resultado, argumento.Value, "*", negativo, false, "", name, true, true, a.Line)
		//gen.AddCodes("goto "+salida+";", ambito)
		gen.NewSalto(salida, false, "", name, true, false, a.Line)
		//gen.AddCodes(positivo+":", ambito)
		gen.NewLabels(positivo, false, "", name, true, true, "")
		//gen.AddCodes(resultado+"="+argumento.Value+";", ambito)
		gen.NewAsignacion(resultado, argumento.Value, false, "", name, true, true, a.Line)
		//gen.AddCodes(salida+":", ambito)
		gen.NewLabels(salida, false, "", name, true, true, "")
		retorno.Value = resultado
		retorno.Type = argumento.Type
		retorno.IsTemp = true
	} else {
		env.(environment.Environment).NewError("SE ESPERABA UNA EXPRESION NUMERICA", a.Line, a.Col)
	}
	return retorno
}
