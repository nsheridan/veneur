package discovery

// Discoverer is an interface for various
// service discovery mechanisms.
type Discoverer interface {
	GetDestinationsForService(string) ([]string, error)
}
