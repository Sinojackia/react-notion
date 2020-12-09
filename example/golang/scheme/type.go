package scheme

import (
	"reflect"
	"src/model"
)

type Scheme struct {
	kindToType map[model.BaseValueKind]reflect.Type
}

func NewScheme() *Scheme {
	s := &Scheme{
		kindToType: map[model.BaseValueKind]reflect.Type{},
	}
	return s
}

func (s *Scheme) AddKnownTypeWithName(name string, obj interface{}) {
}

// AllKnownTypes returns the all known types.
func (s *Scheme) AllKnownTypes() map[model.BaseValueKind]reflect.Type {
	return s.kindToType
}
