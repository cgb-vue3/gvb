package data_mask

import "strings"

func EmailMasking(email string) string {
	emailStr := strings.Split(email, "@")
	if len(emailStr) != 2 {
		return ""
	}
	return emailStr[0][:3] + "****@" + emailStr[1]
}
