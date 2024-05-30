package auxilary

func RemoveUserIfExists(userId string, users []string) []string {
	var newUsers []string
	for _, user := range users {
		if user != userId {
			newUsers = append(newUsers, userId)
		}
	}
	return newUsers
}
