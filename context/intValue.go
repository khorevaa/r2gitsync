package context

func (c *ctx) Int(name string) int {

	return lookupInt(name, c.flagSet)
}

func lookupInt(name string, fs flagSet) int {

	value, ok := fs[name]

	if !ok {
		return 0
	}

	switch val := value.(type) {

	case int:

		return val

	case *int:

		return *val

	default:
		return 0
	}

}
