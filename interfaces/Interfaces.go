package interfaces

import "OLC2/generator"

type Symbol struct {
	Id       string
	Tipo     TipoExpresion
	Tipo2    TipoExpresion
	Posicion int
}

type Value struct {
	Value      string
	IsTemp     bool
	Type       TipoExpresion
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
	Tipo2   string
	Ambito  string
	Fila    string
	Columna string
}

type Errores struct {
	Line  string
	Col   string
	Tipo  string
	Mess  string
	Ambit string
	Fecha string
}

type Expresion interface {
	Ejecutar(env interface{}, gen *generator.Generator) Value
}

type Instruction interface {
	Ejecutar(env interface{}, gen *generator.Generator) interface{}
}
