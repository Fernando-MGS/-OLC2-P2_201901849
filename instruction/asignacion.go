package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
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
	ambito := env.(environment.Environment).DevAmbito()
	fmt.Println("TIPO 1")
	fmt.Println(result.Value)
	fmt.Println("tipo2")
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

			if result.Type == interfaces.VECTOR {

			} else if result.Type == interfaces.ARRAY {

			} else if result.Type == interfaces.STRUCT {

			} else if result.Type == interfaces.BOOLEAN {
				value := "//INICIO DE ASIGNACION"
				l1 := gen.GetTempsB().TrueL
				l2 := gen.GetTempsB().FalseL
				l3 := gen.NewLabel()
				value += l1 + ":\n"
				value += "STACK[" + strconv.Itoa(variable.Posicion) + "]=1;\n"
				value += "goto " + l3 + ";\n"
				value += l2 + ":\n"
				value += "STACK[" + strconv.Itoa(variable.Posicion) + "]=0;\n"
				value += l3 + ":\n"
				gen.AddCodes(value, ambito)
				gen.SetConf()
				//env.(environment.Environment).SaveVariable(p.Line, p.Col, p.Id, simbolo, p.Tipo)
			} else {
				if interfaces.CHAR == result.Type && result.IsTemp == false {
					runes := []rune(result.Value)
					var val string
					for i := 0; i < len(runes); i++ {
						val = strconv.Itoa(int(runes[i]))
					}
					result.Value = val
				}
				value := "//---------INICIANDO ASIGNACION-------\n"
				value += "STACK[" + strconv.Itoa(variable.Posicion) + "]"
				value += "=" + result.Value + ";\n"
				value += "//---------FIN DE ASIGNACION-------"
				gen.AddCodes(value, ambito)
			}
		} else {
			env.(environment.Environment).NewError("LOS TIPOS NO COINCIDEN", p.Line, p.Col)
		}
	}
	result.Type = interfaces.NULL
	return result
}
