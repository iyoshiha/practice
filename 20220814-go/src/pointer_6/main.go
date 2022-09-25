package main

import(
	. "fmt"
)

type Foo struct {
	a int
	b *int
	c int
}

func intp(num int) *int{
	tmp := num
	return &tmp
}

func main() {
	// true 
	var x *int
	Println(x == nil)

	// false
	x = new(int)
	Println(x == nil)
	Println(x)
	Println(*x)

	// false
	// you cant use an & before a primitive time 
	// when you need a primitive type pointer
	// declare a variable and point to it
	var y int
	x = &y
	Println(x == nil)

	xx := &Foo{}
	Println(xx == nil)

	p := Foo{
		a: 1,
		b: intp(2),
		c: 3,
	}
	Println(p )

	// another way to write 
	p = Foo{
		a: 1,
		b: intp(4),
		c: 3,
	}
	Println(p)
	var ap *int
	Println(ap)
}

// dont do this
func badMakeFoo(f *Foo) error {
		f.a = 1
		f.b = intp(2)
	return nil
}

// do this 
func goodMakeFoo() (Foo, error) {
	f := Foo{
		a: 1,
		b: intp(2)),
	}
	return f, nil
}




