package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"

	arrayList "github.com/colegno/arraylist"
)

type For struct {
	Cuerpo    *arrayList.List
	Variable  string
	Iteracion interfaces.For_Range
	line      string
	col       string
}

func NewFor(Cuerpo *arrayList.List, Variable string, Iteracion interfaces.For_Range, line, col int) For {
	instr := For{Cuerpo, Variable, Iteracion, strconv.Itoa(line), strconv.Itoa(col)}
	return instr
}

func (f For) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var result interfaces.Value
	result.Type = interfaces.NULL
	rangoInf := f.Iteracion.Range1.Ejecutar(env, gen)
	rangoSup := f.Iteracion.Range2.Ejecutar(env, gen)
	ambito := env.(environment.Environment).DevAmbito()
	stack := env.(environment.Environment).Control.Stack
	tmpEnv := environment.NewEnvironment(env, env.(environment.Environment).Control.Id, "", "", true, stack)
	if f.Iteracion.Tipo == 0 {
		if rangoInf.Type != interfaces.INTEGER || rangoSup.Type != interfaces.INTEGER && rangoInf.Type != interfaces.USIZE || rangoSup.Type != interfaces.USIZE {
			env.(environment.Environment).NewError("SE ESPERABA UNA EXPRESION NUMERICA", f.line, f.col)
			return result
		}
		tipos := interfaces.TipoSimbolo{Tipo: interfaces.INTEGER, Tipo2: arrayList.New()}
		entrada := gen.NewLabel()
		ciclo := gen.NewLabel()
		salida := gen.NewLabel()
		contador := gen.NewTemp()
		pre_contador := gen.NewTemp()
		p := gen.Stack
		gen.AddCodes("//INICIO DE FOR-----", ambito)
		env_stack := env.(environment.Environment).Control.Stack
		env.(environment.Environment).Control.Stack++
		tam := p + env_stack
		variable := interfaces.Symbol{Id: f.Variable, Tipo: tipos, Posicion: tam, Mutable: true, Line: f.line, Col: f.col}
		tmpEnv.SaveVariable(f.line, f.col, f.Variable, variable, tipos)
		gen.AddCodes("STACK["+strconv.Itoa(tam)+"]="+rangoInf.Value+";", ambito)
		gen.AddCodes(entrada+":", ambito)
		llamada := gen.NewTemp()
		gen.AddCodes(llamada+"=STACK["+strconv.Itoa(tam)+"];", ambito)
		code := "if(" + llamada + "<" + rangoSup.Value + ") goto " + ciclo + ";"
		gen.AddCodes(code, ambito)
		gen.AddCodes("goto "+salida+";", ambito)
		gen.AddCodes(ciclo+":", ambito)
		for _, s := range f.Cuerpo.ToArray() {
			s.(interfaces.Instruction).Ejecutar(tmpEnv, gen)
		}
		contar := pre_contador + "=STACK[" + strconv.Itoa(tam) + "];"
		incrementar := contador + "=" + pre_contador + "+1;"
		modificar := "STACK[" + strconv.Itoa(tam) + "]=" + contador + ";"
		gen.AddCodes(contar, ambito)
		gen.AddCodes(incrementar, ambito)
		gen.AddCodes(modificar, ambito)
		gen.AddCodes("goto "+entrada+";", ambito)
		gen.AddCodes(salida+":", ambito)
		gen.AddCodes("//FIN DE FOR-----", ambito)
	} else {
		fmt.Println("hoal")
	}
	return result
}
