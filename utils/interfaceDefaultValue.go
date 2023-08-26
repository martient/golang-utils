package utils

func IntOrDefault(value interface{}, defaultValue int) int {
	if v, ok := value.(float64); ok {
		return int(v)
	}
	return defaultValue
}

func StringOrDefault(value interface{}, defaultValue string) string {
	if v, ok := value.(string); ok {
		return v
	}
	return defaultValue
}

func BoolOrDefault(value interface{}, defaultValue bool) bool {
	if v, ok := value.(bool); ok {
		return v
	}
	return defaultValue
}

func Float64OrDefault(value interface{}, defaultValue float64) float64 {
	if v, ok := value.(float64); ok {
		return v
	}
	return defaultValue
}
