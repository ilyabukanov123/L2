package main

import (
	"fmt"
	"log"

	"github.com/ilyabukanov123/L2/developer/dev1"
)

func main() {
	time, err := dev1.NtpTime("0.ru.pool.ntp.org")
	if err != nil {
		log.Fatalf("Произошла ошибка по получения времени с Ntp: %s\n", err)
	}
	fmt.Println(time)
}
