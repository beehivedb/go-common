//Package test utils for test
//Copyright (C) 2020 to All Authors. all rights reserved.
//Author Ron
//Date 2020/09/19 22:18:39
//Version 1.0
package test

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/pmezard/go-difflib/difflib"
)

//TB interface to testing.{T,B}.
type TB interface {
	Helper()
	Fatalf(string, ...interface{})
}

// Assert fails the test if the condition is false.
func Assert(tb TB, condition bool, format string, a ...interface{}) {
	tb.Helper()
	if !condition {
		tb.Fatalf("\033[31m"+format+"\033[39m\n", a...)
	}
}

// Ok fails the test if an err is not nil.
func Ok(tb TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("\033[31munexpected error: %v\033[39m\n", err)
	}
}

// NotOk fails the test if an err is nil.
func NotOk(tb TB, err error, a ...interface{}) {
	tb.Helper()
	if err == nil {
		if len(a) != 0 {
			format := a[0].(string)
			tb.Fatalf("\033[31m"+format+": expected error, got none\033[39m", a[1:]...)
		}
		tb.Fatalf("\033[31mexpected error, got none\033[39m")
	}
}

// Equals fails the test if exp is not equal to act.
func Equals(tb TB, exp, act interface{}, msgAndArgs ...interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("\033[31m%s\n\nexp: %#v\n\ngot: %#v%s\033[39m\n", formatMessage(msgAndArgs), exp, act, diff(exp, act))
	}
}

func typeAndKind(v interface{}) (reflect.Type, reflect.Kind) {
	t := reflect.TypeOf(v)
	k := t.Kind()

	if k == reflect.Ptr {
		t = t.Elem()
		k = t.Kind()
	}
	return t, k
}

// diff returns a diff of both values as long as both are of the same type and
// are a struct, map, slice, array or string. Otherwise it returns an empty string.
func diff(expected interface{}, actual interface{}) string {
	if expected == nil || actual == nil {
		return ""
	}

	et, ek := typeAndKind(expected)
	at, _ := typeAndKind(actual)
	if et != at {
		return ""
	}

	if ek != reflect.Struct && ek != reflect.Map && ek != reflect.Slice && ek != reflect.Array && ek != reflect.String {
		return ""
	}

	var e, a string
	c := spew.ConfigState{
		Indent:                  " ",
		DisablePointerAddresses: true,
		DisableCapacities:       true,
		SortKeys:                true,
	}
	if et != reflect.TypeOf("") {
		e = c.Sdump(expected)
		a = c.Sdump(actual)
	} else {
		e = reflect.ValueOf(expected).String()
		a = reflect.ValueOf(actual).String()
	}

	diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(e),
		B:        difflib.SplitLines(a),
		FromFile: "Expected",
		FromDate: "",
		ToFile:   "Actual",
		ToDate:   "",
		Context:  1,
	})
	return "\n\nDiff:\n" + diff
}

// ErrorEqual compares Go errors for equality.
func ErrorEqual(tb TB, left, right error, msgAndArgs ...interface{}) {
	tb.Helper()
	if left == right {
		return
	}

	if left != nil && right != nil {
		Equals(tb, left.Error(), right.Error(), msgAndArgs...)
		return
	}

	tb.Fatalf("\033[31m%s\n\nexp: %#v\n\ngot: %#v\033[39m\n", formatMessage(msgAndArgs), left, right)
}

func formatMessage(msgAndArgs []interface{}) string {
	if len(msgAndArgs) == 0 {
		return ""
	}

	if msg, ok := msgAndArgs[0].(string); ok {
		return fmt.Sprintf("\n\nmsg: "+msg, msgAndArgs[1:]...)
	}
	return ""
}
