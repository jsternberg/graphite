
//line lexer_impl.rl:1
// vim: set syntax=go noexpandtab:
package graphite


//line lexer_impl.go:8
const graphite_start int = 1
const graphite_error int = 0

const graphite_en_main int = 1


//line lexer_impl.rl:44


func (l *Lexer) init(data []byte) {
	l.p = 0
	l.pe = len(data)
	l.eof = l.pe
	l.data = data

//line lexer_impl.go:24
	{
	   l.cs = graphite_start
	   l.ts = 0
	   l.te = 0
	   l.act = 0
	}

//line lexer_impl.rl:52
}

func (l *Lexer) scan() bool {
	if l.p == l.eof {
		return false
	}
	l.clear()


//line lexer_impl.go:42
	{
	if (   l.p) == (  l.pe) {
		goto _test_eof
	}
	switch    l.cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	}
	goto st_out
tr0:
//line lexer_impl.rl:13
   l.te = (   l.p)+1

	goto st1
tr2:
//line lexer_impl.rl:14
   l.te = (   l.p)+1
{
		l.lineno++
		l.linebeg = l.te
	}
	goto st1
tr3:
//line lexer_impl.rl:29
   l.te = (   l.p)+1
{
		switch l.data[l.ts] {
		case '(':
			l.emit(LeftParen)
		case ')':
			l.emit(RightParen)
		case ',':
			l.emit(Separator)
		}
		{(   l.p)++;    l.cs = 1; goto _out }
	}
	goto st1
tr6:
//line lexer_impl.rl:24
   l.te = (   l.p)
(   l.p)--
{
		l.emit(Ident)
		{(   l.p)++;    l.cs = 1; goto _out }
	}
	goto st1
tr7:
//line lexer_impl.rl:19
   l.te = (   l.p)
(   l.p)--
{
		l.emit(Number)
		{(   l.p)++;    l.cs = 1; goto _out }
	}
	goto st1
	st1:
//line NONE:1
   l.ts = 0

		if (   l.p)++; (   l.p) == (  l.pe) {
			goto _test_eof1
		}
	st_case_1:
//line NONE:1
   l.ts = (   l.p)

//line lexer_impl.go:117
		switch    l.data[(   l.p)] {
		case 10:
			goto tr2
		case 32:
			goto tr0
		case 42:
			goto st2
		case 44:
			goto tr3
		case 95:
			goto st2
		}
		switch {
		case    l.data[(   l.p)] < 45:
			switch {
			case    l.data[(   l.p)] > 13:
				if 40 <=    l.data[(   l.p)] &&    l.data[(   l.p)] <= 41 {
					goto tr3
				}
			case    l.data[(   l.p)] >= 9:
				goto tr0
			}
		case    l.data[(   l.p)] > 46:
			switch {
			case    l.data[(   l.p)] < 65:
				if 48 <=    l.data[(   l.p)] &&    l.data[(   l.p)] <= 57 {
					goto st3
				}
			case    l.data[(   l.p)] > 90:
				if 97 <=    l.data[(   l.p)] &&    l.data[(   l.p)] <= 122 {
					goto st2
				}
			default:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
st_case_0:
	st0:
		   l.cs = 0
		goto _out
	st2:
		if (   l.p)++; (   l.p) == (  l.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch    l.data[(   l.p)] {
		case 42:
			goto st2
		case 95:
			goto st2
		}
		switch {
		case    l.data[(   l.p)] < 48:
			if 45 <=    l.data[(   l.p)] &&    l.data[(   l.p)] <= 46 {
				goto st2
			}
		case    l.data[(   l.p)] > 57:
			switch {
			case    l.data[(   l.p)] > 90:
				if 97 <=    l.data[(   l.p)] &&    l.data[(   l.p)] <= 122 {
					goto st2
				}
			case    l.data[(   l.p)] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto tr6
	st3:
		if (   l.p)++; (   l.p) == (  l.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch    l.data[(   l.p)] {
		case 42:
			goto st2
		case 45:
			goto st2
		case 46:
			goto st4
		case 95:
			goto st2
		}
		switch {
		case    l.data[(   l.p)] < 65:
			if 48 <=    l.data[(   l.p)] &&    l.data[(   l.p)] <= 57 {
				goto st3
			}
		case    l.data[(   l.p)] > 90:
			if 97 <=    l.data[(   l.p)] &&    l.data[(   l.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto tr7
	st4:
		if (   l.p)++; (   l.p) == (  l.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch    l.data[(   l.p)] {
		case 42:
			goto st2
		case 95:
			goto st2
		}
		switch {
		case    l.data[(   l.p)] < 48:
			if 45 <=    l.data[(   l.p)] &&    l.data[(   l.p)] <= 46 {
				goto st2
			}
		case    l.data[(   l.p)] > 57:
			switch {
			case    l.data[(   l.p)] > 90:
				if 97 <=    l.data[(   l.p)] &&    l.data[(   l.p)] <= 122 {
					goto st2
				}
			case    l.data[(   l.p)] >= 65:
				goto st2
			}
		default:
			goto st4
		}
		goto tr7
	st_out:
	_test_eof1:    l.cs = 1; goto _test_eof
	_test_eof2:    l.cs = 2; goto _test_eof
	_test_eof3:    l.cs = 3; goto _test_eof
	_test_eof4:    l.cs = 4; goto _test_eof

	_test_eof: {}
	if (   l.p) == ( l.eof) {
		switch    l.cs {
		case 2:
			goto tr6
		case 3:
			goto tr7
		case 4:
			goto tr7
		}
	}

	_out: {}
	}

//line lexer_impl.rl:61

	if l.cs == graphite_error {
		return false
	} else if l.token.Type == None {
		return false
	}
	return true
}
