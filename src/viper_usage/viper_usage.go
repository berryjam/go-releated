package main

import (
	"os"
	"github.com/spf13/viper"
	"fmt"
	"strings"
)

func main() {
	os.Setenv("ORDERER_GENERAL_LOCALMSPDIR", "/var/hyperledger/orderer/msp")
	os.Setenv("ORDERER_FOO", "bar")
	viper.SetEnvPrefix("ORDERER")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading configuration:", err)
	}

	//fmt.Printf("ORDERER_GENERAL_LOCALMSPDIR=%v\n", viper.Get("GENERAL_LOCALMSPDIR"))
	//
	////viper.AutomaticEnv()
	fmt.Printf("foo=%v\n", viper.Get("foo"))

	baseKeys := viper.AllSettings()
	for k, v := range baseKeys {
		fmt.Printf("%v=%v\n", k, v)
	}
}
