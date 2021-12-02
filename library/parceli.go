package parcelilib

import "fmt"

func Parceli(service string, tracking int, verbose bool) {

	if verbose {
		fmt.Printf("Service: %s\tTracking #: %d\n", service, tracking)
	}

	switch service {
	case "ups":
		// UPS specific logic
		break
	case "usps":
		// USPS specific logic
		break
	case "fedex":
		// FEDEX specific logic
		break
	default:
		// Impossible?
		break
	}
}
