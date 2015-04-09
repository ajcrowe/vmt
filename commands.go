package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

var Commmands = []cli.Command{
	cli.Command{
		Name:        "generate",
		Usage:       "generate a new vagrant box metadata file",
		Description: `...`,
		Action:      doGenerateMetadata,
		Flags: []cli.Flag{
			DescFlag(),
			DescShortFlag(),
			BoxNameFlag(),
			OutputFlag(),
		},
	},
	cli.Command{
		Name:        "version",
		Usage:       "modify version data for a box metadata file",
		Description: `...`,
		Subcommands: []cli.Command{
			{
				Name:        "list",
				Usage:       "list versions in metadata",
				Description: `...`,
				Action:      doVersionList,
				Flags: []cli.Flag{
					InputFlag(),
				},
			},
			{
				Name:        "add",
				Usage:       "add new version to existing metadata",
				Description: `...`,
				Action:      doVersionAdd,
				Flags: []cli.Flag{
					QuietFlag(),
					NoopFlag(),
					BoxVerFlag(),
					DescFlag(),
					InputFlag(),
					OutputFlag(),
					BoxFileNameFlag(),
					ProviderFlag(),
				},
			},
			{
				Name:        "remove",
				Usage:       "remove version from existing metadata",
				Description: `...`,
				Action:      doVersionRemove,
				Flags: []cli.Flag{
					QuietFlag(),
					BoxVerFlag(),
					InputFlag(),
					OutputFlag(),
					BoxRmFlag(),
					ProviderFlag(),
				},
			},
		},
	},
}

// generate metadata file
func doGenerateMetadata(c *cli.Context) {
	// get data from flags
	var boxname = c.String("boxname")
	var output = c.String("output")
	var description = c.String("description")
	var shortdesc = c.String("shortdescription")
	// initial variables
	var stdout = false
	var md *Metadata

	if boxname == "" {
		printError("Please specify a boxname")
		os.Exit(1)
	}
	if output == "" {
		stdout = true
	}

	// generate a metadata structure with our data
	md = NewMetadata(boxname, description, shortdesc)

	// if stdout is true output to stdout otherwise write to the file
	if stdout {
		printInfo(md.printJSON())
	} else {
		writeMetaData(output, md)
	}
}

// list versions in metadata
func doVersionList(c *cli.Context) {
	input := c.String("input")
	if input == "" {
		printError("Please specify a filename")
		os.Exit(1)
	}
	// load metadata
	md := readMetaData(input)
	// create tabbed writer
	w := getTabWriter()
	defer w.Flush()

	// write header unless quiet flag
	if !c.Bool("quiet") {
		writeHeader(w, "Version", "Description", "Status")
	}
	// loop over the versions in the metadata and print information
	for _, version := range md.Versions {
		writeLine(w, version.Version, version.DescriptionMarkDown, version.Status)
	}
}

// add a new version to existing box metadata
func doVersionAdd(c *cli.Context) {
	// get data from flags
	var input = c.String("input")
	var version = c.String("version")
	var boxfile = c.String("boxfile")
	var description = c.String("description")
	var provider = c.String("provider")
	// if output is not specified default to replacing input
	output := c.String("output")
	if output == "" {
		output = input
	}
	if provider == "" {
		provider = config.DefaultProvider
	}
	// validate params
	if boxfile == "" {
		printError("No boxfile specified")
		os.Exit(1)
	}
	if input == "" {
		printError("No input file specified")
		os.Exit(1)
	}
	if version == "" {
		printError("No version specified")
		os.Exit(1)
	}

	//  read existing json metadata file
	md := readMetaData(input)

	// check version doesn't exist
	if md.versionExists(version) {
		printError("Version already exists in metadata")
		os.Exit(1)
	}

	// create a new version struct with the passed parameters
	newversion := Version{
		Version:             version,
		Status:              "active",
		DescriptionHTML:     fmt.Sprintf("<p>%s</p>", description),
		DescriptionMarkDown: description,
		Providers: []VersionProvider{
			VersionProvider{
				Name:         provider,
				URL:          fmt.Sprintf("%s/%s", config.BoxURL, boxfile),
				CheckSumType: "sha1",
				CheckSum:     generateCheckSum(boxfile),
			},
		},
	}

	// append new version to existing in metadata
	md.Versions = append(md.Versions, newversion)

	// print and write file
	if !c.Bool("quiet") {
		printInfo(md.printJSON())
	}

	if !c.Bool("noop") {
		writeMetaData(output, md)
	}
}

func doVersionRemove(c *cli.Context) {
	var input = c.String("input")
	var output = c.String("output")
	var version = c.String("version")

	// if output is not specified default to replacing input
	if output == "" {
		output = input
	}

	//  read existing json metadata file
	md := readMetaData(input)

	// check the version exists in the metadata already
	if !md.versionExists(version) {
		printError("version not found in metadata")
		os.Exit(1)
	}

	// create a new slice for the non-matching versions to be copies to
	var newvers []Version
	// interate over the versions present and copy non-matching to new slice
	for _, ver := range md.Versions {
		if ver.Version != version {
			newvers = append(newvers, ver)
		}
	}
	// overwrite the existing versions with the new versions
	md.Versions = newvers

	// print and write file
	if !c.Bool("quiet") {
		printInfo(md.printJSON())
	}
	if !c.Bool("noop") {
		writeMetaData(output, md)
	}

}
