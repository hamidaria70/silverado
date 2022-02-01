package greetings

import (
	"greetings/greetings"
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "jack"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg , err := greetings.Hello("jack")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("jack") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T){
	msg , err := greetings.Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}

}