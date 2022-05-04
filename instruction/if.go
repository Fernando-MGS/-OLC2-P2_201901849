package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"

	//"fmt"

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
	/*if gen.GetConf() == 1 {
		gen.SetConf()
	}*/
	//fmt.Println("HOLA")
	//ret.Value = retorno
	name := env.(environment.Environment).Control.Id
	condicion := p.Condicion.Ejecutar(env, gen)
	//ambito := env.(environment.Environment).DevAmbito()
	conf_condicion := false
	if condicion.Type == interfaces.BOOLEAN {
		salida := ""
		conf_condicion = true
		l2 := condicion.FalseLabel
		if p.Tipo == 0 {
			salida = gen.NewLabel()
			gen.SetSalida(salida)
			//gen.AddCodes("//INICIO DE IF PRIMARIO", ambito)
			gen.NewComentario("INICIO DE IF PRIMARIO", name, true, false, p.line)
			//retorno = gen.NewTemp()
			//fmt.Println("if primario")
			//fmt.Println("la salida se set" + gen.GetSalida() + "-" + salida)
		} else {
			//fmt.Println("else if o else")
			if p.Tipo == 2 {
				//gen.AddCodes("//INICIO DE IF ELSE", ambito)
				gen.NewComentario("INICIO DE IF ELSE", name, true, false, p.line)
			} else {
				//gen.AddCodes("//INICIO DE ELSE", ambito)
				gen.NewComentario("INICIO DE IF ELSE", name, true, false, p.line)
			}
			//retorno = env.(environment.Environment).Control.Retorno
			salida = gen.GetSalida()
		}
		//fmt.Println("EL RETORNO NAME ES " + retorno)
		l1 := condicion.TrueLabel
		//fmt.Println("El conf es " + strconv.Itoa(gen.GetConf()) + l1)
		//gen.SetConf()
		//fmt.Println("El conf luego es " + strconv.Itoa(gen.GetConf()) + l1)
		//value := l1 + ":\n" // si es verdadera irá de nuevo a entrada para repetir el proceso
		gen.NewLabels(l1, false, "", name, true, true, p.line)
		//gen.AddCodes(value, ambito)
		in := env.(environment.Environment).Control.Entrada
		out := env.(environment.Environment).Control.Salida
		loop := env.(environment.Environment).Control.Ciclo
		stack := env.(environment.Environment).Control.Stack
		tmpEnv := environment.NewEnvironment(env.(environment.Environment), env.(environment.Environment).Control.Id, in, out, loop, stack, "")
		for _, s := range p.Cuerpo.ToArray() {
			i := s.(interfaces.OpcionIf)
			//fmt.Println(i)
			var res interface{}
			//fmt.Println("aver si es aqui")
			if i.Tipo == 0 {
				res = s.(interfaces.OpcionIf).Ejecucion.(interfaces.Expresion).Ejecutar(tmpEnv, gen)
			} else {
				res = s.(interfaces.OpcionIf).Ejecucion.(interfaces.Instruction).Ejecutar(tmpEnv, gen)
			}
			//fmt.Println("simon")
			//fmt.Println(res)
			if res.(interfaces.Value).Type != interfaces.NULL {
				//fmt.Println("RES IF")
				//fmt.Println(res)
				retorno := ""
				if p.Tipo == 0 {
					retorno = gen.NewTemp()
					gen.SetRetorno(retorno)
				} else {
					retorno = gen.GetRetorno()
				}
				ret = res.(interfaces.Value)
				//code := retorno + "=" + ret.Value + ";"
				gen.NewAsignacion(retorno, ret.Value, false, "", name, true, true, p.line)
				ret.Value = retorno
				//gen.AddCodes(code, ambito)
			}
		}
		//value = "goto " + salida + ";\n//FIN DE IF\n"
		gen.NewSalto(salida, false, "", name, true, true, p.line)
		//value += l2 + ":"
		gen.NewLabels(l2, false, "", name, true, true, "")
		//gen.AddCodes(value, ambito)
		res := p.Else.Ejecutar(env, gen)

		if res.(interfaces.Value).Type != interfaces.NULL {
			//fmt.Println("RES ELSE")
			//fmt.Println(res)
			ret = res.(interfaces.Value)
		}
	} else {
		env.(environment.Environment).NewError("SE ESPERABA UNA EXPRESION BOOLEANA", p.line, p.col)
		ret.Type = interfaces.NULL
		return ret
	}
	if p.Tipo == 0 && conf_condicion {
		//fmt.Println("LA salida es " + gen.GetSalida())
		//gen.AddCodes("//SALIDA DEL IF", ambito)
		gen.NewComentario("SALIDA DE IF", name, true, false, p.line)
		//gen.AddCodes(gen.GetSalida()+":", ambito)
		gen.NewLabels(gen.GetSalida(), false, "", name, true, true, "")
	}
	/*if gen.GetConf() == 1 {
		gen.SetConf()
	}*/
	if p.Tipo == 0 {
		gen.DelSalida()
		gen.DelRetorno()
	}
	return ret
}
