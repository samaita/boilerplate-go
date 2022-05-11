package dto

type HealthCheckResponse struct {
	DBHealth     string `json:"db_health"`
	DBError      string `json:"db_error,omitempty"`
	DBLatency    string `json:"db_latency"`
	CacheHealth  string `json:"cache_health"`
	CacheError   string `json:"cache_error,omitempty"`
	CacheLatency string `json:"cache_latency"`
}
