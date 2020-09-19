package context

func (c *ctx) Bool(name string) bool {

	return lookupBool(name, c.flagSet)
}

func lookupBool(name string, fs flagSet) bool {

	value, ok := fs[name]

	if !ok {
		return false
	}

	switch val := value.(type) {

	case bool:

		return val

	case *bool:

		return *val

	default:
		return false
	}

}
