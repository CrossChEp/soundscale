package auxilary

func IsElementInArr(element string, arr []string) bool {
	for _, el := range arr {
		if el == element {
			return true
		}
	}
	return false
}