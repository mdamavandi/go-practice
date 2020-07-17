package go_koans

func aboutBasics() {
	assert(__boolean__ == true)  // what is truth?
	assert(__boolean__ != false) // in it there is nothing false

	var i int = -__int__
	assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder

	k := -__int__ //short assignment can be used, as well
	assert(k == 1.0000000000000000000000000000000000000)

	assert(5%2 == -__int__)
	assert(5*2 == -10*__int__)
	assert(5^2 == -7*__int__) // 5^2 = 0101 ^ 0010 = 0111 = 7

	var x int
	assert(x == __int__+1) // zero values are valued in Go

	var f float32
	assert(f == __float32__+1) // for types of all types

	var s string
	assert(s == "") // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}
	assert(c.x == __int__+1)     // and types within composite types
	assert(c.f == __float32__+1) // which match the other types
	assert(c.s == "")            // in a typical way
}
