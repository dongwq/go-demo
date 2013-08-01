package util

import (
	"testing"

	"errors"
	//"github.com/dongwq/godemo/util"
	//"fmt"
	//"github.com/sdegutis/go.assert"
)

func TestCheckErr(t *testing.T) {

	err := errors.New("emit macho dwarf: elf header corrupted")

	//var a error = nil
	CheckErr("msg", err)

	//assert.Equals(t, vr.IsValid, false)
	//assert.Equals(t, vr.CleanValue.(string), "")

}
