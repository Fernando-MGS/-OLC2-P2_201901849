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
	//ambito := env.(environment.Environment).DevAmbito()
	//fmt.Println("VAMO A IMPRIMIR", len(p.Formato))
	//gen.AddCodes("//INICIO DE PRINT", ambito)
	name := env.(environment.Environment).Control.Id
	gen.NewComentario("INICIO DE PRINT", name, true, true, p.line)
	if len(p.Formato) >= 1 {
		index := 0
		textos, formato, conf := Analisis_Print(p.Formato)
		fmt.Println(formato)
		if conf {
			str := textos[index]
			prim := expresion.Primitivo{str, interfaces.STRING, "0", "0"}
			console(prim, env, gen, p.line)
			index++
		}
		for _, s := range p.List_Expresion.ToArray() {
			expr := s.(interfaces.Expresion)
			console(expr, env, gen, p.line)
			if index < len(textos) {
				str := textos[index]
				prim := expresion.Primitivo{str, interfaces.STRING, "0", "0"}
				console(prim, env, gen, p.line)
				index++
			}
		}
		//gen.AddPrintf("c", "10", ambito)
		gen.NewPrint("c", "10", true, "SALTO DE LINEA", name, true, true, p.line)
	} else {
		//fmt.Println("EL LARGO DE LISTA ES ", p.List_Expresion.Len())
		aux := p.List_Expresion.GetValue(0).(interfaces.Expresion)
		tipo := reflect.TypeOf(aux)
		t := fmt.Sprintf("%v", tipo)
		//exec := aux.Ejecutar(env, gen)
		///if exec.Type != interfaces.NULL {
		if t == "expresion.Primitivo" {
			console(aux, env, gen, p.line)
			//gen.AddPrintf("c", "10", ambito)
			gen.NewPrint("c", "10", true, "SALTO DE LINEA", name, true, true, p.line)
		} else {
			env.(environment.Environment).NewError("EL FORMATO NO ES CORRECTO. SE ESPERABA UNA EXPRESION SIMPLE", p.line, p.col)
		}
		//}
	}
	//gen.AddCodes("//FIN DE PRINT", ambito)
	gen.NewComentario("FIN DE PRINT", name, true, true, p.line)
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

