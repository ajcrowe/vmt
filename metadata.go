package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Metadata struct {
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"short_description"`
	Versions         []Version `json:"versions"`
}

type Version struct {
	Version             string            `json:"version"`
	Status              string            `json:"status"`
	DescriptionHTML     string            `json:"description_html"`
	DescriptionMarkDown string            `json:"description_markdown"`
	Providers           []VersionProvider `json:"providers"`
}

type VersionProvider struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	CheckSumType string `json:"checksum_type"`
	CheckSum     string `json:"checksum"`
}

func NewMetadata(name, desc, sdesc string) *Metadata {
	return &Metadata{
		Name:             name,
		Description:      desc,
		ShortDescription: sdesc,
	}
}

func (md *Metadata) printJSON() string {
	jsondata, _ := json.MarshalIndent(md, "", "  ")
	return string(jsondata)
}

func (md *Metadata) versionExists(newver string) bool {
	// interate over all the versions to check for a match
	for _, version := range md.Versions {
		if version.Version == newver {
			return true
		}
	}
	return false
}

func readMetaData(filename string) (md *Metadata) {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		printError("Could not read manifest file")
		os.Exit(1)
	}

	err = json.Unmarshal(text, &md)
	if err != nil {
		printError("Could not parse manifest file")
		os.Exit(1)
	}

	return md
}

func writeMetaData(filename string, md *Metadata) {
	jsondata, _ := json.Marshal(md)
	ioutil.WriteFile(filename, jsondata, 0644)
}
