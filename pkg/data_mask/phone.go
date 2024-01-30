package data_mask

func PhoneMasking(phone string) string {
	if len(phone) != 11 {
		return ""
	}
	return phone[:3] + "****" + phone[7:]
}
