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
		RULE_loops = 12, RULE_ifs = 13, RULE_elses = 14, RULE_breaks = 15, RULE_continues = 16, 
		RULE_expression = 17, RULE_expr_arit = 18, RULE_primitivo = 19;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instrucciones", "instruccion", "asignacion_var", "declaracion_var", 
			"mutable", "types", "tipo_d", "asignar_Array", "dimensiones", "tipo_vector", 
			"vectores", "loops", "ifs", "elses", "breaks", "continues", "expression", 
			"expr_arit", "primitivo"
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
			setState(40);
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
			setState(46);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 33)) & ~0x3f) == 0 && ((1L << (_la - 33)) & ((1L << (CONSOLE - 33)) | (1L << (P_IF - 33)) | (1L << (P_WHILE - 33)) | (1L << (BREAK - 33)) | (1L << (CONTINUE - 33)) | (1L << (LET - 33)) | (1L << (LOOP - 33)) | (1L << (ID - 33)))) != 0)) {
				{
				{
				setState(43);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(48);
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
		public LoopsContext loops;
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
		public LoopsContext loops() {
			return getRuleContext(LoopsContext.class,0);
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
			setState(87);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CONSOLE:
				enterOuterAlt(_localctx, 1);
				{
				setState(51);
				match(CONSOLE);
				setState(52);
				match(PUNTO);
				setState(53);
				match(LOG);
				setState(54);
				match(PARIZQ);
				setState(55);
				((InstruccionContext)_localctx).expression = expression();
				setState(56);
				match(PARDER);
				setState(57);
				match(PTCOMA);
				_localctx.instr = instruction.NewImprimir(((InstruccionContext)_localctx).expression.p)
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(60);
				((InstruccionContext)_localctx).declaracion_var = declaracion_var();
				setState(61);
				match(PTCOMA);
				_localctx.instr=((InstruccionContext)_localctx).declaracion_var.i
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(64);
				((InstruccionContext)_localctx).asignacion_var = asignacion_var();
				setState(65);
				match(PTCOMA);
				_localctx.instr=((InstruccionContext)_localctx).asignacion_var.i
				}
				break;
			case P_WHILE:
				enterOuterAlt(_localctx, 4);
				{
				setState(68);
				match(P_WHILE);
				setState(69);
				((InstruccionContext)_localctx).expression = expression();
				setState(70);
				((InstruccionContext)_localctx).LLAVEIZQ = match(LLAVEIZQ);
				setState(71);
				((InstruccionContext)_localctx).instrucciones = instrucciones();
				setState(72);
				((InstruccionContext)_localctx).LLAVEDER = match(LLAVEDER);
				_localctx.instr = instruction.NewWhile(((InstruccionContext)_localctx).expression.p, ((InstruccionContext)_localctx).instrucciones.l,((InstruccionContext)_localctx).LLAVEIZQ.GetLine(),((InstruccionContext)_localctx).LLAVEDER.GetColumn())
				}
				break;
			case BREAK:
				enterOuterAlt(_localctx, 5);
				{
				setState(75);
				((InstruccionContext)_localctx).breaks = breaks();
				_localctx.instr=((InstruccionContext)_localctx).breaks.i
				}
				break;
			case CONTINUE:
				enterOuterAlt(_localctx, 6);
				{
				setState(78);
				((InstruccionContext)_localctx).continues = continues();
				_localctx.instr=((InstruccionContext)_localctx).continues.i
				}
				break;
			case P_IF:
				enterOuterAlt(_localctx, 7);
				{
				setState(81);
				((InstruccionContext)_localctx).ifs = ifs();
				_localctx.instr=((InstruccionContext)_localctx).ifs.p
				}
				break;
			case LOOP:
				enterOuterAlt(_localctx, 8);
				{
				setState(84);
				((InstruccionContext)_localctx).loops = loops();
				_localctx.instr=((InstruccionContext)_localctx).loops.i
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
			setState(89);
			((Asignacion_varContext)_localctx).id = match(ID);
			setState(90);
			match(IGUAL);
			setState(91);
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
			setState(94);
			((Declaracion_varContext)_localctx).LET = match(LET);
			setState(95);
			((Declaracion_varContext)_localctx).mutable = mutable();
			setState(96);
			((Declaracion_varContext)_localctx).id = match(ID);
			setState(97);
			((Declaracion_varContext)_localctx).types = types();
			setState(98);
			match(IGUAL);
			setState(99);
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
			setState(105);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MUT:
				enterOuterAlt(_localctx, 1);
				{
				setState(102);
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
			setState(118);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(107);
				match(DOSPUNTOS);
				setState(108);
				((TypesContext)_localctx).tipo_vector = tipo_vector();
				_localctx.l=((TypesContext)_localctx).tipo_vector.t
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(111);
				((TypesContext)_localctx).a = asignar_Array();

				    dim:=arrayList.New()
				    dim.Add(((TypesContext)_localctx).a.d)
				    _localctx.l=interfaces.TipoSimbolo{interfaces.ARRAY,dim}
				  
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(114);
				match(DOSPUNTOS);
				setState(115);
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
			setState(136);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(120);
				match(INT);
				_localctx.t=interfaces.INTEGER
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(122);
				match(FLOAT);
				_localctx.t=interfaces.FLOAT
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 3);
				{
				setState(124);
				match(BOOL);
				_localctx.t=interfaces.BOOLEAN
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 4);
				{
				setState(126);
				match(CHAR);
				_localctx.t=interfaces.CHAR
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 5);
				{
				setState(128);
				match(STR);
				_localctx.t=interfaces.STRING
				}
				break;
			case P_STRING:
				enterOuterAlt(_localctx, 6);
				{
				setState(130);
				match(P_STRING);
				_localctx.t=interfaces.STR
				}
				break;
			case USIZE:
				enterOuterAlt(_localctx, 7);
				{
				setState(132);
				match(USIZE);
				_localctx.t=interfaces.USIZE
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 8);
				{
				setState(134);
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
			setState(138);
			match(DOSPUNTOS);
			setState(139);
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
			setState(156);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(142);
				match(CORIZQ);
				setState(143);
				((DimensionesContext)_localctx).tipo_d = tipo_d();
				setState(144);
				match(PTCOMA);
				setState(145);
				((DimensionesContext)_localctx).expression = expression();
				setState(146);
				match(CORDER);

				    list:=arrayList.New()
				    list.Add(((DimensionesContext)_localctx).expression.p)
				    _localctx.d=interfaces.Dimensions{((DimensionesContext)_localctx).tipo_d.t,list}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(149);
				match(CORIZQ);
				setState(150);
				((DimensionesContext)_localctx).dimensiones = dimensiones();
				setState(151);
				match(PTCOMA);
				setState(152);
				((DimensionesContext)_localctx).expression = expression();
				setState(153);
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
			setState(158);
			match(VECT);
			setState(159);
			match(MENOR);
			setState(160);
			((Tipo_vectorContext)_localctx).vectores = vectores();
			setState(161);
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
			setState(186);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(164);
				match(INT);
				_localctx.l=arrayList.New()
				            _localctx.l.Add(interfaces.INTEGER)
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(166);
				match(FLOAT);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.FLOAT)
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 3);
				{
				setState(168);
				match(BOOL);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.BOOLEAN)
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 4);
				{
				setState(170);
				match(CHAR);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.CHAR)
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 5);
				{
				setState(172);
				match(STR);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.STRING)
				}
				break;
			case P_STRING:
				enterOuterAlt(_localctx, 6);
				{
				setState(174);
				match(P_STRING);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.STR)
				}
				break;
			case USIZE:
				enterOuterAlt(_localctx, 7);
				{
				setState(176);
				match(USIZE);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.USIZE)
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 8);
				{
				setState(178);
				((VectoresContext)_localctx).id = match(ID);
				_localctx.l=arrayList.New()
				              _localctx.l.Add((((VectoresContext)_localctx).id!=null?((VectoresContext)_localctx).id.getText():null))
				              _localctx.l.Add(interfaces.STRUCT)
				}
				break;
			case VECT:
				enterOuterAlt(_localctx, 9);
				{
				setState(180);
				match(VECT);
				setState(181);
				match(MENOR);
				setState(182);
				((VectoresContext)_localctx).vectores = vectores();
				setState(183);
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

	public static class LoopsContext extends ParserRuleContext {
		public interfaces.Instruction i;
		public InstruccionesContext instrucciones;
		public TerminalNode LOOP() { return getToken(Chems.LOOP, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Chems.LLAVEIZQ, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Chems.LLAVEDER, 0); }
		public LoopsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loops; }
	}

	public final LoopsContext loops() throws RecognitionException {
		LoopsContext _localctx = new LoopsContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_loops);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(188);
			match(LOOP);
			setState(189);
			match(LLAVEIZQ);
			setState(190);
			((LoopsContext)_localctx).instrucciones = instrucciones();
			setState(191);
			match(LLAVEDER);
			_localctx.i=instruction.NewLoop(((LoopsContext)_localctx).instrucciones.l)
			    fmt.Println("TOY EN EL ANALISIS")
			  
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
		enterRule(_localctx, 26, RULE_ifs);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(194);
			((IfsContext)_localctx).P_IF = match(P_IF);
			setState(195);
			((IfsContext)_localctx).expression = expression();
			setState(196);
			match(LLAVEIZQ);
			setState(197);
			((IfsContext)_localctx).instrucciones = instrucciones();
			setState(198);
			match(LLAVEDER);
			setState(199);
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
		enterRule(_localctx, 28, RULE_elses);
		try {
			setState(218);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(202);
				match(P_ELSE);
				setState(203);
				match(LLAVEIZQ);
				setState(204);
				((ElsesContext)_localctx).instrucciones = instrucciones();
				setState(205);
				match(LLAVEDER);
				_localctx.e=instruction.NewIf(expresion.NewPrimitivo(1,interfaces.BOOLEAN,0,0),((ElsesContext)_localctx).instrucciones.l,instruction.NewElseNull("null"),0,0,3)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(208);
				match(P_ELSE);
				setState(209);
				((ElsesContext)_localctx).P_IF = match(P_IF);
				setState(210);
				((ElsesContext)_localctx).expression = expression();
				setState(211);
				match(LLAVEIZQ);
				setState(212);
				((ElsesContext)_localctx).instrucciones = instrucciones();
				setState(213);
				match(LLAVEDER);
				setState(214);
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
		enterRule(_localctx, 30, RULE_breaks);
		try {
			setState(228);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(220);
				((BreaksContext)_localctx).BREAK = match(BREAK);
				setState(221);
				((BreaksContext)_localctx).expression = expression();
				setState(222);
				match(PTCOMA);
				_localctx.i=instruction.NewBreak(((BreaksContext)_localctx).expression.p,true,((BreaksContext)_localctx).BREAK.GetLine(),((BreaksContext)_localctx).BREAK.GetColumn())
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(225);
				((BreaksContext)_localctx).BREAK = match(BREAK);
				setState(226);
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
		enterRule(_localctx, 32, RULE_continues);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(230);
			((ContinuesContext)_localctx).CONTINUE = match(CONTINUE);
			setState(231);
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
		enterRule(_localctx, 34, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(234);
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
		public LoopsContext loops;
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
		public LoopsContext loops() {
			return getRuleContext(LoopsContext.class,0);
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
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_expr_arit, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(277);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				{
				setState(238);
				match(INT);
				setState(239);
				match(DDPUNTO);
				setState(240);
				match(POW);
				setState(241);
				((Expr_aritContext)_localctx).PARIZQ = match(PARIZQ);
				setState(242);
				((Expr_aritContext)_localctx).opIz = expr_arit(0);
				setState(243);
				match(COMA);
				setState(244);
				((Expr_aritContext)_localctx).opDe = expr_arit(0);
				setState(245);
				match(PARDER);
				_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,"^",((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).PARIZQ.GetLine(),((Expr_aritContext)_localctx).PARIZQ.GetColumn())
				}
				break;
			case FLOAT:
				{
				setState(248);
				match(FLOAT);
				setState(249);
				match(DDPUNTO);
				setState(250);
				match(POWF);
				setState(251);
				((Expr_aritContext)_localctx).PARIZQ = match(PARIZQ);
				setState(252);
				((Expr_aritContext)_localctx).opIz = expr_arit(0);
				setState(253);
				match(COMA);
				setState(254);
				((Expr_aritContext)_localctx).opDe = expr_arit(0);
				setState(255);
				match(PARDER);
				_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,"^",((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).PARIZQ.GetLine(),((Expr_aritContext)_localctx).PARIZQ.GetColumn())
				}
				break;
			case DIFERENTE:
				{
				setState(258);
				((Expr_aritContext)_localctx).DIFERENTE = match(DIFERENTE);
				setState(259);
				((Expr_aritContext)_localctx).opIz = expr_arit(9);
				_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,"!",((Expr_aritContext)_localctx).opIz.p,true,((Expr_aritContext)_localctx).DIFERENTE.GetLine(),((Expr_aritContext)_localctx).DIFERENTE.GetColumn())
				}
				break;
			case SUB:
				{
				setState(262);
				((Expr_aritContext)_localctx).SUB = match(SUB);
				setState(263);
				((Expr_aritContext)_localctx).opIz = expr_arit(8);
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
				setState(266);
				((Expr_aritContext)_localctx).primitivo = primitivo();
				_localctx.p = ((Expr_aritContext)_localctx).primitivo.p
				}
				break;
			case PARIZQ:
				{
				setState(269);
				((Expr_aritContext)_localctx).PARIZQ = match(PARIZQ);
				setState(270);
				((Expr_aritContext)_localctx).expression = expression();
				setState(271);
				match(PARDER);
				_localctx.p = ((Expr_aritContext)_localctx).expression.p
				}
				break;
			case LOOP:
				{
				setState(274);
				((Expr_aritContext)_localctx).loops = loops();
				_localctx.p=expresion.NewDevLoop(((Expr_aritContext)_localctx).loops.i)
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(312);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(310);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(279);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(280);
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
						setState(281);
						((Expr_aritContext)_localctx).opDe = expr_arit(8);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).op.GetLine(),((Expr_aritContext)_localctx).op.GetColumn())
						}
						break;
					case 2:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(284);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(285);
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
						setState(286);
						((Expr_aritContext)_localctx).opDe = expr_arit(7);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).op.GetLine(),((Expr_aritContext)_localctx).op.GetColumn())
						}
						break;
					case 3:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(289);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(290);
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
						setState(291);
						((Expr_aritContext)_localctx).opDe = expr_arit(6);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).op.GetLine(),((Expr_aritContext)_localctx).op.GetColumn())
						}
						break;
					case 4:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(294);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(295);
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
						setState(296);
						((Expr_aritContext)_localctx).opDe = expr_arit(5);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).op.GetLine(),((Expr_aritContext)_localctx).op.GetColumn())
						}
						break;
					case 5:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.exp = _prevctx;
						_localctx.exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(299);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(300);
						match(P_AS);
						setState(301);
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
						setState(304);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(305);
						((Expr_aritContext)_localctx).PUNTO = match(PUNTO);
						setState(306);
						match(T_STRING);
						setState(307);
						((Expr_aritContext)_localctx).PARIZQ = match(PARIZQ);
						setState(308);
						match(PARDER);
						_localctx.p=expresion.NewToString(((Expr_aritContext)_localctx).exp.p,((Expr_aritContext)_localctx).PUNTO.GetLine(),((Expr_aritContext)_localctx).PUNTO.GetColumn())
						}
						break;
					}
					} 
				}
				setState(314);
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
		enterRule(_localctx, 38, RULE_primitivo);
		try {
			setState(329);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NUMBER:
				enterOuterAlt(_localctx, 1);
				{
				setState(315);
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
				setState(317);
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
				setState(319);
				((PrimitivoContext)_localctx).TRUE = match(TRUE);
				            
				        _localctx.p=expresion.NewPrimitivo(1,interfaces.BOOLEAN,((PrimitivoContext)_localctx).TRUE.GetColumn(),((PrimitivoContext)_localctx).TRUE.GetLine())      
				    
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 4);
				{
				setState(321);
				((PrimitivoContext)_localctx).FALSE = match(FALSE);

				        _localctx.p=expresion.NewPrimitivo(0,interfaces.BOOLEAN,((PrimitivoContext)_localctx).FALSE.GetColumn(),((PrimitivoContext)_localctx).FALSE.GetLine())      
				    
				}
				break;
			case DECIMAL:
				enterOuterAlt(_localctx, 5);
				{
				setState(323);
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
				setState(325);
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
				setState(327);
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
		case 18:
			return expr_arit_sempred((Expr_aritContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_arit_sempred(Expr_aritContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 7);
		case 1:
			return precpred(_ctx, 6);
		case 2:
			return precpred(_ctx, 5);
		case 3:
			return precpred(_ctx, 4);
		case 4:
			return precpred(_ctx, 11);
		case 5:
			return precpred(_ctx, 10);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3X\u014e\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\3\2\3\2\3\2\3\3\7\3/\n\3\f\3\16\3\62\13"+
		"\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\5\4Z\n\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\7\3\7\3\7\5\7l\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\5\by\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\5\t\u008b\n\t\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u009f\n\13\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u00bd\n\r\3\16\3\16\3\16\3\16\3\16\3\16\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3"+
		"\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20\u00dd\n\20\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\5\21\u00e7\n\21\3\22\3\22\3\22\3\22"+
		"\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\5\24\u0118\n\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\7\24\u0139\n\24\f\24\16\24\u013c\13"+
		"\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3"+
		"\25\5\25\u014c\n\25\3\25\2\3&\26\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36"+
		" \"$&(\2\6\4\2\24\25\30\30\3\2\26\27\4\2\13\r\17\21\3\2\22\23\2\u0169"+
		"\2*\3\2\2\2\4\60\3\2\2\2\6Y\3\2\2\2\b[\3\2\2\2\n`\3\2\2\2\fk\3\2\2\2\16"+
		"x\3\2\2\2\20\u008a\3\2\2\2\22\u008c\3\2\2\2\24\u009e\3\2\2\2\26\u00a0"+
		"\3\2\2\2\30\u00bc\3\2\2\2\32\u00be\3\2\2\2\34\u00c4\3\2\2\2\36\u00dc\3"+
		"\2\2\2 \u00e6\3\2\2\2\"\u00e8\3\2\2\2$\u00ec\3\2\2\2&\u0117\3\2\2\2(\u014b"+
		"\3\2\2\2*+\5\4\3\2+,\b\2\1\2,\3\3\2\2\2-/\5\6\4\2.-\3\2\2\2/\62\3\2\2"+
		"\2\60.\3\2\2\2\60\61\3\2\2\2\61\63\3\2\2\2\62\60\3\2\2\2\63\64\b\3\1\2"+
		"\64\5\3\2\2\2\65\66\7#\2\2\66\67\7\3\2\2\678\7$\2\289\7\32\2\29:\5$\23"+
		"\2:;\7\33\2\2;<\7\4\2\2<=\b\4\1\2=Z\3\2\2\2>?\5\n\6\2?@\7\4\2\2@A\b\4"+
		"\1\2AZ\3\2\2\2BC\5\b\5\2CD\7\4\2\2DE\b\4\1\2EZ\3\2\2\2FG\7*\2\2GH\5$\23"+
		"\2HI\7\34\2\2IJ\5\4\3\2JK\7\35\2\2KL\b\4\1\2LZ\3\2\2\2MN\5 \21\2NO\b\4"+
		"\1\2OZ\3\2\2\2PQ\5\"\22\2QR\b\4\1\2RZ\3\2\2\2ST\5\34\17\2TU\b\4\1\2UZ"+
		"\3\2\2\2VW\5\32\16\2WX\b\4\1\2XZ\3\2\2\2Y\65\3\2\2\2Y>\3\2\2\2YB\3\2\2"+
		"\2YF\3\2\2\2YM\3\2\2\2YP\3\2\2\2YS\3\2\2\2YV\3\2\2\2Z\7\3\2\2\2[\\\7V"+
		"\2\2\\]\7\t\2\2]^\5$\23\2^_\b\5\1\2_\t\3\2\2\2`a\7\63\2\2ab\5\f\7\2bc"+
		"\7V\2\2cd\5\16\b\2de\7\t\2\2ef\5$\23\2fg\b\6\1\2g\13\3\2\2\2hi\7:\2\2"+
		"il\b\7\1\2jl\b\7\1\2kh\3\2\2\2kj\3\2\2\2l\r\3\2\2\2mn\7\b\2\2no\5\26\f"+
		"\2op\b\b\1\2py\3\2\2\2qr\5\22\n\2rs\b\b\1\2sy\3\2\2\2tu\7\b\2\2uv\5\20"+
		"\t\2vw\b\b\1\2wy\3\2\2\2xm\3\2\2\2xq\3\2\2\2xt\3\2\2\2y\17\3\2\2\2z{\7"+
		"\64\2\2{\u008b\b\t\1\2|}\7\65\2\2}\u008b\b\t\1\2~\177\7\66\2\2\177\u008b"+
		"\b\t\1\2\u0080\u0081\7\67\2\2\u0081\u008b\b\t\1\2\u0082\u0083\78\2\2\u0083"+
		"\u008b\b\t\1\2\u0084\u0085\7\'\2\2\u0085\u008b\b\t\1\2\u0086\u0087\79"+
		"\2\2\u0087\u008b\b\t\1\2\u0088\u0089\7V\2\2\u0089\u008b\b\t\1\2\u008a"+
		"z\3\2\2\2\u008a|\3\2\2\2\u008a~\3\2\2\2\u008a\u0080\3\2\2\2\u008a\u0082"+
		"\3\2\2\2\u008a\u0084\3\2\2\2\u008a\u0086\3\2\2\2\u008a\u0088\3\2\2\2\u008b"+
		"\21\3\2\2\2\u008c\u008d\7\b\2\2\u008d\u008e\5\24\13\2\u008e\u008f\b\n"+
		"\1\2\u008f\23\3\2\2\2\u0090\u0091\7\36\2\2\u0091\u0092\5\20\t\2\u0092"+
		"\u0093\7\4\2\2\u0093\u0094\5$\23\2\u0094\u0095\7\37\2\2\u0095\u0096\b"+
		"\13\1\2\u0096\u009f\3\2\2\2\u0097\u0098\7\36\2\2\u0098\u0099\5\24\13\2"+
		"\u0099\u009a\7\4\2\2\u009a\u009b\5$\23\2\u009b\u009c\7\37\2\2\u009c\u009d"+
		"\b\13\1\2\u009d\u009f\3\2\2\2\u009e\u0090\3\2\2\2\u009e\u0097\3\2\2\2"+
		"\u009f\25\3\2\2\2\u00a0\u00a1\7O\2\2\u00a1\u00a2\7\13\2\2\u00a2\u00a3"+
		"\5\30\r\2\u00a3\u00a4\7\21\2\2\u00a4\u00a5\b\f\1\2\u00a5\27\3\2\2\2\u00a6"+
		"\u00a7\7\64\2\2\u00a7\u00bd\b\r\1\2\u00a8\u00a9\7\65\2\2\u00a9\u00bd\b"+
		"\r\1\2\u00aa\u00ab\7\66\2\2\u00ab\u00bd\b\r\1\2\u00ac\u00ad\7\67\2\2\u00ad"+
		"\u00bd\b\r\1\2\u00ae\u00af\78\2\2\u00af\u00bd\b\r\1\2\u00b0\u00b1\7\'"+
		"\2\2\u00b1\u00bd\b\r\1\2\u00b2\u00b3\79\2\2\u00b3\u00bd\b\r\1\2\u00b4"+
		"\u00b5\7V\2\2\u00b5\u00bd\b\r\1\2\u00b6\u00b7\7O\2\2\u00b7\u00b8\7\13"+
		"\2\2\u00b8\u00b9\5\30\r\2\u00b9\u00ba\7\21\2\2\u00ba\u00bb\b\r\1\2\u00bb"+
		"\u00bd\3\2\2\2\u00bc\u00a6\3\2\2\2\u00bc\u00a8\3\2\2\2\u00bc\u00aa\3\2"+
		"\2\2\u00bc\u00ac\3\2\2\2\u00bc\u00ae\3\2\2\2\u00bc\u00b0\3\2\2\2\u00bc"+
		"\u00b2\3\2\2\2\u00bc\u00b4\3\2\2\2\u00bc\u00b6\3\2\2\2\u00bd\31\3\2\2"+
		"\2\u00be\u00bf\7B\2\2\u00bf\u00c0\7\34\2\2\u00c0\u00c1\5\4\3\2\u00c1\u00c2"+
		"\7\35\2\2\u00c2\u00c3\b\16\1\2\u00c3\33\3\2\2\2\u00c4\u00c5\7(\2\2\u00c5"+
		"\u00c6\5$\23\2\u00c6\u00c7\7\34\2\2\u00c7\u00c8\5\4\3\2\u00c8\u00c9\7"+
		"\35\2\2\u00c9\u00ca\5\36\20\2\u00ca\u00cb\b\17\1\2\u00cb\35\3\2\2\2\u00cc"+
		"\u00cd\7)\2\2\u00cd\u00ce\7\34\2\2\u00ce\u00cf\5\4\3\2\u00cf\u00d0\7\35"+
		"\2\2\u00d0\u00d1\b\20\1\2\u00d1\u00dd\3\2\2\2\u00d2\u00d3\7)\2\2\u00d3"+
		"\u00d4\7(\2\2\u00d4\u00d5\5$\23\2\u00d5\u00d6\7\34\2\2\u00d6\u00d7\5\4"+
		"\3\2\u00d7\u00d8\7\35\2\2\u00d8\u00d9\5\36\20\2\u00d9\u00da\b\20\1\2\u00da"+
		"\u00dd\3\2\2\2\u00db\u00dd\b\20\1\2\u00dc\u00cc\3\2\2\2\u00dc\u00d2\3"+
		"\2\2\2\u00dc\u00db\3\2\2\2\u00dd\37\3\2\2\2\u00de\u00df\7\60\2\2\u00df"+
		"\u00e0\5$\23\2\u00e0\u00e1\7\4\2\2\u00e1\u00e2\b\21\1\2\u00e2\u00e7\3"+
		"\2\2\2\u00e3\u00e4\7\60\2\2\u00e4\u00e5\7\4\2\2\u00e5\u00e7\b\21\1\2\u00e6"+
		"\u00de\3\2\2\2\u00e6\u00e3\3\2\2\2\u00e7!\3\2\2\2\u00e8\u00e9\7\61\2\2"+
		"\u00e9\u00ea\7\4\2\2\u00ea\u00eb\b\22\1\2\u00eb#\3\2\2\2\u00ec\u00ed\5"+
		"&\24\2\u00ed\u00ee\b\23\1\2\u00ee%\3\2\2\2\u00ef\u00f0\b\24\1\2\u00f0"+
		"\u00f1\7\64\2\2\u00f1\u00f2\7\"\2\2\u00f2\u00f3\7/\2\2\u00f3\u00f4\7\32"+
		"\2\2\u00f4\u00f5\5&\24\2\u00f5\u00f6\7\5\2\2\u00f6\u00f7\5&\24\2\u00f7"+
		"\u00f8\7\33\2\2\u00f8\u00f9\b\24\1\2\u00f9\u0118\3\2\2\2\u00fa\u00fb\7"+
		"\65\2\2\u00fb\u00fc\7\"\2\2\u00fc\u00fd\7.\2\2\u00fd\u00fe\7\32\2\2\u00fe"+
		"\u00ff\5&\24\2\u00ff\u0100\7\5\2\2\u0100\u0101\5&\24\2\u0101\u0102\7\33"+
		"\2\2\u0102\u0103\b\24\1\2\u0103\u0118\3\2\2\2\u0104\u0105\7\7\2\2\u0105"+
		"\u0106\5&\24\13\u0106\u0107\b\24\1\2\u0107\u0118\3\2\2\2\u0108\u0109\7"+
		"\27\2\2\u0109\u010a\5&\24\n\u010a\u010b\b\24\1\2\u010b\u0118\3\2\2\2\u010c"+
		"\u010d\5(\25\2\u010d\u010e\b\24\1\2\u010e\u0118\3\2\2\2\u010f\u0110\7"+
		"\32\2\2\u0110\u0111\5$\23\2\u0111\u0112\7\33\2\2\u0112\u0113\b\24\1\2"+
		"\u0113\u0118\3\2\2\2\u0114\u0115\5\32\16\2\u0115\u0116\b\24\1\2\u0116"+
		"\u0118\3\2\2\2\u0117\u00ef\3\2\2\2\u0117\u00fa\3\2\2\2\u0117\u0104\3\2"+
		"\2\2\u0117\u0108\3\2\2\2\u0117\u010c\3\2\2\2\u0117\u010f\3\2\2\2\u0117"+
		"\u0114\3\2\2\2\u0118\u013a\3\2\2\2\u0119\u011a\f\t\2\2\u011a\u011b\t\2"+
		"\2\2\u011b\u011c\5&\24\n\u011c\u011d\b\24\1\2\u011d\u0139\3\2\2\2\u011e"+
		"\u011f\f\b\2\2\u011f\u0120\t\3\2\2\u0120\u0121\5&\24\t\u0121\u0122\b\24"+
		"\1\2\u0122\u0139\3\2\2\2\u0123\u0124\f\7\2\2\u0124\u0125\t\4\2\2\u0125"+
		"\u0126\5&\24\b\u0126\u0127\b\24\1\2\u0127\u0139\3\2\2\2\u0128\u0129\f"+
		"\6\2\2\u0129\u012a\t\5\2\2\u012a\u012b\5&\24\7\u012b\u012c\b\24\1\2\u012c"+
		"\u0139\3\2\2\2\u012d\u012e\f\r\2\2\u012e\u012f\7-\2\2\u012f\u0130\5\20"+
		"\t\2\u0130\u0131\b\24\1\2\u0131\u0139\3\2\2\2\u0132\u0133\f\f\2\2\u0133"+
		"\u0134\7\3\2\2\u0134\u0135\7D\2\2\u0135\u0136\7\32\2\2\u0136\u0137\7\33"+
		"\2\2\u0137\u0139\b\24\1\2\u0138\u0119\3\2\2\2\u0138\u011e\3\2\2\2\u0138"+
		"\u0123\3\2\2\2\u0138\u0128\3\2\2\2\u0138\u012d\3\2\2\2\u0138\u0132\3\2"+
		"\2\2\u0139\u013c\3\2\2\2\u013a\u0138\3\2\2\2\u013a\u013b\3\2\2\2\u013b"+
		"\'\3\2\2\2\u013c\u013a\3\2\2\2\u013d\u013e\7S\2\2\u013e\u014c\b\25\1\2"+
		"\u013f\u0140\7U\2\2\u0140\u014c\b\25\1\2\u0141\u0142\7?\2\2\u0142\u014c"+
		"\b\25\1\2\u0143\u0144\7@\2\2\u0144\u014c\b\25\1\2\u0145\u0146\7T\2\2\u0146"+
		"\u014c\b\25\1\2\u0147\u0148\7W\2\2\u0148\u014c\b\25\1\2\u0149\u014a\7"+
		"V\2\2\u014a\u014c\b\25\1\2\u014b\u013d\3\2\2\2\u014b\u013f\3\2\2\2\u014b"+
		"\u0141\3\2\2\2\u014b\u0143\3\2\2\2\u014b\u0145\3\2\2\2\u014b\u0147\3\2"+
		"\2\2\u014b\u0149\3\2\2\2\u014c)\3\2\2\2\17\60Ykx\u008a\u009e\u00bc\u00dc"+
		"\u00e6\u0117\u0138\u013a\u014b";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}