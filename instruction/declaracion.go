package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"
	"time"
)

type Declaracion struct {
	Id        string
	Tipo      interfaces.TipoSimbolo
	Expresion interfaces.Expresion
	Mutable   bool
	Line      string
	Col       string
}

func NewDeclaracion(id string, tipo interfaces.TipoSimbolo, val interfaces.Expresion, isMutable bool, line, col int) Declaracion {
	//fmt.Println(tipo)
	instr := Declaracion{id, tipo, val, isMutable, strconv.Itoa(line), strconv.Itoa(col)}
	return instr
}

func (p Declaracion) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	//fmt.Println("DECLARACION")
	result := p.Expresion.Ejecutar(env, gen)
	conf := false
	ambito := false
	//index := 0
	if p.Tipo.Tipo == interfaces.USIZE && result.Type == interfaces.INTEGER {
		conf = true
	} else if p.Tipo.Tipo == result.Type {
		conf = true
	} else if result.Type == interfaces.NULL {
		result.Type = interfaces.NULL
		return result.Value
	}
	/*fmt.Println("TIPO 1")
	fmt.Println(result.Type)
	fmt.Println("tipo2")
	fmt.Println(p.Tipo.Tipo)*/
	if conf {
		if p.Tipo.Tipo == interfaces.ARRAY {
			fmt.Println(p.Tipo.Tipo)
			fmt.Println(p.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions).Tipo)
			fmt.Println(p.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions).Dimensions.Clone().ToArray()...)
		} else if p.Tipo.Tipo == interfaces.VECTOR {
			fmt.Println(p.Tipo.Tipo)
			fmt.Println(p.Tipo.Tipo2.ToArray()...)
		} else if p.Tipo.Tipo == interfaces.STR || p.Tipo.Tipo == interfaces.STRING {
			code := "//--------------INICIO DE DECLARACION--------" + p.Id + "\n"
			tam := env.(environment.Environment).Control.Stack
			guia := ""
			incremento := ""
			if env.(environment.Environment).Control.Id == "main" || env.(environment.Environment).Control.Id == "GLOBAL" {
				guia = "(int)P"
				tam = gen.Stack
				incremento = "P=P+1;"
				gen.Stack++
			} else {
				tmp := gen.NewTemp()
				tam = gen.Stack + tam
				gen.AddExpression(tmp, "(int)P", strconv.Itoa(tam), "+", true)
				guia = "(int)" + tmp
				ambito = true
				env.(environment.Environment).Control.Stack++
			}
			simbolo := interfaces.Symbol{Id: p.Id, Tipo: p.Tipo, Posicion: tam, Mutable: p.Mutable}
			env.(environment.Environment).SaveVariable(p.Line, p.Col, p.Id, simbolo, p.Tipo)
			code += "STACK[" + guia + " ]=" + result.Value + ";\n"
			code += incremento + "\n"
			code += "//--------------FIN DE DECLARACION--------" + p.Id + ""
			gen.AddCodes(code, ambito)
		} else {

			codigo := "//--------------INICIO DE DECLARACION--------" + p.Id + "\n"
			tam := env.(environment.Environment).Control.Stack
			guia := ""
			incremento := ""
			if env.(environment.Environment).Control.Id == "main" || env.(environment.Environment).Control.Id == "GLOBAL" {
				guia = "(int)P"
				incremento = "P=P+1;"

				tam = gen.Stack
				gen.Stack++
			} else {
				tmp := gen.NewTemp()
				tam = gen.Stack + tam
				gen.AddExpression(tmp, "(int)P", strconv.Itoa(tam), "+", true)
				guia = "(int)" + tmp
				ambito = true
				env.(environment.Environment).Control.Stack++
			}
			simbolo := interfaces.Symbol{Id: p.Id, Tipo: p.Tipo, Posicion: tam, Mutable: p.Mutable}
			if result.Type == interfaces.INTEGER || result.Type == interfaces.FLOAT || result.Type == interfaces.CHAR || result.Type == interfaces.USIZE {

				if interfaces.CHAR == result.Type {
					runes := []rune(result.Value)
					var val string
					for i := 0; i < len(runes); i++ {
						val = strconv.Itoa(int(runes[i]))
						/*fmt.Println("VAL")
						fmt.Println(val)*/
					}
					result.Value = val
				}
				codigo += "STACK[" + guia + "]=" + fmt.Sprintf("%v", result.Value) + ";\n"
				codigo += incremento + "\n"
				codigo += "//------FIN DE DECLARACION--------" + p.Id
				env.(environment.Environment).SaveVariable(p.Line, p.Col, p.Id, simbolo, p.Tipo)
				if ambito {
					gen.AddFunc(codigo)
				} else {
					gen.AddCode(codigo)
				}
			} else if result.Type == interfaces.BOOLEAN {
				value := ""
				l1 := gen.GetTempsB().TrueL
				l2 := gen.GetTempsB().FalseL
				l3 := gen.NewLabel()
				value += l1 + ":\n"
				value += "STACK[" + guia + "]=1;\n" + incremento + "\n"
				value += "goto " + l3 + ";\n"
				value += l2 + ":\n"
				value += "STACK[" + guia + "]=0;\n" + incremento + "\n"
				value += l3 + ":\n"
				gen.AddCodes(value+"//FIN DE DECLARACION", ambito)

				gen.SetConf()
				env.(environment.Environment).SaveVariable(p.Line, p.Col, p.Id, simbolo, p.Tipo)
			}
		}
	} else {
		t := time.Now()
		fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO CONCUERDAN EN LA DECLARACION", Fecha: fecha}
		env.(environment.Environment).Errores(err)
		//fmt.Println(err)
	}
	result.Type = interfaces.NULL
	return result.Value
}
