package collect

import "io"

func Check(err *error) {
	if *err != nil && *err != io.EOF {
		panic(*err)
	}
}
