package main

import (
    "flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Manifest struct {
	DefaultEvents []string `yaml:"default_events,omitempty"`
	DefaultPermissions map[string]string `yaml:"default_permissions,omitempty"`
}

func main() {
    manifestPtr := flag.String("manifest", "app.yml", "Manifest file path relative to current directory")
    flag.Parse()
	fmt.Printf("manifestPtr: %s\n", *manifestPtr)
	
	var manifest Manifest
	source, err := ioutil.ReadFile(*manifestPtr)
	if err != nil {
        panic(err)
    }
	// fmt.Printf("Source: %s\n", source);
    err = yaml.Unmarshal(source, &manifest)
    if err != nil {
        panic(err)
    }
    fmt.Printf("DefaultEvents: %v\nDefaultPermissions: %v\n", manifest.DefaultEvents, manifest.DefaultPermissions)
}