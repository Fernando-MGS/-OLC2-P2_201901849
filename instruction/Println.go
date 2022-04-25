package instruction

import (
	"OLC2/environment"
	"OLC2/expresion"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/colegno/arraylist"
)

type Print struct {
	List_Expresion *arraylist.List
	Formato        string
	line           string
	col            string
}

func NewPrint(expr *arraylist.List, format string, line, col int) Print {
	exp := Print{expr, format, strconv.Itoa(line), strconv.Itoa(col)}
	return exp
}

func (p Print) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var ret interfaces.Value
	ret.Type = interfaces.NULL
	ambito := env.(environment.Environment).DevAmbito()
	gen.AddCodes("//INICIO DE PRINT", ambito)
	if len(p.Formato) >= 1 {
		index := 0
		textos, formato, conf := Analisis_Print(p.Formato)
		fmt.Println(formato)
		if conf {
			str := textos[index]
			prim := expresion.Primitivo{str, interfaces.STRING, "0", "0"}
			console(prim, env, gen)
			index++
		}
		for _, s := range p.List_Expresion.ToArray() {
			expr := s.(interfaces.Expresion)
			console(expr, env, gen)
			if index < len(textos) {
				str := textos[index]
				prim := expresion.Primitivo{str, interfaces.STRING, "0", "0"}
				console(prim, env, gen)
				index++
			}
		}
		gen.AddPrintf("c", "10", ambito)
	} else {
		aux := p.List_Expresion.GetValue(0).(interfaces.Expresion)
		tipo := reflect.TypeOf(aux)
		t := fmt.Sprintf("%v", tipo)
		//exec := aux.Ejecutar(env, gen)
		///if exec.Type != interfaces.NULL {
		if t == "expresion.Primitivo" {
			console(aux, env, gen)
			gen.AddPrintf("c", "10", ambito)
		} else {
			env.(environment.Environment).NewError("EL FORMATO NO ES CORRECTO. SE ESPERABA UNA EXPRESION SIMPLE", p.line, p.col)
		}
		//}
	}
	gen.AddCodes("//FIN DE PRINT", ambito)
	return ret
}

func Analisis_Print(form string) ([]string, string, bool) {
	conf := false
	inicio := 0
	var cadenas []string
	nueva := ""
	formato := ""
	auxs := 0
	cadena := form
	largo := len(cadena)
	k := 0
	texto := ""
	for k < largo {
		if string(cadena[k]) == "{" {
			break
		}
		texto += string(cadena[k])
		k++
	}
	//largo := len(cadena)
	//fmt.Println(texto)

	if posicion := strings.Index(cadena, "{}"); posicion != -1 {
		inicio = posicion
	}
	if posicion := strings.Index(cadena, "{:?}"); posicion != -1 {
		auxs = posicion
	}
	if inicio < auxs {
		inicio = auxs
	} else {
		auxs = inicio
	}
	if inicio != 0 {
		cadenas = append(cadenas, texto)
		conf = true
	}
	texto = ""
	for inicio < largo {
		if string(cadena[inicio]) == "{" {
			if texto != "" {
				cadenas = append(cadenas, texto)
			}
			texto = ""
			inicio++
			if string(cadena[inicio]) == "}" {
				inicio++
				formato += "a"
			} else if string(cadena[inicio]) == ":" {
				inicio++
				if string(cadena[inicio]) == "?" {
					inicio++
					if string(cadena[inicio]) == "}" {
						inicio++
						formato += "b"
					} else {
						formato = "n"
						break
					}
				} else {
					formato = "n"
					break
				}
			} else {
				formato = "n"
				break
			}
		} else {
			//fmt.Println(texto+"texto")
			texto += string(cadena[inicio])
			inicio++
		}
		//inicio++
	}
	cadenas = append(cadenas, texto)
	j := 0
	for j < auxs {
		nueva += string(cadena[j])
		j++
	}
	//fmt.Println(formato)
	/*for _, s := range cadenas {
		fmt.Println(s)
	}*/
	return cadenas, formato, conf
}

func console(p interfaces.Expresion, env interface{}, gen *generator.Generator) {
	result := p.Ejecutar(env, gen)
	ambito := env.(environment.Environment).DevAmbito()
	if result.Type == interfaces.INTEGER {
		gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value), ambito)
	} else if result.Type == interfaces.FLOAT {
		gen.AddPrintf("f", fmt.Sprintf("%v", result.Value), ambito)
	} else if result.Type == interfaces.CHAR {
		gen.AddPrintf("c", "(int)"+fmt.Sprintf("%v", result.Value), ambito)
	} else if result.Type == interfaces.BOOLEAN {
		value := ""
		l1 := gen.GetTempsB().TrueL  //ln	true
		l2 := gen.GetTempsB().FalseL //ln+1	false
		l3 := gen.NewLabel()         //ln+2	exit
		/*if result.IsTemp {
		value = "if (" + result.Value + "==1) goto " + l1 + ";\n"
		value += "goto " + l2 + ";\n"*/
		value += l1 + ":\n"
		value += "proc_extra_printTrue();\n"
		value += "goto " + l3 + ";\n"
		value += l2 + ":\n"
		value += "proc_extra_printFalse();\n"
		value += l3 + ":\n"
		gen.AddCodes(value, ambito)
		gen.SetConf()
		gen.AddFuncExtra("PRINTBOOL")
	} else if result.Type == interfaces.STR || result.Type == interfaces.STRING {
		l1 := gen.NewLabel()
		l2 := gen.NewLabel()
		code := "if (" + result.Value + "==-1) goto " + l1 + ";\n"
		code += "P=P+1;\n"
		gen.Stack++
		t1 := gen.NewTemp()
		code += t1 + "=P+1;\n"
		code += "STACK[(int)" + t1 + "] =" + result.Value + ";\n"
		code += "proc_printString();\n"
		code += "P=P-1;\n"
		gen.Stack--
		code += "goto " + l2 + ";\n"
		code += l1 + ":\n"
		code += l2 + ":\n"
		gen.AddCodes(code, ambito)
		gen.AddFuncExtra("PRINTSTR")
	} else if result.Type == interfaces.ARRAY {
		dimension := result.Tipo2.GetValue(0).(interfaces.Dimensions)

		tmp1 := gen.NewTemp()
		gen.AddCodes(tmp1+"=1+P;", ambito)
		tmp2 := gen.NewTemp()
		gen.AddCodes(tmp2+"=2+P;", ambito)
		gen.AddCodes("STACK[(int)"+tmp1+"]="+result.Value+";", ambito)
		gen.AddCodes("STACK[(int)"+tmp2+"]="+result.TrueLabel+";", ambito)
		if dimension.Tipo == interfaces.INTEGER {
			gen.AddFuncExtra("PRINTINT")
			gen.AddCodes("proc_printSetNumbers();", ambito)
		} else if dimension.Tipo == interfaces.FLOAT {
			gen.AddFuncExtra("PRINTFLOAT")
			gen.AddCodes("proc_printSetFloat();", ambito)
		} else if dimension.Tipo == interfaces.CHAR {
			gen.AddFuncExtra("PRINTCHAR")
			gen.AddCodes("proc_printSetChar();", ambito)
		} else if dimension.Tipo == interfaces.STR || dimension.Tipo == interfaces.STRING {
			gen.AddFuncExtra("PRINTSTRING")
			gen.AddCodes("proc_printSetString();", ambito)
		}

		//gen.AddCodes(tmp2+"=2+P;",ambito)
		fmt.Println(result.Value)
		fmt.Println(result.TrueLabel)

	}
}
