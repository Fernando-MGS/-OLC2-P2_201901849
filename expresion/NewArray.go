package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"reflect"
	"strconv"

	arrayList "github.com/colegno/arraylist"
)

type Array struct {
	ListExp *arrayList.List
}

func NewArray(list *arrayList.List) Array {
	exp := Array{list}
	return exp
}

func (p Array) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {

	//tempExp := arrayList.New()
	//arr := p.ListExp.ToArray()
	/*fmt.Println(arr[len(arr)-1])
	fmt.Println("=======")*/
	/*for _, s := range p.ListExp.ToArray() {
		//tempExp.Add(s.(interfaces.Expresion).Ejecutar(env,gen))
	}*/
	ambito := env.(environment.Environment).DevAmbito()
	valor := gen.NewTemp()
	var ret interfaces.Value
	//tipo := WatchTipos(p.ListExp, env, gen)
	gen.AddCodes("//CREANDO ARRAY", ambito)
	dimensiones := ContarLista(p.ListExp)
	dimensiones.Add(strconv.Itoa(len(p.ListExp.ToArray())))
	//fmt.Println(dimensiones.ToArray()...)

	tipo_def, conf := ExecExpresiones(p.ListExp, arrayList.New(), env, gen)
	//fmt.Println(conf.ToArray()...)
	init := true
	for _, s := range conf.ToArray() {
		if init {
			gen.AddCodes(valor+"=H;", ambito)
			init = false
		}
		val := fmt.Sprintf("%v", s)
		gen.AddCodes("HEAP[(int)H]="+val+";", ambito)
		gen.AddCodes("H=H+1;", ambito)
		gen.Heap++
	}
	dimension := interfaces.Dimensions{Tipo: tipo_def, Dimensions: dimensiones}
	gen.AddCodes("HEAP[(int)H]=-123456;", ambito)
	gen.AddCodes("H=H+1;", ambito)
	gen.Heap++
	ret.Type = interfaces.ARRAY
	ret.IsTemp = true
	//ret.Tipo2 = FixLista(tipo, tipo_def)
	ret.Tipo2 = arrayList.New()
	ret.Tipo2.Add(dimension)
	fmt.Println("ret.Tipo2.ToArray()...")
	fmt.Println(ret.Tipo2.ToArray()...)
	ret.Value = valor
	ret.TrueLabel = DevSize(dimensiones)
	gen.AddCodes("//ARRAY CREADO", ambito)
	return ret
}

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
