package interfaces

type TipoExpresion int

const (
	INTEGER TipoExpresion = iota
	FLOAT
	STRING
	STR
	BOOLEAN
	CHAR
	USIZE
	VECTOR
	ARRAY
	STRUCT
	VOID
	NULL
	ERROR
)
