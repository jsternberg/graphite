// vim: set syntax=go noexpandtab:
package graphite

%%{
  machine graphite;
  access   l.;
  variable p   l.p;
  variable pe  l.pe;
  variable eof l.eof;

  main := |*

	(space - '\n');
	'\n' {
		l.lineno++
		l.linebeg = l.te
	};

	digit+ . ('.' . digit*)? {
		l.emit(Number)
		fbreak;
	};

	(alnum | [_\-.*])+ {
		l.emit(Ident)
		fbreak;
	};

	[(),] {
		switch l.data[l.ts] {
		case '(':
			l.emit(LeftParen)
		case ')':
			l.emit(RightParen)
		case ',':
			l.emit(Separator)
		}
		fbreak;
	};

  *|;

  write data nofinal;
}%%

func (l *Lexer) init(data []byte) {
	l.p = 0
	l.pe = len(data)
	l.eof = l.pe
	l.data = data
%% write init;
}

func (l *Lexer) scan() bool {
	if l.p == l.eof {
		return false
	}
	l.clear()

%% write exec;

	if l.cs == graphite_error {
		return false
	} else if l.token.Type == None {
		return false
	}
	return true
}
