package utils

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"time"
)

// JSONResponse is a utility function to send JSON responses
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("Error encoding JSON response: %v", err)
		}
	}
}

// FormatTime formats a time.Time into a standard string format
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// RoundToDecimal rounds a float64 to the specified decimal places
func RoundToDecimal(val float64, places int) float64 {
	if places < 0 {
		places = 0
	}

	factor := math.Pow(10, float64(places))
	return math.Round(val*factor) / factor
}

// IsValidTemperature checks if a temperature value is within a reasonable range
func IsValidTemperature(temp float64) bool {
	// Assuming temperature in Celsius
	// Lowest recorded on Earth: -89.2°C, Highest: 56.7°C
	// Adding some margin to account for extremes
	return temp >= -100 && temp <= 70
}

// IsValidHumidity checks if a humidity value is valid (between 0-100%)
func IsValidHumidity(humidity float64) bool {
	return humidity >= 0 && humidity <= 100
}

// TimeAgo returns a human-readable string for how long ago a time was
func TimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	switch {
	case diff < time.Minute:
		return "just now"
	case diff < time.Hour:
		return formatDuration(diff.Minutes(), "minute")
	case diff < 24*time.Hour:
		return formatDuration(diff.Hours(), "hour")
	case diff < 30*24*time.Hour:
		return formatDuration(diff.Hours()/24, "day")
	case diff < 365*24*time.Hour:
		return formatDuration(diff.Hours()/(24*30), "month")
	default:
		return formatDuration(diff.Hours()/(24*365), "year")
	}
}

// formatDuration helper function for TimeAgo
func formatDuration(value float64, unit string) string {
	value = math.Floor(value)
	if value == 1 {
		return "1 " + unit + " ago"
	}
	return RoundToString(value) + " " + unit + "s ago"
}

// RoundToString rounds a float to an integer string
func RoundToString(value float64) string {
	return string(rune(int(math.Round(value)) + '0'))
}
