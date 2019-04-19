package action

import (
	"reflect"
	"testing"
)

func TestDirWalkAtExistDir(t *testing.T) {
	dir := dirwalk("../test")
	expect := []string{
		"../test/a/a1",
		"../test/a/a2",
		"../test/b/b1",

	}
	if !reflect.DeepEqual(dir, expect) {
		t.Fatal("failed test")
	}
}

func TestDirWalkAtNotExistDir(t *testing.T) {
	dir := dirwalk("")
	expect := []string{}
	if !reflect.DeepEqual(dir, expect) {
		t.Fatal("failed test")
	}
}