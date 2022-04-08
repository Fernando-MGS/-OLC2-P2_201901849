package environment

import (
	"OLC2/interfaces"
	"fmt"
	"time"

	"github.com/colegno/arraylist"
)

type Environment struct {
	id       string
	father   interface{}
	variable map[string]interfaces.Symbol
	size     int
	bases    *arraylist.List
	simbolos *arraylist.List
	errores  *arraylist.List
}

func NewEnvironment(father interface{}, id string) Environment {
	env := Environment{id, father, make(map[string]interfaces.Symbol), 0, arraylist.New(), arraylist.New(), arraylist.New()}
	return env
}

func (env Environment) SaveVariable(id string, value interfaces.Symbol, tipo interfaces.TipoSimbolo) {
	if variable, ok := env.variable[id]; ok {
		t := time.Now()
		fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		err := interfaces.Errores{Line: "0", Col: "0", Mess: "LA VARIABLE " + variable.Id + " YA EXISTE", Fecha: fecha}
		env.Errores(err)
		return
	}
	env.variable[id] = value
	simbolos := interfaces.Simbolos{ID: id, Tipo: Tipo2(value.Tipo.Tipo), Ambito: "GLOBAL", Fila: "0", Columna: "0"}
	env.LogSimbolo(simbolos)
	//env.size = env.size + 1
}

func (env Environment) GetVariable(id string) interfaces.Symbol {

	var tmpEnv Environment
	tmpEnv = env

	for {
		if variable, ok := tmpEnv.variable[id]; ok {
			return variable
		}

		if tmpEnv.father == nil {
			break
		} else {
			tmpEnv = tmpEnv.father.(Environment)
		}
	}

	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	err := interfaces.Errores{Line: "0", Col: "0", Mess: "LA VARIABLE " + id + " NO EXISTE", Fecha: fecha}
	env.Errores(err)
	tipos := interfaces.TipoSimbolo{Tipo: interfaces.NULL}
	return interfaces.Symbol{Id: "", Tipo: tipos, Posicion: 0}
}

func (env Environment) LogBase(s interfaces.Bases) {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if tmpEnv.father == nil {
			break
		} else {
			tmpEnv = tmpEnv.father.(Environment)
		}
	}
	fmt.Println(s)
	tmpEnv.bases.Add(s)
}

func (env Environment) LogSimbolo(s interfaces.Simbolos) {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if tmpEnv.father == nil {
			break
		} else {
			tmpEnv = tmpEnv.father.(Environment)
		}
	}
	tmpEnv.simbolos.Add(s)
}

func (env Environment) Errores(err interfaces.Errores) {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if tmpEnv.father == nil {
			break
		} else {
			tmpEnv = tmpEnv.father.(Environment)
		}
	}
	tmpEnv.errores.Add(err)
}

func (env Environment) DevSimbolos() []interface{} {
	return env.simbolos.ToArray()
}

func (env Environment) DevBases() []interface{} {
	return env.bases.ToArray()
}

func (env Environment) DevErrores() []interface{} {
	return env.errores.ToArray()
}

func Tipo2(t interfaces.TipoExpresion) string {
	if t == interfaces.ARRAY {
		return "ARRAY"
	} else if t == interfaces.CHAR {
		return "CHAR"
	} else if t == interfaces.BOOLEAN {
		return "BOOLEAN"
	} else if t == interfaces.FLOAT {
		return "FLOAT"
	} else if t == interfaces.INTEGER {
		return "INTEGER"
	} else if t == interfaces.STR {
		return "STR"
	} else if t == interfaces.STRING {
		return "STRING"
	} else if t == interfaces.STRUCT {
		return "STRUCT"
	} else if t == interfaces.USIZE {
		return "USIZE"
	} else if t == interfaces.VECTOR {
		return "VECTOR"
	}
	return "VOID"
}
