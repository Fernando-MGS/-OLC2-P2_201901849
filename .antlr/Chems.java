// Generated from c:\Users\ferna\OneDrive\Documentos\Go\src\github.com\Fernando-MGS\-OLC2-P2_201901849\Chems.g4 by ANTLR 4.8

    import "OLC2/interfaces"
    import "OLC2/expresion"
    import "OLC2/instruction"
    import arrayList "github.com/colegno/arraylist"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Chems extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PUNTO=1, PTCOMA=2, COMA=3, COMILLAS=4, DIFERENTE=5, DOSPUNTOS=6, IGUAL=7, 
		AMPER=8, MENOR=9, MAYORIGUAL=10, MENORIGUAL=11, OPMATCH=12, D_IGUAL=13, 
		NOT_E=14, MAYOR=15, OR=16, AND=17, MUL=18, DIV=19, ADD=20, SUB=21, MOD=22, 
		POT=23, PARIZQ=24, PARDER=25, LLAVEIZQ=26, LLAVEDER=27, CORIZQ=28, CORDER=29, 
		INTDER=30, OR_COND=31, DDPUNTO=32, CONSOLE=33, LOG=34, PRINTLN=35, P_NUMBER=36, 
		P_STRING=37, P_IF=38, P_ELSE=39, P_WHILE=40, P_FOR=41, P_IN=42, P_AS=43, 
		POWF=44, POW=45, BREAK=46, CONTINUE=47, RETURN=48, LET=49, INT=50, FLOAT=51, 
		BOOL=52, CHAR=53, STR=54, USIZE=55, MUT=56, P_STRUCT=57, FUNCTION=58, 
		MODULE=59, PUB=60, TRUE=61, FALSE=62, MATCH=63, LOOP=64, PUSH=65, T_STRING=66, 
		F_ABS=67, F_SQRT=68, F_CLONE=69, INSERT=70, REMOVE=71, CONTAINS=72, LEN=73, 
		CAPACITY=74, DEF=75, VEC=76, VECT=77, CAP=78, NEW=79, COMENTARIO=80, NUMBER=81, 
		DECIMAL=82, STRING=83, ID=84, CARACTER=85, WHITESPACE=86;
	public static final int
		RULE_start = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_asignacion_var = 3, 
		RULE_declaracion_var = 4, RULE_mutable = 5, RULE_types = 6, RULE_tipo_d = 7, 
		RULE_asignar_Array = 8, RULE_dimensiones = 9, RULE_tipo_vector = 10, RULE_vectores = 11, 
		RULE_ifs = 12, RULE_elses = 13, RULE_breaks = 14, RULE_continues = 15, 
		RULE_expression = 16, RULE_expr_arit = 17, RULE_primitivo = 18;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instrucciones", "instruccion", "asignacion_var", "declaracion_var", 
			"mutable", "types", "tipo_d", "asignar_Array", "dimensiones", "tipo_vector", 
			"vectores", "ifs", "elses", "breaks", "continues", "expression", "expr_arit", 
			"primitivo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'.'", "';'", "','", "'\"'", "'!'", "':'", "'='", "'&'", "'<'", 
			"'>='", "'<='", "'=>'", "'=='", "'!='", "'>'", "'||'", "'&&'", "'*'", 
			"'/'", "'+'", "'-'", "'%'", "'^'", "'('", "')'", "'{'", "'}'", "'['", 
			"']'", "'?'", "'|'", "'::'", "'console'", "'log'", "'println!'", "'number'", 
			"'String'", "'if'", "'else'", "'while'", "'for'", "'in'", "'as'", "'powf'", 
			"'pow'", "'break'", "'continue'", "'return'", "'let'", "'i64'", "'f64'", 
			"'bool'", "'char'", "'&str'", "'usize'", "'mut'", "'struct'", "'fn'", 
			"'mod'", "'pub'", "'true'", "'false'", "'match'", "'loop'", "'push'", 
			"'to_string'", "'abs'", "'sqrt'", "'clone'", "'insert'", "'remove'", 
			"'contains'", "'len'", "'capacity'", "'_'", "'vec!'", "'Vec'", "'with_capacity'", 
			"'new'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PUNTO", "PTCOMA", "COMA", "COMILLAS", "DIFERENTE", "DOSPUNTOS", 
			"IGUAL", "AMPER", "MENOR", "MAYORIGUAL", "MENORIGUAL", "OPMATCH", "D_IGUAL", 
			"NOT_E", "MAYOR", "OR", "AND", "MUL", "DIV", "ADD", "SUB", "MOD", "POT", 
			"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORIZQ", "CORDER", "INTDER", 
			"OR_COND", "DDPUNTO", "CONSOLE", "LOG", "PRINTLN", "P_NUMBER", "P_STRING", 
			"P_IF", "P_ELSE", "P_WHILE", "P_FOR", "P_IN", "P_AS", "POWF", "POW", 
			"BREAK", "CONTINUE", "RETURN", "LET", "INT", "FLOAT", "BOOL", "CHAR", 
			"STR", "USIZE", "MUT", "P_STRUCT", "FUNCTION", "MODULE", "PUB", "TRUE", 
			"FALSE", "MATCH", "LOOP", "PUSH", "T_STRING", "F_ABS", "F_SQRT", "F_CLONE", 
			"INSERT", "REMOVE", "CONTAINS", "LEN", "CAPACITY", "DEF", "VEC", "VECT", 
			"CAP", "NEW", "COMENTARIO", "NUMBER", "DECIMAL", "STRING", "ID", "CARACTER", 
			"WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Chems.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Chems(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public *arrayList.List lista;
		public InstruccionesContext instrucciones;
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(38);
			((StartContext)_localctx).instrucciones = instrucciones();
			_localctx.lista = ((StartContext)_localctx).instrucciones.l
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionesContext extends ParserRuleContext {
		public *arrayList.List l;
		public InstruccionContext instruccion;
		public List<InstruccionContext> e = new ArrayList<InstruccionContext>();
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_instrucciones);

		    _localctx.l =  arrayList.New()
		  
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 33)) & ~0x3f) == 0 && ((1L << (_la - 33)) & ((1L << (CONSOLE - 33)) | (1L << (P_IF - 33)) | (1L << (P_WHILE - 33)) | (1L << (BREAK - 33)) | (1L << (CONTINUE - 33)) | (1L << (LET - 33)) | (1L << (ID - 33)))) != 0)) {
				{
				{
				setState(41);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(46);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}

			      listInt := localctx.(*InstruccionesContext).GetE()
			      		for _, e := range listInt {
			            _localctx.l.Add(e.GetInstr())
			          }
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionContext extends ParserRuleContext {
		public interfaces.Instruction instr;
		public ExpressionContext expression;
		public Declaracion_varContext declaracion_var;
		public Asignacion_varContext asignacion_var;
		public Token LLAVEIZQ;
		public InstruccionesContext instrucciones;
		public Token LLAVEDER;
		public BreaksContext breaks;
		public ContinuesContext continues;
		public IfsContext ifs;
		public TerminalNode CONSOLE() { return getToken(Chems.CONSOLE, 0); }
		public TerminalNode PUNTO() { return getToken(Chems.PUNTO, 0); }
		public TerminalNode LOG() { return getToken(Chems.LOG, 0); }
		public TerminalNode PARIZQ() { return getToken(Chems.PARIZQ, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Chems.PARDER, 0); }
		public TerminalNode PTCOMA() { return getToken(Chems.PTCOMA, 0); }
		public Declaracion_varContext declaracion_var() {
			return getRuleContext(Declaracion_varContext.class,0);
		}
		public Asignacion_varContext asignacion_var() {
			return getRuleContext(Asignacion_varContext.class,0);
		}
		public TerminalNode P_WHILE() { return getToken(Chems.P_WHILE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Chems.LLAVEIZQ, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Chems.LLAVEDER, 0); }
		public BreaksContext breaks() {
			return getRuleContext(BreaksContext.class,0);
		}
		public ContinuesContext continues() {
			return getRuleContext(ContinuesContext.class,0);
		}
		public IfsContext ifs() {
			return getRuleContext(IfsContext.class,0);
		}
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruccion);
		try {
			setState(82);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CONSOLE:
				enterOuterAlt(_localctx, 1);
				{
				setState(49);
				match(CONSOLE);
				setState(50);
				match(PUNTO);
				setState(51);
				match(LOG);
				setState(52);
				match(PARIZQ);
				setState(53);
				((InstruccionContext)_localctx).expression = expression();
				setState(54);
				match(PARDER);
				setState(55);
				match(PTCOMA);
				_localctx.instr = instruction.NewImprimir(((InstruccionContext)_localctx).expression.p)
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(58);
				((InstruccionContext)_localctx).declaracion_var = declaracion_var();
				setState(59);
				match(PTCOMA);
				_localctx.instr=((InstruccionContext)_localctx).declaracion_var.i
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(62);
				((InstruccionContext)_localctx).asignacion_var = asignacion_var();
				setState(63);
				match(PTCOMA);
				_localctx.instr=((InstruccionContext)_localctx).asignacion_var.i
				}
				break;
			case P_WHILE:
				enterOuterAlt(_localctx, 4);
				{
				setState(66);
				match(P_WHILE);
				setState(67);
				((InstruccionContext)_localctx).expression = expression();
				setState(68);
				((InstruccionContext)_localctx).LLAVEIZQ = match(LLAVEIZQ);
				setState(69);
				((InstruccionContext)_localctx).instrucciones = instrucciones();
				setState(70);
				((InstruccionContext)_localctx).LLAVEDER = match(LLAVEDER);
				_localctx.instr = instruction.NewWhile(((InstruccionContext)_localctx).expression.p, ((InstruccionContext)_localctx).instrucciones.l,((InstruccionContext)_localctx).LLAVEIZQ.GetLine(),((InstruccionContext)_localctx).LLAVEDER.GetColumn())
				}
				break;
			case BREAK:
				enterOuterAlt(_localctx, 5);
				{
				setState(73);
				((InstruccionContext)_localctx).breaks = breaks();
				_localctx.instr=((InstruccionContext)_localctx).breaks.i
				}
				break;
			case CONTINUE:
				enterOuterAlt(_localctx, 6);
				{
				setState(76);
				((InstruccionContext)_localctx).continues = continues();
				_localctx.instr=((InstruccionContext)_localctx).continues.i
				}
				break;
			case P_IF:
				enterOuterAlt(_localctx, 7);
				{
				setState(79);
				((InstruccionContext)_localctx).ifs = ifs();
				_localctx.instr=((InstruccionContext)_localctx).ifs.p
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Asignacion_varContext extends ParserRuleContext {
		public interfaces.Instruction i;
		public Token id;
		public ExpressionContext expression;
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public Asignacion_varContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion_var; }
	}

	public final Asignacion_varContext asignacion_var() throws RecognitionException {
		Asignacion_varContext _localctx = new Asignacion_varContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_asignacion_var);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(84);
			((Asignacion_varContext)_localctx).id = match(ID);
			setState(85);
			match(IGUAL);
			setState(86);
			((Asignacion_varContext)_localctx).expression = expression();

			    linea:=_localctx.id.GetLine()
			    col:=_localctx.id.GetColumn()
			    _localctx.i=instruction.NewAssignment((((Asignacion_varContext)_localctx).id!=null?((Asignacion_varContext)_localctx).id.getText():null),((Asignacion_varContext)_localctx).expression.p,linea,col)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Declaracion_varContext extends ParserRuleContext {
		public interfaces.Instruction i;
		public Token LET;
		public MutableContext mutable;
		public Token id;
		public TypesContext types;
		public ExpressionContext expression;
		public TerminalNode LET() { return getToken(Chems.LET, 0); }
		public MutableContext mutable() {
			return getRuleContext(MutableContext.class,0);
		}
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public Declaracion_varContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion_var; }
	}

	public final Declaracion_varContext declaracion_var() throws RecognitionException {
		Declaracion_varContext _localctx = new Declaracion_varContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_declaracion_var);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			((Declaracion_varContext)_localctx).LET = match(LET);
			setState(90);
			((Declaracion_varContext)_localctx).mutable = mutable();
			setState(91);
			((Declaracion_varContext)_localctx).id = match(ID);
			setState(92);
			((Declaracion_varContext)_localctx).types = types();
			setState(93);
			match(IGUAL);
			setState(94);
			((Declaracion_varContext)_localctx).expression = expression();
			_localctx.i=instruction.NewDeclaracion((((Declaracion_varContext)_localctx).id!=null?((Declaracion_varContext)_localctx).id.getText():null),((Declaracion_varContext)_localctx).types.l,((Declaracion_varContext)_localctx).expression.p,((Declaracion_varContext)_localctx).mutable.mut,((Declaracion_varContext)_localctx).LET.GetLine(),((Declaracion_varContext)_localctx).LET.GetColumn())
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MutableContext extends ParserRuleContext {
		public bool mut;
		public TerminalNode MUT() { return getToken(Chems.MUT, 0); }
		public MutableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mutable; }
	}

	public final MutableContext mutable() throws RecognitionException {
		MutableContext _localctx = new MutableContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_mutable);
		try {
			setState(100);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MUT:
				enterOuterAlt(_localctx, 1);
				{
				setState(97);
				match(MUT);
				_localctx.mut=true
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 2);
				{
				_localctx.mut=false
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypesContext extends ParserRuleContext {
		public interfaces.TipoSimbolo l;
		public Tipo_vectorContext tipo_vector;
		public Asignar_ArrayContext a;
		public Tipo_dContext tipo_d;
		public TerminalNode DOSPUNTOS() { return getToken(Chems.DOSPUNTOS, 0); }
		public Tipo_vectorContext tipo_vector() {
			return getRuleContext(Tipo_vectorContext.class,0);
		}
		public Asignar_ArrayContext asignar_Array() {
			return getRuleContext(Asignar_ArrayContext.class,0);
		}
		public Tipo_dContext tipo_d() {
			return getRuleContext(Tipo_dContext.class,0);
		}
		public TypesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_types; }
	}

	public final TypesContext types() throws RecognitionException {
		TypesContext _localctx = new TypesContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_types);
		try {
			setState(113);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(102);
				match(DOSPUNTOS);
				setState(103);
				((TypesContext)_localctx).tipo_vector = tipo_vector();
				_localctx.l=((TypesContext)_localctx).tipo_vector.t
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(106);
				((TypesContext)_localctx).a = asignar_Array();

				    dim:=arrayList.New()
				    dim.Add(((TypesContext)_localctx).a.d)
				    _localctx.l=interfaces.TipoSimbolo{interfaces.ARRAY,dim}
				  
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(109);
				match(DOSPUNTOS);
				setState(110);
				((TypesContext)_localctx).tipo_d = tipo_d();
				_localctx.l=interfaces.TipoSimbolo{((TypesContext)_localctx).tipo_d.t,arrayList.New()}
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tipo_dContext extends ParserRuleContext {
		public interfaces.TipoExpresion t;
		public TerminalNode INT() { return getToken(Chems.INT, 0); }
		public TerminalNode FLOAT() { return getToken(Chems.FLOAT, 0); }
		public TerminalNode BOOL() { return getToken(Chems.BOOL, 0); }
		public TerminalNode CHAR() { return getToken(Chems.CHAR, 0); }
		public TerminalNode STR() { return getToken(Chems.STR, 0); }
		public TerminalNode P_STRING() { return getToken(Chems.P_STRING, 0); }
		public TerminalNode USIZE() { return getToken(Chems.USIZE, 0); }
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public Tipo_dContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo_d; }
	}

	public final Tipo_dContext tipo_d() throws RecognitionException {
		Tipo_dContext _localctx = new Tipo_dContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_tipo_d);
		try {
			setState(131);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(115);
				match(INT);
				_localctx.t=interfaces.INTEGER
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(117);
				match(FLOAT);
				_localctx.t=interfaces.FLOAT
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 3);
				{
				setState(119);
				match(BOOL);
				_localctx.t=interfaces.BOOLEAN
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 4);
				{
				setState(121);
				match(CHAR);
				_localctx.t=interfaces.CHAR
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 5);
				{
				setState(123);
				match(STR);
				_localctx.t=interfaces.STRING
				}
				break;
			case P_STRING:
				enterOuterAlt(_localctx, 6);
				{
				setState(125);
				match(P_STRING);
				_localctx.t=interfaces.STR
				}
				break;
			case USIZE:
				enterOuterAlt(_localctx, 7);
				{
				setState(127);
				match(USIZE);
				_localctx.t=interfaces.USIZE
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 8);
				{
				setState(129);
				match(ID);
				_localctx.t=interfaces.STRUCT
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Asignar_ArrayContext extends ParserRuleContext {
		public interfaces.Dimensions d;
		public DimensionesContext dimensiones;
		public TerminalNode DOSPUNTOS() { return getToken(Chems.DOSPUNTOS, 0); }
		public DimensionesContext dimensiones() {
			return getRuleContext(DimensionesContext.class,0);
		}
		public Asignar_ArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignar_Array; }
	}

	public final Asignar_ArrayContext asignar_Array() throws RecognitionException {
		Asignar_ArrayContext _localctx = new Asignar_ArrayContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_asignar_Array);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
			match(DOSPUNTOS);
			setState(134);
			((Asignar_ArrayContext)_localctx).dimensiones = dimensiones();
			_localctx.d=((Asignar_ArrayContext)_localctx).dimensiones.d
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DimensionesContext extends ParserRuleContext {
		public interfaces.Dimensions d;
		public Tipo_dContext tipo_d;
		public ExpressionContext expression;
		public DimensionesContext dimensiones;
		public TerminalNode CORIZQ() { return getToken(Chems.CORIZQ, 0); }
		public Tipo_dContext tipo_d() {
			return getRuleContext(Tipo_dContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Chems.PTCOMA, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(Chems.CORDER, 0); }
		public DimensionesContext dimensiones() {
			return getRuleContext(DimensionesContext.class,0);
		}
		public DimensionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dimensiones; }
	}

	public final DimensionesContext dimensiones() throws RecognitionException {
		DimensionesContext _localctx = new DimensionesContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_dimensiones);
		try {
			setState(151);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(137);
				match(CORIZQ);
				setState(138);
				((DimensionesContext)_localctx).tipo_d = tipo_d();
				setState(139);
				match(PTCOMA);
				setState(140);
				((DimensionesContext)_localctx).expression = expression();
				setState(141);
				match(CORDER);

				    list:=arrayList.New()
				    list.Add(((DimensionesContext)_localctx).expression.p)
				    _localctx.d=interfaces.Dimensions{((DimensionesContext)_localctx).tipo_d.t,list}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(144);
				match(CORIZQ);
				setState(145);
				((DimensionesContext)_localctx).dimensiones = dimensiones();
				setState(146);
				match(PTCOMA);
				setState(147);
				((DimensionesContext)_localctx).expression = expression();
				setState(148);
				match(CORDER);
				((DimensionesContext)_localctx).dimensiones.d.Dimensions.Add(((DimensionesContext)_localctx).expression.p)
				                                                  ((DimensionesContext)_localctx).d = ((DimensionesContext)_localctx).dimensiones.d;
				                                                
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tipo_vectorContext extends ParserRuleContext {
		public interfaces.TipoSimbolo t;
		public VectoresContext vectores;
		public TerminalNode VECT() { return getToken(Chems.VECT, 0); }
		public TerminalNode MENOR() { return getToken(Chems.MENOR, 0); }
		public VectoresContext vectores() {
			return getRuleContext(VectoresContext.class,0);
		}
		public TerminalNode MAYOR() { return getToken(Chems.MAYOR, 0); }
		public Tipo_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo_vector; }
	}

	public final Tipo_vectorContext tipo_vector() throws RecognitionException {
		Tipo_vectorContext _localctx = new Tipo_vectorContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_tipo_vector);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(153);
			match(VECT);
			setState(154);
			match(MENOR);
			setState(155);
			((Tipo_vectorContext)_localctx).vectores = vectores();
			setState(156);
			match(MAYOR);
			_localctx.t=interfaces.TipoSimbolo{interfaces.VECTOR,((Tipo_vectorContext)_localctx).vectores.l}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VectoresContext extends ParserRuleContext {
		public *arrayList.List l;
		public Token id;
		public VectoresContext vectores;
		public TerminalNode INT() { return getToken(Chems.INT, 0); }
		public TerminalNode FLOAT() { return getToken(Chems.FLOAT, 0); }
		public TerminalNode BOOL() { return getToken(Chems.BOOL, 0); }
		public TerminalNode CHAR() { return getToken(Chems.CHAR, 0); }
		public TerminalNode STR() { return getToken(Chems.STR, 0); }
		public TerminalNode P_STRING() { return getToken(Chems.P_STRING, 0); }
		public TerminalNode USIZE() { return getToken(Chems.USIZE, 0); }
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public TerminalNode VECT() { return getToken(Chems.VECT, 0); }
		public TerminalNode MENOR() { return getToken(Chems.MENOR, 0); }
		public VectoresContext vectores() {
			return getRuleContext(VectoresContext.class,0);
		}
		public TerminalNode MAYOR() { return getToken(Chems.MAYOR, 0); }
		public VectoresContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectores; }
	}

	public final VectoresContext vectores() throws RecognitionException {
		VectoresContext _localctx = new VectoresContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_vectores);
		try {
			setState(181);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(159);
				match(INT);
				_localctx.l=arrayList.New()
				            _localctx.l.Add(interfaces.INTEGER)
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(161);
				match(FLOAT);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.FLOAT)
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 3);
				{
				setState(163);
				match(BOOL);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.BOOLEAN)
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 4);
				{
				setState(165);
				match(CHAR);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.CHAR)
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 5);
				{
				setState(167);
				match(STR);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.STRING)
				}
				break;
			case P_STRING:
				enterOuterAlt(_localctx, 6);
				{
				setState(169);
				match(P_STRING);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.STR)
				}
				break;
			case USIZE:
				enterOuterAlt(_localctx, 7);
				{
				setState(171);
				match(USIZE);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.USIZE)
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 8);
				{
				setState(173);
				((VectoresContext)_localctx).id = match(ID);
				_localctx.l=arrayList.New()
				              _localctx.l.Add((((VectoresContext)_localctx).id!=null?((VectoresContext)_localctx).id.getText():null))
				              _localctx.l.Add(interfaces.STRUCT)
				}
				break;
			case VECT:
				enterOuterAlt(_localctx, 9);
				{
				setState(175);
				match(VECT);
				setState(176);
				match(MENOR);
				setState(177);
				((VectoresContext)_localctx).vectores = vectores();
				setState(178);
				match(MAYOR);

				    ((VectoresContext)_localctx).vectores.l.Add(interfaces.VECTOR)
				    _localctx.l=((VectoresContext)_localctx).vectores.l
				  
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IfsContext extends ParserRuleContext {
		public interfaces.Instruction p;
		public Token P_IF;
		public ExpressionContext expression;
		public InstruccionesContext instrucciones;
		public ElsesContext elses;
		public TerminalNode P_IF() { return getToken(Chems.P_IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(Chems.LLAVEIZQ, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Chems.LLAVEDER, 0); }
		public ElsesContext elses() {
			return getRuleContext(ElsesContext.class,0);
		}
		public IfsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifs; }
	}

	public final IfsContext ifs() throws RecognitionException {
		IfsContext _localctx = new IfsContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_ifs);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			((IfsContext)_localctx).P_IF = match(P_IF);
			setState(184);
			((IfsContext)_localctx).expression = expression();
			setState(185);
			match(LLAVEIZQ);
			setState(186);
			((IfsContext)_localctx).instrucciones = instrucciones();
			setState(187);
			match(LLAVEDER);
			setState(188);
			((IfsContext)_localctx).elses = elses();
			_localctx.p=instruction.NewIf(((IfsContext)_localctx).expression.p,((IfsContext)_localctx).instrucciones.l,((IfsContext)_localctx).elses.e,((IfsContext)_localctx).P_IF.GetLine(),((IfsContext)_localctx).P_IF.GetColumn(),0)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ElsesContext extends ParserRuleContext {
		public interfaces.Instruction e;
		public InstruccionesContext instrucciones;
		public Token P_IF;
		public ExpressionContext expression;
		public ElsesContext elses;
		public TerminalNode P_ELSE() { return getToken(Chems.P_ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Chems.LLAVEIZQ, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Chems.LLAVEDER, 0); }
		public TerminalNode P_IF() { return getToken(Chems.P_IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ElsesContext elses() {
			return getRuleContext(ElsesContext.class,0);
		}
		public ElsesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elses; }
	}

	public final ElsesContext elses() throws RecognitionException {
		ElsesContext _localctx = new ElsesContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_elses);
		try {
			setState(207);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(191);
				match(P_ELSE);
				setState(192);
				match(LLAVEIZQ);
				setState(193);
				((ElsesContext)_localctx).instrucciones = instrucciones();
				setState(194);
				match(LLAVEDER);
				_localctx.e=instruction.NewIf(expresion.NewPrimitivo(1,interfaces.BOOLEAN,0,0),((ElsesContext)_localctx).instrucciones.l,instruction.NewElseNull("null"),0,0,3)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(197);
				match(P_ELSE);
				setState(198);
				((ElsesContext)_localctx).P_IF = match(P_IF);
				setState(199);
				((ElsesContext)_localctx).expression = expression();
				setState(200);
				match(LLAVEIZQ);
				setState(201);
				((ElsesContext)_localctx).instrucciones = instrucciones();
				setState(202);
				match(LLAVEDER);
				setState(203);
				((ElsesContext)_localctx).elses = elses();
				_localctx.e=instruction.NewIf(((ElsesContext)_localctx).expression.p,((ElsesContext)_localctx).instrucciones.l,((ElsesContext)_localctx).elses.e,((ElsesContext)_localctx).P_IF.GetLine(),((ElsesContext)_localctx).P_IF.GetColumn(),2)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				_localctx.e= instruction.NewElseNull("null")
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BreaksContext extends ParserRuleContext {
		public interfaces.Instruction i;
		public Token BREAK;
		public ExpressionContext expression;
		public TerminalNode BREAK() { return getToken(Chems.BREAK, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Chems.PTCOMA, 0); }
		public BreaksContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breaks; }
	}

	public final BreaksContext breaks() throws RecognitionException {
		BreaksContext _localctx = new BreaksContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_breaks);
		try {
			setState(217);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(209);
				((BreaksContext)_localctx).BREAK = match(BREAK);
				setState(210);
				((BreaksContext)_localctx).expression = expression();
				setState(211);
				match(PTCOMA);
				_localctx.i=instruction.NewBreak(((BreaksContext)_localctx).expression.p,true,((BreaksContext)_localctx).BREAK.GetLine(),((BreaksContext)_localctx).BREAK.GetColumn())
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(214);
				((BreaksContext)_localctx).BREAK = match(BREAK);
				setState(215);
				match(PTCOMA);
				_localctx.i=instruction.NewBreak(expresion.NewPrimitivo(1,interfaces.INTEGER,0,0),false,((BreaksContext)_localctx).BREAK.GetLine(),((BreaksContext)_localctx).BREAK.GetColumn())
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ContinuesContext extends ParserRuleContext {
		public interfaces.Instruction i;
		public Token CONTINUE;
		public TerminalNode CONTINUE() { return getToken(Chems.CONTINUE, 0); }
		public TerminalNode PTCOMA() { return getToken(Chems.PTCOMA, 0); }
		public ContinuesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continues; }
	}

	public final ContinuesContext continues() throws RecognitionException {
		ContinuesContext _localctx = new ContinuesContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_continues);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(219);
			((ContinuesContext)_localctx).CONTINUE = match(CONTINUE);
			setState(220);
			match(PTCOMA);
			_localctx.i=instruction.NewContinue("continue",((ContinuesContext)_localctx).CONTINUE.GetLine(),((ContinuesContext)_localctx).CONTINUE.GetColumn())
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public interfaces.Expresion p;
		public Expr_aritContext expr_arit;
		public Expr_aritContext expr_arit() {
			return getRuleContext(Expr_aritContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(223);
			((ExpressionContext)_localctx).expr_arit = expr_arit(0);
			_localctx.p = ((ExpressionContext)_localctx).expr_arit.p
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Expr_aritContext extends ParserRuleContext {
		public interfaces.Expresion p;
		public Expr_aritContext exp;
		public Expr_aritContext opIz;
		public Token PARIZQ;
		public Expr_aritContext opDe;
		public Token DIFERENTE;
		public Token SUB;
		public PrimitivoContext primitivo;
		public ExpressionContext expression;
		public Token op;
		public Tipo_dContext tipo_d;
		public Token PUNTO;
		public TerminalNode INT() { return getToken(Chems.INT, 0); }
		public TerminalNode DDPUNTO() { return getToken(Chems.DDPUNTO, 0); }
		public TerminalNode POW() { return getToken(Chems.POW, 0); }
		public TerminalNode PARIZQ() { return getToken(Chems.PARIZQ, 0); }
		public TerminalNode COMA() { return getToken(Chems.COMA, 0); }
		public TerminalNode PARDER() { return getToken(Chems.PARDER, 0); }
		public List<Expr_aritContext> expr_arit() {
			return getRuleContexts(Expr_aritContext.class);
		}
		public Expr_aritContext expr_arit(int i) {
			return getRuleContext(Expr_aritContext.class,i);
		}
		public TerminalNode FLOAT() { return getToken(Chems.FLOAT, 0); }
		public TerminalNode POWF() { return getToken(Chems.POWF, 0); }
		public TerminalNode DIFERENTE() { return getToken(Chems.DIFERENTE, 0); }
		public TerminalNode SUB() { return getToken(Chems.SUB, 0); }
		public PrimitivoContext primitivo() {
			return getRuleContext(PrimitivoContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode MUL() { return getToken(Chems.MUL, 0); }
		public TerminalNode DIV() { return getToken(Chems.DIV, 0); }
		public TerminalNode MOD() { return getToken(Chems.MOD, 0); }
		public TerminalNode ADD() { return getToken(Chems.ADD, 0); }
		public TerminalNode D_IGUAL() { return getToken(Chems.D_IGUAL, 0); }
		public TerminalNode NOT_E() { return getToken(Chems.NOT_E, 0); }
		public TerminalNode MENOR() { return getToken(Chems.MENOR, 0); }
		public TerminalNode MENORIGUAL() { return getToken(Chems.MENORIGUAL, 0); }
		public TerminalNode MAYORIGUAL() { return getToken(Chems.MAYORIGUAL, 0); }
		public TerminalNode MAYOR() { return getToken(Chems.MAYOR, 0); }
		public TerminalNode OR() { return getToken(Chems.OR, 0); }
		public TerminalNode AND() { return getToken(Chems.AND, 0); }
		public TerminalNode P_AS() { return getToken(Chems.P_AS, 0); }
		public Tipo_dContext tipo_d() {
			return getRuleContext(Tipo_dContext.class,0);
		}
		public TerminalNode PUNTO() { return getToken(Chems.PUNTO, 0); }
		public TerminalNode T_STRING() { return getToken(Chems.T_STRING, 0); }
		public Expr_aritContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_arit; }
	}

	public final Expr_aritContext expr_arit() throws RecognitionException {
		return expr_arit(0);
	}

	private Expr_aritContext expr_arit(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expr_aritContext _localctx = new Expr_aritContext(_ctx, _parentState);
		Expr_aritContext _prevctx = _localctx;
		int _startState = 34;
		enterRecursionRule(_localctx, 34, RULE_expr_arit, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(263);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				{
				setState(227);
				match(INT);
				setState(228);
				match(DDPUNTO);
				setState(229);
				match(POW);
				setState(230);
				((Expr_aritContext)_localctx).PARIZQ = match(PARIZQ);
				setState(231);
				((Expr_aritContext)_localctx).opIz = expr_arit(0);
				setState(232);
				match(COMA);
				setState(233);
				((Expr_aritContext)_localctx).opDe = expr_arit(0);
				setState(234);
				match(PARDER);
				_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,"^",((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).PARIZQ.GetLine(),((Expr_aritContext)_localctx).PARIZQ.GetColumn())
				}
				break;
			case FLOAT:
				{
				setState(237);
				match(FLOAT);
				setState(238);
				match(DDPUNTO);
				setState(239);
				match(POWF);
				setState(240);
				((Expr_aritContext)_localctx).PARIZQ = match(PARIZQ);
				setState(241);
				((Expr_aritContext)_localctx).opIz = expr_arit(0);
				setState(242);
				match(COMA);
				setState(243);
				((Expr_aritContext)_localctx).opDe = expr_arit(0);
				setState(244);
				match(PARDER);
				_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,"^",((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).PARIZQ.GetLine(),((Expr_aritContext)_localctx).PARIZQ.GetColumn())
				}
				break;
			case DIFERENTE:
				{
				setState(247);
				((Expr_aritContext)_localctx).DIFERENTE = match(DIFERENTE);
				setState(248);
				((Expr_aritContext)_localctx).opIz = expr_arit(8);
				_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,"!",((Expr_aritContext)_localctx).opIz.p,true,((Expr_aritContext)_localctx).DIFERENTE.GetLine(),((Expr_aritContext)_localctx).DIFERENTE.GetColumn())
				}
				break;
			case SUB:
				{
				setState(251);
				((Expr_aritContext)_localctx).SUB = match(SUB);
				setState(252);
				((Expr_aritContext)_localctx).opIz = expr_arit(7);
				_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,"°",((Expr_aritContext)_localctx).opIz.p,true,((Expr_aritContext)_localctx).SUB.GetLine(),((Expr_aritContext)_localctx).SUB.GetColumn())
				}
				break;
			case TRUE:
			case FALSE:
			case NUMBER:
			case DECIMAL:
			case STRING:
			case ID:
			case CARACTER:
				{
				setState(255);
				((Expr_aritContext)_localctx).primitivo = primitivo();
				_localctx.p = ((Expr_aritContext)_localctx).primitivo.p
				}
				break;
			case PARIZQ:
				{
				setState(258);
				((Expr_aritContext)_localctx).PARIZQ = match(PARIZQ);
				setState(259);
				((Expr_aritContext)_localctx).expression = expression();
				setState(260);
				match(PARDER);
				_localctx.p = ((Expr_aritContext)_localctx).expression.p
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(298);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(296);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(265);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(266);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MUL) | (1L << DIV) | (1L << MOD))) != 0)) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(267);
						((Expr_aritContext)_localctx).opDe = expr_arit(7);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).op.GetLine(),((Expr_aritContext)_localctx).op.GetColumn())
						}
						break;
					case 2:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(270);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(271);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(272);
						((Expr_aritContext)_localctx).opDe = expr_arit(6);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).op.GetLine(),((Expr_aritContext)_localctx).op.GetColumn())
						}
						break;
					case 3:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(275);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(276);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MENOR) | (1L << MAYORIGUAL) | (1L << MENORIGUAL) | (1L << D_IGUAL) | (1L << NOT_E) | (1L << MAYOR))) != 0)) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(277);
						((Expr_aritContext)_localctx).opDe = expr_arit(5);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).op.GetLine(),((Expr_aritContext)_localctx).op.GetColumn())
						}
						break;
					case 4:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(280);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(281);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OR || _la==AND) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(282);
						((Expr_aritContext)_localctx).opDe = expr_arit(4);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).op.GetLine(),((Expr_aritContext)_localctx).op.GetColumn())
						}
						break;
					case 5:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.exp = _prevctx;
						_localctx.exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(285);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(286);
						match(P_AS);
						setState(287);
						((Expr_aritContext)_localctx).tipo_d = tipo_d();
						_localctx.p=expresion.NewCast(((Expr_aritContext)_localctx).exp.p,((Expr_aritContext)_localctx).tipo_d.t)
						}
						break;
					case 6:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.exp = _prevctx;
						_localctx.exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(290);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(291);
						((Expr_aritContext)_localctx).PUNTO = match(PUNTO);
						setState(292);
						match(T_STRING);
						setState(293);
						((Expr_aritContext)_localctx).PARIZQ = match(PARIZQ);
						setState(294);
						match(PARDER);
						_localctx.p=expresion.NewToString(((Expr_aritContext)_localctx).exp.p,((Expr_aritContext)_localctx).PUNTO.GetLine(),((Expr_aritContext)_localctx).PUNTO.GetColumn())
						}
						break;
					}
					} 
				}
				setState(300);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class PrimitivoContext extends ParserRuleContext {
		public interfaces.Expresion p;
		public Token NUMBER;
		public Token STRING;
		public Token TRUE;
		public Token FALSE;
		public Token DECIMAL;
		public Token CARACTER;
		public Token id;
		public TerminalNode NUMBER() { return getToken(Chems.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(Chems.STRING, 0); }
		public TerminalNode TRUE() { return getToken(Chems.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(Chems.FALSE, 0); }
		public TerminalNode DECIMAL() { return getToken(Chems.DECIMAL, 0); }
		public TerminalNode CARACTER() { return getToken(Chems.CARACTER, 0); }
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public PrimitivoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitivo; }
	}

	public final PrimitivoContext primitivo() throws RecognitionException {
		PrimitivoContext _localctx = new PrimitivoContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_primitivo);
		try {
			setState(315);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NUMBER:
				enterOuterAlt(_localctx, 1);
				{
				setState(301);
				((PrimitivoContext)_localctx).NUMBER = match(NUMBER);

				            	num,err := strconv.Atoi((((PrimitivoContext)_localctx).NUMBER!=null?((PrimitivoContext)_localctx).NUMBER.getText():null))
				                if err!= nil{
				                    fmt.Println(err)
				                }
				            linea:=_localctx._NUMBER.GetLine()
				            col:=_localctx._NUMBER.GetColumn()
				            _localctx.p = expresion.NewPrimitivo (num,interfaces.INTEGER,col,linea)
				            
				       
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 2);
				{
				setState(303);
				((PrimitivoContext)_localctx).STRING = match(STRING);
				 
				      linea:=_localctx._STRING.GetLine()
				      col:=_localctx._STRING.GetColumn()
				      str:= (((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null)[1:len((((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null))-1]
				      _localctx.p = expresion.NewPrimitivo(str,interfaces.STRING,col,linea)
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(305);
				((PrimitivoContext)_localctx).TRUE = match(TRUE);
				            
				        _localctx.p=expresion.NewPrimitivo(1,interfaces.BOOLEAN,((PrimitivoContext)_localctx).TRUE.GetColumn(),((PrimitivoContext)_localctx).TRUE.GetLine())      
				    
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 4);
				{
				setState(307);
				((PrimitivoContext)_localctx).FALSE = match(FALSE);

				        _localctx.p=expresion.NewPrimitivo(0,interfaces.BOOLEAN,((PrimitivoContext)_localctx).FALSE.GetColumn(),((PrimitivoContext)_localctx).FALSE.GetLine())      
				    
				}
				break;
			case DECIMAL:
				enterOuterAlt(_localctx, 5);
				{
				setState(309);
				((PrimitivoContext)_localctx).DECIMAL = match(DECIMAL);

				            linea:=_localctx._DECIMAL.GetLine()
				            col:=_localctx._DECIMAL.GetColumn()
				              num,err:=strconv.ParseFloat((((PrimitivoContext)_localctx).DECIMAL!=null?((PrimitivoContext)_localctx).DECIMAL.getText():null),64)
				              if err!=nil{ 
				                fmt.Println(err)
				              }
				              _localctx.p=expresion.NewPrimitivo(num,interfaces.FLOAT,col,linea)
				}
				break;
			case CARACTER:
				enterOuterAlt(_localctx, 6);
				{
				setState(311);
				((PrimitivoContext)_localctx).CARACTER = match(CARACTER);

				      linea:=_localctx._CARACTER.GetLine()
				      col:=_localctx._CARACTER.GetColumn()
				      str:= (((PrimitivoContext)_localctx).CARACTER!=null?((PrimitivoContext)_localctx).CARACTER.getText():null)[1:len((((PrimitivoContext)_localctx).CARACTER!=null?((PrimitivoContext)_localctx).CARACTER.getText():null))-1]
				      _localctx.p = expresion.NewPrimitivo(str,interfaces.CHAR,col,linea)
				    
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 7);
				{
				setState(313);
				((PrimitivoContext)_localctx).id = match(ID);

				      linea:=_localctx.id.GetLine()
				      col:=_localctx.id.GetColumn()
				      _localctx.p=expresion.NewCallVariable((((PrimitivoContext)_localctx).id!=null?((PrimitivoContext)_localctx).id.getText():null),linea,col)
				    
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 17:
			return expr_arit_sempred((Expr_aritContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_arit_sempred(Expr_aritContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 6);
		case 1:
			return precpred(_ctx, 5);
		case 2:
			return precpred(_ctx, 4);
		case 3:
			return precpred(_ctx, 3);
		case 4:
			return precpred(_ctx, 10);
		case 5:
			return precpred(_ctx, 9);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3X\u0140\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\3\2\3\2\3\2\3\3\7\3-\n\3\f\3\16\3\60\13\3\3\3\3\3"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4U\n"+
		"\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\5\7"+
		"g\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\bt\n\b\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u0086\n\t\3\n"+
		"\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\5\13\u009a\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\r\5\r\u00b8\n\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17"+
		"\u00d2\n\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20\u00dc\n\20\3"+
		"\21\3\21\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3"+
		"\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3"+
		"\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3"+
		"\23\3\23\5\23\u010a\n\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\7\23\u012b\n\23\f\23\16\23\u012e"+
		"\13\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\5\24\u013e\n\24\3\24\2\3$\25\2\4\6\b\n\f\16\20\22\24\26\30\32\34"+
		"\36 \"$&\2\6\4\2\24\25\30\30\3\2\26\27\4\2\13\r\17\21\3\2\22\23\2\u015a"+
		"\2(\3\2\2\2\4.\3\2\2\2\6T\3\2\2\2\bV\3\2\2\2\n[\3\2\2\2\ff\3\2\2\2\16"+
		"s\3\2\2\2\20\u0085\3\2\2\2\22\u0087\3\2\2\2\24\u0099\3\2\2\2\26\u009b"+
		"\3\2\2\2\30\u00b7\3\2\2\2\32\u00b9\3\2\2\2\34\u00d1\3\2\2\2\36\u00db\3"+
		"\2\2\2 \u00dd\3\2\2\2\"\u00e1\3\2\2\2$\u0109\3\2\2\2&\u013d\3\2\2\2()"+
		"\5\4\3\2)*\b\2\1\2*\3\3\2\2\2+-\5\6\4\2,+\3\2\2\2-\60\3\2\2\2.,\3\2\2"+
		"\2./\3\2\2\2/\61\3\2\2\2\60.\3\2\2\2\61\62\b\3\1\2\62\5\3\2\2\2\63\64"+
		"\7#\2\2\64\65\7\3\2\2\65\66\7$\2\2\66\67\7\32\2\2\678\5\"\22\289\7\33"+
		"\2\29:\7\4\2\2:;\b\4\1\2;U\3\2\2\2<=\5\n\6\2=>\7\4\2\2>?\b\4\1\2?U\3\2"+
		"\2\2@A\5\b\5\2AB\7\4\2\2BC\b\4\1\2CU\3\2\2\2DE\7*\2\2EF\5\"\22\2FG\7\34"+
		"\2\2GH\5\4\3\2HI\7\35\2\2IJ\b\4\1\2JU\3\2\2\2KL\5\36\20\2LM\b\4\1\2MU"+
		"\3\2\2\2NO\5 \21\2OP\b\4\1\2PU\3\2\2\2QR\5\32\16\2RS\b\4\1\2SU\3\2\2\2"+
		"T\63\3\2\2\2T<\3\2\2\2T@\3\2\2\2TD\3\2\2\2TK\3\2\2\2TN\3\2\2\2TQ\3\2\2"+
		"\2U\7\3\2\2\2VW\7V\2\2WX\7\t\2\2XY\5\"\22\2YZ\b\5\1\2Z\t\3\2\2\2[\\\7"+
		"\63\2\2\\]\5\f\7\2]^\7V\2\2^_\5\16\b\2_`\7\t\2\2`a\5\"\22\2ab\b\6\1\2"+
		"b\13\3\2\2\2cd\7:\2\2dg\b\7\1\2eg\b\7\1\2fc\3\2\2\2fe\3\2\2\2g\r\3\2\2"+
		"\2hi\7\b\2\2ij\5\26\f\2jk\b\b\1\2kt\3\2\2\2lm\5\22\n\2mn\b\b\1\2nt\3\2"+
		"\2\2op\7\b\2\2pq\5\20\t\2qr\b\b\1\2rt\3\2\2\2sh\3\2\2\2sl\3\2\2\2so\3"+
		"\2\2\2t\17\3\2\2\2uv\7\64\2\2v\u0086\b\t\1\2wx\7\65\2\2x\u0086\b\t\1\2"+
		"yz\7\66\2\2z\u0086\b\t\1\2{|\7\67\2\2|\u0086\b\t\1\2}~\78\2\2~\u0086\b"+
		"\t\1\2\177\u0080\7\'\2\2\u0080\u0086\b\t\1\2\u0081\u0082\79\2\2\u0082"+
		"\u0086\b\t\1\2\u0083\u0084\7V\2\2\u0084\u0086\b\t\1\2\u0085u\3\2\2\2\u0085"+
		"w\3\2\2\2\u0085y\3\2\2\2\u0085{\3\2\2\2\u0085}\3\2\2\2\u0085\177\3\2\2"+
		"\2\u0085\u0081\3\2\2\2\u0085\u0083\3\2\2\2\u0086\21\3\2\2\2\u0087\u0088"+
		"\7\b\2\2\u0088\u0089\5\24\13\2\u0089\u008a\b\n\1\2\u008a\23\3\2\2\2\u008b"+
		"\u008c\7\36\2\2\u008c\u008d\5\20\t\2\u008d\u008e\7\4\2\2\u008e\u008f\5"+
		"\"\22\2\u008f\u0090\7\37\2\2\u0090\u0091\b\13\1\2\u0091\u009a\3\2\2\2"+
		"\u0092\u0093\7\36\2\2\u0093\u0094\5\24\13\2\u0094\u0095\7\4\2\2\u0095"+
		"\u0096\5\"\22\2\u0096\u0097\7\37\2\2\u0097\u0098\b\13\1\2\u0098\u009a"+
		"\3\2\2\2\u0099\u008b\3\2\2\2\u0099\u0092\3\2\2\2\u009a\25\3\2\2\2\u009b"+
		"\u009c\7O\2\2\u009c\u009d\7\13\2\2\u009d\u009e\5\30\r\2\u009e\u009f\7"+
		"\21\2\2\u009f\u00a0\b\f\1\2\u00a0\27\3\2\2\2\u00a1\u00a2\7\64\2\2\u00a2"+
		"\u00b8\b\r\1\2\u00a3\u00a4\7\65\2\2\u00a4\u00b8\b\r\1\2\u00a5\u00a6\7"+
		"\66\2\2\u00a6\u00b8\b\r\1\2\u00a7\u00a8\7\67\2\2\u00a8\u00b8\b\r\1\2\u00a9"+
		"\u00aa\78\2\2\u00aa\u00b8\b\r\1\2\u00ab\u00ac\7\'\2\2\u00ac\u00b8\b\r"+
		"\1\2\u00ad\u00ae\79\2\2\u00ae\u00b8\b\r\1\2\u00af\u00b0\7V\2\2\u00b0\u00b8"+
		"\b\r\1\2\u00b1\u00b2\7O\2\2\u00b2\u00b3\7\13\2\2\u00b3\u00b4\5\30\r\2"+
		"\u00b4\u00b5\7\21\2\2\u00b5\u00b6\b\r\1\2\u00b6\u00b8\3\2\2\2\u00b7\u00a1"+
		"\3\2\2\2\u00b7\u00a3\3\2\2\2\u00b7\u00a5\3\2\2\2\u00b7\u00a7\3\2\2\2\u00b7"+
		"\u00a9\3\2\2\2\u00b7\u00ab\3\2\2\2\u00b7\u00ad\3\2\2\2\u00b7\u00af\3\2"+
		"\2\2\u00b7\u00b1\3\2\2\2\u00b8\31\3\2\2\2\u00b9\u00ba\7(\2\2\u00ba\u00bb"+
		"\5\"\22\2\u00bb\u00bc\7\34\2\2\u00bc\u00bd\5\4\3\2\u00bd\u00be\7\35\2"+
		"\2\u00be\u00bf\5\34\17\2\u00bf\u00c0\b\16\1\2\u00c0\33\3\2\2\2\u00c1\u00c2"+
		"\7)\2\2\u00c2\u00c3\7\34\2\2\u00c3\u00c4\5\4\3\2\u00c4\u00c5\7\35\2\2"+
		"\u00c5\u00c6\b\17\1\2\u00c6\u00d2\3\2\2\2\u00c7\u00c8\7)\2\2\u00c8\u00c9"+
		"\7(\2\2\u00c9\u00ca\5\"\22\2\u00ca\u00cb\7\34\2\2\u00cb\u00cc\5\4\3\2"+
		"\u00cc\u00cd\7\35\2\2\u00cd\u00ce\5\34\17\2\u00ce\u00cf\b\17\1\2\u00cf"+
		"\u00d2\3\2\2\2\u00d0\u00d2\b\17\1\2\u00d1\u00c1\3\2\2\2\u00d1\u00c7\3"+
		"\2\2\2\u00d1\u00d0\3\2\2\2\u00d2\35\3\2\2\2\u00d3\u00d4\7\60\2\2\u00d4"+
		"\u00d5\5\"\22\2\u00d5\u00d6\7\4\2\2\u00d6\u00d7\b\20\1\2\u00d7\u00dc\3"+
		"\2\2\2\u00d8\u00d9\7\60\2\2\u00d9\u00da\7\4\2\2\u00da\u00dc\b\20\1\2\u00db"+
		"\u00d3\3\2\2\2\u00db\u00d8\3\2\2\2\u00dc\37\3\2\2\2\u00dd\u00de\7\61\2"+
		"\2\u00de\u00df\7\4\2\2\u00df\u00e0\b\21\1\2\u00e0!\3\2\2\2\u00e1\u00e2"+
		"\5$\23\2\u00e2\u00e3\b\22\1\2\u00e3#\3\2\2\2\u00e4\u00e5\b\23\1\2\u00e5"+
		"\u00e6\7\64\2\2\u00e6\u00e7\7\"\2\2\u00e7\u00e8\7/\2\2\u00e8\u00e9\7\32"+
		"\2\2\u00e9\u00ea\5$\23\2\u00ea\u00eb\7\5\2\2\u00eb\u00ec\5$\23\2\u00ec"+
		"\u00ed\7\33\2\2\u00ed\u00ee\b\23\1\2\u00ee\u010a\3\2\2\2\u00ef\u00f0\7"+
		"\65\2\2\u00f0\u00f1\7\"\2\2\u00f1\u00f2\7.\2\2\u00f2\u00f3\7\32\2\2\u00f3"+
		"\u00f4\5$\23\2\u00f4\u00f5\7\5\2\2\u00f5\u00f6\5$\23\2\u00f6\u00f7\7\33"+
		"\2\2\u00f7\u00f8\b\23\1\2\u00f8\u010a\3\2\2\2\u00f9\u00fa\7\7\2\2\u00fa"+
		"\u00fb\5$\23\n\u00fb\u00fc\b\23\1\2\u00fc\u010a\3\2\2\2\u00fd\u00fe\7"+
		"\27\2\2\u00fe\u00ff\5$\23\t\u00ff\u0100\b\23\1\2\u0100\u010a\3\2\2\2\u0101"+
		"\u0102\5&\24\2\u0102\u0103\b\23\1\2\u0103\u010a\3\2\2\2\u0104\u0105\7"+
		"\32\2\2\u0105\u0106\5\"\22\2\u0106\u0107\7\33\2\2\u0107\u0108\b\23\1\2"+
		"\u0108\u010a\3\2\2\2\u0109\u00e4\3\2\2\2\u0109\u00ef\3\2\2\2\u0109\u00f9"+
		"\3\2\2\2\u0109\u00fd\3\2\2\2\u0109\u0101\3\2\2\2\u0109\u0104\3\2\2\2\u010a"+
		"\u012c\3\2\2\2\u010b\u010c\f\b\2\2\u010c\u010d\t\2\2\2\u010d\u010e\5$"+
		"\23\t\u010e\u010f\b\23\1\2\u010f\u012b\3\2\2\2\u0110\u0111\f\7\2\2\u0111"+
		"\u0112\t\3\2\2\u0112\u0113\5$\23\b\u0113\u0114\b\23\1\2\u0114\u012b\3"+
		"\2\2\2\u0115\u0116\f\6\2\2\u0116\u0117\t\4\2\2\u0117\u0118\5$\23\7\u0118"+
		"\u0119\b\23\1\2\u0119\u012b\3\2\2\2\u011a\u011b\f\5\2\2\u011b\u011c\t"+
		"\5\2\2\u011c\u011d\5$\23\6\u011d\u011e\b\23\1\2\u011e\u012b\3\2\2\2\u011f"+
		"\u0120\f\f\2\2\u0120\u0121\7-\2\2\u0121\u0122\5\20\t\2\u0122\u0123\b\23"+
		"\1\2\u0123\u012b\3\2\2\2\u0124\u0125\f\13\2\2\u0125\u0126\7\3\2\2\u0126"+
		"\u0127\7D\2\2\u0127\u0128\7\32\2\2\u0128\u0129\7\33\2\2\u0129\u012b\b"+
		"\23\1\2\u012a\u010b\3\2\2\2\u012a\u0110\3\2\2\2\u012a\u0115\3\2\2\2\u012a"+
		"\u011a\3\2\2\2\u012a\u011f\3\2\2\2\u012a\u0124\3\2\2\2\u012b\u012e\3\2"+
		"\2\2\u012c\u012a\3\2\2\2\u012c\u012d\3\2\2\2\u012d%\3\2\2\2\u012e\u012c"+
		"\3\2\2\2\u012f\u0130\7S\2\2\u0130\u013e\b\24\1\2\u0131\u0132\7U\2\2\u0132"+
		"\u013e\b\24\1\2\u0133\u0134\7?\2\2\u0134\u013e\b\24\1\2\u0135\u0136\7"+
		"@\2\2\u0136\u013e\b\24\1\2\u0137\u0138\7T\2\2\u0138\u013e\b\24\1\2\u0139"+
		"\u013a\7W\2\2\u013a\u013e\b\24\1\2\u013b\u013c\7V\2\2\u013c\u013e\b\24"+
		"\1\2\u013d\u012f\3\2\2\2\u013d\u0131\3\2\2\2\u013d\u0133\3\2\2\2\u013d"+
		"\u0135\3\2\2\2\u013d\u0137\3\2\2\2\u013d\u0139\3\2\2\2\u013d\u013b\3\2"+
		"\2\2\u013e\'\3\2\2\2\17.Tfs\u0085\u0099\u00b7\u00d1\u00db\u0109\u012a"+
		"\u012c\u013d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}