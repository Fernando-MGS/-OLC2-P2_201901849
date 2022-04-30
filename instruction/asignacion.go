package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"reflect"

	//"fmt"
	"strconv"
)

type Assignment struct {
	Id        string
	Expresion interfaces.Expresion
	Line      string
	Col       string
}

func NewAssignment(id string, val interfaces.Expresion, line, col int) Assignment {
	instr := Assignment{id, val, strconv.Itoa(line), strconv.Itoa(col)}
	return instr
}

func (p Assignment) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	result := p.Expresion.Ejecutar(env, gen)
	tipo := reflect.TypeOf(p.Expresion)
	t := fmt.Sprintf("%v", tipo)
	fmt.Println(t)
	fmt.Println(result)
	ambito := env.(environment.Environment).DevAmbito()
	/*fmt.Println("TIPO 1")
	fmt.Println(result.Value)
	fmt.Println("tipo2")*/
	if result.Type == interfaces.NULL {
		err := "LA EXPRESION NO ES VALIDA"
		env.(environment.Environment).NewError(err, p.Line, p.Col)
	} else {
		variable := env.(environment.Environment).GetVariable(p.Id, p.Line, p.Col)
		if !variable.Mutable {
			env.(environment.Environment).NewError("LA VARIABLE \""+variable.Id+"\" NO ES MUTABLE", p.Line, p.Col)
			result.Type = interfaces.NULL
			return result
		}
		if variable.Tipo.Tipo == result.Type || variable.Tipo.Tipo == interfaces.USIZE && result.Type == interfaces.INTEGER {
			gen.AddCodes("//INICIANDO ASIGNACION DE "+p.Id, ambito)
			if result.Type == interfaces.VECTOR {
				gen.AddCodes(variable.Posicion2+"="+result.Value+";", ambito)
			} else if result.Type == interfaces.ARRAY {
				//posicion2:=""

				if t == "expresion.CallVariable" {
					entrada := gen.NewLabel()
					salida := gen.NewLabel()
					//gen.AddCodes("if("+bounds.TrueLabel+") ",ambito)
					iterador := gen.NewTemp()
					contador := gen.NewTemp()
					contador2 := gen.NewTemp()
					aux := gen.NewTemp()
					gen.AddCodes(iterador+"=0;//iterador", ambito)
					gen.AddCodes(contador+"="+variable.Posicion2+";//contador1", ambito)
					gen.AddCodes(contador2+"="+result.Value+";//contador2", ambito)
					gen.AddCodes(entrada+":", ambito)
					gen.AddCodes("if("+iterador+"=="+variable.Longitud+") goto "+salida+";", ambito)
					gen.AddCodes(aux+"=HEAP[(int)"+contador2+"];", ambito)
					gen.AddCodes("HEAP[(int)"+contador+"]="+aux+";", ambito)
					gen.AddCodes(contador+"="+contador+"+1;", ambito)
					gen.AddCodes(contador2+"="+contador2+"+1;", ambito)
					gen.AddCodes(iterador+"="+iterador+"+1;", ambito)
					gen.AddCodes("goto "+entrada+";", ambito)
					gen.AddCodes(salida+":", ambito)
					//posicion2 = contador
				} else {
					//posicion2=result.Value
					gen.AddCodes(variable.Posicion2+"="+result.Value+";", ambito)
				}
				gen.AddCodes("//FINALIZANDO ASIGNACION DE"+p.Id, ambito)
				//env.(environment.Environment).
			} else if result.Type == interfaces.STRUCT {

			} else if result.Type == interfaces.BOOLEAN {
				value := "//INICIO DE ASIGNACION" + p.Id + "\n"
				l1 := result.TrueLabel
				l2 := result.FalseLabel
				l3 := gen.NewLabel()
				value += l1 + ":\n"
				value += "STACK[(int)" + variable.Posicion + "]=1;\n"
				value += "goto " + l3 + ";\n"
				value += l2 + ":\n"
				value += "STACK[(int)" + variable.Posicion + "]=0;\n"
				value += l3 + ":\n"
				gen.AddCodes(value, ambito)
				//gen.SetConf()*/
				//env.(environment.Environment).SaveVariable(p.Line, p.Col, p.Id, simbolo, p.Tipo)
			} else {
				if interfaces.CHAR == result.Type && !result.IsTemp {
					runes := []rune(result.Value)
					var val string
					for i := 0; i < len(runes); i++ {
						val = strconv.Itoa(int(runes[i]))
					}
					result.Value = val
				}
				value := "//---------INICIANDO ASIGNACION-------" + p.Id + "\n"
				value += "STACK[(int)" + variable.Posicion + "]"
				value += "=" + result.Value + ";\n"
				value += "//---------FIN DE ASIGNACION-------" + p.Id + ""
				gen.AddCodes(value, ambito)
			}
		} else {
			env.(environment.Environment).NewError("LOS TIPOS NO COINCIDEN", p.Line, p.Col)
		}
	}
	result.Type = interfaces.NULL
	return result
}
