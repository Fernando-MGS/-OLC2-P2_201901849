package environment

import (
	"OLC2/interfaces"
	"fmt"

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
		fmt.Println("La variable " + variable.Id + " ya existe")
		return
	}
	env.variable[id] = value
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

	fmt.Println("La variable no existe")
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
