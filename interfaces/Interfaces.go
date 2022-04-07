package interfaces

import (
	"OLC2/generator"

	"github.com/colegno/arraylist"
)

type Symbol struct {
	Id       string
	Tipo     TipoSimbolo
	Posicion int
	Mutable  bool
}

type Value struct {
	Value      string
	IsTemp     bool
	Type       TipoExpresion
	Tipo2      *arraylist.List
	TrueLabel  string
	FalseLabel string
}

type Tablas struct {
	Nombre string
	Father string
	Linea  string
}

type Bases struct {
	Nombre   string
	NoTables string
	Linea    string
}

type Optimizacion struct {
	Tipo          string
	Regla         string
	Expr_Original string
	Expr_Optima   string
	Fila          string
}

type Simbolos struct {
	ID      string
	Tipo    string
	Ambito  string
	Fila    string
	Columna string
}

type Errores struct {
	Line  string
	Col   string
	Mess  string
	Fecha string
}

type Expresion interface {
	Ejecutar(env interface{}, gen *generator.Generator) Value
}

type Instruction interface {
	Ejecutar(env interface{}, gen *generator.Generator) interface{}
}

type TipoSimbolo struct {
	Tipo  TipoExpresion
	Tipo2 *arraylist.List
}

type Dimensions struct {
	Tipo       TipoExpresion
	Dimensions *arraylist.List
}
