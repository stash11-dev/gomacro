// this file was generated by gomacro command: import "crypto/rc4"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "crypto/rc4"
	. "reflect"
)

func init() {
	Binds["crypto/rc4"] = map[string]Value{
		"NewCipher":	ValueOf(pkg.NewCipher),
	}
	Types["crypto/rc4"] = map[string]Type{
		"Cipher":	TypeOf((*pkg.Cipher)(nil)).Elem(),
		"KeySizeError":	TypeOf((*pkg.KeySizeError)(nil)).Elem(),
	}
}