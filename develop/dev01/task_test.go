package main

import (
	"testing"
)

// Тест на пустое значение для текущего времени
func TestNtpTimeEmpty(t *testing.T) {
	time, err := NtpTime("")
	if err != nil {
		// fatalf выводит сообщение и выбрасывает панику
		return
	}
	t.Fatalf("Должна была произойти ошибка, но вернулось время: %v", time)
}

func TestNtpTimeNotEmpty(t *testing.T) {
	time, err := NtpTime("0.ru.pool.ntp.org")
	if err != nil {
		t.Fatalf("Должно было вернутся время, но вернулась ошибка: %v", err)
	}
	t.Logf("Полученное время: %v", time)
}
