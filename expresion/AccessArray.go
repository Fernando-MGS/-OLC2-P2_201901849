package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"

	"github.com/colegno/arraylist"
)

type ArrayAccess struct {
	Id     interfaces.Expresion
	Access *arraylist.List
	Line   string
	Col    string
}

func NewArrayAccess(id interfaces.Expresion, access *arraylist.List, line, col int) ArrayAccess {
	expr := ArrayAccess{id, access, strconv.Itoa(line), strconv.Itoa(col)}
	return expr
}

func (p ArrayAccess) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	array := p.Id.Ejecutar(env, gen)
	if array.Type != interfaces.ARRAY {
		env.(environment.Environment).NewError("EL ELEMENTO AL QUE SE QUIERE ACCEDER NO ES UN ARRAY", p.Line, p.Col)
		retorno.Type = interfaces.NULL
		return retorno
	}
	ambito := env.(environment.Environment).DevAmbito()
	gen.AddCodes("//ACCEDIENDO A ELEMENTO ", ambito)
	//array := env.(environment.Environment).GetVariable(p.Id, p.Line, p.Col)
	if array.Type != interfaces.ARRAY && array.Type != interfaces.VECTOR {
		retorno.Type = interfaces.NULL
		env.(environment.Environment).NewError("EL ELEMENTO NO ES UN ARRAY NI UN VECTOR", p.Line, p.Col)
		return retorno
	}
	tipos := array.Tipo2.GetValue(0).(interfaces.Dimensions)
	if len(p.Access.ToArray()) > tipos.Dimensions.Len() {
		retorno.Type = interfaces.NULL
		env.(environment.Environment).NewError("ACCESO FUERA DE LOS LIMITES DEL ARRAY", p.Line, p.Col)
		gen.AddFuncExtra("BOUNDS")
		gen.AddCodes("proc_boundsError();", env.(environment.Environment).DevAmbito())
		return retorno
	} else if p.Access.Len() == tipos.Dimensions.Len() {
		retorno.Type = tipos.Tipo
	} else {
		retorno.Type = array.Type
	}
	fmt.Println("LA POS ORIGINAL ES " + array.Value)
	list_access := EjecutarAccesos(p.Access, env, gen)
	dimensions := array.Tipo2.GetValue(0).(interfaces.Dimensions)
	result_access := AccessVector(array.Value, "", dimensions.Dimensions, list_access, gen, ambito)
	retorno.Value = result_access.Value
	retorno.TrueLabel = result_access.TrueLabel
	retorno.IsTemp = true
	new_dimension := interfaces.Dimensions{Tipo: dimensions.Tipo, Dimensions: result_access.Tipo2}
	retorno.Tipo2 = arraylist.New()
	retorno.Tipo2.Add(new_dimension)
	fmt.Println("SALIO DEL ARRAY ACCESS")
	return retorno
}

func EjecutarAccesos(lista *arraylist.List, env interface{}, gen *generator.Generator) *arraylist.List {
	access_list := arraylist.New()
	for _, s := range lista.ToArray() {
		valor := s.(interfaces.Expresion).Ejecutar(env, gen)
		access_list.Add(valor)
	}
	return access_list
}

func AccessVector(origen, posHeap string, tipo2, acceso *arraylist.List, gen *generator.Generator, ambito bool) interfaces.Value {
	var retorno interfaces.Value
	if acceso.Len() > 0 {
		index := acceso.GetValue(0).(interfaces.Value).Value
		tmp_size := gen.NewTemp()
		tmp_pos := gen.NewTemp()
		tmp_value := gen.NewTemp()
		gen.AddCodes("//ACCEDIENDO A VALOR", ambito)
		gen.AddCodes(tmp_size+"="+index+"+1;", ambito) //se le suma 1 al valor porque en el 0 está el tamaño del vector
		gen.AddCodes(tmp_pos+"="+tmp_size+"+"+origen+";", ambito)
		gen.AddCodes(tmp_value+"=HEAP[(int)"+tmp_pos+"];", ambito)
		_tipo2 := Delete_element(tipo2)
		_acceso := Delete_element(acceso)
		retorno = AccessVector(tmp_value, tmp_pos, _tipo2, _acceso, gen, ambito)
	} else {
		retorno.Value = origen
		retorno.Tipo2 = tipo2
		retorno.IsTemp = true
		retorno.TrueLabel = posHeap
	}
	return retorno
}

func Delete_element(lista *arraylist.List) *arraylist.List {
	new_lista := arraylist.New()
	conf := true
	for _, s := range lista.ToArray() {
		if conf {
			conf = false
		} else {
			new_lista.Add(s)
		}
	}
	return new_lista
}
