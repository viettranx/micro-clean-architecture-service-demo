package entity

func checkTitle(s string) error {
	if s == "" {
		return ErrTitleIsBlank
	}

	return nil
}

func checkStatus(s Status) error {
	if s != StatusDoing && s != StatusDone && s != StatusDeleted {
		return ErrStatusInvalid
	}

	return nil
}
