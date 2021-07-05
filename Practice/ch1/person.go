package main

import "time"

func main() {
	type Employee struct {
		ID        int
		Name      string
		Address   string
		DoB       time.Time
		Position  string
		Salary    int
		ManagerID int
	}

	var jwang Employee

	jwang.Address = "Ah"
	jwang.Salary = 12000
	jwang.ID = 1

	pos := &jwang.Position

	*pos = "beijin" + *pos

	jack := &jwang

	jack.ID = 5

	println(jwang.Address)
	println(jwang.Position)
	println(jwang.ID)
	println(jack.ID)

}
