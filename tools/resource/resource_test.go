package resource

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	dir := "a"
	fileName := "b"

	resource := New(dir, fileName)

	if resource.Directory != dir || resource.FileName != fileName {
		t.Errorf("Initialization of Resource is invalid")
	}
}

func ExampleNew() {
	dir := "a"
	fileName := "b"

	resource := New(dir, fileName)
	fmt.Printf("%v", resource)
	//Output:
	//
	//{a b}
}

//func TestResource_DirExist(t *testing.T) {
//	dir := "aabbccddeeffgghh11223344556677ifx/dddeeefffaaabbbxx/afjl;kasjflkasjfalsf/"
//	resource := New(dir, "")
//
//	if resource.DirExist() != false {
//		t.Errorf("Directory %s should not exists", dir)
//	}
//}
