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
  |asignacion_var PTCOMA  {$instr=$asignacion_var.i}
  | P_WHILE  expression  LLAVEIZQ instrucciones LLAVEDER  {$instr = instruction.NewWhile($expression.p, $instrucciones.l,$LLAVEIZQ.GetLine(),$LLAVEDER.GetColumn())}
  |breaks {$instr=$breaks.i}
  |continues  {$instr=$continues.i}
  |ifs  {$instr=$ifs.p}
  |loops  {$instr=$loops.i}
  |impr PTCOMA{$instr=$impr.p}
  |matches    {$instr=$matches.m}
  |rfor {$instr=$rfor.p}
  |mod_Array PTCOMA {$instr=$mod_Array.p}
  |accesoArr PUNTO PUSH PARIZQ expression PARDER PTCOMA  {$instr=instruction.NewPush($accesoArr.p,$expression.p,$PARIZQ.GetLine(),$PARIZQ.GetColumn())}
  |accesoArr PUNTO INSERT PARIZQ exp1=expression COMA exp2=expression PARDER PTCOMA {$instr=instruction.NewInsert($accesoArr.p,$exp1.p,$exp2.p,$COMA.GetLine(),$COMA.GetColumn())}
  |funcs {$instr=$funcs.i}
  |t_struct {$instr=$t_struct.i}
  |llamadas PTCOMA{$instr=$llamadas.i}
;
instruccion_wc returns [interfaces.Instruction instr]:
  CONSOLE '.' LOG PARIZQ expression PARDER {$instr = instruction.NewImprimir($expression.p)}
  |declaracion_var  {$instr=$declaracion_var.i}
  |asignacion_var   {$instr=$asignacion_var.i}
  | P_WHILE  expression  LLAVEIZQ instrucciones LLAVEDER  {$instr = instruction.NewWhile($expression.p, $instrucciones.l,$LLAVEIZQ.GetLine(),$LLAVEDER.GetColumn())}
  |breaks {$instr=$breaks.i}
  |continues  {$instr=$continues.i}
  |ifs  {$instr=$ifs.p}
  |loops  {$instr=$loops.i}
  |impr {$instr=$impr.p}
  |rfor {$instr=$rfor.p}
  |mod_Array  {$instr=$mod_Array.p}
  |accesoArr PUNTO PUSH PARIZQ expression PARDER  {$instr=instruction.NewPush($accesoArr.p,$expression.p,$PARIZQ.GetLine(),$PARIZQ.GetColumn())}
;

value_ref returns [bool t]:
  AMPER MUT   {$t=true}
  |     {$t=false}
;

llamadas returns [interfaces.Instruction i]:
  id=ID PARIZQ p=param_call PARDER{
    $i=instruction.NewCallF($id.text,$p.l,$id.GetLine(),$id.GetColumn())    
  }
;

param_call returns [*arrayList.List l]:
  value_ref expression  {$l=arrayList.New()
                $l.Add($expression.p)
  }
  |p=param_call COMA value_ref expression {
              $p.l.Add($expression.p)
              $l=$p.l
  }
  |           {$l=arrayList.New()}
;

asignacion_var  returns [interfaces.Instruction i]:
  id=ID IGUAL expression {
    linea:=$ctx.id.GetLine()
    col:=$ctx.id.GetColumn()
    $i=instruction.NewAssignment($id.text,$expression.p,linea,col)}
;

declaracion_var returns [interfaces.Instruction i]:
  LET mutable id=ID types IGUAL expression {$i=instruction.NewDeclaracion($id.text,$types.l,$expression.p,$mutable.mut,$LET.GetLine(),$LET.GetColumn())}
;
mutable returns [bool mut]
  :MUT  {$mut=true}
  |     {$mut=false}
;
types returns [interfaces.TipoSimbolo l]:
  DOSPUNTOS tipo_vector {$l=$tipo_vector.t}
  |DOSPUNTOS  a=asignar_Array  {
    dim:=arrayList.New()
    dim.Add($a.d)
    $l=interfaces.TipoSimbolo{interfaces.ARRAY,dim}
  }
  |DOSPUNTOS tipo_d {$l=interfaces.TipoSimbolo{$tipo_d.t,arrayList.New()}}
  |DOSPUNTOS id=ID {
    dim:=arrayList.New()
    dim.Add($id.text)
    $l=interfaces.TipoSimbolo{interfaces.STRUCT,dim}
  }
