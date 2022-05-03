package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"

	"github.com/colegno/arraylist"
)

type CallFunction struct {
	Id     string
	Params *arraylist.List
	Line   string
	Col    string
}

func NewCallF(id string, p *arraylist.List, line, col int) CallFunction {
	return CallFunction{id, p, strconv.Itoa(line), strconv.Itoa(col)}
}

func (p CallFunction) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var retorno interfaces.Value
	name := env.(environment.Environment).Control.Id
	conf, funct := env.(environment.Environment).GetFunc(p.Id, p.Line, p.Col)
	if conf {
		list := arraylist.New()
		gen.NewComentario("EJECUTANDO PARAMETROS", name, true, false, "")
		for _, s := range p.Params.ToArray() {
			val := s.(interfaces.Expresion).Ejecutar(env, gen)
			list.Add(val)
		}
		conf = confParametros(list, funct.Params)
		fmt.Println("PASO POR AQUI")
		if conf {
			p_save := gen.NewTemp()
			gen.NewComentario("INICIO DE TRASLADO DE PARAMETROS", name, true, false, "")
			i := 1
			gen.NewAsignacion(p_save, "P", true, "GUARDADO DEL VALOR DE P", name, true, false, "")
			for _, s := range list.ToArray() {
				param := gen.NewTemp()
				gen.NewOperacion(param, "P", "+", strconv.Itoa(i), false, "", name, true, false, "")
				gen.NewStack(param, s.(interfaces.Value).Value, false, "", name, true, false, "")
				i++
			}
			gen.NewLlamada(p.Id, false, "", name, true, false, "")
			fmt.Println(funct.Tipo.Tipo)
		} else {
			env.(environment.Environment).NewError("LOS PARAMETROS NO CONCUERDAN EN SUS TIPOS", p.Line, p.Col)
		}
		//gen.NewLlamada(p.Id, false, "", name, true, false, "")
		retorno.Type = funct.Tipo.Tipo
		retorno.Tipo2 = funct.Tipo.Tipo2
	}
	//gen.NewLlamada(p.Id,)
	return retorno
}

func confParametros(p1, p2 *arraylist.List) bool {
	conf := false
	if p1.Len() == p2.Len() {
		i := 0
		for _, s := range p2.ToArray() {
			val := s.(interfaces.Parametros)
			val2 := p1.GetValue(i).(interfaces.Value)
			if val.Tipo.Tipo != val2.Type {
				return false
			}
			i++
		}
		conf = true
	}

	return conf
}
