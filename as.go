package as

func Dummy() {
}

func Map(ss []string, f func(string) (string, error)) ([]string, error) {
	result := make([]string, len(ss))
	for i, s := range ss {
		r, err := f(s)
		if err != nil {
			return nil, err
		}
		result[i] = r
	}
	return result, nil
}

func MustMap(ss []string, f func(string) string) []string {
	result, _ := Map(ss, func(s string) (string, error) { return f(s), nil })
	return result
}
