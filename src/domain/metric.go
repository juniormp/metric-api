package domain

type Metric struct {
	Name      string  `json:"name"`
	Value     float64 `json:"score"`
	ExpiredAt string  `json:"expiredAt"`
}
