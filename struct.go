package utils

import (
	"github.com/fatih/structs"
)

//SetStructValue funct
func SetStructValue(s interface{}, m map[string]interface{}) {
	v := structs.New(s)

	for _, f := range v.Fields() {

		tag := f.Tag("json")

		if val, ok := m[tag]; ok {
			f.Set(val)
		}
	}
}
