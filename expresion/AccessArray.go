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
	tipos := array.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions)
	if len(p.Access.ToArray()) > tipos.Dimensions.Len() {
		retorno.Type = interfaces.NULL
		env.(environment.Environment).NewError("ACCESO FUERA DE LOS LIMITES DEL ARRAY", p.Line, p.Col)
		gen.AddFuncExtra("BOUNDS")
		gen.AddCodes("proc_boundsError();", env.(environment.Environment).DevAmbito())
		return retorno
	} else if p.Access.Len() == tipos.Dimensions.Len() {
		retorno.Type = tipos.Tipo
	} else {
		retorno.Type = array.Tipo.Tipo
	}
	//largo:=gen.NewTemp()
	first_iter := gen.NewTemp()
	//i:=0
	x := ""
	y := ""
	z := ""
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
	conf := false
	if p.Access.Len() == tipos.Dimensions.Len() {
		conf = true
		//valor_d := fmt.Sprintf("%v", tipos.Dimensions.GetValue(0))
		//retorno.Type=tipos.Tipo
		gen.AddCodes(first_iter+"=1;", ambito)
	} else if tipos.Dimensions.Len()-p.Access.Len() == 1 {
		valor_d := fmt.Sprintf("%v", tipos.Dimensions.GetValue(0))
		//gen.AddCodes(first_iter+"="+valor_d+";",ambito)
		//valor_d1 := fmt.Sprintf("%v", tipos.Dimensions.GetValue(1))
		//tmp1:=gen.NewTemp()
		//tmp2:=gen.NewTemp()
		//gen.AddCodes(tmp1+"="+valor_d+";",ambito)
		gen.AddCodes(first_iter+"="+valor_d+";", ambito)
		//first_iter=tmp1
	} else {
		/*l1:=tipos.Dimensions.Len()-2
		l2:=tipos.Dimensions.Len()-1*/
		valor_d := fmt.Sprintf("%v", tipos.Dimensions.GetValue(0))
		valor_d1 := fmt.Sprintf("%v", tipos.Dimensions.GetValue(1))
		//valor_d2 := fmt.Sprintf("%v", tipos.Dimensions.GetValue(2))
		//tmp1:=gen.NewTemp()
		//tmp2 := gen.NewTemp()
		//gen.AddCodes(tmp1+"="+valor_d+";",ambito)
		gen.AddCodes(first_iter+"="+valor_d1+"*"+valor_d+";", ambito)
		//gen.AddCodes(first_iter+"="+tmp2+"*"+valor_d2+";", ambito)
		//first_iter=tmp1
	}

	//j:=tipos.Dimensions.Len()-1
	//expr:=arraylist.New()
	gen.AddCodes("//CALCULANDO POSICION", ambito)
	coordenada := gen.NewTemp()
	if p.Access.Len() == 1 {
		valor := p.Access.GetValue(0).(interfaces.Expresion).Ejecutar(env, gen).Value
		if tipos.Dimensions.Len() == 1 {
			gen.AddCodes(coordenada+"="+valor+";", ambito)
		} else {
			//valor_d:=p.Access.GetValue(1).(interfaces.Expresion).Ejecutar(env,gen).Value
			if tipos.Dimensions.Len() == 2 {
				gen.AddCodes(coordenada+"="+y+"*"+valor+";", ambito)
			} else {
				//valor_d1:=p.Access.GetValue(2).(interfaces.Expresion).Ejecutar(env,gen).Value
				tmp := gen.NewTemp()
				gen.AddCodes(tmp+"="+y+"*"+valor+";", ambito)
				gen.AddCodes(coordenada+"="+tmp+"*"+z+";", ambito)
			}

		}

	} else if p.Access.Len() == 2 {
		i := p.Access.GetValue(0).(interfaces.Expresion).Ejecutar(env, gen).Value
		if tipos.Dimensions.Len() == 2 {
			gen.AddCodes(coordenada+"="+i+";", ambito)
		} else {
			j := p.Access.GetValue(1).(interfaces.Expresion).Ejecutar(env, gen).Value
			tmp := gen.NewTemp()
			if tipos.Dimensions.Len() == 2 {
				gen.AddCodes(tmp+"="+y+"*"+i+";", ambito)
				gen.AddCodes(coordenada+"="+tmp+"+"+j+";", ambito)
			} else {
				tmp2 := gen.NewTemp()
				tmp3 := gen.NewTemp()
				//k:=p.Access.GetValue(2).(interfaces.Expresion).Ejecutar(env,gen).Value
				gen.AddCodes(tmp+"="+y+"*"+i+";", ambito)
				gen.AddCodes(tmp2+"="+tmp+"*"+z+";", ambito)
				gen.AddCodes(tmp3+"="+j+"*"+z+";", ambito)
				gen.AddCodes(coordenada+"="+tmp3+"+"+tmp2+";", ambito)
				/*gen.AddCodes(tmp2+"="+z+"*"+j+";",ambito)
				gen.AddCodes(tmp3+"="+tmp+"+"+tmp2+";",ambito)
				gen.AddCodes(coordenada+tmp3+"+"+k+";",ambito)*/
			}
		}
	} else {
		tmp := gen.NewTemp()
		tmp2 := gen.NewTemp()
		tmp3 := gen.NewTemp()
		tmp4 := gen.NewTemp()
		i := p.Access.GetValue(0).(interfaces.Expresion).Ejecutar(env, gen).Value
		j := p.Access.GetValue(1).(interfaces.Expresion).Ejecutar(env, gen).Value
		k := p.Access.GetValue(2).(interfaces.Expresion).Ejecutar(env, gen).Value
		gen.AddCodes(tmp+"="+y+"*"+i+";", ambito)
		gen.AddCodes(tmp2+"="+tmp+"*"+z+";", ambito)
		gen.AddCodes(tmp3+"="+j+"*"+z+";", ambito)
		gen.AddCodes(tmp4+"="+tmp3+"+"+tmp2+";", ambito)
		gen.AddCodes(coordenada+"="+tmp4+"+"+k+";", ambito)
	}
	pos := gen.NewTemp()
	if conf {
		gen.AddCodes(pos+"="+array.Posicion2+"+"+coordenada+";", ambito)
		resultado := gen.NewTemp()
		gen.AddCodes(resultado+"=HEAP[(int)"+pos+"];", ambito)
		retorno.TrueLabel = pos
		pos = resultado
	} else {
		gen.AddCodes(pos+"="+array.Posicion2+"+"+coordenada+";", ambito)
		retorno.TrueLabel = first_iter
	}
	retorno.Value = pos // si es un array o un vector manda la posicion dentro del heap

	//retorno.Type=array.Tipo.Tipo
	fmt.Println("lA POS ES " + pos)
	fmt.Println("La long es " + first_iter)
	gen.AddCodes("//ACCESO CREADO", ambito)
	return retorno
}
