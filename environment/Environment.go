package environment

import (
	"OLC2/interfaces"
	"fmt"
	"time"

	"github.com/colegno/arraylist"
)

type Environment struct {
	Control   *EnvControl
	father    interface{}
	variable  map[string]interfaces.Symbol
	bases     *arraylist.List
	simbolos  *arraylist.List
	errores   *arraylist.List
	functions map[string]interfaces.Functions
	structs   map[string]interfaces.Structs
}

type EnvControl struct {
	Id        string
	Stack     int
	Heap      int
	Entrada   string
	Salida    string
	Ciclo     bool
	Retorno   string
	Variables *arraylist.List
}

func NewEnvironment(father interface{}, id, in, out string, ciclo bool, stack int, retorn string) Environment {
	env := Environment{&EnvControl{Id: id, Stack: stack, Heap: 1, Retorno: retorn, Entrada: in, Salida: out, Ciclo: ciclo, Variables: arraylist.New()}, father, make(map[string]interfaces.Symbol), arraylist.New(), arraylist.New(), arraylist.New(), make(map[string]interfaces.Functions), make(map[string]interfaces.Structs)}
	return env
}

func (env Environment) SaveFunc(line, col, id string, function interfaces.Functions) bool {
	if _, ok := env.variable[id]; ok {
		env.NewError("LA FUNCION \""+id+"\" YA EXISTE", line, col)
		return false
	}
	env.functions[id] = function
	simbolos := interfaces.Simbolos{ID: id, Tipo: Tipo2(function.Tipo.Tipo), Ambito: env.Control.Id, Fila: line, Columna: col}
	env.LogSimbolo(simbolos)
	return true
	//env.size = env.size + 1
}

func (env Environment) DevSalida() string {
	var tmpEnv Environment
	tmpEnv = env
	//salida:=""
	for {
		if tmpEnv.father.(Environment).Control.Id == "GLOBAL" {
			break
		} else {
			tmpEnv = tmpEnv.father.(Environment)
		}
	}
	return tmpEnv.Control.Retorno
}

func (env Environment) SaveVariable(line, col, id string, value interfaces.Symbol, Tipo interfaces.TipoSimbolo) {
	if variable, ok := env.variable[id]; ok {
		t := time.Now()
		fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		err := interfaces.Errores{Line: line, Col: col, Mess: "LA VARIABLE \"" + variable.Id + "\" YA EXISTE", Fecha: fecha}
		env.Errores(err)
		return
	}
	env.variable[id] = value
	if Tipo.Tipo == interfaces.BOOLEAN || Tipo.Tipo == interfaces.CHAR || Tipo.Tipo == interfaces.FLOAT || Tipo.Tipo == interfaces.INTEGER || Tipo.Tipo == interfaces.USIZE {
		env.Control.Variables.Add(value)
	}
	simbolos := interfaces.Simbolos{ID: id, Tipo: Tipo2(value.Tipo.Tipo), Ambito: env.Control.Id, Fila: line, Columna: col}
	env.LogSimbolo(simbolos)
	//env.size = env.size + 1
}

func (env Environment) DevVariables() *arraylist.List {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if tmpEnv.father.(Environment).Control.Id == "GLOBAL" {
			break
		} else {
			tmpEnv = tmpEnv.father.(Environment)
		}
	}
	//fmt.Println("EL SIZE LEN ES ", tmpEnv.Control.Variables.Len(), env.Control.Id)
	return tmpEnv.Control.Variables
}

func (env Environment) SaveStruct(line, col, id string, structs interfaces.Structs) {
	if variable, ok := env.variable[id]; ok {
		env.NewError("EL STRUCT \""+variable.Id+"\" YA HA SIDO DECLARADO", line, col)
		return
	}
	env.structs[id] = structs
	simbolos := interfaces.Simbolos{ID: id, Tipo: "STRUCT", Ambito: env.Control.Id, Fila: line, Columna: col}
	env.LogSimbolo(simbolos)
	//env.size = env.size + 1
}

func (env Environment) GetFunc(id, line, col string) (bool, interfaces.Functions) {
	var tmpEnv Environment
	tmpEnv = env

	for {
		if variable, ok := tmpEnv.functions[id]; ok {
			return true, variable
		}

		if tmpEnv.father == nil {
			break
		} else {
			tmpEnv = tmpEnv.father.(Environment)
		}
	}
	env.NewError("LA FUNCION "+id+" NO HA SIDO ENCONTRADA", line, col)
	return false, interfaces.Functions{}
}

func inutil(v interfaces.Symbol) {

}

func (env Environment) AlterVariable(id string, value interfaces.Symbol) {
	var tmpEnv Environment
	tmpEnv = env

	for {
		if variable, ok := tmpEnv.variable[id]; ok {
			inutil(variable)
			tmpEnv.variable[id] = value
			return
		}

		if tmpEnv.father == nil {
			break
		} else {
			tmpEnv = tmpEnv.father.(Environment)
		}
	}
}

func (env Environment) GetVariable(id, line, col string) interfaces.Symbol {

	var tmpEnv Environment
	tmpEnv = env

	for {
		if variable, ok := tmpEnv.variable[id]; ok {
			fmt.Println(id)
			fmt.Println("El id es " + tmpEnv.Control.Id)
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
	err := interfaces.Errores{Line: line, Col: col, Mess: "LA VARIABLE \"" + id + "\" NO EXISTE", Fecha: fecha}
	env.Errores(err)
	tipos := interfaces.TipoSimbolo{Tipo: interfaces.NULL}
	return interfaces.Symbol{Id: "", Tipo: tipos, Posicion: "0"}
}

func (env Environment) GetStruct(id, line, col string) (bool, interfaces.Structs) {
	var tmpEnv Environment
	tmpEnv = env

	for {
		if tmpStruct, ok := tmpEnv.structs[id]; ok {
			return true, tmpStruct
		}

		if tmpEnv.father == nil {
			break
		} else {
			tmpEnv = tmpEnv.father.(Environment)
		}
	}
	env.NewError("EL STRUCT \""+id+"\" NO FUE ENCONTRADO", line, col)
	return false, interfaces.Structs{}
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

func (env Environment) DevAmbito() bool {
	if env.Control.Id == "main" {
		return false
	}
	return true
}

func (env Environment) NewError(mess, line, col string) {
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	err := interfaces.Errores{Line: line, Col: col, Mess: mess, Fecha: fecha}
	env.Errores(err)
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
