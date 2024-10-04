package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// loadConfig loads configuration from defaults or a configuration file.
// If a configuration file is specified it exits the program if the
// reading or parsing fails for any reason.
func loadConfig() {

	// Check for config file and load it into viper, handle errors

	// We only read the config file if the flag is set and exists
	if flag.Lookup("config") != nil && *aConfigFile != "" {
		viper.SetConfigFile(*aConfigFile)
		err := viper.ReadInConfig()
		if err != nil {

			// We got a read error, but it could just be some parsing error
			// This is useful feedback to the user
			if _, ok := err.(viper.ConfigParseError); ok {
				fmt.Fprintf(os.Stderr, "ERROR - Unable to parse config file '%s': %s", *aConfigFile, err)
				os.Exit(1)
			}

			// Some other read error
			fmt.Fprintf(os.Stderr, "ERROR - Unable to read config file '%s': %s\n", *aConfigFile, err)
			os.Exit(1)
		}
	}

	// Load default settings in case some key has not been defined in the config file
	// Note that viper defaults will not override configuration set in the config file
	// hence we can safely set them after reading in the file.
	setConfigDefaults()

	// Allow environment variables to override the config file
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	return
}

// setConfigDefaults will set the configuration defaults to take care of any missing keys
// in the config file
func setConfigDefaults() {

	viper.SetDefault("debug", false)

	viper.SetDefault("http_server.cert_file", "")
	viper.SetDefault("http_server.key_file", "")
	viper.SetDefault("http_server.address", "")
	viper.SetDefault("http_server.port", 8080)
	viper.SetDefault("http_server.concurrency", 0)
	viper.SetDefault("http_server.cache_ttl", -1)
	viper.SetDefault("http_server.read_header_timeout", 20)
	viper.SetDefault("http_server.read_timeout", 60)
	viper.SetDefault("http_server.write_timeout", 60)
	viper.SetDefault("http_server.idle_timeout", 120)
	viper.SetDefault("http_server.max_header_bytes", 1000000) // 1M

}
