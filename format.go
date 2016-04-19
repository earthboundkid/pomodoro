package main

import (
	"fmt"
	"math"
	"time"
)

func formatDays(timeLeft time.Duration) string {
	days := int(timeLeft.Hours() / 24)
	hours := int(timeLeft.Hours()) % 24
	minutes := int(timeLeft.Minutes()) % 60
	seconds := int(timeLeft.Seconds()) % 60
	return fmt.Sprintf("%d:%02d:%02d:%02d",
		days, hours, minutes, seconds)
}

func formatHours(timeLeft time.Duration) string {
	hours := int(timeLeft.Hours())
	minutes := int(timeLeft.Minutes()) % 60
	seconds := int(timeLeft.Seconds()) % 60
	return fmt.Sprintf("%d:%02d:%02d",
		hours, minutes, seconds)
}

func formatMinutes(timeLeft time.Duration) string {
	minutes := int(timeLeft.Minutes())
	seconds := int(timeLeft.Seconds()) % 60
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}

func formatSeconds(timeLeft time.Duration) string {
	return fmt.Sprintf("%02.1f", math.Abs(timeLeft.Seconds()))
}
