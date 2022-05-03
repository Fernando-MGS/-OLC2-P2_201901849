package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"

	"github.com/colegno/arraylist"
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
	var fragmento generator.Fragment
	name := env.(environment.Environment).Control.Id
	var code interface{}
	result := p.Expresion.Ejecutar(env, gen)
	conf := false
	//ambito := env.(environment.Environment).DevAmbito()
	//index := 0
	//tipo := reflect.TypeOf(p.Expresion)
	//t := fmt.Sprintf("%v", tipo)
	if p.Tipo.Tipo == interfaces.USIZE && result.Type == interfaces.INTEGER {
		conf = true
	} else if p.Tipo.Tipo == result.Type {
		conf = true
	} else if result.Type == interfaces.NULL {
		result.Type = interfaces.NULL
		return result
	}
	/*fmt.Println("TIPO 1")
	fmt.Println(result.Type)
	fmt.Println("tipo2")
	fmt.Println(p.Tipo.Tipo)*/
	if conf {
		if p.Tipo.Tipo == interfaces.ARRAY {
			//gen.AddCodes("//INICIO DE DECLARACION "+p.Id, ambito)
			code = generator.Comentario{Comentario: "INICIO DE DECLARACION " + p.Id}
			fragmento = generator.Fragment{Valor: code, Tipo: 5, Valido: true, Mod: false, Line: p.Line}
			gen.AddFragment(name, fragmento)
			//fmt.Println(p.Tipo.Tipo2.ToArray()...)
			tipoArr := p.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions).Tipo
			//fmt.Println("simon")
			if tipoArr != result.Tipo2.GetValue(0).(interfaces.Dimensions).Tipo {
				//fmt.Println("NEL1")
				conf = false
			}
			if len(p.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions).Dimensions.ToArray()) != len(result.Tipo2.GetValue(0).(interfaces.Dimensions).Dimensions.ToArray()) {
				//fmt.Println("NEL2")
				conf = false
				env.(environment.Environment).NewError("LAS DIMENSIONES DEL ARREGLO NO CONCUERDAN EN LA DECLARACION", p.Line, p.Col)
				result.Type = interfaces.NULL
				return result
			}
			//fmt.Println(conf)
			if conf {
				posicion2 := ""
				posicion2 = result.Value
				dimension := interfaces.Dimensions{Tipo: tipoArr, Dimensions: result.Tipo2.GetValue(0).(interfaces.Dimensions).Dimensions}
				dimensions_list := arraylist.New()
				dimensions_list.Add(dimension)
				tipoSimbolo := interfaces.TipoSimbolo{Tipo: interfaces.ARRAY, Tipo2: dimensions_list}
				variable := interfaces.Symbol{Id: p.Id, Posicion2: posicion2, Mutable: p.Mutable, Line: p.Line, Col: p.Col, Tipo: tipoSimbolo, Longitud: result.TrueLabel}
				env.(environment.Environment).SaveVariable(p.Line, p.Col, p.Id, variable, variable.Tipo)
			} else {
				env.(environment.Environment).NewError("LOS TIPOS NO CONCUERDAN EN LA DECLARACION", p.Line, p.Col)
				result.Type = interfaces.NULL
				return result
			}
		} else if p.Tipo.Tipo == interfaces.VECTOR {
			fmt.Println("hola 0")
			dimension_res := result.Tipo2.GetValue(0).(interfaces.Dimensions)
			fmt.Println("hola 1")
			fmt.Println(p.Tipo.Tipo2.GetValue(0))
			dimension_tipo := p.Tipo.Tipo2.GetValue(0).(interfaces.TipoExpresion)
			fmt.Println("hola 2")
			if dimension_res.Tipo == interfaces.VOID || dimension_res.Tipo == dimension_tipo {
				if dimension_res.Dimensions.Len() != p.Tipo.Tipo2.Len() {
					env.(environment.Environment).NewError("LOS TAMAÑOS NO CONCUERDAN", p.Line, p.Col)
					result.Type = interfaces.NULL
					return result
				} else {
					posicion2 := ""
					posicion2 = result.Value
					dimension := interfaces.Dimensions{Tipo: dimension_tipo, Dimensions: p.Tipo.Tipo2}
					dimensions_list := arraylist.New()
					dimensions_list.Add(dimension)
					tipoSimbolo := interfaces.TipoSimbolo{Tipo: interfaces.VECTOR, Tipo2: dimensions_list}
					variable := interfaces.Symbol{Id: p.Id, Posicion2: posicion2, Mutable: p.Mutable, Line: p.Line, Col: p.Col, Tipo: tipoSimbolo, Longitud: result.TrueLabel}
					env.(environment.Environment).SaveVariable(p.Line, p.Col, p.Id, variable, variable.Tipo)
					//fmt.Println("se guardo EN " + posicion2)
				}
			} else {
				env.(environment.Environment).NewError("LOS TIPOS NO CONCUERDAN EN LA DECLARACION", p.Line, p.Col)
				result.Type = interfaces.NULL
				return result
			}
		} else if p.Tipo.Tipo == interfaces.STR || p.Tipo.Tipo == interfaces.STRING {
			//code := "//--------------INICIO DE DECLARACION--------" + p.Id + "\n"
			code = generator.Comentario{Comentario: "--------------INICIO DE DECLARACION--------" + p.Id}
			fragmento = generator.Fragment{Valor: code, Tipo: 5, Valido: true, Mod: false, Line: p.Line}
			gen.AddFragment(name, fragmento)
			//tam := gen.NewTemp()
			//guia := ""
			tmp_pos := gen.NewTemp()
			simbolo := interfaces.Symbol{Id: p.Id, Tipo: p.Tipo, Posicion: tmp_pos, Mutable: p.Mutable}
			env.(environment.Environment).SaveVariable(p.Line, p.Col, p.Id, simbolo, p.Tipo)
			//code += tmp_pos + "=P;"
			code = generator.Asignacion{Destino: tmp_pos, Op: "P", Comentario: false, Comentario_Cont: ""}
			fragmento = generator.Fragment{Valor: code, Tipo: 1, Valido: true, Mod: true, Line: p.Line}
			gen.AddFragment(name, fragmento)
			//code += "STACK[(int)" + tmp_pos + "]=" + result.Value + ";\n"
			code = generator.Stack{Index: tmp_pos, Value: result.Value, Comentario: false, Comentario_Cont: ""}
			fragmento = generator.Fragment{Valor: code, Tipo: 4, Valido: true, Mod: true, Line: p.Line}
			gen.AddFragment(name, fragmento)
			//code += "P=P+1;\n"
			code = generator.Operacion{Destino: "P", Op1: "P", Signo: "+", Op2: "1", Comentario: false, Comentario_Cont: ""}
			fragmento = generator.Fragment{Valor: code, Tipo: 0, Valido: true, Mod: false, Line: p.Line}
			gen.AddFragment(name, fragmento)
			//code += "//--------------FIN DE DECLARACION--------" + p.Id + ""
			code = generator.Comentario{Comentario: "--------------FIN DE DECLARACION--------" + p.Id}
			fragmento = generator.Fragment{Valor: code, Tipo: 5, Valido: true, Mod: false, Line: p.Line}
			gen.AddFragment(name, fragmento)
			//gen.AddCodes(code, ambito)
		} else if interfaces.STRUCT == result.Type {
			tipo := interfaces.TipoSimbolo{result.Type, result.Tipo2}
			simbolo := interfaces.Symbol{Id: p.Id, Tipo: tipo, Posicion: result.Value, Mutable: p.Mutable}
			env.(environment.Environment).SaveVariable(p.Line, p.Col, p.Id, simbolo, p.Tipo)
		} else {

			//codigo := "//--------------INICIO DE DECLARACION--------" + p.Id + "\n"
			code = generator.Comentario{Comentario: "--------------INICIO DE DECLARACION--------" + p.Id}
			fragmento = generator.Fragment{Valor: code, Tipo: 5, Valido: true, Mod: false, Line: p.Line}
			gen.AddFragment(name, fragmento)
			tam := gen.NewTemp()

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
				//gen.AddCodes(tam+"=P;", ambito)
				code = generator.Asignacion{Destino: tam, Op: "P", Comentario: false, Comentario_Cont: ""}
				fragmento = generator.Fragment{Valor: code, Tipo: 1, Valido: true, Mod: true, Line: p.Line}
				gen.AddFragment(name, fragmento)
				//codigo += "STACK[(int)P]=" + fmt.Sprintf("%v", result.Value) + ";\n"
				code = generator.Stack{Index: "P", Value: fmt.Sprintf("%v", result.Value), Comentario: false, Comentario_Cont: ""}
				fragmento = generator.Fragment{Valor: code, Tipo: 4, Valido: true, Mod: true, Line: p.Line}
				gen.AddFragment(name, fragmento)
				//codigo += "P=P+1;\n"
				code = generator.Operacion{Destino: "P", Op1: "P", Signo: "+", Op2: "1", Comentario: false, Comentario_Cont: ""}
				fragmento = generator.Fragment{Valor: code, Tipo: 0, Valido: true, Mod: false, Line: p.Line}
				gen.AddFragment(name, fragmento)
				//codigo += "//------FIN DE DECLARACION--------" + p.Id
				code = generator.Comentario{Comentario: "------FIN DE DECLARACION--------" + p.Id}
				fragmento = generator.Fragment{Valor: code, Tipo: 5, Valido: true, Mod: false, Line: p.Line}
				gen.AddFragment(name, fragmento)
				env.(environment.Environment).SaveVariable(p.Line, p.Col, p.Id, simbolo, p.Tipo)
				/*if ambito {
					gen.AddFunc(codigo)
				} else {
					gen.AddCode(codigo)
				}*/
			} else if result.Type == interfaces.BOOLEAN {
				//value := ""
				l1 := result.TrueLabel
				l2 := result.FalseLabel
				l3 := gen.NewLabel()
				//value += l1 + ":\n"
				code = generator.Label{Name: l1, Comentario: true, Comentario_Cont: "-----VERDADERO"}
				fragmento = generator.Fragment{Valor: code, Tipo: 8, Valido: true, Mod: false, Line: p.Line}
				gen.AddFragment(name, fragmento)
				//value += "STACK[(int)" + tam + "]=1;\n" +
				code = generator.Stack{Index: tam, Value: "1", Comentario: false, Comentario_Cont: ""}
				fragmento = generator.Fragment{Valor: code, Tipo: 4, Valido: true, Mod: true, Line: p.Line}
				gen.AddFragment(name, fragmento)
				//"P=P+1;\n"
				code = generator.Operacion{Destino: "P", Op1: "P", Signo: "+", Op2: "1", Comentario: false, Comentario_Cont: ""}
				fragmento = generator.Fragment{Valor: code, Tipo: 0, Valido: true, Mod: false, Line: p.Line}
				gen.AddFragment(name, fragmento)
				//value += "goto " + l3 + ";\n"
				code = generator.Salto{Label: l3, Comentario: false, Comentario_Cont: "//----FALSO"}
				fragmento = generator.Fragment{Valor: code, Tipo: 9, Valido: true, Mod: false, Line: p.Line}
				gen.AddFragment(name, fragmento)
				//value += l2 + ":\n"
				code = generator.Label{Name: l2, Comentario: false, Comentario_Cont: "-----SALIDA"}
				fragmento = generator.Fragment{Valor: code, Tipo: 8, Valido: true, Mod: false, Line: p.Line}
				gen.AddFragment(name, fragmento)
				//value += "STACK[(int)" + tam + "]=0;\n" +
				code = generator.Stack{Index: tam, Value: "0", Comentario: false, Comentario_Cont: ""}
				fragmento = generator.Fragment{Valor: code, Tipo: 4, Valido: true, Mod: true, Line: p.Line}
				gen.AddFragment(name, fragmento)
				//"P=P+1;\n"
				code = generator.Operacion{Destino: "P", Op1: "P", Signo: "+", Op2: "1", Comentario: false, Comentario_Cont: ""}
				fragmento = generator.Fragment{Valor: code, Tipo: 0, Valido: true, Mod: false, Line: p.Line}
				gen.AddFragment(name, fragmento)
				//value += l3 + ":\n"
				code = generator.Label{Name: l3, Comentario: false, Comentario_Cont: "-----SALIDA"}
				fragmento = generator.Fragment{Valor: code, Tipo: 8, Valido: true, Mod: false, Line: p.Line}
				gen.AddFragment(name, fragmento)
				//value += "P=P+1;"
				//gen.AddCodes(value+"//FIN DE DECLARACION", ambito)
				code = generator.Comentario{Comentario: "FIN DE DECLARACION " + p.Id}
				fragmento = generator.Fragment{Valor: code, Tipo: 5, Valido: true, Mod: false, Line: p.Line}
				gen.AddFragment(name, fragmento)
				//gen.SetConf()
				env.(environment.Environment).SaveVariable(p.Line, p.Col, p.Id, simbolo, p.Tipo)
			}
		}
	} else {
		env.(environment.Environment).NewError("LOS TIPOS NO CONCUERDAN", p.Line, p.Col)
		//fmt.Println(err)
	}
	result.Type = interfaces.NULL
	return result
}

