package main

func main() {
	purchase := 3333_33
	percent := 1
	limit := 100

	bonus := (purchase / 100) / 100 * percent // сначала приводим копейки к рублям
	if bonus > limit {
		bonus = limit
	}
	println(bonus) // должно быть 33* (см. объяснение ниже)
}
