package test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/antoniofmoliveira/tri/cmd"
)

func TestAdd(t *testing.T) {
	cmd := cmd.RootCmd
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"add", "test", "--datafile", "./.tridos.json"})
	cmd.Execute()
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}

func TestList(t *testing.T) {

	cmd := cmd.RootCmd
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list", "--datafile", "./.tridos.json"})
	cmd.Execute()
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != `[{"Text":"test","Priority":2,"Done":false,"DueDate":"","Created":""}]` {
		t.Fatalf("expected \"%s\" got \"%s\"", "testisawesome", string(out))
	}
}
