package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"

	"github.com/colegno/arraylist"
)

type VectorC struct {
	hola string
}

func NewVectorC() VectorC {
	return VectorC{""}
}

func (p VectorC) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	name := env.(environment.Environment).Control.Id
	//ambito := env.(environment.Environment).DevAmbito()
	//gen.AddCodes("//CREANDO VECTOR VACIO", ambito)
	gen.NewComentario("CREANDO VECTOR VACIO", name, true, false, "")
	retorno.Type = interfaces.VECTOR
	retorno.Tipo2 = arraylist.New()
	dim_long := arraylist.New()
	dim_long.Add(interfaces.VOID)
	dimension := interfaces.Dimensions{Tipo: interfaces.VOID, Dimensions: dim_long}
	retorno.Tipo2.Add(dimension)
	pos := gen.NewTemp()
	//gen.AddCodes(pos+"=H;", ambito)
	gen.NewAsignacion(pos, "H", false, "", name, true, true, "")
	//gen.AddCodes("HEAP[(int)H]=0;", ambito)
	gen.NewHeap("H", "0", false, "", name, true, false, "")
	//gen.AddCodes("H=H+1;", ambito)
	gen.NewOperacion("H", "H", "+", "1", false, "", name, true, false, "")
	retorno.Value = pos

	return retorno
}
