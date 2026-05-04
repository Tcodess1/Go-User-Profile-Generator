package main

func comparePhone(phone float64) string {
	const standardPhone = 6.43

	if phone > standardPhone {
		return "Phone is considered premium level"
	} else if phone < standardPhone {
		return "Phone is below premium standard"
	}
	return "Phone is standard"
}