package main

import "testing"

func Test_sayHello(t *testing.T) {
    name := "Cyril"
    want := "Hello Cyril"

    if got := sayHello(name); got != want {
        t.Errorf("hello() = %q, want %q", got, want)
    }
}