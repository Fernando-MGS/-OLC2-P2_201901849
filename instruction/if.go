package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"

	arrayList "github.com/colegno/arraylist"
)

type If struct {
	Condicion interfaces.Expresion
	Cuerpo    *arrayList.List
	Else      interfaces.Instruction
	line      string
	col       string
	Tipo      int //para verificar si se trata del if principal o de un else
}

func NewIf(condition interfaces.Expresion, bloque *arrayList.List, Else interfaces.Instruction, line, col, tipo int) If {
	ifInstr := If{condition, bloque, Else, strconv.Itoa(line), strconv.Itoa(col), tipo}
	return ifInstr
}

func (p If) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var ret interfaces.Value
	ret.Type = interfaces.NULL

	condicion := p.Condicion.Ejecutar(env, gen)
	ambito := env.(environment.Environment).DevAmbito()
	gen.AddCodes("//INICIO DE IF", ambito)
	if condicion.Type == interfaces.BOOLEAN {
		salida := ""
		l2 := gen.GetTempsB().FalseL
		if p.Tipo == 0 {
			salida = gen.NewLabel()
			gen.SetSalida(salida)
			fmt.Println("la salida se set" + gen.GetSalida() + "-" + salida)
		} else {
			salida = gen.GetSalida()
		}
		l1 := gen.GetTempsB().TrueL

		gen.SetConf()
		value := l1 + ":\n" // si es verdadera irá de nuevo a entrada para repetir el proceso
		gen.AddCodes(value, ambito)
		in := env.(environment.Environment).Control.Entrada
		out := env.(environment.Environment).Control.Salida
		loop := env.(environment.Environment).Control.Ciclo
		tmpEnv := environment.NewEnvironment(env.(environment.Environment), env.(environment.Environment).Control.Id, in, out, loop)
		for _, s := range p.Cuerpo.ToArray() {
			s.(interfaces.Instruction).Ejecutar(tmpEnv, gen)
		}
		value = "goto " + salida + ";\n//FIN DE IF\n"
		value += l2 + ":"
		gen.AddCodes(value, ambito)
		p.Else.Ejecutar(env, gen)
	} else {
		env.(environment.Environment).NewError("SE ESPERABA UNA EXPRESION BOOLEANA", p.line, p.col)
	}
	if p.Tipo == 0 {
		fmt.Println("LA salida es " + gen.GetSalida())
		gen.AddCodes("//SALIDA DEL IF", ambito)
		gen.AddCodes(gen.GetSalida()+":", ambito)
	}
	return ret
}
