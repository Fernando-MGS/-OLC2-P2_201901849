package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"
	"time"
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
	if result.Type == interfaces.NULL {
		t := time.Now()
		fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LA EXPRESION NO ES VALIDA", Fecha: fecha}
		env.(environment.Environment).Errores(err)
	} else if result.Type == interfaces.VOID {
		t := time.Now()
		fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LAS FUNCIONES DE TIPO VOID NO DEVUELVEN NINGUN VALOR", Fecha: fecha}
		env.(environment.Environment).Errores(err)
	} else {
		variable := env.(environment.Environment).GetVariable(p.Id, p.Line, p.Col)
		if variable.Tipo.Tipo == result.Type {
			ambito := true
			if env.(environment.Environment).Control.Id == "GLOBAL" || env.(environment.Environment).Control.Id == "main" {
				ambito = false
			}
			if result.Type == interfaces.STRING {

			} else if result.Type == interfaces.STR {

			} else if result.Type == interfaces.ARRAY {

			} else if result.Type == interfaces.VECTOR {

			} else if result.Type == interfaces.STRUCT {

			} else {
				value := "//---------INICIANDO ASIGNACION-------"
				value += "STACK[" + strconv.Itoa(variable.Posicion) + "]"
				value += "=" + result.Value
				value += "//---------FIN DE ASIGNACION-------"
				if ambito {
					gen.AddFunc(value)
				} else {
					gen.AddCode(value)
				}
			}
		} else {
			t := time.Now()
			fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())
			err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO COINCIDEN", Fecha: fecha}
			env.(environment.Environment).Errores(err)
		}
	}
	result.Type = interfaces.NULL
	return result.Value
}
