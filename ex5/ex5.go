// Given a lichess username retrieve its profile through the REST API Get
// service api/user/{username}
//
// See https://lichess.org/api#operation/playerTopNbPerfType
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
)

// globals
// ----------------------------------------------------------------------------
const EXIT_SUCCESS = 0
const EXIT_FAILURE = 1

const version = "0.1"

var user string
var want_version bool

// functions
// ----------------------------------------------------------------------------

// init module
//
// setup the flag environment for the on-line help
func init() {

	// first, create a command-line argument for parsing the username
	flag.StringVar(&user, "user", "", "username")

	// also, create an additional flag for showing the version
	flag.BoolVar(&want_version, "version", false, "shows version info and exits")
}

// showVersion
//
// show the current version of this program and exits with the given signal
func showVersion(signal int) {

	fmt.Printf(" %v %v\n", os.Args[0], version)
	os.Exit(signal)
}

// JSONtoMap
//
// converts the data in a map from strings to arbitrary types returned by
// unmarshalling JSON data into a map of strings to arbitrary types preseving
// the nested structures
func JSONtoMap(dataJSON map[string]interface{}) map[string]interface{} {

	// --initialization
	result := make(map[string]interface{})

	// process the input dictionary
	for key, value := range dataJSON {

		if section, ok := value.(map[string]interface{}); ok {

			// if this value can be effectively casted into a nested dictionary
			// then proceed recursively
			result[key] = JSONtoMap(section)
		} else {

			// otherwise, assign this value to this key
			result[key] = value
		}
	}

	// and return the result
	return result
}

// main function
//
// given a number decide whether it is divisible by 7 or not
func main() {

	// first things first, parse the flags
	flag.Parse()

	// if the current version is requested, then show it on the standard output
	// and exit
	if want_version {
		showVersion(EXIT_SUCCESS)
	}

	// create a query to retrieve user information from Lichess
	response, err := http.Get("https://lichess.org/api/user/" + user)
	if err != nil {
		log.Fatal("It was not possible to connect to lichess")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Bad Response code '%v'", response.StatusCode)
	}

	// Unmarshalling JSON data. The result is represented as a map of strings
	// (all keywords are strings in the Body Response from lichess) to instances
	// of an arbitrary type
	var userInfo map[string]interface{}
	if err = json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		log.Fatal("Error Unmarshalling the JSON response")
	}

	// cast now data into a map of strings to arbitrary types
	userMap := JSONtoMap(userInfo)

	// extract now information of the different section in dedicated maps to the
	// correct types

	// -- profile
	profile := make(map[string]string)
	if _, ok := userMap["profile"]; ok {
		for key, value := range userMap["profile"].(map[string]interface{}) {
			if content, ok := value.(string); !ok {
				log.Print("It was not possible to cast one 'profile' into a string")
				continue
			} else {
				profile[key] = content
			}
		}
	}

	// -- count
	// Note that the the numerical type supported by the JSON package is float64
	count := make(map[string]float64)
	for key, value := range userMap["count"].(map[string]interface{}) {
		if content, ok := value.(float64); !ok {
			log.Fatal("It was not possible to cast one 'count' into a floating-point number")
		} else {
			count[key] = content
		}
	}

	// -- playTime
	// Note that the the numerical type supported by the JSON package is float64
	playTime := make(map[string]float64)
	for key, value := range userMap["playTime"].(map[string]interface{}) {
		if content, ok := value.(float64); !ok {
			log.Fatal("It was not possible to cast one 'playTime' into a floating-point number")
		} else {
			playTime[key] = content
		}
	}

	// -- perfs
	//
	// in the following, only the numerical fields are considered. Thus, "prov"
	// is ignored
	perfs := make(map[string]map[string]float64)
	for variant, stats := range userMap["perfs"].(map[string]interface{}) {
		variantData := make(map[string]float64)
		for stat, value := range stats.(map[string]interface{}) {
			if stat != "prov" {
				if content, ok := value.(float64); !ok {
					log.Fatal("It was not possible to cast one 'perf' into a floating-point number")
				} else {
					variantData[stat] = content
				}
			}
		}
		perfs[variant] = variantData
	}

	// show data

	// -- identity
	fmt.Println()
	fmt.Printf(" %v ", userMap["username"])
	firstName, okFirstName := profile["firstName"]
	lastName, okLastName := profile["lastName"]
	if okFirstName || okLastName {
		fmt.Printf("(")
		if okLastName {
			fmt.Printf("%v, ", lastName)
		}
		if okFirstName {
			fmt.Printf(firstName)
		}
		fmt.Printf(")")
	}
	fmt.Println()
	fmt.Println()

	// -- Geographical data
	country, okCountry := profile["country"]
	location, okLocation := profile["location"]
	if okCountry || okLocation {
		fmt.Print(" * ")
		if okLocation {
			fmt.Printf("%v", location)
			if okCountry {
				fmt.Print(",")
			}
		}
		if okCountry {
			fmt.Print(country)
		}
		fmt.Println()
		fmt.Println()
	}

	// -- time spent on Lichess
	tabber := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintf(tabber, "\tTime spent\t\n")
	for key, value := range playTime {
		fmt.Fprintf(tabber, "%s\t%d\t\n", key, int(value))
	}
	tabber.Flush()
	fmt.Println()

	// -- number of games played
	//
	// for this a tabwriter is used
	fmt.Fprint(tabber, "\t# Games\t\n")
	for category, nbgames := range count {
		fmt.Fprintf(tabber, "%s\t%d\t\n", category, int(nbgames))
	}
	tabber.Flush()
	fmt.Println()

	// -- stats of the different variants
	//
	// again, a tabber is used for this
	fmt.Fprintf(tabber, "Variant\tRating\tIncr.\tRd\t# Games\t\n")
	for variant, stats := range perfs {
		fmt.Fprintf(tabber, "%s\t%d\t%d\t%d\t%d\t\n", variant,
			int(stats["rating"]), int(stats["prog"]), int(stats["rd"]), int(stats["games"]))
	}
	tabber.Flush()
	fmt.Printf("\n * Completion rate: %v%%\n", userMap["completionRate"])
	fmt.Println()

	// -- social
	fmt.Printf(" * Following: %v\n", userMap["nbFollowing"])
	fmt.Printf(" * Followers: %v\n\n", userMap["nbFollowers"])

	// -- online
	fmt.Printf(" * Online: %v\n\n", userMap["online"])
}
