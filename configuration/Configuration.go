package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

/**
Define a simple database configuration
*/
type Configuration struct {
	//Load in the params from json
	params map[string]interface{}
}

//Provide a function to create a new one
func NewConfiguration(configFiles ...string) *Configuration {
	//Define a Configuration
	config := Configuration{}

	//Now march over each file
	for _, configFile := range configFiles {
		//Load in the file
		configFileStream, err := os.Open(configFile)
		defer configFileStream.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		//Get the json and add to the params
		jsonParser := json.NewDecoder(configFileStream)
		jsonParser.Decode(&config.params)
	}

	//Return it
	return &config
}

/**
 * Add function to get item
 */
func (config *Configuration) Get(key string) interface{} {
	//Get the key from the file
	param := config.params[key]

	//Now see if it is specified in the env
	systemEnvVar := os.Getenv(key)

	//If it is not empty set it
	if len(systemEnvVar) != 0 {
		param = systemEnvVar
	}

	return param

}

/**
 * Add function to get item
 */
func (config *Configuration) GetString(key string) string {
	//Get the key from the
	return fmt.Sprint(config.Get(key))

}

/**
 * Add function to get item
 */
func (config *Configuration) GetInt(key string) (int, error) {
	//Get the key from the
	res, err := strconv.Atoi(config.GetString(key))

	return res, err

}
