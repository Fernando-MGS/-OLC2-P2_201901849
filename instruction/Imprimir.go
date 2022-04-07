package instruction

import (
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
)

type Imprimir struct {
	Expresion interfaces.Expresion
}

func NewImprimir(val interfaces.Expresion) Imprimir {
	exp := Imprimir{val}
	return exp
}

func (p Imprimir) Ejecutar(env interface{}, gen *generator.Generator) interface{} {

	result := p.Expresion.Ejecutar(env, gen)
	//fmt.Println(result)
	if result.Type == interfaces.INTEGER || result.Type == interfaces.USIZE {
		gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))
	} else if result.Type == interfaces.FLOAT {
		gen.AddPrintf("f", fmt.Sprintf("%v", result.Value))
	} else if result.Type == interfaces.CHAR {
		gen.AddPrintf("c", fmt.Sprintf("%v", result.Value))
	} else if result.Type == interfaces.STR || result.Type == interfaces.STRING {

	}
	//gen.AddPrintf("c", fmt.Sprintf("%v", result.Value))
	//salto de línea
	gen.AddCode("printf(\"\\n\");")
	var retorno interfaces.Symbol
	retorno.Tipo.Tipo = interfaces.NULL
	return retorno
}
