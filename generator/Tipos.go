package generator

//0
type Operacion struct { //0
	Destino         string
	Op1             string
	Signo           string
	Op2             string
	Comentario      bool
	Comentario_Cont string
}

//1
type Asignacion struct { //1
	Destino         string
	Op              string
	Comentario      bool
	Comentario_Cont string
}

//2
type If struct { //2
	Op1             string
	Signo           string
	Op2             string
	Salto           string
	Comentario      bool
	Comentario_Cont string
}

//3
type Heap struct { //3
	Index           string
	Value           string
	Comentario      bool
	Comentario_Cont string
}

//4
type Stack struct { //4
	Index           string
	Value           string
	Comentario      bool
	Comentario_Cont string
}

//5
type Comentario struct {
	Comentario string
}

//6
type Llamada struct { //6
	Nombre          string
	Comentario      bool
	Comentario_Cont string
}

//7
type Parametros struct { //7
	Index           string
	Valor           string
	Comentario      bool
	Comentario_Cont string
}

//8
type Label struct { //8
	Name            string
	Comentario      bool
	Comentario_Cont string
}

//9
type Salto struct { //9
	Label           string
	Comentario      bool
	Comentario_Cont string
}

//10
type Return struct { //10
	Valor           string
	Comentario      bool
	Comentario_Cont string
}

//11
type Break struct { //11
	Valor           bool
	Value           string
	Comentario      bool
	Comentario_Cont string
}

//12
type Varios struct { //12
	Valor           string
	Comentario      bool
	Comentario_Cont string
}

//13
type Obligatorio struct { //13
	Valor           string
	Comentario      bool
	Comentario_Cont string
}

//14
type Print struct {
	Tipo  string //solo c,f,d
	Valor string
}
