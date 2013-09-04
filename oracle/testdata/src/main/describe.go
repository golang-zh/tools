package describe // @describe pkgdecl "describe"

// Tests of 'describe' query.
// See go.tools/oracle/oracle_test.go for explanation.
// See describe.golden for expected query results.

// TODO(adonovan): more coverage of the (extensive) logic.

type cake float64 // @describe type-ref-builtin "float64"

const pi = 3.141     // @describe const-def-pi "pi"
const pie = cake(pi) // @describe const-def-pie "pie"
const _ = pi         // @describe const-ref-pi "pi"

func main() { // @describe func-def-main "main"
	// func objects
	_ = main   // @describe func-ref-main "main"
	_ = (*C).f // @describe func-ref-*C.f "..C..f"
	_ = D.f    // @describe func-ref-D.f "D.f"
	_ = I.f    // @describe func-ref-I.f "I.f"
	var d D    // @describe type-D "D"
	var i I    // @describe type-I "I"
	_ = d.f    // @describe func-ref-d.f "d.f"
	_ = i.f    // @describe func-ref-i.f "i.f"

	// var objects
	anon := func() {
		_ = d // @describe ref-lexical-d "d"
	}
	_ = anon // @describe ref-anon "anon"

	// SSA affords some local flow sensitivity.
	var a, b int
	var x = &a // @describe var-def-x-1 "x"
	_ = x      // @describe var-ref-x-1 "x"
	x = &b     // @describe var-def-x-2 "x"
	_ = x      // @describe var-ref-x-2 "x"

	// const objects
	const localpi = 3.141     // @describe const-local-pi "localpi"
	const localpie = cake(pi) // @describe const-local-pie "localpie"
	const _ = localpi         // @describe const-ref-localpi "localpi"

	// type objects
	type T int      // @describe type-def-T "T"
	var three T = 3 // @describe type-ref-T "T"

	print(1 + 2*3)        // @describe const-expr " 2.3"
	print(real(1+2i) - 3) // @describe const-expr2 "real.*3"

	m := map[string]*int{"a": &a}
	// TODO(adonovan): fix spurious error in map-lookup,ok result.
	mapval, _ := m["a"] // @describe map-lookup,ok "m..a.."
	_ = mapval          // @describe mapval "mapval"
	_ = m               // @describe m "m"

	defer main() // @describe defer-stmt "defer"
	go main()    // @describe go-stmt "go"
}

func deadcode() {
	var a int // @describe var-decl-stmt "var a int"
	// Pointer analysis can't run on dead code.
	var b = &a // @describe b "b"
}

type I interface { // @describe def-iface-I "I"
	f() // @describe def-imethod-I.f "f"
}

type C int
type D struct{}

func (c *C) f() {}
func (d D) f()  {}