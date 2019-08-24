package main

import (
	"testing"
	"fmt"
)

func TestPrint(t *testing.T){
	res:=Printto();
	fmt.Println("hey");
	if res != 210{
		t.Errorf("wrong");
	}else{
		fmt.Println("right");
	}
}