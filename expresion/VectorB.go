package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"

	"github.com/colegno/arraylist"
)

type VectorB struct {
	list *arraylist.List
}

func NewVectorB(lista *arraylist.List) VectorB {
	expr := VectorB{lista}
	return expr
}

func (p VectorB) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	retorno.Type = interfaces.VECTOR
	//ambito := env.(environment.Environment).DevAmbito()
	name := env.(environment.Environment).Control.Id
	valores := arraylist.New()
	conf := true
	//gen.AddCodes("//CREANDO VECTOR CON LISTA DE VALORES", ambito)
	gen.NewComentario("CREANDO VECTOR CON LISTA DE VALORES", name, true, false, "")
	var tipo interfaces.TipoExpresion
	for _, s := range p.list.ToArray() {
		res := s.(interfaces.Expresion).Ejecutar(env, gen)
		if res.Type == interfaces.VECTOR && conf {
			retorno.Tipo2 = res.Tipo2
			conf = false
		} else {
			tipo = res.Type
		}
		valores.Add(res)
		//fmt.Println(res)
	}
	largo := strconv.Itoa(p.list.Len())
	long := gen.NewTemp()
	//pos := gen.NewTemp()
	//gen.AddCodes(long+"=H;//POS DE INICIO VECTOR B", ambito)
	gen.NewAsignacion(long, "H", false, "", name, true, false, "")
	//gen.AddCodes("HEAP[(int)H]="+largo+";", ambito)
	gen.NewHeap("H", largo, false, "", name, true, false, "")
	//gen.AddCodes("H=H+1;", ambito)
	gen.NewOperacion("H", "H", "+", "1", false, "", name, true, false, "")
	for _, s := range valores.ToArray() {
		valors := s.(interfaces.Value)
		//gen.AddCodes("HEAP[(int)H]="+valors.Value+";", ambito)
		gen.NewHeap("H", valors.Value, false, "", name, true, false, "")
		//gen.AddCodes("H=H+1;", ambito)
		gen.NewOperacion("H", "H", "+", "1", false, "", name, true, true, "")
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
		tipo_d.Dimensions.Add(interfaces.VECTOR)
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
