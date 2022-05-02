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
	//Valor           bool
	Label           string
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
	Tipo            string //solo c,f,d
	Valor           string
	Comentario      bool
	Comentario_Cont string
}

//15
type CallStack struct {
	Destino         string
	Index           string
	Comentario      bool
	Comentario_Cont string
}

//16
type CallHeap struct {
	Destino         string
	Index           string
	Comentario      bool
	Comentario_Cont string
}

//17
type Fmod struct {
	Destino         string
	Div1            string
	Div2            string
	Comentario      bool
	Comentario_Cont string
}

func (p Operacion) Ejecutar() string {
	code := p.Destino + "=" + p.Op1 + p.Signo + p.Op2 + ";"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p Asignacion) Ejecutar() string {
	code := p.Destino + "=" + p.Op + ";"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p If) Ejecutar() string {
	code := "if(" + p.Op1 + p.Signo + p.Op2 + ") goto " + p.Salto + ";"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p Heap) Ejecutar() string {
	code := "HEAP[(int)" + p.Index + "]=" + p.Value + ";"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p Stack) Ejecutar() string {
	code := "STACK[(int)" + p.Index + "]=" + p.Value + ";"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p Comentario) Ejecutar() string {
	code := "//" + p.Comentario
	return code
}

func (p Llamada) Ejecutar() string {
	code := p.Nombre + "();"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p Parametros) Ejecutar() string {
	code := "STACK[(int)+" + p.Index + "]=" + p.Valor + ";"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p Label) Ejecutar() string {
	code := p.Name + ":"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p Salto) Ejecutar() string {
	code := "goto " + p.Label + ";"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p Return) Ejecutar() string {
	code := "();"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p Break) Ejecutar() string {
	code := "goto " + p.Label + ";"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p Varios) Ejecutar() string {
	code := p.Valor
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p Obligatorio) Ejecutar() string {
	code := p.Valor
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p Print) Ejecutar() string {
	code := "printf(\"%" + p.Tipo + "\"," + p.Valor + ");"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p CallStack) Ejecutar() string {
	code := p.Destino + "=STACK[(int)" + p.Index + "];"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p CallHeap) Ejecutar() string {
	code := p.Destino + "=HEAP[(int)" + p.Index + "];"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}

func (p Fmod) Ejecutar() string {
	code := p.Destino + "=fmod(" + p.Div1 + "," + p.Div2 + ");"
	if p.Comentario {
		code += "//" + p.Comentario_Cont
	}
	return code
}