func CompararTipos(TipoIzq *arraylist.List, TipoDer *arraylist.List) bool {
	var conf bool
	fmt.Println("COMPARARTIPOS")
	//fmt.Println(len(TipoIzq.GetValue()))
	tipos := TipoIzq.GetValue(0).(interfaces.Dimensions)
	for _, s := range tipos.Dimensions.ToArray() {
		fmt.Println(s)
	}
	return conf
}

/*func declaracionStack(env interface{},gen *generator.Generator){
	if !env.(environment.Environment).Control.Ciclo {
		if env.(environment.Environment).Control.Id == "main" || env.(environment.Environment).Control.Id == "GLOBAL" {
			guia = "(int)P"
			incremento = "P=P+1;"

			tam = gen.Stack
			gen.Stack++
		} else {
			tmp := gen.NewTemp()
			tam = gen.Stack + tam
			gen.AddExpression(tmp, "(int)P", strconv.Itoa(tam), "+", ambito)
			guia = "(int)" + tmp
			//ambito = true
			env.(environment.Environment).Control.Stack++
		}
	} else {
		tmp := gen.NewTemp()
		tam = gen.Stack + tam
		gen.AddExpression(tmp, strconv.Itoa(gen.Stack), strconv.Itoa(tam), "+", ambito)
		guia = "(int)" + tmp
		gen.Stack++
		//ambito = true
		env.(environment.Environment).Control.Stack++
	}
}*/
