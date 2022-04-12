package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"
	"strings"
)

type Primitivo struct {
	Valor interface{}
	Tipo  interfaces.TipoExpresion
	Linea string
	Col   string
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
	} else if p.Tipo == interfaces.BOOLEAN {
		name := env.(environment.Environment).Control.Id
		ambito := true
		if name == "GLOBAL" || name == "main" {
			ambito = false
		}
		l1 := gen.NewLabel()
		l2 := ""

		conf := gen.GetConf()
		if conf == 0 {
			l2 = gen.NewLabel()
			if p.Valor == 0 {
				gen.AddCodes("goto "+l2+";\n", ambito)
			} else {
				gen.AddCodes("goto "+l1+";\n", ambito)
			}
			gen.AddTempBool(l1, l2)
			gen.SetConf()
		} else {
			l1 = gen.NewLabel()
			gen.AddTempBool(l1, "goto ")
		}

		/*gen.AddTempBool(l1, l2)
		gen.SetConf()*/

	}
	return interfaces.Value{
		Value:      fmt.Sprintf("%v", p.Valor),
		IsTemp:     false,
		Type:       p.Tipo,
		TrueLabel:  "",
		FalseLabel: "",
	}
}

func NewPrimitivo(val interface{}, tipo interfaces.TipoExpresion, col, line int) Primitivo {

	exp := Primitivo{val, tipo, strconv.Itoa(line), strconv.Itoa(col)}
	return exp
}
