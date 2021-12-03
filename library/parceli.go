package parcelilib

import (
	"bytes"
	"fmt"
)

var (
	usps_tracking = []byte("94")
	usps_priority_mail = []byte("9205")
	usps_certified_mail = []byte("9407")
	usps_express_global = []byte("82")
	usps_hold_del_num = []byte("9303")
	usps_express_international = []byte("EC")
	usps_express_number = []byte("9270")
	usps_priority_international = []byte("CP")
	usps_registered_mail = []byte("9208")
	usps_signature = []byte("9202")

	ups = []byte("1Z")
)

const (
	UPS = 1 << 0
	USPS = 1 << 1
	FEDEX = 1 << 2
)

func Parceli(tracking string, verbose bool) {

	// Determine service here
	// Fedex
	//

	// Slice off first 4 bytes

	var service int


	trackingSlice := []byte(tracking)

	id := trackingSlice[:4]
	if bytes.Equal(id, usps_priority_mail[:]) || bytes.Equal(id, usps_certified_mail[:]) || bytes.Equal(id, usps_signature[:]) || bytes.Equal(id, usps_signature[:]) || bytes.Equal(id, usps_hold_del_num[:]) || bytes.Equal(id, usps_express_number[:]) || bytes.Equal(id, usps_registered_mail[:]){
		service = USPS
	} else {
		// Two number IDS
		id = trackingSlice[:2]
		// ups
		if bytes.Equal(id, ups[:]) {
			service = UPS
		} else if bytes.Equal(id, usps_express_global[:]) || bytes.Equal(id, usps_express_international[:]) || bytes.Equal(id, usps_tracking[:]) || bytes.Equal(id, usps_priority_international[:]) {
			service = USPS
		} else {
			service = FEDEX
		}
	}

	switch service {
	case UPS :
		fmt.Println("UPS")
		break
	case USPS :
		fmt.Println("USPS")
		break
	case FEDEX :
		fmt.Println("FEDEX")
		break
	default :
		fmt.Println("e")
		break
	}


	// if verbose {
	// 	fmt.Printf("Service: %s\tTracking #: %d\n", service, tracking)
	// }
}
