package erratum

var testVersion = 2

func Use(ro ResourceOpener, input string) error {
	res, err := open(ro)
	if err != nil {
		return err
	}
	defer res.Close()

	if err = frob(res, input); err != nil {
		return err
	}
	return nil
}

func open(ro ResourceOpener) (Resource, error) {
	for {
		res, err := ro()
		switch err.(type) {
		case TransientError:
			continue
		default:
			return res, err
		}
	}
}

func frob(res Resource, input string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				err = e
				res.Defrob(e.defrobTag)
			case error:
				err = e
			}
		}
	}()
	res.Frob(input)
	return
}