;

tipo_d returns [interfaces.TipoExpresion t]:
  INT       {$t=interfaces.INTEGER}  
  |FLOAT    {$t=interfaces.FLOAT}
  |BOOL     {$t=interfaces.BOOLEAN}
  |CHAR     {$t=interfaces.CHAR}
  |STR      {$t=interfaces.STRING}
  |P_STRING {$t=interfaces.STR}
  |USIZE    {$t=interfaces.USIZE}
  |ID       {$t=interfaces.STRUCT}
;

asignar_Array returns [interfaces.Dimensions d]:
  dimensiones {$d=$dimensiones.d}
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
              $l.Add(interfaces.STRUCT)
              }
  |VECT MENOR vectores MAYOR{
    $vectores.l.Add(interfaces.VECTOR)
    $l=$vectores.l
  }
;

loops returns [interfaces.Instruction i]:
  LOOP  LLAVEIZQ instrucciones LLAVEDER {$i=instruction.NewLoop($instrucciones.l)}
;

ifs returns[interfaces.Instruction p]:
  P_IF  expression LLAVEIZQ  llaves_if LLAVEDER elses {$p=instruction.NewIf($expression.p,$llaves_if.l,$elses.e,$P_IF.GetLine(),$P_IF.GetColumn(),0)}
;

llaves_if returns[*arrayList.List l]:
  expression              {$l=arrayList.New()
                            i:=interfaces.OpcionIf{0,$expression.p}
                            $l.Add(i)}
  |instruccion            {$l=arrayList.New()
                            i:=interfaces.OpcionIf{1,$instruccion.instr}
                            $l.Add(i)}
  |k=llaves_if expression {
                            i:=interfaces.OpcionIf{0,$expression.p}
                            $k.l.Add(i)
                            $l=$k.l}
  |k=llaves_if instruccion{
                            i:=interfaces.OpcionIf{1,$instruccion.instr}
                            $k.l.Add(i)
                            $l=$k.l}
;

elses returns[interfaces.Instruction e]:
  P_ELSE LLAVEIZQ  llaves_if LLAVEDER {$e=instruction.NewIf(expresion.NewPrimitivo(1,interfaces.BOOLEAN,0,0),$llaves_if.l,instruction.NewElseNull("null"),0,0,3)}
  |P_ELSE P_IF  expression LLAVEIZQ  llaves_if LLAVEDER elses {$e=instruction.NewIf($expression.p,$llaves_if.l,$elses.e,$P_IF.GetLine(),$P_IF.GetColumn(),2)}
  |             {$e= instruction.NewElseNull("null")}
;

breaks returns [interfaces.Instruction i]:
  BREAK expression  PTCOMA  {$i=instruction.NewBreak($expression.p,true,$BREAK.GetLine(),$BREAK.GetColumn())}
  |BREAK PTCOMA {$i=instruction.NewBreak(expresion.NewPrimitivo(1,interfaces.INTEGER,0,0),false,$BREAK.GetLine(),$BREAK.GetColumn())}
;

continues returns [interfaces.Instruction i]:
  CONTINUE PTCOMA {$i=instruction.NewContinue("continue",$CONTINUE.GetLine(),$CONTINUE.GetColumn())}
;

impr returns[interfaces.Instruction p]:
  PRINTLN PARIZQ formato listValues PARDER   {$p=instruction.NewPrint($listValues.l,$formato.s,$PARIZQ.GetLine(),$PARIZQ.GetColumn())}
;

formato returns [string s]:
  STRING COMA {str:= $STRING.text[1:len($STRING.text)-1]
      $s=str
   }
  |         {$s=""}  
;
matches returns [interfaces.Instruction m]:
  MATCH expression LLAVEIZQ  casos LLAVEDER {$m=instruction.NewMatch($expression.p,$casos.l,$MATCH.GetLine(),$MATCH.GetColumn())}
;
casos returns [*arrayList.List l]:
  cs=casos cases      {$cs.l.Add($cases.c)
                        $l=$cs.l}
  |cs=casos defaults  {$cs.l.Add($defaults.c)
                        $l=$cs.l}
  |cases      {$l=arrayList.New()
                $l.Add($cases.c)}
  |defaults   {$l=arrayList.New()
                $l.Add($defaults.c)}    
