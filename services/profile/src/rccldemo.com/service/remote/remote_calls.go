package remote

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)





const K8_SHIP_SERVICE_NAME = "ships"
const K8_SHIP_SERVICE_PORT = "8070"
var RemoteServiceUri string = "http://" + K8_SHIP_SERVICE_NAME + ":" + K8_SHIP_SERVICE_PORT + "/royal/api/ships/AL"

// var LocalServiceUri string = "http://127.0.0.1:8072/royal/api/ships/AL"


//
// Fetches data from remote
//
func CallRemoteShips() (string, error) {

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// uri := LocalServiceUri
	uri := RemoteServiceUri
	req, err := http.NewRequest("GET", uri, nil)
	if (err != nil) {
		return "Problem!", err
	}

	// set headers required by remote service
	//req.Header.Set(helpers.TRACE_ID_HEADER, traceId)
	//req.Header.Set(helpers.VDSID_HEADER, vdsId )


	resp, err := client.Do(req)
	if err != nil {
		return"Problem!", fmt.Errorf("Remote service call to Content Service failed: %s. Error: ", uri, err)
	}
	defer resp.Body.Close()

	// Parse data from remote service
	if (resp.StatusCode != http.StatusOK) {
		//var result map[string]interface{}
		//json.NewDecoder(resp.Body).Decode(&result)
		return "Problem!", fmt.Errorf("Failure from remote service. StatusCode=%v, body=%v", resp.StatusCode, "none")
	}

	// read response as html
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return"Problem!", fmt.Errorf("Remote service call to Content Service failed: %s. Error: ", uri, err)
	}

	// print out
	fmt.Println(os.Stdout, string(htmlData))

	//var result models.ShipInfo
	//json.NewDecoder(resp.Body).Decode(&result)

	//log.Println("Found: ", result)

	return string(htmlData), nil

}
