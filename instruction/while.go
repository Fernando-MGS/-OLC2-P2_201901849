package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"

	/*"fmt"
	"os/exec"
	"reflect"*/
	"strconv"
	//"strings"

	//"time"

	arrayList "github.com/colegno/arraylist"
)

type While struct {
	Expresion interfaces.Expresion
	Bloque    *arrayList.List
	Line      string
	Col       string
}

func NewWhile(condition interfaces.Expresion, bloque *arrayList.List, line, col int) While {
	whileInstr := While{condition, bloque, strconv.Itoa(line), strconv.Itoa(col)}
	return whileInstr
}

func (p While) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var result interfaces.Value
	result.Type = interfaces.NULL
	entrada := gen.NewLabel()
	ambito := env.(environment.Environment).DevAmbito()
	gen.AddCodes("//INICIO DE WHILE", ambito)
	gen.AddCodes(entrada+":", ambito)
	//salida:=gen.NewLabel()
	condition := p.Expresion.Ejecutar(env, gen)
	if condition.Type == interfaces.BOOLEAN {
		l1 := gen.GetTempsB().TrueL
		l2 := gen.GetTempsB().FalseL
		gen.SetConf()
		value := l1 + ":\n" // si es verdadera irá de nuevo a entrada para repetir el proceso
		gen.AddCodes(value, ambito)
		tmpEnv := environment.NewEnvironment(env.(environment.Environment), env.(environment.Environment).Control.Id, entrada, l2, true)
		for _, s := range p.Bloque.ToArray() {
			s.(interfaces.Instruction).Ejecutar(tmpEnv, gen)
		}
		value = "goto " + entrada + ";\n"
		value += l2 + ":"
		gen.AddCodes(value, ambito)

	} else {
		env.(environment.Environment).NewError("SE ESPERABA UNA EXPRESION BOOLEANA", p.Line, p.Col)
	}
	//gen.AddCodes("goto "+salida+";",ambito)
	return result
}