;
cases returns [interfaces.Cases c]:
  conditions OPMATCH set_match {$c=interfaces.Cases{$conditions.l,$set_match.cs.Cuerpo,$set_match.cs.Retorno,0,false}}
;
conditions returns [*arrayList.List l]:
 cond=conditions OR_COND expression {$cond.l.Add($expression.p)
                                      $l=$cond.l}               
  |expression                       {$l=arrayList.New() 
                                      $l.Add($expression.p) }
;
defaults returns [interfaces.Cases c]:
  DEF OPMATCH set_match {$c=interfaces.Cases{$set_match.cs.Condicion,$set_match.cs.Cuerpo,$set_match.cs.Retorno,1,false}}
;
set_match returns [interfaces.Cases cs]:
  expression COMA      {arr:=arrayList.New()
    arr.Add(instruction.NewElseNull("null"))
    $cs=interfaces.Cases{ arrayList.New(),arr,$expression.p,0,false}}
  |LLAVEIZQ instrucciones LLAVEDER  {$cs=interfaces.Cases{ arrayList.New(),$instrucciones.l,expresion.NewPrimitivo (1,interfaces.INTEGER,0,0),0,false}}
  |instruccion_wc COMA {arr:=arrayList.New()
    arr.Add($instruccion_wc.instr)
    $cs=interfaces.Cases{ arrayList.New(),arr,expresion.NewPrimitivo (1,interfaces.INTEGER,0,0),0,false}}
;

rfor returns[interfaces.Instruction p]:
  P_FOR id=ID P_IN iter_for LLAVEIZQ instrucciones LLAVEDER {$p=instruction.NewFor($instrucciones.l,$id.text,$iter_for.p,$P_FOR.GetLine(),$P_FOR.GetColumn())}
;
iter_for returns[interfaces.For_Range p]:
  exp1=expression PUNTO PUNTO exp2=expression{$p=interfaces.For_Range{$exp1.p,$exp2.p,0}}
  |exp1=expression {$p=interfaces.For_Range{$exp1.p,expresion.NewPrimitivo (1,interfaces.INTEGER,0,0),1}}
;
mod_Array returns[interfaces.Instruction p]:
  exp1=accesoArr IGUAL expression {$p=instruction.NewModArray($exp1.p,$expression.p,$IGUAL.GetLine(),$IGUAL.GetColumn())} 
;

funcs returns [interfaces.Instruction i]:
  FUNCTION id=ID PARIZQ p=param_dec PARDER tr=type_ret LLAVEIZQ instrucciones LLAVEDER{
    $i=instruction.NewFunctions($id.text,$tr.l,$p.l,$instrucciones.l,$FUNCTION.GetLine(),$FUNCTION.GetColumn())
  }
;

type_ret  returns [interfaces.TipoSimbolo l]:
  SUB MAYOR type_func {$l=$type_func.l}
  |                   {$l=interfaces.TipoSimbolo{interfaces.NULL,arrayList.New()}}
;

param_dec returns [*arrayList.List l]:
  id=ID t=type_param  {
    par:=interfaces.Parametros{$id.text,$t.t}
    $l=arrayList.New()
    $l.Add(par)}
  |p=param_dec COMA id=ID t=type_param  {
    par:=interfaces.Parametros{$id.text,$t.t}  
    $p.l.Add(par)
    $l=$p.l
  }
  |     {$l=arrayList.New()}
;

type_param returns [interfaces.TipoSimbolo t]:
  DOSPUNTOS AMPER MUT tipo_vector {$t=$tipo_vector.t}
  |DOSPUNTOS AMPER MUT a=asignar_Array  {
    dim:=arrayList.New()
    dim.Add($a.d)
    $t=interfaces.TipoSimbolo{interfaces.ARRAY,dim}
  }
  |DOSPUNTOS tipo_d {$t=interfaces.TipoSimbolo{$tipo_d.t,arrayList.New()}} 
  |DOSPUNTOS AMPER MUT id=ID {
    dim:=arrayList.New()
    dim.Add($id.text)
    $t=interfaces.TipoSimbolo{interfaces.STRUCT,dim}
  }
;

type_func returns [interfaces.TipoSimbolo l]:
  tipo_vector {$l=$tipo_vector.t}
  |a=asignar_Array  {
    dim:=arrayList.New()
    dim.Add($a.d)
    $l=interfaces.TipoSimbolo{interfaces.ARRAY,dim}
  }
  |tipo_d {$l=interfaces.TipoSimbolo{$tipo_d.t,arrayList.New()}}
  |id=ID {
    dim:=arrayList.New()
    dim.Add($id.text)
    $l=interfaces.TipoSimbolo{interfaces.STRUCT,dim}
  }
