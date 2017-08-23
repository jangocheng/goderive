// Code generated by goderive DO NOT EDIT.

package gostring

import (
	"bytes"
	"fmt"
)

func deriveGoString(this *MyStruct) string {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "func() *gostring.MyStruct {\n")
	if this == nil {
		fmt.Fprintf(buf, "return nil\n")
	} else {
		fmt.Fprintf(buf, "this := &gostring.MyStruct{}\n")
		fmt.Fprintf(buf, "this.Int64 = %#v\n", this.Int64)
		if this.StringPtr != nil {
			fmt.Fprintf(buf, "this.StringPtr = %s\n", deriveGoString_(this.StringPtr))
		}
		fmt.Fprintf(buf, "return this\n")
	}
	fmt.Fprintf(buf, "}()\n")
	return buf.String()
}

func deriveGoString_(this *string) string {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "func() *string {\n")
	if this == nil {
		fmt.Fprintf(buf, "return nil\n")
	} else {
		fmt.Fprintf(buf, "this := new(string)\n")
		fmt.Fprintf(buf, "*this = %#v\n", *this)
		fmt.Fprintf(buf, "return this\n")
	}
	fmt.Fprintf(buf, "}()\n")
	return buf.String()
}
