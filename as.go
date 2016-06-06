package as

import "regexp"

// eias - everything is a string

// operations on slice of strings

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

func IndexOf(ss []string, elem string) int {
	for i, s := range ss {
		if s == elem {
			return i
		}
	}
	return -1
}

func IndexOfPattern(ss []string, pattern string) int {
	re := regexp.MustCompile(pattern)
	for i, s := range ss {
		if re.MatchString(s) {
			return i
		}
	}
	return -1
}

func IndicesOf(ss []string, elem string) []int {
	var result = make([]int, 0)
	for i, s := range ss {
		if s == elem {
			result = append(result, i)
		}
	}
	return result
}

func IndicesOfPattern(ss []string, pattern string) []int {
	re := regexp.MustCompile(pattern)
	var result = make([]int, 0)
	for i, s := range ss {
		if re.MatchString(s) {
			result = append(result, i)
		}
	}
	return result
}

// split slice of strings into multiple slices of strings
// splitpoints provide the line numbers at which to split
// return (len(splitpoints) + 1) slices/sections of slices of strings
// if splitpoints is empty return the orginal ss as a single element of external slice
func SplitAt(ss []string, splitpoints []int) [][]string {
	result := make([][]string, len(splitpoints)+1)
	prevP := 0
	for i, p := range splitpoints {
		result[i] = ss[prevP:p]
		prevP = p
	}
	result[len(splitpoints)] = ss[prevP:]
	return result
}

//////////////////////////////////////////////////////////////////////////////
// This is actually for imath package
//////////////////////////////////////////////////////////////////////////////

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
