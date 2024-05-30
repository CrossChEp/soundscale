package auxiliary

func DeleteUserFromArray(userId string, array []string) []string {
	var deletedUserArray []string
	for _, user := range array {
		if user == userId {
			continue
		}
		deletedUserArray = append(deletedUserArray, user)
	}
	return deletedUserArray
}
