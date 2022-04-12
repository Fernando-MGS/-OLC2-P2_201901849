package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"
	"time"
	//"fmt"
)

type ToString struct {
	expresion interfaces.Expresion
	line      string
	col       string
}

func NewToString(expr interfaces.Expresion, line, col int) ToString {
	exp := ToString{expr, strconv.Itoa(line), strconv.Itoa(col)}
	return exp
}

func (t ToString) Ejecutar(env interface{}, g *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	retorno = t.expresion.Ejecutar(env, g)
	if retorno.Type == interfaces.STRING {
		retorno.Type = interfaces.STR
		return retorno
	}

	tm := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		tm.Year(), tm.Month(), tm.Day(),
		tm.Hour(), tm.Minute(), tm.Second())
	err := interfaces.Errores{Line: t.line, Col: t.col, Mess: "SE ESPERABA UNA EXPRESION &str", Fecha: fecha}
	env.(environment.Environment).Errores(err)
	retorno.IsTemp = false
	retorno.Type = interfaces.NULL
	return retorno
}
