package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"reflect"
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
	result := p.Expresion.Ejecutar(env, gen)
	conf := false
	ambito := env.(environment.Environment).DevAmbito()
	//index := 0
	tipo := reflect.TypeOf(p.Expresion)
	t := fmt.Sprintf("%v", tipo)
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
			gen.AddCodes("//INICIO DE DECLARACION "+p.Id, ambito)
			fmt.Println(p.Tipo.Tipo2.ToArray()...)
			tipoArr := p.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions).Tipo
			fmt.Println("simon")
			if tipoArr != result.Tipo2.GetValue(0).(interfaces.Dimensions).Tipo {
				fmt.Println("NEL1")
				conf = false
			}
			if len(p.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions).Dimensions.ToArray()) != len(result.Tipo2.GetValue(0).(interfaces.Dimensions).Dimensions.ToArray()) {
				fmt.Println("NEL2")
				conf = false
				env.(environment.Environment).NewError("LAS DIMENSIONES DEL ARREGLO NO CONCUERDAN EN LA DECLARACION", p.Line, p.Col)
				result.Type = interfaces.NULL
				return result
			}
			fmt.Println(conf)
			if conf {
				/*size := arraylist.New()
				for _, s := range p.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions).Dimensions.ToArray() {
					tam := s.(interfaces.Expresion).Ejecutar(env, gen)
					size.Add(tam.Value)
				}
				size.Add(p.Tipo.Tipo)*/
				//fmt.Println(size.ToArray()...)
				posicion2 := ""
				if t == "expresion.CallVariable" {
					entrada := gen.NewLabel()
					salida := gen.NewLabel()
					//gen.AddCodes("if("+bounds.TrueLabel+") ",ambito)
					iterador := gen.NewTemp()
					contador := gen.NewTemp()
					contador2 := gen.NewTemp()
					aux := gen.NewTemp()
					gen.AddCodes(iterador+"=0;//iterador", ambito)
					gen.AddCodes(contador+"=H;//contador1", ambito)
					gen.AddCodes(contador2+"="+result.Value+";//contador2", ambito)
					gen.AddCodes(entrada+":", ambito)
					gen.AddCodes("if("+iterador+"=="+result.TrueLabel+") goto "+salida+";", ambito)
					gen.AddCodes(aux+"=HEAP[(int)"+contador2+"];", ambito)
					gen.AddCodes("HEAP[(int)H]="+aux+";", ambito)
					gen.AddCodes("H=H+1;", ambito)
					//gen.AddCodes(contador+"="+contador+"+1;", ambito)
					gen.AddCodes(contador2+"="+contador2+"+1;", ambito)
					gen.AddCodes(iterador+"="+iterador+"+1;", ambito)
					gen.AddCodes("goto "+entrada+";", ambito)
					gen.AddCodes(salida+":", ambito)
					posicion2 = contador
				} else {
					posicion2 = result.Value
				}
				tipoSimbolo := interfaces.TipoSimbolo{Tipo: interfaces.ARRAY, Tipo2: result.Tipo2}
				variable := interfaces.Symbol{Id: p.Id, Posicion2: posicion2, Mutable: p.Mutable, Line: p.Line, Col: p.Col, Tipo: tipoSimbolo, Longitud: result.TrueLabel}
				env.(environment.Environment).SaveVariable(p.Line, p.Col, p.Id, variable, variable.Tipo)
				fmt.Println("se guardo")
			} else {
				env.(environment.Environment).NewError("LOS TIPOS NO CONCUERDAN EN LA DECLARACION", p.Line, p.Col)
				result.Type = interfaces.NULL
				return result
			}
			/*fmt.Println(p.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions).Tipo)
			fmt.Print(len(p.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions).Dimensions.ToArray()))
			fmt.Print("vs")
			fmt.Println(len(result.Tipo2.ToArray()))
			/*fmt.Println(p.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions).Dimensions.Clone().ToArray()...)*/
			//CompararTipos(p.Tipo.Tipo2, result.Tipo2)
		} else if p.Tipo.Tipo == interfaces.VECTOR {
			fmt.Print("P.tipo   ")
			fmt.Println(result.Tipo2.ToArray()...)
			dimension_res := result.Tipo2.GetValue(0).(interfaces.Dimensions)
			fmt.Print("res.dimension   ")
			fmt.Println(dimension_res.Dimensions.ToArray()...)
			dimension_tipo := p.Tipo.Tipo2.GetValue(0).(interfaces.TipoExpresion)
			fmt.Println("hola 2")
			if dimension_res.Tipo == interfaces.VOID || dimension_res.Tipo == dimension_tipo {
				posicion2 := ""
				posicion2 = result.Value
				dimension := interfaces.Dimensions{Tipo: dimension_tipo, Dimensions: p.Tipo.Tipo2}
				dimensions_list := arraylist.New()
				dimensions_list.Add(dimension)
				tipoSimbolo := interfaces.TipoSimbolo{Tipo: interfaces.VECTOR, Tipo2: dimensions_list}
				variable := interfaces.Symbol{Id: p.Id, Posicion2: posicion2, Mutable: p.Mutable, Line: p.Line, Col: p.Col, Tipo: tipoSimbolo, Longitud: result.TrueLabel}
				env.(environment.Environment).SaveVariable(p.Line, p.Col, p.Id, variable, variable.Tipo)
				fmt.Println("se guardo EN " + posicion2)
			} else {
				env.(environment.Environment).NewError("LOS TIPOS NO CONCUERDAN EN LA DECLARACION", p.Line, p.Col)
				result.Type = interfaces.NULL
				return result
			}
		} else if p.Tipo.Tipo == interfaces.STR || p.Tipo.Tipo == interfaces.STRING {
			code := "//--------------INICIO DE DECLARACION--------" + p.Id + "\n"
			tam := env.(environment.Environment).Control.Stack
			guia := ""
			incremento := ""
			if !env.(environment.Environment).Control.Ciclo {
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
					//ambito = true
					env.(environment.Environment).Control.Stack++
				}
			} else {
				tmp := gen.NewTemp()
				tam = gen.Stack + tam
				gen.AddExpression(tmp, "(int)P", strconv.Itoa(tam), "+", true)
				guia = "(int)" + tmp
				//ambito = true
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
				fmt.Println("=========")
				fmt.Println(env.(environment.Environment).Control.Stack)
				fmt.Println(gen.Stack)
				tmp := gen.NewTemp()
				tam = gen.Stack + env.(environment.Environment).Control.Stack
				fmt.Println(tam)
				gen.AddExpression(tmp, strconv.Itoa(gen.Stack), strconv.Itoa(env.(environment.Environment).Control.Stack), "+", ambito)
				guia = "(int)" + tmp
				gen.Stack++
				fmt.Println("////////" + p.Id)
				fmt.Println(env.(environment.Environment).Control.Stack)
				fmt.Println(gen.Stack)
				//ambito = true
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
