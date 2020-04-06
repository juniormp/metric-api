package usecases

type MetricRequest struct {
	Value float64 `json:"value" binding:"required"`
	Name  string
}