;

t_struct returns [interfaces.Instruction i]:
  P_STRUCT id=ID LLAVEIZQ lista=lista_att LLAVEDER  {$i = instruction.NewStruct($id.text, $lista.l,$P_STRUCT.GetLine(),$P_STRUCT.GetColumn())}
;

lista_att returns [*arrayList.List l]:
  list=lista_att COMA id=ID DOSPUNTOS type_func{
    att:=interfaces.Atributos{$id.text,$type_func.l}
    $list.l.Add(att);
    $l = $list.l;
  }
  |id=ID DOSPUNTOS type_func{
    att:=interfaces.Atributos{$id.text,$type_func.l}
    $l = arrayList.New();
    $l.Add(att);
  }
;

//EXPRESIONES
expression returns[interfaces.Expresion p]
    : expr_arit    {$p = $expr_arit.p}
    |id=ID LLAVEIZQ listAtt LLAVEDER  {$p=expresion.NewStruct($id.text,$listAtt.l,$LLAVEIZQ.GetLine(),$LLAVEIZQ.GetColumn())}
;

expr_arit returns[interfaces.Expresion p]
    : INT DDPUNTO POW PARIZQ opIz=expr_arit COMA opDe=expr_arit PARDER {$p = expresion.NewOperacion($opIz.p,"^",$opDe.p,false,$PARIZQ.GetLine(),$PARIZQ.GetColumn())}
    |FLOAT DDPUNTO POWF PARIZQ opIz=expr_arit COMA opDe=expr_arit PARDER {$p = expresion.NewOperacion($opIz.p,"^",$opDe.p,false,$PARIZQ.GetLine(),$PARIZQ.GetColumn())}
    | exp=expr_arit P_AS tipo_d  {$p=expresion.NewCast($exp.p,$tipo_d.t)} 
    |exp=expr_arit PUNTO T_STRING PARIZQ PARDER  {$p=expresion.NewToString($exp.p,$PUNTO.GetLine(),$PUNTO.GetColumn())}
    |DIFERENTE  opIz = expr_arit  {$p = expresion.NewOperacion($opIz.p,"!",$opIz.p,true,$DIFERENTE.GetLine(),$DIFERENTE.GetColumn())}
    |SUB  opIz = expr_arit  {$p = expresion.NewOperacion($opIz.p,"°",$opIz.p,true,$SUB.GetLine(),$SUB.GetColumn())}
    | opIz = expr_arit op=('*'|'/'|'%') opDe = expr_arit {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false,$op.GetLine(),$op.GetColumn())}
    | opIz = expr_arit op=('+'|'-') opDe = expr_arit {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false,$op.GetLine(),$op.GetColumn())}     
    | opIz = expr_arit op=('=='|'!='|'<'|'<='|'>='|'>') opDe = expr_arit {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false,$op.GetLine(),$op.GetColumn())}     
    | opIz = expr_arit op=('||'|'&&') opDe = expr_arit {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false,$op.GetLine(),$op.GetColumn())} 
    | primitivo {$p = $primitivo.p} 
    | PARIZQ expression PARDER {$p = $expression.p}
    | loops {$p=expresion.NewDevLoop($loops.i)}
    | ifs   {$p=expresion.NewDevLoop($ifs.p)}
    |matches {$p=expresion.NewDevLoop($matches.m)}
    | CORIZQ listValues CORDER { $p = expresion.NewArray($listValues.l) }
    |acc=accesoArr {$p=$acc.p}
    |creatArray {$p=$creatArray.p}
    |VEC CORIZQ listValues CORDER   {$p=expresion.NewVectorB($listValues.l)}
    |VEC CORIZQ exp1=expression PTCOMA exp2=expression CORDER   {$p=expresion.NewVectorA($exp1.p,$exp2.p)}
    |VECT DDPUNTO CAP PARIZQ expression PARDER  {$p=expresion.New_CapacityV($expression.p)}
    |VECT DDPUNTO NEW PARIZQ PARDER {$p=expresion.NewVectorC()}    
    |id=accesoArr PUNTO CAPACITY PARIZQ PARDER {$p=expresion.NewLen($id.p,$PUNTO.GetLine(),$PUNTO.GetColumn())}
    |id=accesoArr PUNTO LEN PARIZQ PARDER {$p=expresion.NewLenT($id.p,$PUNTO.GetLine(),$PUNTO.GetColumn())}
    |exp=expr_arit PUNTO F_ABS PARIZQ PARDER {$p=expresion.NewAbsoluto($exp.p,$PUNTO.GetLine(),$PUNTO.GetColumn())}
    |accesoArr PUNTO REMOVE PARIZQ expression PARDER PTCOMA  {$p=expresion.NewRemove($accesoArr.p,$expression.p,$PARIZQ.GetLine(),$PARIZQ.GetColumn())}
    |accesoArr PUNTO CONTAINS PARIZQ AMPER expression  PARDER  {$p=expresion.NewContains($accesoArr.p,$expression.p,$PUNTO.GetLine(),$PUNTO.GetColumn())}    
    |llamadas {$p=expresion.NewDevLoop($llamadas.i)}
