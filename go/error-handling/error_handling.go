// Package erratum is a solution to lerning #17 (exercism.io)
package erratum

// Use opens a resource, calls Frob(input) on the result resource and then closes that resource (in all cases)
func Use(o ResourceOpener, input string) (err error) {
	res, err := o()
	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return
	}
	defer res.Close()

	defer func() error {
		if r := recover(); r != nil {
			switch rec := r.(type) {
			case FrobError:
				res.Defrob(rec.defrobTag)
				err = rec.inner
				return rec
			case error:
				err = rec
			}
		}
		return err
	}()
	res.Frob(input)

	return
}
