package checker

func IsUserInArray(userId string, likesArray []string) bool {
	for _, user := range likesArray {
		if user == userId {
			return true
		}
	}
	return false
}
