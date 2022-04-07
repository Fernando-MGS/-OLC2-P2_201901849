package expresion

import (
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strings"
)

type Primitivo struct {
	Valor interface{}
	Tipo  interfaces.TipoExpresion
}

func (p Primitivo) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	if p.Tipo == interfaces.STRING || p.Tipo == interfaces.CHAR {
		cadena := fmt.Sprintf("%v", p.Valor)
		cadena = strings.Replace(cadena, "\\n", "\n", -1)
		cadena = strings.Replace(cadena, "\\t", "\t", -1)
		cadena = strings.Replace(cadena, "\\\"", "\"", -1)
		cadena = strings.Replace(cadena, "\\'", "'", -1)
		//cadena = strings.Replace(cadena, "\\\\", "\\", -1)
		p.Valor = cadena
		//fmt.Println(p.Valor)
	}
	if p.Tipo == interfaces.CHAR {
		r1 := fmt.Sprintf("%v", p.Valor)
		fmt.Println(r1)
		//p.Valor=(int)p.Valor.(char)
	}
	return interfaces.Value{
		Value:      fmt.Sprintf("%v", p.Valor),
		IsTemp:     false,
		Type:       p.Tipo,
		TrueLabel:  "",
		FalseLabel: "",
	}
}

func NewPrimitivo(val interface{}, tipo interfaces.TipoExpresion) Primitivo {
	exp := Primitivo{val, tipo}
	return exp
}
