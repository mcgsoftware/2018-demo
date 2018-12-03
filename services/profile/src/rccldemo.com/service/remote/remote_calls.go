package remote

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)




// http://35.245.49.124/royal/api/bookings/bjm100
const K8_SHIP_SERVICE_NAME = "booking"
const K8_SHIP_SERVICE_PORT = "8070"
var RemoteServiceUri string = "http://" + K8_SHIP_SERVICE_NAME + ":" + K8_SHIP_SERVICE_PORT + "/royal/api/bookings/"

// var LocalServiceUri string = "http://127.0.0.1:8072/royal/api/ships/AL"


//
// Fetches data from remote
//
func CallRemoteBookingService( vdsId string) ([]byte, error) {


	// Log service metric for remote service call
	//defer helpers.LogServiceMetric(time.Now(), vdsId, traceId , service, operation, method, status )

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// uri := LocalServiceUri
	uri := RemoteServiceUri + vdsId
	req, err := http.NewRequest("GET", uri, nil)
	if (err != nil) {
		return []byte("Problem!"), err
	}

	// set headers required by remote service
	//req.Header.Set(helpers.TRACE_ID_HEADER, traceId)
	//req.Header.Set(helpers.VDSID_HEADER, vdsId )


	resp, err := client.Do(req)
	if err != nil {
		return []byte("Problem!"), fmt.Errorf("Remote service call to Content Service failed: %s. Error: ", uri, err)
	}
	defer resp.Body.Close()

	// Parse data from remote service
	if (resp.StatusCode != http.StatusOK) {
		return []byte("Problem!"), fmt.Errorf("Failure from remote service. StatusCode=%v, body=%v", resp.StatusCode, "none")
	}

	// read response as html
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("Problem!"), fmt.Errorf("Remote service call to Content Service failed: %s. Error: ", uri, err)
	}

	// debug out
	//fmt.Println(os.Stdout, string(responseData))


	return responseData, nil

}
