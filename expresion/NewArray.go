package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"

	"github.com/colegno/arraylist"
)

type Array struct {
	ListExp *arraylist.List
}

func NewArray(list *arraylist.List) Array {
	exp := Array{list}
	return exp
}

func (p Array) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	retorno.Type = interfaces.ARRAY
	ambito := env.(environment.Environment).DevAmbito()
	valores := arraylist.New()
	conf := true
	gen.AddCodes("//CREANDO ARRAY CON LISTA DE VALORES", ambito)
	var tipo interfaces.TipoExpresion
	for _, s := range p.ListExp.ToArray() {
		res := s.(interfaces.Expresion).Ejecutar(env, gen)
		if res.Type == interfaces.ARRAY && conf {
			retorno.Tipo2 = res.Tipo2
			conf = false
		} else {
			tipo = res.Type
		}
		valores.Add(res)
		//fmt.Println(res)
	}
	largo := strconv.Itoa(p.ListExp.Len())
	long := gen.NewTemp()
	//pos := gen.NewTemp()
	gen.AddCodes(long+"=H;//POS DE INICIO NEWARRAY", ambito)
	gen.AddCodes("HEAP[(int)H]="+largo+";", ambito)
	gen.AddCodes("H=H+1;", ambito)
	for _, s := range valores.ToArray() {
		valors := s.(interfaces.Value)
		gen.AddCodes("HEAP[(int)H]="+valors.Value+";", ambito)
		gen.AddCodes("H=H+1;", ambito)
	}
	//gen.AddCodes(long+"="+largo+";", ambito)
	/*gen.AddCodes(pos+"=H;", ambito)
	gen.AddCodes("HEAP[(int)H]="+largo+";", ambito)
	gen.AddCodes("H=H+1;", ambito)*/
	if conf {
		dimension := arraylist.New()
		dimension.Add(tipo)
		tipo_d := interfaces.Dimensions{Tipo: tipo, Dimensions: dimension}
		retorno.Tipo2 = arraylist.New()
		retorno.Tipo2.Add(tipo_d)
	} else {
		tipo_d := retorno.Tipo2.GetValue(0).(interfaces.Dimensions)
		tipo_d.Dimensions.Add(interfaces.ARRAY)
		retorno.Tipo2 = arraylist.New()
		retorno.Tipo2.Add(tipo_d)
	}
	fmt.Println("========")
	fmt.Println(retorno.Tipo2.GetValue(0))
	retorno.TrueLabel = long
	retorno.Value = long
	fmt.Println(tipo)
	return retorno
}

/*
func WatchTipos(l *arrayList.List, env interface{}, gen *generator.Generator) *arrayList.List {
	lista := arrayList.New()
	//fmt.Println("======")

	var tipo_var interfaces.TipoExpresion
	for _, s := range l.ToArray() {
		tipo := reflect.TypeOf(s)
		t := fmt.Sprintf("%v", tipo)
		if t == "expresion.Array" {
			//fmt.Println("H=H1")
			lista = WatchTipos(s.(Array).ListExp, env, gen)
			lista.Add(interfaces.ARRAY)
			return lista

		} else {
			exp := s.(interfaces.Expresion).Ejecutar(env, gen)
			tipo_var = exp.Type
			//fmt.Println(exp.Value)
			//lista.Add(interfaces.INTEGER)
			//fmt.Println(t)
			//return lista
		}

	}
	lista.Add(tipo_var)
	return lista
}

func ExecExpresiones(l, values *arrayList.List, env interface{}, gen *generator.Generator) (interfaces.TipoExpresion, *arrayList.List) {
	//lista := arrayList.New()
	//fmt.Println("======")
	vals := values
	var tipo_var interfaces.TipoExpresion
	for _, s := range l.ToArray() {
		tipo := reflect.TypeOf(s)
		t := fmt.Sprintf("%v", tipo)
		if t == "expresion.Array" {
			//fmt.Println(t)
			tipo_var, vals = ExecExpresiones(s.(Array).ListExp, vals, env, gen)
			//lista.Add(interfaces.ARRAY)

		} else {
			exp := s.(interfaces.Expresion).Ejecutar(env, gen)
			tipo_var = exp.Type
			vals.Add(exp.Value)
			//fmt.Println(exp.Value)
			//lista.Add(interfaces.INTEGER)
			//fmt.Println(t)
			//return lista
		}

	}
	return tipo_var, vals
}

func FixLista(l *arrayList.List, tipo interfaces.TipoExpresion) *arrayList.List {
	i := 0
	newLista := arrayList.New()
	for _, s := range l.ToArray() {
		if i == 0 {
			newLista.Add(tipo)
		} else {
			newLista.Add(s)
		}
		i++
	}
	return newLista
}

func ContarLista(l *arrayList.List) *arrayList.List {
	//largo:=0
	lista := arrayList.New()
	fmt.Println("======")

	//var tipo_var interfaces.TipoExpresion
	for _, s := range l.ToArray() {
		tipo := reflect.TypeOf(s)
		t := fmt.Sprintf("%v", tipo)
		if t == "expresion.Array" {
			//fmt.Println("H=H1")
			lista = ContarLista(s.(Array).ListExp)
			lista.Add(strconv.Itoa(len(s.(Array).ListExp.ToArray())))
			//return lista

		}
	}
	return lista
}

func DevSize(lens *arrayList.List) string {
	x := fmt.Sprintf("%v", lens.GetValue(0))
	size := 1
	i, err := strconv.Atoi(x)

	if err != nil {
		return "1"
	} else {
		size = i
	}
	if lens.Len() > 1 {
		y := fmt.Sprintf("%v", lens.GetValue(1))
		j, err := strconv.Atoi(y)

		if err != nil {
			return "1"
		} else {
			size = size * j
		}
		if lens.Len() > 2 {
			z := fmt.Sprintf("%v", lens.GetValue(1))
			k, err := strconv.Atoi(z)
			if err != nil {
				return "1"
			} else {
				size = size * k
			}
		}
	}
	return strconv.Itoa(size)
}



*/
