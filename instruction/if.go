package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"

	//"fmt"

	//"fmt"
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
	//retorno := ""
	if gen.GetConf() == 1 {
		gen.SetConf()
	}
	fmt.Println("HOLA")
	//ret.Value = retorno
	condicion := p.Condicion.Ejecutar(env, gen)
	ambito := env.(environment.Environment).DevAmbito()

	if condicion.Type == interfaces.BOOLEAN {
		salida := ""

		l2 := gen.GetTempsB().FalseL
		if p.Tipo == 0 {
			salida = gen.NewLabel()
			gen.SetSalida(salida)
			gen.AddCodes("//INICIO DE IF PRIMARIO", ambito)
			//retorno = gen.NewTemp()
			//fmt.Println("if primario")
			//fmt.Println("la salida se set" + gen.GetSalida() + "-" + salida)
		} else {
			//fmt.Println("else if o else")
			if p.Tipo == 2 {
				gen.AddCodes("//INICIO DE IF ELSE", ambito)
			} else {
				gen.AddCodes("//INICIO DE ELSE", ambito)
			}
			//retorno = env.(environment.Environment).Control.Retorno
			salida = gen.GetSalida()
		}
		//fmt.Println("EL RETORNO NAME ES " + retorno)
		l1 := gen.GetTempsB().TrueL
		//fmt.Println("El conf es " + strconv.Itoa(gen.GetConf()) + l1)
		gen.SetConf()
		//fmt.Println("El conf luego es " + strconv.Itoa(gen.GetConf()) + l1)
		value := l1 + ":\n" // si es verdadera irá de nuevo a entrada para repetir el proceso
		gen.AddCodes(value, ambito)
		in := env.(environment.Environment).Control.Entrada
		out := env.(environment.Environment).Control.Salida
		loop := env.(environment.Environment).Control.Ciclo
		stack := env.(environment.Environment).Control.Stack
		tmpEnv := environment.NewEnvironment(env.(environment.Environment), env.(environment.Environment).Control.Id, in, out, loop, stack)
		for _, s := range p.Cuerpo.ToArray() {
			i := s.(interfaces.OpcionIf)
			fmt.Println(i)
			var res interface{}
			fmt.Println("aver si es aqui")
			if i.Tipo == 0 {
				res = s.(interfaces.OpcionIf).Ejecucion.(interfaces.Expresion).Ejecutar(tmpEnv, gen)
			} else {
				res = s.(interfaces.OpcionIf).Ejecucion.(interfaces.Instruction).Ejecutar(tmpEnv, gen)
			}
			fmt.Println("simon")
			fmt.Println(res)
			if res.(interfaces.Value).Type != interfaces.NULL {
				fmt.Println("RES IF")
				fmt.Println(res)
				retorno := ""
				if p.Tipo == 0 {
					retorno = gen.NewTemp()
					gen.SetRetorno(retorno)
				} else {
					retorno = gen.GetRetorno()
				}
				ret = res.(interfaces.Value)
				code := retorno + "=" + ret.Value + ";"
				ret.Value = retorno
				gen.AddCodes(code, ambito)
			}
		}
		value = "goto " + salida + ";\n//FIN DE IF\n"
		value += l2 + ":"
		gen.AddCodes(value, ambito)
		res := p.Else.Ejecutar(env, gen)

		if res.(interfaces.Value).Type != interfaces.NULL {
			fmt.Println("RES ELSE")
			fmt.Println(res)
			ret = res.(interfaces.Value)
		}
	} else {
		env.(environment.Environment).NewError("SE ESPERABA UNA EXPRESION BOOLEANA", p.line, p.col)
	}
	if p.Tipo == 0 {
		fmt.Println("LA salida es " + gen.GetSalida())
		gen.AddCodes("//SALIDA DEL IF", ambito)
		gen.AddCodes(gen.GetSalida()+":", ambito)
	}
	if gen.GetConf() == 1 {
		gen.SetConf()
	}
	if p.Tipo == 0 {
		gen.DelSalida()
		gen.DelRetorno()
	}
	return ret
}
