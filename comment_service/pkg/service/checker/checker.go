// Package checker package for checking elements
package checker

import "comment_service/pkg/factory"

func IsConstantInArr(constant string, arr []string) bool {
	for _, el := range arr {
		if el == constant {
			return true
		}
	}
	return false
}

func IsEntityExists(entityType string, entityId string) bool {
	state := factory.GetState(entityType)
	if state == nil {
		return false
	}
	if err := state.Get(entityId); err == nil {
		return true
	}
	return false
}
