package expresion

import (
	"OLC2/generator"
	"OLC2/interfaces"
	//"fmt"
)

type DevLoop struct {
	loop interfaces.Instruction
}

func NewDevLoop(loop interfaces.Instruction) DevLoop {
	expr := DevLoop{loop: loop}
	return expr
}

func (l DevLoop) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	res := l.loop.Ejecutar(env, gen).(interfaces.Value)
	/*fmt.Println("//////")
	fmt.Println("EL RET EN EXPR")
	fmt.Println(res)
	fmt.Println("//////")*/
	if res.Type != interfaces.NULL {
		retorno = res
	}
	return retorno
}
