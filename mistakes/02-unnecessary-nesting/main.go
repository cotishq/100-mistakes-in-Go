package main

import "errors"


func joinNested(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

func joinClean(s1,s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}
	if s2 == "" {
		return "", errors.New("s2 is empty")
	}

	concat, err := concatenate(s1,s2)
	if err != nil {
		return "" , err
	}

	if len(concat) > max {
		return concat[:max], nil
	}

	return concat, nil
}

func concatenate(s1, s2 string) (string, error) {
	// In real world this might do something more complex
	// Here it's simple concatenation with a space
	return s1 + " " + s2, nil
}

