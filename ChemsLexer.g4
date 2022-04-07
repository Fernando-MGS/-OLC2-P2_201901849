lexer grammar ChemsLexer;
// Tokens

PUNTO:          '.';
PTCOMA:         ';';
COMA:           ',';
COMILLAS:       '"';
DIFERENTE:      '!';
DOSPUNTOS:      ':';
IGUAL:          '=';
AMPER:          '&';
MENOR:          '<';
MAYORIGUAL:     '>=';
MENORIGUAL:     '<=';
OPMATCH:        '=>';
D_IGUAL:        '==';
NOT_E:          '!=';
MAYOR:          '>';
OR:             '||';
AND:            '&&';
MUL:            '*';
DIV:            '/';
ADD:            '+';
SUB:            '-';
MOD:            '%';
POT:            '^';
PARIZQ:         '(';
PARDER:         ')';
LLAVEIZQ:       '{';
LLAVEDER:       '}';
CORIZQ:         '[';
CORDER:         ']';
INTDER:         '?';
OR_COND:        '|';
DDPUNTO:        '::';

CONSOLE:    'console';
LOG:        'log';
PRINTLN:    'println!';
P_NUMBER:   'number';
P_STRING:   'String';
P_IF:       'if';
P_ELSE:     'else';
P_WHILE:    'while';
P_FOR:      'for';
P_IN:       'in';
P_AS:       'as';
POWF:       'powf';
POW:        'pow';
BREAK:      'break';
CONTINUE:   'continue';
RETURN:     'return';
LET:        'let';
INT:        'i64';
FLOAT:      'f64';
BOOL:       'bool';
CHAR:       'char';
STR:        '&str';
USIZE:      'usize';
MUT:        'mut';
P_STRUCT:   'struct';
FUNCTION:   'fn';
MODULE:     'mod';
PUB:        'pub';
TRUE:       'true';
FALSE:      'false';
MATCH:      'match';
LOOP:       'loop';
PUSH:       'push';
T_STRING:   'to_string';
F_ABS:        'abs';
F_SQRT:       'sqrt';
F_CLONE:      'clone';
INSERT:     'insert';
REMOVE:     'remove';
CONTAINS:   'contains';
LEN:        'len';
CAPACITY:   'capacity';
DEF:        '_';
VEC:        'vec!';
VECT:       'Vec';
CAP:        'with_capacity';
NEW:        'new';


COMENTARIO: ('/')('/')(.)*? '\n'->skip;
NUMBER: [0-9]+;
DECIMAL: [0-9]+[.]([0-9]+);
STRING: '"'~["]*'"';
ID: ([a-zA-Z_])[a-zA-Z0-9_]*;
CARACTER: ['][\\]?.['];

WHITESPACE: [ \\\r\n\t]+ -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;
