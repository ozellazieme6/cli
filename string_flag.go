package cli

import (
	"flag"
	"os"
)

// Apply populates the flagSet and the destination pointer
func (f *StringFlag) Apply(set *flag.FlagSet) error {
	val := f.Value
	for _, envVar := range f.EnvVars {
		if envVal, ok := os.LookupEnv(envVar); ok {
			val = envVal
			break
		}
	}

	if f.Destination != nil {
		*f.Destination = val
	}

	set.String(f.Name, val, f.Usage)
	return nil
}