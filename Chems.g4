parser grammar Chems;



options {
  tokenVocab = ChemsLexer;
}

@header {
    import "OLC2/interfaces"
    import "OLC2/expresion"
    import "OLC2/instruction"
    import arrayList "github.com/colegno/arraylist"
}

start returns [*arrayList.List lista]
  : instrucciones {$lista = $instrucciones.l}
;

instrucciones returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e +=instruccion*  {
      listInt := localctx.(*InstruccionesContext).GetE()
      		for _, e := range listInt {
            $l.Add(e.GetInstr())
          }
    }
;


instruccion returns [interfaces.Instruction instr]
  : CONSOLE '.' LOG PARIZQ expression PARDER ';' {$instr = instruction.NewImprimir($expression.p)}
  |declaracion_var PTCOMA {$instr=$declaracion_var.i}
;

declaracion_var returns [interfaces.Instruction i]:
  LET mutable id=ID types IGUAL expression {$i=instruction.NewDeclaracion($id.text,$types.l,$expression.p,$mutable.mut)}
;
mutable returns [bool mut]
  :MUT  {$mut=true}
  |     {$mut=false}
;
types returns [interfaces.TipoSimbolo l]:
  DOSPUNTOS tipo_vector {$l=$tipo_vector.t}
  |a=asignar_Array  {
    dim:=arrayList.New()
    dim.Add($a.d)
    $l=interfaces.TipoSimbolo{interfaces.ARRAY,dim}
  }
  |DOSPUNTOS tipo_d {$l=interfaces.TipoSimbolo{$tipo_d.t,arrayList.New()}}
;

tipo_d returns [interfaces.TipoExpresion t]:
  INT       {$t=interfaces.INTEGER}  
  |FLOAT    {$t=interfaces.FLOAT}
  |BOOL     {$t=interfaces.BOOLEAN}
  |CHAR     {$t=interfaces.CHAR}
  |STR      {$t=interfaces.STRING}
  |P_STRING {$t=interfaces.STR}
  |USIZE    {$t=interfaces.ARRAY}
  |ID       {$t=interfaces.STRUCT}
;

asignar_Array returns [interfaces.Dimensions d]:
  DOSPUNTOS dimensiones {$d=$dimensiones.d}
  //|           {$d=interfaces.Dimensions{interfaces.ALL,arrayList.New()}}
;
dimensiones  returns [interfaces.Dimensions d]:
  CORIZQ tipo_d PTCOMA expression CORDER  {
    list:=arrayList.New()
    list.Add($expression.p)
    $d=interfaces.Dimensions{$tipo_d.t,list}}
  |CORIZQ dimensiones PTCOMA expression CORDER  {$dimensiones.d.Dimensions.Add($expression.p)
                                                  $d=$dimensiones.d;
                                                } 
;

tipo_vector returns [interfaces.TipoSimbolo t]:
  VECT MENOR vectores MAYOR {$t=interfaces.TipoSimbolo{interfaces.VECTOR,$vectores.l}}
;
vectores returns [*arrayList.List l]:
  INT      {$l=arrayList.New()
            $l.Add(interfaces.INTEGER)}  
  |FLOAT    {$l=arrayList.New()
              $l.Add(interfaces.FLOAT)}  
  |BOOL     {$l=arrayList.New()
              $l.Add(interfaces.BOOLEAN)}  
  |CHAR     {$l=arrayList.New()
              $l.Add(interfaces.CHAR)}  
  |STR      {$l=arrayList.New()
              $l.Add(interfaces.STRING)}  
  |P_STRING {$l=arrayList.New()
              $l.Add(interfaces.STR)}  
  |USIZE    {$l=arrayList.New()
              $l.Add(interfaces.USIZE)}  
  |id=ID       {$l=arrayList.New()
              $l.Add($id.text)
              $l.Add(interfaces.STRUCT)}
  |VECT MENOR vectores MAYOR{
    $vectores.l.Add(interfaces.VECTOR)
    $l=$vectores.l
  }
;




//EXPRESIONES
expression returns[interfaces.Expresion p]
    : expr_arit    {$p = $expr_arit.p}
;

expr_arit returns[interfaces.Expresion p]
    : opIz = expr_arit op=('*'|'/') opDe = expr_arit {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false)}
    | opIz = expr_arit op=('+'|'-') opDe = expr_arit {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false)}     
    | opIz = expr_arit op=('<'|'<='|'>='|'>') opDe = expr_arit {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false)}     
    | primitivo {$p = $primitivo.p} 
    | PARIZQ expression PARDER {$p = $expression.p}
    | exp=expr_arit P_AS tipo_d  {$p=expresion.NewCast($exp.p,$tipo_d.t)
      fmt.Println("line")
    } 
;

primitivo returns[interfaces.Expresion p]
    :NUMBER {
            	num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expresion.NewPrimitivo (num,interfaces.INTEGER)
            linea:=$ctx._NUMBER.GetLine()
            col:=$ctx._NUMBER.GetColumn()
            fmt.Println(linea)
            fmt.Println(col)
       } 
    | STRING { 
      str:= $STRING.text[1:len($STRING.text)-1]
      $p = expresion.NewPrimitivo(str,interfaces.STRING)}
    |TRUE{
        $p=expresion.NewPrimitivo(true,interfaces.BOOLEAN)      
    }              
    |FALSE{
        $p=expresion.NewPrimitivo(false,interfaces.BOOLEAN)      
    }
    |DECIMAL{
              num,err:=strconv.ParseFloat($DECIMAL.text,64)
              if err!=nil{ 
                fmt.Println(err)
              }
              $p=expresion.NewPrimitivo(num,interfaces.FLOAT)}  
    |CARACTER{
      str:= $CARACTER.text[1:len($CARACTER.text)-1]
      $p = expresion.NewPrimitivo(str,interfaces.CHAR)
    }     
    |id=ID{
      $p=expresion.NewCallVariable($id.text)
    }
;