;

accesoArr returns[interfaces.Expresion p]:
     id=ID  {$p=expresion.NewCallVariable($id.text,$id.GetLine(),$id.GetColumn())}
     |a=accesoArr PUNTO id=ID {$p=expresion.NewStructAccess($a.p,$id.text,$PUNTO.GetLine(),$PUNTO.GetColumn())}
     |a=accesoArr list=listArray {$p=expresion.NewArrayAccess($a.p,$list.l,0,0)}
;

primitivo returns[interfaces.Expresion p]
    :NUMBER {
            	num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }
            linea:=$ctx._NUMBER.GetLine()
            col:=$ctx._NUMBER.GetColumn()
            $p = expresion.NewPrimitivo (num,interfaces.INTEGER,col,linea)
            
       } 
    | STRING { 
      linea:=$ctx._STRING.GetLine()
      col:=$ctx._STRING.GetColumn()
      str:= $STRING.text[1:len($STRING.text)-1]
      $p = expresion.NewPrimitivo(str,interfaces.STRING,col,linea)}
    |TRUE{            
        $p=expresion.NewPrimitivo(1,interfaces.BOOLEAN,$TRUE.GetColumn(),$TRUE.GetLine())      
    }              
    |FALSE{
        $p=expresion.NewPrimitivo(0,interfaces.BOOLEAN,$FALSE.GetColumn(),$FALSE.GetLine())      
    }
    |DECIMAL{
            linea:=$ctx._DECIMAL.GetLine()
            col:=$ctx._DECIMAL.GetColumn()
              num,err:=strconv.ParseFloat($DECIMAL.text,64)
              if err!=nil{ 
                fmt.Println(err)
              }
              $p=expresion.NewPrimitivo(num,interfaces.FLOAT,col,linea)}  
    |CARACTER{
      linea:=$ctx._CARACTER.GetLine()
      col:=$ctx._CARACTER.GetColumn()
      str:= $CARACTER.text[1:len($CARACTER.text)-1]
      $p = expresion.NewPrimitivo(str,interfaces.CHAR,col,linea)
    }     
    |id=ID{
      linea:=$ctx.id.GetLine()
      col:=$ctx.id.GetColumn()
      $p=expresion.NewCallVariable($id.text,linea,col)
    }
;

listValues returns[*arrayList.List l]
    : list=listValues ',' expression { 
                                        $list.l.Add($expression.p)
                                        $l = $list.l
                                    }
    | expression { 
                    $l = arrayList.New()
                    $l.Add($expression.p)
                }
;

listAtt returns [*arrayList.List l]:
  id=ID DOSPUNTOS expression  {
    $l=arrayList.New()
    att:=interfaces.Atributo{$id.text,$expression.p}
    $l.Add(att)
  }
  |list=listAtt COMA id=ID DOSPUNTOS expression{
    att:=interfaces.Atributo{$id.text,$expression.p}
    $list.l.Add(att)
    $l=$list.l
  }
;


listArray returns[*arrayList.List l]:
  CORIZQ expression CORDER {
    $l=arrayList.New()
    $l.Add($expression.p)
  }
  |lista=listArray CORIZQ expression CORDER{
    $lista.l.Add($expression.p)
    $l=$lista.l
  }
;

creatArray returns [interfaces.Expresion p]:
  CORIZQ exp1=expression PTCOMA exp2=expression CORDER {$p=expresion.NewArray2($exp1.p,$exp2.p)}
;


