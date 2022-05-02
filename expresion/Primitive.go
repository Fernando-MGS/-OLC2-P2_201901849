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
	//ambito := true
	true_l := ""
	false_l := ""
	/*	if name == "GLOBAL" || name == "main" {
		ambito = false
	}*/
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
			comentario := generator.Comentario{Comentario: "INICIO DE STRING"}
			fragmento := generator.Fragment{Valor: comentario, Tipo: 5, Valido: true, Mod: false, Line: p.Linea}
			gen.AddFragment(name, fragmento)
			//declarar := "//INICIO DE STRING\n" + tmp + "=H;"
			//gen.AddCodes(declarar, ambito)
			code := generator.Asignacion{Destino: tmp, Op: "H", Comentario: false, Comentario_Cont: ""}
			fragmento = generator.Fragment{Valor: code, Tipo: 1, Valido: true, Mod: true, Line: p.Linea}
			gen.AddFragment(name, fragmento)
			chars := gen.Array_char(cadena)
			for _, s := range chars {
				val := strconv.Itoa(int(s))
				//save := "HEAP[(int)H]=" + val + ";//" + string(s) + "\n"
				code := generator.Heap{Index: "H", Value: val, Comentario: true, Comentario_Cont: string(s)}
				fragmento = generator.Fragment{Valor: code, Tipo: 3, Valido: true, Mod: false, Line: p.Linea}
				gen.AddFragment(name, fragmento)
				//save += "H=H+1;"
				operacion := generator.Operacion{Destino: "H", Op1: "H", Signo: "+", Op2: "1", Comentario: false, Comentario_Cont: ""}
				fragmento = generator.Fragment{Valor: operacion, Tipo: 0, Valido: true, Mod: false, Line: p.Linea}
				gen.AddFragment(name, fragmento)
				gen.Heap++
				//gen.AddCodes(save, ambito)

			}
			code1 := generator.Heap{Index: "H", Value: "-1", Comentario: false, Comentario_Cont: ""}
			fragmento = generator.Fragment{Valor: code1, Tipo: 3, Valido: true, Mod: false, Line: p.Linea}
			gen.AddFragment(name, fragmento)
			//save := "HEAP[(int)H]=-1;\n"
			operacion := generator.Operacion{Destino: "H", Op1: "H", Signo: "+", Op2: "1", Comentario: false, Comentario_Cont: ""}
			fragmento = generator.Fragment{Valor: operacion, Tipo: 0, Valido: true, Mod: false, Line: p.Linea}
			gen.AddFragment(name, fragmento)
			//save += "H=H+1;"
			gen.Heap++
			//gen.AddCodes(save, ambito)
			comentario = generator.Comentario{Comentario: "FIN DE STRING"}
			fragmento = generator.Fragment{Valor: comentario, Tipo: 5, Valido: true, Mod: false, Line: p.Linea}
			gen.AddFragment(name, fragmento)
			//gen.AddCodes("//FIN DE STRING", ambito)
		}
		//fmt.Println(p.Valor)
	} else if p.Tipo == interfaces.BOOLEAN {
		var code interface{}
		true_l = gen.NewLabel()
		false_l = gen.NewLabel()
		val := fmt.Sprintf("%v", p.Valor)
		code = generator.Comentario{Comentario: "INICIO DE BOOLEANO"}
		fragmento := generator.Fragment{Valor: code, Tipo: 5, Valido: true, Mod: false, Line: p.Linea}
		gen.AddFragment(name, fragmento)
		//fmt.Println("EL VALOR DEL BOOL ES ", val)
		//gen.AddCodes("//INICIO DE BOOLEANO", ambito)
		//gen.AddCodes("if("+val+"==1) goto "+true_l+";", ambito)
		code = generator.If{Op1: val, Signo: "==", Op2: "1", Salto: true_l, Comentario: false, Comentario_Cont: ""}
		fragmento = generator.Fragment{Valor: code, Tipo: 2, Valido: true, Mod: false, Line: p.Linea}
		gen.AddFragment(name, fragmento)
		//gen.AddCodes("goto "+false_l+";", ambito)
		code = generator.Salto{Label: false_l, Comentario: false, Comentario_Cont: ""}
		fragmento = generator.Fragment{Valor: code, Tipo: 9, Valido: true, Mod: false, Line: p.Linea}
		gen.AddFragment(name, fragmento)
		//gen.AddCodes("//FIN DE BOOLEANO "+p.Linea+"-"+p.Col, ambito)
		code = generator.Comentario{Comentario: "FIN DE BOOLEANO " + p.Linea + "-" + p.Col}
		fragmento = generator.Fragment{Valor: code, Tipo: 5, Valido: true, Mod: false, Line: p.Linea}
		gen.AddFragment(name, fragmento)
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
