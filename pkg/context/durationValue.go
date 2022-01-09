package context

import "time"

// Duration looks up the value of a local DurationFlag, returns
// 0 if not found
func (c *ctx) Duration(name string) time.Duration {

	return lookupDuration(name, c.flagSet)
}

func lookupDuration(name string, fs flagSet) time.Duration {

	value, ok := fs[name]

	if !ok {
		return 0
	}

	switch val := value.(type) {

	case time.Duration:

		return val

	case *time.Duration:

		return *val

	default:
		return 0
	}

}
