package cache

var (
	// AuthToken - Cache table for auth token
	AuthToken = "authToken"
)

func UserDefinedCacheTable() []string {
	return []string{
		AuthToken,
	}
}
