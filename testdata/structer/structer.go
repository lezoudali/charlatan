// generated by "charlatan -dir=testdata/structer -output=testdata/structer/structer.go Structer".  DO NOT EDIT.

package main

import "reflect"

// StructerStructInvocation represents a single call of FakeStructer.Struct
type StructerStructInvocation struct {
	Parameters struct {
		Ident1 struct {
			a string
			b string
		}
	}
	Results struct {
		Ident2 struct {
			c string
			d string
		}
	}
}

// NewStructerStructInvocation creates a new instance of StructerStructInvocation
func NewStructerStructInvocation(ident1 struct {
	a string
	b string
}, ident2 struct {
	c string
	d string
}) *StructerStructInvocation {
	invocation := new(StructerStructInvocation)

	invocation.Parameters.Ident1 = ident1

	invocation.Results.Ident2 = ident2

	return invocation
}

// StructerNamedStructInvocation represents a single call of FakeStructer.NamedStruct
type StructerNamedStructInvocation struct {
	Parameters struct {
		A struct {
			a string
			b string
		}
	}
	Results struct {
		Z struct {
			c string
			d string
		}
	}
}

// NewStructerNamedStructInvocation creates a new instance of StructerNamedStructInvocation
func NewStructerNamedStructInvocation(a struct {
	a string
	b string
}, z struct {
	c string
	d string
}) *StructerNamedStructInvocation {
	invocation := new(StructerNamedStructInvocation)

	invocation.Parameters.A = a

	invocation.Results.Z = z

	return invocation
}

// StructerTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type StructerTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeStructer is a mock implementation of Structer for testing.
Use it in your tests as in this example:

	package example

	func TestWithStructer(t *testing.T) {
		f := &main.FakeStructer{
			StructHook: func(ident1 struct {
	a string
	b string
}) (ident2 struct {
	c string
	d string
}) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeStruct ...
		f.AssertStructCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeStruct.
*/
type FakeStructer struct {
	StructHook func(struct {
		a string
		b string
	}) struct {
		c string
		d string
	}
	NamedStructHook func(struct {
		a string
		b string
	}) struct {
		c string
		d string
	}

	StructCalls      []*StructerStructInvocation
	NamedStructCalls []*StructerNamedStructInvocation
}

// NewFakeStructerDefaultPanic returns an instance of FakeStructer with all hooks configured to panic
func NewFakeStructerDefaultPanic() *FakeStructer {
	return &FakeStructer{
		StructHook: func(struct {
			a string
			b string
		}) (ident2 struct {
			c string
			d string
		}) {
			panic("Unexpected call to Structer.Struct")
		},
		NamedStructHook: func(struct {
			a string
			b string
		}) (z struct {
			c string
			d string
		}) {
			panic("Unexpected call to Structer.NamedStruct")
		},
	}
}

// NewFakeStructerDefaultFatal returns an instance of FakeStructer with all hooks configured to call t.Fatal
func NewFakeStructerDefaultFatal(t_sym1 StructerTestingT) *FakeStructer {
	return &FakeStructer{
		StructHook: func(struct {
			a string
			b string
		}) (ident2 struct {
			c string
			d string
		}) {
			t_sym1.Fatal("Unexpected call to Structer.Struct")
			return
		},
		NamedStructHook: func(struct {
			a string
			b string
		}) (z struct {
			c string
			d string
		}) {
			t_sym1.Fatal("Unexpected call to Structer.NamedStruct")
			return
		},
	}
}

// NewFakeStructerDefaultError returns an instance of FakeStructer with all hooks configured to call t.Error
func NewFakeStructerDefaultError(t_sym2 StructerTestingT) *FakeStructer {
	return &FakeStructer{
		StructHook: func(struct {
			a string
			b string
		}) (ident2 struct {
			c string
			d string
		}) {
			t_sym2.Error("Unexpected call to Structer.Struct")
			return
		},
		NamedStructHook: func(struct {
			a string
			b string
		}) (z struct {
			c string
			d string
		}) {
			t_sym2.Error("Unexpected call to Structer.NamedStruct")
			return
		},
	}
}

func (f *FakeStructer) Reset() {
	f.StructCalls = []*StructerStructInvocation{}
	f.NamedStructCalls = []*StructerNamedStructInvocation{}
}

func (f_sym3 *FakeStructer) Struct(ident1 struct {
	a string
	b string
}) (ident2 struct {
	c string
	d string
}) {
	if f_sym3.StructHook == nil {
		panic("Structer.Struct() called but FakeStructer.StructHook is nil")
	}

	invocation_sym3 := new(StructerStructInvocation)
	f_sym3.StructCalls = append(f_sym3.StructCalls, invocation_sym3)

	invocation_sym3.Parameters.Ident1 = ident1

	ident2 = f_sym3.StructHook(ident1)

	invocation_sym3.Results.Ident2 = ident2

	return
}

// SetStructStub configures Structer.Struct to always return the given values
func (f_sym4 *FakeStructer) SetStructStub(ident2 struct {
	c string
	d string
}) {
	f_sym4.StructHook = func(struct {
		a string
		b string
	}) struct {
		c string
		d string
	} {
		return ident2
	}
}

// SetStructInvocation configures Structer.Struct to return the given results when called with the given parameters
// If no match is found for an invocation the result(s) of the fallback function are returned
func (f_sym5 *FakeStructer) SetStructInvocation(calls_sym5 []*StructerStructInvocation, fallback_sym5 func() struct {
	c string
	d string
}) {
	f_sym5.StructHook = func(ident1 struct {
		a string
		b string
	}) (ident2 struct {
		c string
		d string
	}) {
		for _, call_sym5 := range calls_sym5 {
			if reflect.DeepEqual(call_sym5.Parameters.Ident1, ident1) {
				ident2 = call_sym5.Results.Ident2

				return
			}
		}

		return fallback_sym5()
	}
}

// StructCalled returns true if FakeStructer.Struct was called
func (f *FakeStructer) StructCalled() bool {
	return len(f.StructCalls) != 0
}

// AssertStructCalled calls t.Error if FakeStructer.Struct was not called
func (f *FakeStructer) AssertStructCalled(t StructerTestingT) {
	t.Helper()
	if len(f.StructCalls) == 0 {
		t.Error("FakeStructer.Struct not called, expected at least one")
	}
}

// StructNotCalled returns true if FakeStructer.Struct was not called
func (f *FakeStructer) StructNotCalled() bool {
	return len(f.StructCalls) == 0
}

// AssertStructNotCalled calls t.Error if FakeStructer.Struct was called
func (f *FakeStructer) AssertStructNotCalled(t StructerTestingT) {
	t.Helper()
	if len(f.StructCalls) != 0 {
		t.Error("FakeStructer.Struct called, expected none")
	}
}

// StructCalledOnce returns true if FakeStructer.Struct was called exactly once
func (f *FakeStructer) StructCalledOnce() bool {
	return len(f.StructCalls) == 1
}

// AssertStructCalledOnce calls t.Error if FakeStructer.Struct was not called exactly once
func (f *FakeStructer) AssertStructCalledOnce(t StructerTestingT) {
	t.Helper()
	if len(f.StructCalls) != 1 {
		t.Errorf("FakeStructer.Struct called %d times, expected 1", len(f.StructCalls))
	}
}

// StructCalledN returns true if FakeStructer.Struct was called at least n times
func (f *FakeStructer) StructCalledN(n int) bool {
	return len(f.StructCalls) >= n
}

// AssertStructCalledN calls t.Error if FakeStructer.Struct was called less than n times
func (f *FakeStructer) AssertStructCalledN(t StructerTestingT, n int) {
	t.Helper()
	if len(f.StructCalls) < n {
		t.Errorf("FakeStructer.Struct called %d times, expected >= %d", len(f.StructCalls), n)
	}
}

// StructCalledWith returns true if FakeStructer.Struct was called with the given values
func (f_sym6 *FakeStructer) StructCalledWith(ident1 struct {
	a string
	b string
}) bool {
	for _, call_sym6 := range f_sym6.StructCalls {
		if reflect.DeepEqual(call_sym6.Parameters.Ident1, ident1) {
			return true
		}
	}

	return false
}

// AssertStructCalledWith calls t.Error if FakeStructer.Struct was not called with the given values
func (f_sym7 *FakeStructer) AssertStructCalledWith(t StructerTestingT, ident1 struct {
	a string
	b string
}) {
	t.Helper()
	var found_sym7 bool
	for _, call_sym7 := range f_sym7.StructCalls {
		if reflect.DeepEqual(call_sym7.Parameters.Ident1, ident1) {
			found_sym7 = true
			break
		}
	}

	if !found_sym7 {
		t.Error("FakeStructer.Struct not called with expected parameters")
	}
}

// StructCalledOnceWith returns true if FakeStructer.Struct was called exactly once with the given values
func (f_sym8 *FakeStructer) StructCalledOnceWith(ident1 struct {
	a string
	b string
}) bool {
	var count_sym8 int
	for _, call_sym8 := range f_sym8.StructCalls {
		if reflect.DeepEqual(call_sym8.Parameters.Ident1, ident1) {
			count_sym8++
		}
	}

	return count_sym8 == 1
}

// AssertStructCalledOnceWith calls t.Error if FakeStructer.Struct was not called exactly once with the given values
func (f_sym9 *FakeStructer) AssertStructCalledOnceWith(t StructerTestingT, ident1 struct {
	a string
	b string
}) {
	t.Helper()
	var count_sym9 int
	for _, call_sym9 := range f_sym9.StructCalls {
		if reflect.DeepEqual(call_sym9.Parameters.Ident1, ident1) {
			count_sym9++
		}
	}

	if count_sym9 != 1 {
		t.Errorf("FakeStructer.Struct called %d times with expected parameters, expected one", count_sym9)
	}
}

// StructResultsForCall returns the result values for the first call to FakeStructer.Struct with the given values
func (f_sym10 *FakeStructer) StructResultsForCall(ident1 struct {
	a string
	b string
}) (ident2 struct {
	c string
	d string
}, found_sym10 bool) {
	for _, call_sym10 := range f_sym10.StructCalls {
		if reflect.DeepEqual(call_sym10.Parameters.Ident1, ident1) {
			ident2 = call_sym10.Results.Ident2
			found_sym10 = true
			break
		}
	}

	return
}

func (f_sym11 *FakeStructer) NamedStruct(a struct {
	a string
	b string
}) (z struct {
	c string
	d string
}) {
	if f_sym11.NamedStructHook == nil {
		panic("Structer.NamedStruct() called but FakeStructer.NamedStructHook is nil")
	}

	invocation_sym11 := new(StructerNamedStructInvocation)
	f_sym11.NamedStructCalls = append(f_sym11.NamedStructCalls, invocation_sym11)

	invocation_sym11.Parameters.A = a

	z = f_sym11.NamedStructHook(a)

	invocation_sym11.Results.Z = z

	return
}

// SetNamedStructStub configures Structer.NamedStruct to always return the given values
func (f_sym12 *FakeStructer) SetNamedStructStub(z struct {
	c string
	d string
}) {
	f_sym12.NamedStructHook = func(struct {
		a string
		b string
	}) struct {
		c string
		d string
	} {
		return z
	}
}

// SetNamedStructInvocation configures Structer.NamedStruct to return the given results when called with the given parameters
// If no match is found for an invocation the result(s) of the fallback function are returned
func (f_sym13 *FakeStructer) SetNamedStructInvocation(calls_sym13 []*StructerNamedStructInvocation, fallback_sym13 func() struct {
	c string
	d string
}) {
	f_sym13.NamedStructHook = func(a struct {
		a string
		b string
	}) (z struct {
		c string
		d string
	}) {
		for _, call_sym13 := range calls_sym13 {
			if reflect.DeepEqual(call_sym13.Parameters.A, a) {
				z = call_sym13.Results.Z

				return
			}
		}

		return fallback_sym13()
	}
}

// NamedStructCalled returns true if FakeStructer.NamedStruct was called
func (f *FakeStructer) NamedStructCalled() bool {
	return len(f.NamedStructCalls) != 0
}

// AssertNamedStructCalled calls t.Error if FakeStructer.NamedStruct was not called
func (f *FakeStructer) AssertNamedStructCalled(t StructerTestingT) {
	t.Helper()
	if len(f.NamedStructCalls) == 0 {
		t.Error("FakeStructer.NamedStruct not called, expected at least one")
	}
}

// NamedStructNotCalled returns true if FakeStructer.NamedStruct was not called
func (f *FakeStructer) NamedStructNotCalled() bool {
	return len(f.NamedStructCalls) == 0
}

// AssertNamedStructNotCalled calls t.Error if FakeStructer.NamedStruct was called
func (f *FakeStructer) AssertNamedStructNotCalled(t StructerTestingT) {
	t.Helper()
	if len(f.NamedStructCalls) != 0 {
		t.Error("FakeStructer.NamedStruct called, expected none")
	}
}

// NamedStructCalledOnce returns true if FakeStructer.NamedStruct was called exactly once
func (f *FakeStructer) NamedStructCalledOnce() bool {
	return len(f.NamedStructCalls) == 1
}

// AssertNamedStructCalledOnce calls t.Error if FakeStructer.NamedStruct was not called exactly once
func (f *FakeStructer) AssertNamedStructCalledOnce(t StructerTestingT) {
	t.Helper()
	if len(f.NamedStructCalls) != 1 {
		t.Errorf("FakeStructer.NamedStruct called %d times, expected 1", len(f.NamedStructCalls))
	}
}

// NamedStructCalledN returns true if FakeStructer.NamedStruct was called at least n times
func (f *FakeStructer) NamedStructCalledN(n int) bool {
	return len(f.NamedStructCalls) >= n
}

// AssertNamedStructCalledN calls t.Error if FakeStructer.NamedStruct was called less than n times
func (f *FakeStructer) AssertNamedStructCalledN(t StructerTestingT, n int) {
	t.Helper()
	if len(f.NamedStructCalls) < n {
		t.Errorf("FakeStructer.NamedStruct called %d times, expected >= %d", len(f.NamedStructCalls), n)
	}
}

// NamedStructCalledWith returns true if FakeStructer.NamedStruct was called with the given values
func (f_sym14 *FakeStructer) NamedStructCalledWith(a struct {
	a string
	b string
}) bool {
	for _, call_sym14 := range f_sym14.NamedStructCalls {
		if reflect.DeepEqual(call_sym14.Parameters.A, a) {
			return true
		}
	}

	return false
}

// AssertNamedStructCalledWith calls t.Error if FakeStructer.NamedStruct was not called with the given values
func (f_sym15 *FakeStructer) AssertNamedStructCalledWith(t StructerTestingT, a struct {
	a string
	b string
}) {
	t.Helper()
	var found_sym15 bool
	for _, call_sym15 := range f_sym15.NamedStructCalls {
		if reflect.DeepEqual(call_sym15.Parameters.A, a) {
			found_sym15 = true
			break
		}
	}

	if !found_sym15 {
		t.Error("FakeStructer.NamedStruct not called with expected parameters")
	}
}

// NamedStructCalledOnceWith returns true if FakeStructer.NamedStruct was called exactly once with the given values
func (f_sym16 *FakeStructer) NamedStructCalledOnceWith(a struct {
	a string
	b string
}) bool {
	var count_sym16 int
	for _, call_sym16 := range f_sym16.NamedStructCalls {
		if reflect.DeepEqual(call_sym16.Parameters.A, a) {
			count_sym16++
		}
	}

	return count_sym16 == 1
}

// AssertNamedStructCalledOnceWith calls t.Error if FakeStructer.NamedStruct was not called exactly once with the given values
func (f_sym17 *FakeStructer) AssertNamedStructCalledOnceWith(t StructerTestingT, a struct {
	a string
	b string
}) {
	t.Helper()
	var count_sym17 int
	for _, call_sym17 := range f_sym17.NamedStructCalls {
		if reflect.DeepEqual(call_sym17.Parameters.A, a) {
			count_sym17++
		}
	}

	if count_sym17 != 1 {
		t.Errorf("FakeStructer.NamedStruct called %d times with expected parameters, expected one", count_sym17)
	}
}

// NamedStructResultsForCall returns the result values for the first call to FakeStructer.NamedStruct with the given values
func (f_sym18 *FakeStructer) NamedStructResultsForCall(a struct {
	a string
	b string
}) (z struct {
	c string
	d string
}, found_sym18 bool) {
	for _, call_sym18 := range f_sym18.NamedStructCalls {
		if reflect.DeepEqual(call_sym18.Parameters.A, a) {
			z = call_sym18.Results.Z
			found_sym18 = true
			break
		}
	}

	return
}
