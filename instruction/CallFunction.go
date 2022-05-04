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
	true_l := ""
	false_l := ""
	if conf {
		list := arraylist.New()
		gen.NewComentario("EJECUTANDO PARAMETROS", name, true, false, "")
		for _, s := range p.Params.ToArray() {
			val := s.(interfaces.Expresion).Ejecutar(env, gen)
			list.Add(val)
		}
		valor_return := ""
		conf = confParametros(list, funct.Params)
		fmt.Println("PASO POR AQUI")
		if conf {
			saves := env.(environment.Environment).DevVariables()
			gen.NewComentario("INICIANDO EL RESGUARDO DE LAS POSICIONES", name, true, false, "")
			i := 1
			for _, s := range saves.ToArray() {
				sym := s.(interfaces.Symbol)
				pos := gen.NewTemp()
				size := gen.NewTemp()
				//j := strconv.Itoa(i)
				//CALCULO DE LA POSICION EN EL STACK DONDE SE DEBE GUARDAR
				gen.NewCallStack(size, "81999", false, "", name, true, false, "")
				gen.NewOperacion(size, size, "+", "1", false, "", name, true, false, "")
				gen.NewOperacion(pos, "81999", "-", size, false, "", name, true, false, "")
				pos_symbol := sym.Posicion
				value_symbol := gen.NewTemp()
				//OBTENCION DEL VALOR DEL SIMBOLO
				gen.NewCallStack(value_symbol, pos_symbol, false, "", name, true, false, "")
				//AUMENTAR EL TAMAÑO DE LA PILA
				gen.NewStack("81999", size, false, "", name, true, false, "")
				//GUARDAR EL VALOR EN LA PILA
				gen.NewStack(pos, value_symbol, false, "", name, true, false, "")
				i++
			}
			gen.NewComentario("FIN DEL RESGUARDO DE LAS POSICIONES", name, true, false, "")
			p_save := gen.NewTemp()
			gen.NewComentario("INICIO DE TRASLADO DE PARAMETROS", name, true, false, "")
			i = 1
			gen.NewAsignacion(p_save, "P", true, "GUARDADO DEL VALOR DE P", name, true, false, "")
			for _, s := range list.ToArray() {
				param := gen.NewTemp()
				gen.NewOperacion(param, "P", "+", strconv.Itoa(i), false, "", name, true, false, "")
				gen.NewStack(param, s.(interfaces.Value).Value, false, "", name, true, false, "")
				i++
			}
			fmt.Println("name es " + name)

			gen.NewLlamada(p.Id, false, "", name, true, false, "")
			i = 1
			gen.NewComentario("INICIO DE LA RECUPERACION DE LAS POSICIONES", name, true, false, "")
			for _, s := range saves.ToArray() {
				sym := s.(interfaces.Symbol)
				size_stack := gen.NewTemp()
				index_stack := gen.NewTemp()
				tmp_value := gen.NewTemp()
				gen.NewCallStack(size_stack, "81999", false, "", name, true, false, "")
				gen.NewOperacion(index_stack, "81999", "-", size_stack, false, "", name, true, false, "")
				gen.NewCallStack(tmp_value, index_stack, false, "", name, true, false, "")
				gen.NewStack(sym.Posicion, tmp_value, false, "", name, true, false, "")
				gen.NewOperacion(size_stack, size_stack, "-", "1", false, "", name, true, false, "")
				gen.NewStack("81999", size_stack, false, "", name, true, false, "")
			}
			gen.NewComentario("FIN DE LA RECUPERACION DE LAS POSICIONES", name, true, false, "")
			if funct.Tipo.Tipo != interfaces.NULL {
				valor_return = gen.NewTemp()
				gen.NewCallStack(valor_return, "P", true, "RETORNO DE "+p.Id, name, true, false, "")
				gen.NewAsignacion("P", p_save, true, "RECUPERACION DEL VALOR DE P", name, true, false, "")
				if funct.Tipo.Tipo == interfaces.BOOLEAN {
					true_l = gen.NewLabel()
					false_l = gen.NewLabel()
					gen.NewIf(valor_return, "==", "1", true_l, false, "", name, true, true, "")
					gen.NewSalto(false_l, false, "", name, true, true, "")
					retorno.TrueLabel = true_l
					retorno.FalseLabel = false_l
				}
			}

			//fmt.Println(funct.Tipo.Tipo)
		} else {
			env.(environment.Environment).NewError("LOS PARAMETROS NO CONCUERDAN EN SUS TIPOS", p.Line, p.Col)
		}
		//gen.NewLlamada(p.Id, false, "", name, true, false, "")
		retorno.Type = funct.Tipo.Tipo
		retorno.Tipo2 = funct.Tipo.Tipo2
		retorno.Value = valor_return
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
