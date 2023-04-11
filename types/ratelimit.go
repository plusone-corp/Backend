package types

type RateLimit struct {
	Route            string
	RequestPerHour   int64
	IncludeSubRoutes bool
}