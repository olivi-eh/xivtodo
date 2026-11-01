package models

import (
	"reflect"
	"testing"
)

// Firestore's encoder panics on nil embedded struct pointers, which godestone produces for unresolved entities
func TestNoEmbeddedPointers(t *testing.T) {
	pkgPath := reflect.TypeOf(User{}).PkgPath()
	seen := map[reflect.Type]bool{}
	var walk func(reflect.Type, string)
	walk = func(typ reflect.Type, path string) {
		for typ.Kind() == reflect.Pointer || typ.Kind() == reflect.Slice || typ.Kind() == reflect.Map {
			typ = typ.Elem()
		}
		if typ.Kind() != reflect.Struct || seen[typ] || typ.PkgPath() != pkgPath {
			return
		}
		seen[typ] = true
		for i := 0; i < typ.NumField(); i++ {
			f := typ.Field(i)
			if f.Anonymous && f.Type.Kind() == reflect.Pointer {
				t.Errorf("%s.%s embeds a pointer", path, f.Name)
			}
			walk(f.Type, path+"."+f.Name)
		}
	}
	walk(reflect.TypeOf(User{}), "User")
}
