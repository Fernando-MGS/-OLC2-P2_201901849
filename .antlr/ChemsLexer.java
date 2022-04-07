// Generated from c:\Users\ferna\OneDrive\Documentos\Go\src\github.com\Fernando-MGS\-OLC2-P2_201901849\ChemsLexer.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ChemsLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PUNTO", "PTCOMA", "COMA", "COMILLAS", "DIFERENTE", "DOSPUNTOS", "IGUAL", 
			"AMPER", "MENOR", "MAYORIGUAL", "MENORIGUAL", "OPMATCH", "D_IGUAL", "NOT_E", 
			"MAYOR", "OR", "AND", "MUL", "DIV", "ADD", "SUB", "MOD", "POT", "PARIZQ", 
			"PARDER", "LLAVEIZQ", "LLAVEDER", "CORIZQ", "CORDER", "INTDER", "OR_COND", 
			"DDPUNTO", "CONSOLE", "LOG", "PRINTLN", "P_NUMBER", "P_STRING", "P_IF", 
			"P_ELSE", "P_WHILE", "P_FOR", "P_IN", "P_AS", "POWF", "POW", "BREAK", 
			"CONTINUE", "RETURN", "LET", "INT", "FLOAT", "BOOL", "CHAR", "STR", "USIZE", 
			"MUT", "P_STRUCT", "FUNCTION", "MODULE", "PUB", "TRUE", "FALSE", "MATCH", 
			"LOOP", "PUSH", "T_STRING", "F_ABS", "F_SQRT", "F_CLONE", "INSERT", "REMOVE", 
			"CONTAINS", "LEN", "CAPACITY", "DEF", "VEC", "VECT", "CAP", "NEW", "COMENTARIO", 
			"NUMBER", "DECIMAL", "STRING", "ID", "CARACTER", "WHITESPACE", "ESC_SEQ"
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


	public ChemsLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ChemsLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2X\u023c\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\4M\tM\4N\tN\4O\tO\4P\tP\4Q\tQ\4R\tR\4S\tS\4T\tT"+
		"\4U\tU\4V\tV\4W\tW\4X\tX\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3"+
		"\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\r\3\16"+
		"\3\16\3\16\3\17\3\17\3\17\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\22\3\23"+
		"\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32"+
		"\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3"+
		"!\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3$\3$\3$\3$\3$\3$\3$\3$"+
		"\3$\3%\3%\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3&\3&\3\'\3\'\3\'\3(\3(\3(\3("+
		"\3(\3)\3)\3)\3)\3)\3)\3*\3*\3*\3*\3+\3+\3+\3,\3,\3,\3-\3-\3-\3-\3-\3."+
		"\3.\3.\3.\3/\3/\3/\3/\3/\3/\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3"+
		"\60\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\62\3\62\3\62\3\62\3\63\3\63\3"+
		"\63\3\63\3\64\3\64\3\64\3\64\3\65\3\65\3\65\3\65\3\65\3\66\3\66\3\66\3"+
		"\66\3\66\3\67\3\67\3\67\3\67\3\67\38\38\38\38\38\38\39\39\39\39\3:\3:"+
		"\3:\3:\3:\3:\3:\3;\3;\3;\3<\3<\3<\3<\3=\3=\3=\3=\3>\3>\3>\3>\3>\3?\3?"+
		"\3?\3?\3?\3?\3@\3@\3@\3@\3@\3@\3A\3A\3A\3A\3A\3B\3B\3B\3B\3B\3C\3C\3C"+
		"\3C\3C\3C\3C\3C\3C\3C\3D\3D\3D\3D\3E\3E\3E\3E\3E\3F\3F\3F\3F\3F\3F\3G"+
		"\3G\3G\3G\3G\3G\3G\3H\3H\3H\3H\3H\3H\3H\3I\3I\3I\3I\3I\3I\3I\3I\3I\3J"+
		"\3J\3J\3J\3K\3K\3K\3K\3K\3K\3K\3K\3K\3L\3L\3M\3M\3M\3M\3M\3N\3N\3N\3N"+
		"\3O\3O\3O\3O\3O\3O\3O\3O\3O\3O\3O\3O\3O\3O\3P\3P\3P\3P\3Q\3Q\3Q\7Q\u0203"+
		"\nQ\fQ\16Q\u0206\13Q\3Q\3Q\3Q\3Q\3R\6R\u020d\nR\rR\16R\u020e\3S\6S\u0212"+
		"\nS\rS\16S\u0213\3S\3S\6S\u0218\nS\rS\16S\u0219\3T\3T\7T\u021e\nT\fT\16"+
		"T\u0221\13T\3T\3T\3U\3U\7U\u0227\nU\fU\16U\u022a\13U\3V\3V\5V\u022e\n"+
		"V\3V\3V\3V\3W\6W\u0234\nW\rW\16W\u0235\3W\3W\3X\3X\3X\3\u0204\2Y\3\3\5"+
		"\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21"+
		"!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!"+
		"A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s"+
		";u<w=y>{?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089F\u008bG\u008dH\u008f"+
		"I\u0091J\u0093K\u0095L\u0097M\u0099N\u009bO\u009dP\u009fQ\u00a1R\u00a3"+
		"S\u00a5T\u00a7U\u00a9V\u00abW\u00adX\u00af\2\3\2\13\3\2\62;\3\2\60\60"+
		"\3\2$$\5\2C\\aac|\6\2\62;C\\aac|\3\2))\3\2^^\6\2\13\f\17\17\"\"^^\t\2"+
		"\"#%%--/\60<<BB]_\2\u0242\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2"+
		"\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2"+
		"\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3"+
		"\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2"+
		"\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67"+
		"\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2"+
		"\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2"+
		"\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]"+
		"\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2"+
		"\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2"+
		"\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2"+
		"\2\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2"+
		"\u008b\3\2\2\2\2\u008d\3\2\2\2\2\u008f\3\2\2\2\2\u0091\3\2\2\2\2\u0093"+
		"\3\2\2\2\2\u0095\3\2\2\2\2\u0097\3\2\2\2\2\u0099\3\2\2\2\2\u009b\3\2\2"+
		"\2\2\u009d\3\2\2\2\2\u009f\3\2\2\2\2\u00a1\3\2\2\2\2\u00a3\3\2\2\2\2\u00a5"+
		"\3\2\2\2\2\u00a7\3\2\2\2\2\u00a9\3\2\2\2\2\u00ab\3\2\2\2\2\u00ad\3\2\2"+
		"\2\3\u00b1\3\2\2\2\5\u00b3\3\2\2\2\7\u00b5\3\2\2\2\t\u00b7\3\2\2\2\13"+
		"\u00b9\3\2\2\2\r\u00bb\3\2\2\2\17\u00bd\3\2\2\2\21\u00bf\3\2\2\2\23\u00c1"+
		"\3\2\2\2\25\u00c3\3\2\2\2\27\u00c6\3\2\2\2\31\u00c9\3\2\2\2\33\u00cc\3"+
		"\2\2\2\35\u00cf\3\2\2\2\37\u00d2\3\2\2\2!\u00d4\3\2\2\2#\u00d7\3\2\2\2"+
		"%\u00da\3\2\2\2\'\u00dc\3\2\2\2)\u00de\3\2\2\2+\u00e0\3\2\2\2-\u00e2\3"+
		"\2\2\2/\u00e4\3\2\2\2\61\u00e6\3\2\2\2\63\u00e8\3\2\2\2\65\u00ea\3\2\2"+
		"\2\67\u00ec\3\2\2\29\u00ee\3\2\2\2;\u00f0\3\2\2\2=\u00f2\3\2\2\2?\u00f4"+
		"\3\2\2\2A\u00f6\3\2\2\2C\u00f9\3\2\2\2E\u0101\3\2\2\2G\u0105\3\2\2\2I"+
		"\u010e\3\2\2\2K\u0115\3\2\2\2M\u011c\3\2\2\2O\u011f\3\2\2\2Q\u0124\3\2"+
		"\2\2S\u012a\3\2\2\2U\u012e\3\2\2\2W\u0131\3\2\2\2Y\u0134\3\2\2\2[\u0139"+
		"\3\2\2\2]\u013d\3\2\2\2_\u0143\3\2\2\2a\u014c\3\2\2\2c\u0153\3\2\2\2e"+
		"\u0157\3\2\2\2g\u015b\3\2\2\2i\u015f\3\2\2\2k\u0164\3\2\2\2m\u0169\3\2"+
		"\2\2o\u016e\3\2\2\2q\u0174\3\2\2\2s\u0178\3\2\2\2u\u017f\3\2\2\2w\u0182"+
		"\3\2\2\2y\u0186\3\2\2\2{\u018a\3\2\2\2}\u018f\3\2\2\2\177\u0195\3\2\2"+
		"\2\u0081\u019b\3\2\2\2\u0083\u01a0\3\2\2\2\u0085\u01a5\3\2\2\2\u0087\u01af"+
		"\3\2\2\2\u0089\u01b3\3\2\2\2\u008b\u01b8\3\2\2\2\u008d\u01be\3\2\2\2\u008f"+
		"\u01c5\3\2\2\2\u0091\u01cc\3\2\2\2\u0093\u01d5\3\2\2\2\u0095\u01d9\3\2"+
		"\2\2\u0097\u01e2\3\2\2\2\u0099\u01e4\3\2\2\2\u009b\u01e9\3\2\2\2\u009d"+
		"\u01ed\3\2\2\2\u009f\u01fb\3\2\2\2\u00a1\u01ff\3\2\2\2\u00a3\u020c\3\2"+
		"\2\2\u00a5\u0211\3\2\2\2\u00a7\u021b\3\2\2\2\u00a9\u0224\3\2\2\2\u00ab"+
		"\u022b\3\2\2\2\u00ad\u0233\3\2\2\2\u00af\u0239\3\2\2\2\u00b1\u00b2\7\60"+
		"\2\2\u00b2\4\3\2\2\2\u00b3\u00b4\7=\2\2\u00b4\6\3\2\2\2\u00b5\u00b6\7"+
		".\2\2\u00b6\b\3\2\2\2\u00b7\u00b8\7$\2\2\u00b8\n\3\2\2\2\u00b9\u00ba\7"+
		"#\2\2\u00ba\f\3\2\2\2\u00bb\u00bc\7<\2\2\u00bc\16\3\2\2\2\u00bd\u00be"+
		"\7?\2\2\u00be\20\3\2\2\2\u00bf\u00c0\7(\2\2\u00c0\22\3\2\2\2\u00c1\u00c2"+
		"\7>\2\2\u00c2\24\3\2\2\2\u00c3\u00c4\7@\2\2\u00c4\u00c5\7?\2\2\u00c5\26"+
		"\3\2\2\2\u00c6\u00c7\7>\2\2\u00c7\u00c8\7?\2\2\u00c8\30\3\2\2\2\u00c9"+
		"\u00ca\7?\2\2\u00ca\u00cb\7@\2\2\u00cb\32\3\2\2\2\u00cc\u00cd\7?\2\2\u00cd"+
		"\u00ce\7?\2\2\u00ce\34\3\2\2\2\u00cf\u00d0\7#\2\2\u00d0\u00d1\7?\2\2\u00d1"+
		"\36\3\2\2\2\u00d2\u00d3\7@\2\2\u00d3 \3\2\2\2\u00d4\u00d5\7~\2\2\u00d5"+
		"\u00d6\7~\2\2\u00d6\"\3\2\2\2\u00d7\u00d8\7(\2\2\u00d8\u00d9\7(\2\2\u00d9"+
		"$\3\2\2\2\u00da\u00db\7,\2\2\u00db&\3\2\2\2\u00dc\u00dd\7\61\2\2\u00dd"+
		"(\3\2\2\2\u00de\u00df\7-\2\2\u00df*\3\2\2\2\u00e0\u00e1\7/\2\2\u00e1,"+
		"\3\2\2\2\u00e2\u00e3\7\'\2\2\u00e3.\3\2\2\2\u00e4\u00e5\7`\2\2\u00e5\60"+
		"\3\2\2\2\u00e6\u00e7\7*\2\2\u00e7\62\3\2\2\2\u00e8\u00e9\7+\2\2\u00e9"+
		"\64\3\2\2\2\u00ea\u00eb\7}\2\2\u00eb\66\3\2\2\2\u00ec\u00ed\7\177\2\2"+
		"\u00ed8\3\2\2\2\u00ee\u00ef\7]\2\2\u00ef:\3\2\2\2\u00f0\u00f1\7_\2\2\u00f1"+
		"<\3\2\2\2\u00f2\u00f3\7A\2\2\u00f3>\3\2\2\2\u00f4\u00f5\7~\2\2\u00f5@"+
		"\3\2\2\2\u00f6\u00f7\7<\2\2\u00f7\u00f8\7<\2\2\u00f8B\3\2\2\2\u00f9\u00fa"+
		"\7e\2\2\u00fa\u00fb\7q\2\2\u00fb\u00fc\7p\2\2\u00fc\u00fd\7u\2\2\u00fd"+
		"\u00fe\7q\2\2\u00fe\u00ff\7n\2\2\u00ff\u0100\7g\2\2\u0100D\3\2\2\2\u0101"+
		"\u0102\7n\2\2\u0102\u0103\7q\2\2\u0103\u0104\7i\2\2\u0104F\3\2\2\2\u0105"+
		"\u0106\7r\2\2\u0106\u0107\7t\2\2\u0107\u0108\7k\2\2\u0108\u0109\7p\2\2"+
		"\u0109\u010a\7v\2\2\u010a\u010b\7n\2\2\u010b\u010c\7p\2\2\u010c\u010d"+
		"\7#\2\2\u010dH\3\2\2\2\u010e\u010f\7p\2\2\u010f\u0110\7w\2\2\u0110\u0111"+
		"\7o\2\2\u0111\u0112\7d\2\2\u0112\u0113\7g\2\2\u0113\u0114\7t\2\2\u0114"+
		"J\3\2\2\2\u0115\u0116\7U\2\2\u0116\u0117\7v\2\2\u0117\u0118\7t\2\2\u0118"+
		"\u0119\7k\2\2\u0119\u011a\7p\2\2\u011a\u011b\7i\2\2\u011bL\3\2\2\2\u011c"+
		"\u011d\7k\2\2\u011d\u011e\7h\2\2\u011eN\3\2\2\2\u011f\u0120\7g\2\2\u0120"+
		"\u0121\7n\2\2\u0121\u0122\7u\2\2\u0122\u0123\7g\2\2\u0123P\3\2\2\2\u0124"+
		"\u0125\7y\2\2\u0125\u0126\7j\2\2\u0126\u0127\7k\2\2\u0127\u0128\7n\2\2"+
		"\u0128\u0129\7g\2\2\u0129R\3\2\2\2\u012a\u012b\7h\2\2\u012b\u012c\7q\2"+
		"\2\u012c\u012d\7t\2\2\u012dT\3\2\2\2\u012e\u012f\7k\2\2\u012f\u0130\7"+
		"p\2\2\u0130V\3\2\2\2\u0131\u0132\7c\2\2\u0132\u0133\7u\2\2\u0133X\3\2"+
		"\2\2\u0134\u0135\7r\2\2\u0135\u0136\7q\2\2\u0136\u0137\7y\2\2\u0137\u0138"+
		"\7h\2\2\u0138Z\3\2\2\2\u0139\u013a\7r\2\2\u013a\u013b\7q\2\2\u013b\u013c"+
		"\7y\2\2\u013c\\\3\2\2\2\u013d\u013e\7d\2\2\u013e\u013f\7t\2\2\u013f\u0140"+
		"\7g\2\2\u0140\u0141\7c\2\2\u0141\u0142\7m\2\2\u0142^\3\2\2\2\u0143\u0144"+
		"\7e\2\2\u0144\u0145\7q\2\2\u0145\u0146\7p\2\2\u0146\u0147\7v\2\2\u0147"+
		"\u0148\7k\2\2\u0148\u0149\7p\2\2\u0149\u014a\7w\2\2\u014a\u014b\7g\2\2"+
		"\u014b`\3\2\2\2\u014c\u014d\7t\2\2\u014d\u014e\7g\2\2\u014e\u014f\7v\2"+
		"\2\u014f\u0150\7w\2\2\u0150\u0151\7t\2\2\u0151\u0152\7p\2\2\u0152b\3\2"+
		"\2\2\u0153\u0154\7n\2\2\u0154\u0155\7g\2\2\u0155\u0156\7v\2\2\u0156d\3"+
		"\2\2\2\u0157\u0158\7k\2\2\u0158\u0159\78\2\2\u0159\u015a\7\66\2\2\u015a"+
		"f\3\2\2\2\u015b\u015c\7h\2\2\u015c\u015d\78\2\2\u015d\u015e\7\66\2\2\u015e"+
		"h\3\2\2\2\u015f\u0160\7d\2\2\u0160\u0161\7q\2\2\u0161\u0162\7q\2\2\u0162"+
		"\u0163\7n\2\2\u0163j\3\2\2\2\u0164\u0165\7e\2\2\u0165\u0166\7j\2\2\u0166"+
		"\u0167\7c\2\2\u0167\u0168\7t\2\2\u0168l\3\2\2\2\u0169\u016a\7(\2\2\u016a"+
		"\u016b\7u\2\2\u016b\u016c\7v\2\2\u016c\u016d\7t\2\2\u016dn\3\2\2\2\u016e"+
		"\u016f\7w\2\2\u016f\u0170\7u\2\2\u0170\u0171\7k\2\2\u0171\u0172\7|\2\2"+
		"\u0172\u0173\7g\2\2\u0173p\3\2\2\2\u0174\u0175\7o\2\2\u0175\u0176\7w\2"+
		"\2\u0176\u0177\7v\2\2\u0177r\3\2\2\2\u0178\u0179\7u\2\2\u0179\u017a\7"+
		"v\2\2\u017a\u017b\7t\2\2\u017b\u017c\7w\2\2\u017c\u017d\7e\2\2\u017d\u017e"+
		"\7v\2\2\u017et\3\2\2\2\u017f\u0180\7h\2\2\u0180\u0181\7p\2\2\u0181v\3"+
		"\2\2\2\u0182\u0183\7o\2\2\u0183\u0184\7q\2\2\u0184\u0185\7f\2\2\u0185"+
		"x\3\2\2\2\u0186\u0187\7r\2\2\u0187\u0188\7w\2\2\u0188\u0189\7d\2\2\u0189"+
		"z\3\2\2\2\u018a\u018b\7v\2\2\u018b\u018c\7t\2\2\u018c\u018d\7w\2\2\u018d"+
		"\u018e\7g\2\2\u018e|\3\2\2\2\u018f\u0190\7h\2\2\u0190\u0191\7c\2\2\u0191"+
		"\u0192\7n\2\2\u0192\u0193\7u\2\2\u0193\u0194\7g\2\2\u0194~\3\2\2\2\u0195"+
		"\u0196\7o\2\2\u0196\u0197\7c\2\2\u0197\u0198\7v\2\2\u0198\u0199\7e\2\2"+
		"\u0199\u019a\7j\2\2\u019a\u0080\3\2\2\2\u019b\u019c\7n\2\2\u019c\u019d"+
		"\7q\2\2\u019d\u019e\7q\2\2\u019e\u019f\7r\2\2\u019f\u0082\3\2\2\2\u01a0"+
		"\u01a1\7r\2\2\u01a1\u01a2\7w\2\2\u01a2\u01a3\7u\2\2\u01a3\u01a4\7j\2\2"+
		"\u01a4\u0084\3\2\2\2\u01a5\u01a6\7v\2\2\u01a6\u01a7\7q\2\2\u01a7\u01a8"+
		"\7a\2\2\u01a8\u01a9\7u\2\2\u01a9\u01aa\7v\2\2\u01aa\u01ab\7t\2\2\u01ab"+
		"\u01ac\7k\2\2\u01ac\u01ad\7p\2\2\u01ad\u01ae\7i\2\2\u01ae\u0086\3\2\2"+
		"\2\u01af\u01b0\7c\2\2\u01b0\u01b1\7d\2\2\u01b1\u01b2\7u\2\2\u01b2\u0088"+
		"\3\2\2\2\u01b3\u01b4\7u\2\2\u01b4\u01b5\7s\2\2\u01b5\u01b6\7t\2\2\u01b6"+
		"\u01b7\7v\2\2\u01b7\u008a\3\2\2\2\u01b8\u01b9\7e\2\2\u01b9\u01ba\7n\2"+
		"\2\u01ba\u01bb\7q\2\2\u01bb\u01bc\7p\2\2\u01bc\u01bd\7g\2\2\u01bd\u008c"+
		"\3\2\2\2\u01be\u01bf\7k\2\2\u01bf\u01c0\7p\2\2\u01c0\u01c1\7u\2\2\u01c1"+
		"\u01c2\7g\2\2\u01c2\u01c3\7t\2\2\u01c3\u01c4\7v\2\2\u01c4\u008e\3\2\2"+
		"\2\u01c5\u01c6\7t\2\2\u01c6\u01c7\7g\2\2\u01c7\u01c8\7o\2\2\u01c8\u01c9"+
		"\7q\2\2\u01c9\u01ca\7x\2\2\u01ca\u01cb\7g\2\2\u01cb\u0090\3\2\2\2\u01cc"+
		"\u01cd\7e\2\2\u01cd\u01ce\7q\2\2\u01ce\u01cf\7p\2\2\u01cf\u01d0\7v\2\2"+
		"\u01d0\u01d1\7c\2\2\u01d1\u01d2\7k\2\2\u01d2\u01d3\7p\2\2\u01d3\u01d4"+
		"\7u\2\2\u01d4\u0092\3\2\2\2\u01d5\u01d6\7n\2\2\u01d6\u01d7\7g\2\2\u01d7"+
		"\u01d8\7p\2\2\u01d8\u0094\3\2\2\2\u01d9\u01da\7e\2\2\u01da\u01db\7c\2"+
		"\2\u01db\u01dc\7r\2\2\u01dc\u01dd\7c\2\2\u01dd\u01de\7e\2\2\u01de\u01df"+
		"\7k\2\2\u01df\u01e0\7v\2\2\u01e0\u01e1\7{\2\2\u01e1\u0096\3\2\2\2\u01e2"+
		"\u01e3\7a\2\2\u01e3\u0098\3\2\2\2\u01e4\u01e5\7x\2\2\u01e5\u01e6\7g\2"+
		"\2\u01e6\u01e7\7e\2\2\u01e7\u01e8\7#\2\2\u01e8\u009a\3\2\2\2\u01e9\u01ea"+
		"\7X\2\2\u01ea\u01eb\7g\2\2\u01eb\u01ec\7e\2\2\u01ec\u009c\3\2\2\2\u01ed"+
		"\u01ee\7y\2\2\u01ee\u01ef\7k\2\2\u01ef\u01f0\7v\2\2\u01f0\u01f1\7j\2\2"+
		"\u01f1\u01f2\7a\2\2\u01f2\u01f3\7e\2\2\u01f3\u01f4\7c\2\2\u01f4\u01f5"+
		"\7r\2\2\u01f5\u01f6\7c\2\2\u01f6\u01f7\7e\2\2\u01f7\u01f8\7k\2\2\u01f8"+
		"\u01f9\7v\2\2\u01f9\u01fa\7{\2\2\u01fa\u009e\3\2\2\2\u01fb\u01fc\7p\2"+
		"\2\u01fc\u01fd\7g\2\2\u01fd\u01fe\7y\2\2\u01fe\u00a0\3\2\2\2\u01ff\u0200"+
		"\7\61\2\2\u0200\u0204\7\61\2\2\u0201\u0203\13\2\2\2\u0202\u0201\3\2\2"+
		"\2\u0203\u0206\3\2\2\2\u0204\u0205\3\2\2\2\u0204\u0202\3\2\2\2\u0205\u0207"+
		"\3\2\2\2\u0206\u0204\3\2\2\2\u0207\u0208\7\f\2\2\u0208\u0209\3\2\2\2\u0209"+
		"\u020a\bQ\2\2\u020a\u00a2\3\2\2\2\u020b\u020d\t\2\2\2\u020c\u020b\3\2"+
		"\2\2\u020d\u020e\3\2\2\2\u020e\u020c\3\2\2\2\u020e\u020f\3\2\2\2\u020f"+
		"\u00a4\3\2\2\2\u0210\u0212\t\2\2\2\u0211\u0210\3\2\2\2\u0212\u0213\3\2"+
		"\2\2\u0213\u0211\3\2\2\2\u0213\u0214\3\2\2\2\u0214\u0215\3\2\2\2\u0215"+
		"\u0217\t\3\2\2\u0216\u0218\t\2\2\2\u0217\u0216\3\2\2\2\u0218\u0219\3\2"+
		"\2\2\u0219\u0217\3\2\2\2\u0219\u021a\3\2\2\2\u021a\u00a6\3\2\2\2\u021b"+
		"\u021f\7$\2\2\u021c\u021e\n\4\2\2\u021d\u021c\3\2\2\2\u021e\u0221\3\2"+
		"\2\2\u021f\u021d\3\2\2\2\u021f\u0220\3\2\2\2\u0220\u0222\3\2\2\2\u0221"+
		"\u021f\3\2\2\2\u0222\u0223\7$\2\2\u0223\u00a8\3\2\2\2\u0224\u0228\t\5"+
		"\2\2\u0225\u0227\t\6\2\2\u0226\u0225\3\2\2\2\u0227\u022a\3\2\2\2\u0228"+
		"\u0226\3\2\2\2\u0228\u0229\3\2\2\2\u0229\u00aa\3\2\2\2\u022a\u0228\3\2"+
		"\2\2\u022b\u022d\t\7\2\2\u022c\u022e\t\b\2\2\u022d\u022c\3\2\2\2\u022d"+
		"\u022e\3\2\2\2\u022e\u022f\3\2\2\2\u022f\u0230\13\2\2\2\u0230\u0231\t"+
		"\7\2\2\u0231\u00ac\3\2\2\2\u0232\u0234\t\t\2\2\u0233\u0232\3\2\2\2\u0234"+
		"\u0235\3\2\2\2\u0235\u0233\3\2\2\2\u0235\u0236\3\2\2\2\u0236\u0237\3\2"+
		"\2\2\u0237\u0238\bW\2\2\u0238\u00ae\3\2\2\2\u0239\u023a\7^\2\2\u023a\u023b"+
		"\t\n\2\2\u023b\u00b0\3\2\2\2\13\2\u0204\u020e\u0213\u0219\u021f\u0228"+
		"\u022d\u0235\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}