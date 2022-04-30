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
	name := env.(environment.Environment).Control.Id
	ambito := true
	true_l := ""
	false_l := ""
	if name == "GLOBAL" || name == "main" {
		ambito = false
	}
	if p.Tipo == interfaces.STRING || p.Tipo == interfaces.CHAR {

		cadena := fmt.Sprintf("%v", p.Valor)
		cadena = strings.Replace(cadena, "\\n", "\n", -1)
		cadena = strings.Replace(cadena, "\\t", "\t", -1)
		cadena = strings.Replace(cadena, "\\\"", "\"", -1)
		cadena = strings.Replace(cadena, "\\'", "'", -1)
		//cadena = strings.Replace(cadena, "\\\\", "\\", -1)
		p.Valor = cadena
		if p.Tipo == interfaces.STRING {
			tmp := gen.NewTemp()
			p.Valor = tmp
			declarar := "//INICIO DE STRING\n" + tmp + "=H;"
			gen.AddCodes(declarar, ambito)
			chars := gen.Array_char(cadena)
			for _, s := range chars {
				val := strconv.Itoa(int(s))
				save := "HEAP[(int)H]=" + val + ";//" + string(s) + "\n"
				save += "H=H+1;"
				gen.Heap++
				gen.AddCodes(save, ambito)

			}
			save := "HEAP[(int)H]=-1;\n"
			save += "H=H+1;"
			gen.Heap++
			gen.AddCodes(save, ambito)
			gen.AddCodes("//FIN DE STRING", ambito)
		}
		//fmt.Println(p.Valor)
	} else if p.Tipo == interfaces.BOOLEAN {

		true_l = gen.NewLabel()
		false_l = gen.NewLabel()
		val := fmt.Sprintf("%v", p.Valor)
		fmt.Println("EL VALOR DEL BOOL ES ", val)
		gen.AddCodes("//INICIO DE BOOLEANO", ambito)
		gen.AddCodes("if("+val+"==1) goto "+true_l+";", ambito)
		gen.AddCodes("goto "+false_l+";", ambito)
		gen.AddCodes("//FIN DE BOOLEANO "+p.Linea+"-"+p.Col, ambito)
		/*gen.AddTempBool(l1, l2)
		gen.SetConf()*/

	}
	return interfaces.Value{
		Value:      fmt.Sprintf("%v", p.Valor),
		IsTemp:     false,
		Type:       p.Tipo,
		TrueLabel:  true_l,
		FalseLabel: false_l,
	}
}

func NewPrimitivo(val interface{}, tipo interfaces.TipoExpresion, col, line int) Primitivo {

	exp := Primitivo{val, tipo, strconv.Itoa(line), strconv.Itoa(col)}
	return exp
}
