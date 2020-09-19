package context

func (c *ctx) String(name string) string {

	return lookupString(name, c.flagSet)
}

func lookupString(name string, fs flagSet) string {

	value, ok := fs[name]

	if !ok {
		return ""
	}

	switch val := value.(type) {

	case string:

		return val

	case *string:

		return *val

	default:
		return "false"
	}

}
