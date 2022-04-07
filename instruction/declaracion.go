package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"time"
)

type Declaracion struct {
	Id        string
	Tipo      interfaces.TipoSimbolo
	Expresion interfaces.Expresion
	Mutable   bool
}

func NewDeclaracion(id string, tipo interfaces.TipoSimbolo, val interfaces.Expresion, isMutable bool) Declaracion {
	//fmt.Println(tipo)
	instr := Declaracion{id, tipo, val, isMutable}
	return instr
}

func (p Declaracion) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	fmt.Println("DECLARACION")
	result := p.Expresion.Ejecutar(env, gen)
	if result.Type == p.Tipo.Tipo {
		if p.Tipo.Tipo == interfaces.ARRAY {
			fmt.Println(p.Tipo.Tipo)
			fmt.Println(p.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions).Tipo)
			fmt.Println(p.Tipo.Tipo2.GetValue(0).(interfaces.Dimensions).Dimensions.Clone().ToArray()...)
		} else if p.Tipo.Tipo == interfaces.VECTOR {
			fmt.Println(p.Tipo.Tipo)
			fmt.Println(p.Tipo.Tipo2.ToArray()...)
		} else {
			if result.Type == interfaces.INTEGER {
				simbolo := interfaces.Symbol{Id: p.Id, Tipo: p.Tipo, Posicion: gen.Stack, Mutable: p.Mutable}
				if result.IsTemp {
					gen.AddCode("STACK[(int)P]=" + fmt.Sprintf("%v", result.Value) + ";")
				} else {
					gen.AddCode("STACK[(int)P]=(float)" + fmt.Sprintf("%v", result.Value) + ";")
				}
				gen.AddCode("P=P+1;")
				gen.Stack++
				env.(environment.Environment).SaveVariable(p.Id, simbolo, p.Tipo)
			} else if result.Type == interfaces.FLOAT {
				simbolo := interfaces.Symbol{Id: p.Id, Tipo: p.Tipo, Posicion: gen.Stack, Mutable: p.Mutable}
				gen.AddCode("STACK[(int)P]=" + fmt.Sprintf("%v", result.Value) + ";")
				gen.AddCode("P=P+1;")
				gen.Stack++
				env.(environment.Environment).SaveVariable(p.Id, simbolo, p.Tipo)
			} /*else if result.Type == interfaces.CHAR {
				simbolo := interfaces.Symbol{Id: "", Tipo: p.Tipo, Posicion: gen.Stack, Mutable: p.Mutable}
				gen.AddCode("stack[(int)P]=" + fmt.Sprintf("%v", result.Value))
				gen.AddCode("P=P+1;")
				gen.Stack++
				env.(environment.Environment).SaveVariable(p.Id, simbolo, p.Tipo)
			}*/
		}
	} else {
		t := time.Now()
		fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		err := interfaces.Errores{Line: "0", Col: "0", Mess: "LOS TIPOS NO CONCUERDAN EN LA DECLARACION", Fecha: fecha}
		env.(environment.Environment).Errores(err)
	}

	result.Type = interfaces.NULL
	return result.Value
}