func console(p interfaces.Expresion, env interface{}, gen *generator.Generator, linea string) {
	fmt.Println("console")
	result := p.Ejecutar(env, gen)
	//ambito := env.(environment.Environment).DevAmbito()
	//fmt.Println(result)
	name := env.(environment.Environment).Control.Id
	if result.Type == interfaces.INTEGER {
		//gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value), ambito)
		gen.NewPrint("d", "(int)"+fmt.Sprintf("%v", result.Value), false, "", name, true, true, linea)
	} else if result.Type == interfaces.FLOAT {
		//gen.AddPrintf("f", fmt.Sprintf("%v", result.Value), ambito)
		gen.NewPrint("f", fmt.Sprintf("%v", result.Value), false, "", name, true, true, linea)
	} else if result.Type == interfaces.CHAR {
		//gen.AddPrintf("c", "(int)"+fmt.Sprintf("%v", result.Value), ambito)
		gen.NewPrint("c", "(int)"+fmt.Sprintf("%v", result.Value), false, "", name, true, true, linea)
	} else if result.Type == interfaces.BOOLEAN {
		//value := ""
		l1 := result.TrueLabel  //ln	true
		l2 := result.FalseLabel //ln+1	false
		l3 := gen.NewLabel()    //ln+2	exit
		/*if result.IsTemp {
		value = "if (" + result.Value + "==1) goto " + l1 + ";\n"
		value += "goto " + l2 + ";\n"*/
		//value += l1 + ":\n"
		gen.NewLabels(l1, false, "", name, true, true, linea)
		//value += "proc_extra_printTrue();\n"
		gen.NewLlamada("proc_extra_printTrue", true, "IMPRIMIR BOOLEANOS", name, true, false, linea)
		//value += "goto " + l3 + ";\n"
		gen.NewSalto(l3, false, "", name, true, true, linea)
		//value += l2 + ":\n"
		gen.NewLabels(l2, false, "", name, true, true, linea)
		//value += "proc_extra_printFalse();\n"
		gen.NewLlamada("proc_extra_printFalse", true, "IMPRIMIR BOOLEANOS", name, true, false, linea)
		//value += l3 + ":\n"
		gen.NewLabels(l3, false, "", name, true, true, linea)
		//gen.AddCodes(value, ambito)
		//gen.SetConf()
		gen.AddFuncExtra("PRINTBOOL")
	} else if result.Type == interfaces.STR || result.Type == interfaces.STRING {
		l1 := gen.NewLabel()
		l2 := gen.NewLabel()
		//code := "if (" + result.Value + "==-1) goto " + l1 + ";\n"
		gen.NewIf(result.Value, "==", "-1", l1, false, "", name, true, true, linea)
		//code += "P=P+1;\n"
		gen.NewOperacion("P", "P", "+", "1", false, "", name, true, false, linea)
		env.(environment.Environment).Control.Stack++
		gen.Stack++
		t1 := gen.NewTemp()
		//code += t1 + "=P+1;\n"
		gen.NewOperacion(t1, "P", "+", "1", false, "", name, true, true, linea)
		//code += "STACK[(int)" + t1 + "] =" + result.Value + ";\n"
		gen.NewStack(t1, result.Value, false, "", name, true, true, linea)
		//code += "proc_printString();\n"
		gen.NewLlamada("proc_printString", false, "", name, true, false, linea)
		//code += "P=P-1;\n"
		gen.NewOperacion("P", "P", "-", "1", false, "", name, true, false, linea)
		/*env.(environment.Environment).Control.Stack--
		gen.Stack--*/
		//code += "goto " + l2 + ";\n"
		gen.NewSalto(l2, false, "", name, true, true, linea)
		//code += l1 + ":\n"
		gen.NewLabels(l1, false, "", name, true, true, linea)
		//code += l2 + ":\n"
		gen.NewLabels(l2, false, "", name, true, true, linea)
		//gen.AddCodes(code, ambito)
		gen.AddFuncExtra("PRINTSTR")
	} else if result.Type == interfaces.ARRAY {
		dimension := result.Tipo2.GetValue(0).(interfaces.Dimensions)
		largo := strconv.Itoa(dimension.Dimensions.Len())
		tmp1 := gen.NewTemp()
		//gen.AddCodes(tmp1+"=1+P;", ambito)
		gen.NewOperacion(tmp1, "1", "+", "P", false, "", name, true, true, linea)
		tmp2 := gen.NewTemp()
		//gen.AddCodes(tmp2+"=2+P;", ambito)
		gen.NewOperacion(tmp2, "2", "+", "P", false, "", name, true, true, linea)
		//gen.AddCodes("STACK[(int)"+tmp1+"]="+largo+";", ambito)
		gen.NewStack(tmp1, largo, false, "", name, true, true, linea)
		//gen.AddCodes("STACK[(int)"+tmp2+"]="+result.Value+";", ambito)
		gen.NewStack(tmp2, result.Value, false, "", name, true, true, linea)
		//fmt.Println("EL LARGO ES " + largo)
		if dimension.Tipo == interfaces.INTEGER {
			//gen.AddCodes("//IMPRIMIENDO ARRAY DE INT", ambito)
			gen.NewComentario("IMPRIMIENDO ARRAY DE INT", name, true, false, linea)
			gen.AddFuncExtra("SETINT")
			//gen.AddCodes("proc_printInt();", ambito)
			gen.NewLlamada("proc_printInt", true, "IMPRIMIENDO SET DE ENTEROS", name, true, false, linea)
		} else if dimension.Tipo == interfaces.FLOAT {
			//gen.AddCodes("//IMPRIMIENDO ARRAY DE FLOAT", ambito)
			gen.NewComentario("IMPRIMIENDO ARRAY DE FLOAT", name, true, false, linea)
			gen.AddFuncExtra("SETFLOAT")
			//gen.AddCodes("proc_printFloat();", ambito)
			gen.NewLlamada("proc_printFloat", true, "IMPRIMIENDO SET DE FLOAT", name, true, false, linea)
		} else if dimension.Tipo == interfaces.CHAR {
			//gen.AddCodes("//IMPRIMIENDO ARRAY DE CHAR", ambito)
			gen.NewComentario("IMPRIMIENDO ARRAY DE CHAR", name, true, false, linea)
			gen.AddFuncExtra("SETCHAR")
			//gen.AddCodes("proc_printChar();", ambito)
			gen.NewLlamada("proc_printChar", true, "IMPRIMIENDO SET DE FLOAT", name, true, false, linea)
		} else if dimension.Tipo == interfaces.STR || dimension.Tipo == interfaces.STRING {
			//gen.AddCodes("//IMPRIMIENDO ARRAY DE INT", ambito)
			gen.NewComentario("IMPRIMIENDO ARRAY DE STRING|| STR", name, true, false, linea)
			gen.AddFuncExtra("SETSTR")
			//gen.AddCodes("proc_printStr();", ambito)
			gen.NewLlamada("proc_printStr", true, "IMPRIMIENDO SET DE STRING", name, true, false, linea)
		}

		//gen.AddCodes(tmp2+"=2+P;",ambito)
		/*fmt.Println(result.Value)
		fmt.Println(result.TrueLabel)*/
	} else if result.Type == interfaces.VECTOR {

		//fmt.Println(result.Tipo2.ToArray()...)
		dimension := result.Tipo2.GetValue(0).(interfaces.Dimensions)
		largo := strconv.Itoa(dimension.Dimensions.Len())
		//fmt.Println("EL LARGO ES ", largo)
		tmp1 := gen.NewTemp()
		//gen.AddCodes(tmp1+"=1+P;", ambito)
		gen.NewOperacion(tmp1, "1", "+", "P", false, "", name, true, true, linea)
		tmp2 := gen.NewTemp()
		//gen.AddCodes(tmp2+"=2+P;", ambito)
		gen.NewOperacion(tmp2, "2", "+", "P", true, "//2+9", name, true, true, linea)
		//gen.AddCodes("STACK[(int)"+tmp1+"]="+largo+";", ambito)
		gen.NewStack(tmp1, largo, false, "", name, true, true, linea)
		//gen.AddCodes("STACK[(int)"+tmp2+"]="+result.Value+";", ambito)
		gen.NewStack(tmp2, result.Value, false, "", name, true, true, linea)
		//fmt.Println("EL LARGO ES " + largo)
		if dimension.Tipo == interfaces.INTEGER {
			//gen.AddCodes("//IMPRIMIENDO VECTOR DE INT", ambito)
			gen.NewComentario("IMPRIMIENDO VECTOR DE INT", name, true, false, linea)
			gen.AddFuncExtra("SETINT")
			//gen.AddCodes("proc_printInt();", ambito)
			gen.NewLlamada("proc_printInt", true, "IMPRIMIENDO SET DE ENTEROS", name, true, false, linea)
		} else if dimension.Tipo == interfaces.FLOAT {
			//gen.AddCodes("//IMPRIMIENDO VECTOR DE FLOAT", ambito)
			gen.NewComentario("IMPRIMIENDO VECTOR DE FLOAT", name, true, false, linea)
			gen.AddFuncExtra("SETFLOAT")
			//gen.AddCodes("proc_printFloat();", ambito)
			gen.NewLlamada("proc_printFloat", true, "IMPRIMIENDO SET DE FLOAT", name, true, false, linea)
		} else if dimension.Tipo == interfaces.CHAR {
			//gen.AddCodes("//IMPRIMIENDO FLOAT DE CHAR", ambito)
			gen.NewComentario("IMPRIMIENDO VECTOR DE CHAR", name, true, false, linea)
			gen.AddFuncExtra("SETCHAR")
			//gen.AddCodes("proc_printChar();", ambito)
			gen.NewLlamada("proc_printChar", true, "IMPRIMIENDO SET DE FLOAT", name, true, false, linea)
		} else if dimension.Tipo == interfaces.STR || dimension.Tipo == interfaces.STRING {
			//gen.AddCodes("//IMPRIMIENDO VECTOR DE STRING", ambito)
			gen.NewComentario("IMPRIMIENDO ARRAY DE STRING|| STR", name, true, false, linea)
			gen.AddFuncExtra("SETSTR")
			//gen.AddCodes("proc_printStr();", ambito)
			gen.NewLlamada("proc_printStr", true, "IMPRIMIENDO SET DE STRING", name, true, false, linea)
		}
	}
}
