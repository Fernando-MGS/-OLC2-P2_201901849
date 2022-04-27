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
	Id     string
	Access *arraylist.List
	Line   string
	Col    string
}

func NewArrayAccess(id string, access *arraylist.List, line, col int) ArrayAccess {
	expr := ArrayAccess{id, access, strconv.Itoa(line), strconv.Itoa(col)}
	return expr
}

func (p ArrayAccess) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value

	ambito := env.(environment.Environment).DevAmbito()
	gen.AddCodes("//ACCEDIENDO A ARRAY "+p.Id, ambito)
	array := env.(environment.Environment).GetVariable(p.Id, p.Line, p.Col)
	if array.Tipo.Tipo != interfaces.ARRAY && array.Tipo.Tipo != interfaces.VECTOR {
		retorno.Type = interfaces.NULL
		env.(environment.Environment).NewError("EL ELEMENTO NO ES UN ARRAY NI UN VECTOR", p.Line, p.Col)
		return retorno
	}
	first_iter := ""
	x := ""
	y := ""
	z := ""
	tipos := array.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions)
	conf := false
	if len(p.Access.ToArray()) > tipos.Dimensions.Len() {
		retorno.Type = interfaces.NULL
		env.(environment.Environment).NewError("ACCESO FUERA DE LOS LIMITES DEL ARRAY", p.Line, p.Col)
		gen.AddFuncExtra("BOUNDS")
		gen.AddCodes("proc_boundsError();", env.(environment.Environment).DevAmbito())
		return retorno
	} else if p.Access.Len() == tipos.Dimensions.Len() {
		retorno.Type = tipos.Tipo
		conf = true
	} else {
		retorno.Type = array.Tipo.Tipo
	}
	if array.Tipo.Tipo == interfaces.ARRAY {
		//largo:=gen.NewTemp()
		first_iter = gen.NewTemp()
		//i:=0
		fmt.Println(x)
		gen.AddCodes("//CALCULANDO LONGITUD", ambito)
		if tipos.Dimensions.Len() == 1 {
			x = fmt.Sprintf("%v", tipos.Dimensions.GetValue(0))
			//x=valor_d
		} else if tipos.Dimensions.Len() == 2 {
			y = fmt.Sprintf("%v", tipos.Dimensions.GetValue(0))
			x = fmt.Sprintf("%v", tipos.Dimensions.GetValue(1))
		} else {
			z = fmt.Sprintf("%v", tipos.Dimensions.GetValue(0))
			y = fmt.Sprintf("%v", tipos.Dimensions.GetValue(1))
			x = fmt.Sprintf("%v", tipos.Dimensions.GetValue(2))
		}
		if p.Access.Len() == tipos.Dimensions.Len() {
			//conf = true
			gen.AddCodes(first_iter+"=1;", ambito)
		} else if tipos.Dimensions.Len()-p.Access.Len() == 1 {
			valor_d := fmt.Sprintf("%v", tipos.Dimensions.GetValue(0))
			gen.AddCodes(first_iter+"="+valor_d+";", ambito)
		} else {
			valor_d := fmt.Sprintf("%v", tipos.Dimensions.GetValue(0))
			valor_d1 := fmt.Sprintf("%v", tipos.Dimensions.GetValue(1))
			gen.AddCodes(first_iter+"="+valor_d1+"*"+valor_d+";", ambito)
		}

		//FIN DEL CALCULO DE LONGITUD

		gen.AddCodes("//CALCULANDO POSICION", ambito)
		coordenada := gen.NewTemp()
		if p.Access.Len() == 1 {
			valor := p.Access.GetValue(0).(interfaces.Expresion).Ejecutar(env, gen).Value //El valor del indice en el acceso id[valor]
			if array.Tipo.Tipo == interfaces.ARRAY {
				if tipos.Dimensions.Len() == 1 {
					gen.AddCodes(coordenada+"="+valor+";", ambito)
				} else {
					if tipos.Dimensions.Len() == 2 {
						if array.Tipo.Tipo == interfaces.ARRAY {
							gen.AddCodes(coordenada+"="+y+"*"+valor+";", ambito)
						}
					} else {
						if array.Tipo.Tipo == interfaces.ARRAY {
							tmp := gen.NewTemp()
							gen.AddCodes(tmp+"="+y+"*"+valor+";", ambito)
							gen.AddCodes(coordenada+"="+tmp+"*"+z+";", ambito)
						}
					}

				}
			} else {

				tmp_size := gen.NewTemp()
				//tmp_pos:=gen.NewTemp()
				gen.AddCodes("//calculo de 1 dimension", ambito)
				gen.AddCodes(tmp_size+"="+valor+"+1;", ambito)                        //se le suma 1 al valor porque en el 0 está el tamaño del vector
				gen.AddCodes(coordenada+"="+tmp_size+"+"+array.Posicion2+";", ambito) // en la coordenada está la posición del vector donde está el vector
				//gen.AddCodes(coordenada+"=",ambito)
				//coordenada = array.Posicion2
			}
		} else if p.Access.Len() == 2 {
			i := p.Access.GetValue(0).(interfaces.Expresion).Ejecutar(env, gen).Value
			j := p.Access.GetValue(1).(interfaces.Expresion).Ejecutar(env, gen).Value
			if tipos.Dimensions.Len() == 2 {
				if interfaces.ARRAY == array.Tipo.Tipo {
					tmp := gen.NewTemp()
					if tipos.Dimensions.Len() == 2 {
						gen.AddCodes(tmp+"="+y+"*"+i+";", ambito)
						gen.AddCodes(coordenada+"="+tmp+"+"+j+";", ambito)
					} else {
						tmp2 := gen.NewTemp()
						tmp3 := gen.NewTemp()
						gen.AddCodes(tmp+"="+y+"*"+i+";", ambito)
						gen.AddCodes(tmp2+"="+tmp+"*"+z+";", ambito)
						gen.AddCodes(tmp3+"="+j+"*"+z+";", ambito)
						gen.AddCodes(coordenada+"="+tmp3+"+"+tmp2+";", ambito)
					}
				} else {
					tmp_size := gen.NewTemp()
					tmp_pos := gen.NewTemp()
					tmp_value := gen.NewTemp()
					gen.AddCodes(tmp_size+"="+i+"+1;", ambito)                         //se le suma 1 al valor porque en el 0 está el tamaño del vector
					gen.AddCodes(tmp_pos+"="+tmp_size+"+"+array.Posicion2+";", ambito) // en la coordenada está la posición del vector donde está el vector
					gen.AddCodes(tmp_value+"=HEAP[(int)"+tmp_pos+"];", ambito)         //La posicion del segundo vector, también es la longitud de este segundo
					gen.AddCodes(coordenada+"="+tmp_value+"+"+j+";", ambito)
				}
			}
		} else {
			i := p.Access.GetValue(0).(interfaces.Expresion).Ejecutar(env, gen).Value
			j := p.Access.GetValue(1).(interfaces.Expresion).Ejecutar(env, gen).Value
			k := p.Access.GetValue(2).(interfaces.Expresion).Ejecutar(env, gen).Value
			if interfaces.ARRAY == array.Tipo.Tipo {
				tmp := gen.NewTemp()
				tmp2 := gen.NewTemp()
				tmp3 := gen.NewTemp()
				tmp4 := gen.NewTemp()

				gen.AddCodes(tmp+"="+y+"*"+i+";", ambito)
				gen.AddCodes(tmp2+"="+tmp+"*"+z+";", ambito)
				gen.AddCodes(tmp3+"="+j+"*"+z+";", ambito)
				gen.AddCodes(tmp4+"="+tmp3+"+"+tmp2+";", ambito)
				gen.AddCodes(coordenada+"="+tmp4+"+"+k+";", ambito)
			} else {
				tmp_size := gen.NewTemp()
				tmp_pos := gen.NewTemp()
				tmp_value := gen.NewTemp()
				tmp_value2 := gen.NewTemp()
				tmp_pos2 := gen.NewTemp()
				gen.AddCodes(tmp_size+"="+i+"+1;", ambito)                         //se le suma 1 al valor porque en el 0 está el tamaño del vector
				gen.AddCodes(tmp_pos+"="+tmp_size+"+"+array.Posicion2+";", ambito) // en la coordenada está la posición del vector donde está el vector
				gen.AddCodes(tmp_value+"=HEAP[(int)"+tmp_pos+"];", ambito)         //La posicion del segundo vector, también es la longitud de este segundo
				gen.AddCodes(tmp_pos2+"="+tmp_value+"+"+j+";", ambito)
				gen.AddCodes(tmp_value2+"=HEAP[(int)"+tmp_pos2+"];", ambito) //La posicion del segundo vector, también es la longitud de este segundo
				gen.AddCodes(coordenada+"="+tmp_value2+"+"+k+";", ambito)
			}
		}
		//FIN DEL CALCULO DE LA POSICION
		pos := gen.NewTemp()
		if conf {
			if array.Tipo.Tipo == interfaces.ARRAY {
				gen.AddCodes(pos+"="+array.Posicion2+"+"+coordenada+";", ambito) //sumamos la posicion original a la posicion calculada
				resultado := gen.NewTemp()
				gen.AddCodes(resultado+"=HEAP[(int)"+pos+"];", ambito) //le asignamos a un temporal el valor en el heap
				retorno.TrueLabel = pos
				pos = resultado
			} else {
				//gen.AddCodes(pos+"="+array.Posicion2+"+"+coordenada+";", ambito)
				pos = coordenada
				resultado := gen.NewTemp()
				gen.AddCodes(resultado+"=HEAP[(int)"+pos+"];", ambito)
				//retorno.TrueLabel = pos
				pos = resultado
			}
		} else {
			if array.Tipo.Tipo == interfaces.ARRAY {
				gen.AddCodes(pos+"="+array.Posicion2+"+"+coordenada+";", ambito)
				retorno.TrueLabel = first_iter
			} else {
				pos = coordenada
			}
		}
		retorno.Value = pos // si es un array o un vector manda la posicion dentro del heap

		fmt.Println("lA POS ES " + pos)
		fmt.Println("La long es " + first_iter)
		gen.AddCodes("//ACCESO CREADO", ambito)
	} else {
		fmt.Println("LA POS ORIGINAL ES " + array.Posicion2)
		list_access := EjecutarAccesos(p.Access, env, gen)
		dimensions := array.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions).Dimensions
		result_access := AccessVector(array.Posicion2, dimensions, list_access, gen, ambito)
		retorno.Value = result_access.Value
		retorno.IsTemp = true
	}

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

func AccessVector(origen string, tipo2, acceso *arraylist.List, gen *generator.Generator, ambito bool) interfaces.Value {
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
		retorno = AccessVector(tmp_value, _tipo2, _acceso, gen, ambito)
	} else {
		retorno.Value = origen
		retorno.Tipo2 = tipo2
		retorno.IsTemp = true
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
