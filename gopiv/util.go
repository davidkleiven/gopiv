package gopiv

import "strings"

func stringSliceEqual(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// return all items in candidates that is present at least once in expression
func identifyParameters(candidates []string, expression string) []string {
	res := make(map[string]interface{})
	for _,  str := range candidates {
		if strings.Contains(expression, str) {
			res[str] = nil
		}
	}

	resultSet := []string{}
	for k := range res {
		resultSet = append(resultSet, k)
	}
	return resultSet
}