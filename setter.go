package defaults

// Setter is an interface for setting default values
type Setter interface {
	SetDefaults()
}

func callSetter(v interface{}) {
	if ds, ok := v.(Setter); ok {
		ds.SetDefaults()
	}
}

// SetterWithError is another interface for setting default values
type SetterWithError interface {
	SetDefaults() error
}

func callSetterWithError(v interface{}) (bool, error) {
	if ds, ok := v.(SetterWithError); ok {
		err := ds.SetDefaults()
		return true, err
	}
	return false, nil
}
