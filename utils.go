package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"text/tabwriter"
)

var config *Config

type Config struct {
	BoxURL          string `json:"box_url"`
	BoxFileRoot     string `json:"box_file_root"`
	DefaultProvider string `json:"default_provider"`
}

// load the json configuration and override if environment variables are present
func loadConfig() {
	// set config to defaults
	config = new(Config)
	config.BoxURL = "http://localhost"
	config.BoxFileRoot = "."
	config.DefaultProvider = "virtualbox"
	// read rc file if it exists
	text, err := ioutil.ReadFile(os.Getenv("HOME") + "/.vmtrc")
	if err == nil {
		// unmarshall json to config struct
		json.Unmarshal(text, &config)
	}
	// check if env vars are set and overwrite config if they are
	if os.Getenv("VMT_BOX_URL") != "" {
		config.BoxURL = os.Getenv("VMT_BOX_URL")
	}
	if os.Getenv("VMT_BOX_FILE_ROOT") != "" {
		config.BoxFileRoot = os.Getenv("VMT_BOX_FILE_ROOT")
	}
	if os.Getenv("VMT_DEFAULT_PROVIDER") != "" {
		config.DefaultProvider = os.Getenv("VMT_DEFAULT_PROVIDER")
	}

}

const filechunk = 8192

// generate a sha1 checksum for the specified file
func generateCheckSum(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading box file: %s", filename)
		os.Exit(1)
	}

	defer file.Close()

	info, _ := file.Stat()
	filesize := info.Size()
	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

	hash := sha1.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}

// return a tabbed writter to stdout
func getTabWriter() (w *tabwriter.Writer) {
	w = new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 0, 3, ' ', 0)
	return w
}

// write a line with the specified strings
func writeLine(w *tabwriter.Writer, values ...string) {

	line := ""
	for _, v := range values {
		line += v + "\t"
	}
	fmt.Fprintln(w, line)
}

// write header before the output
func writeHeader(w *tabwriter.Writer, values ...string) {

	header := ""
	for _, v := range values {
		header += v + "\t"
	}
	fmt.Fprintln(w, header)
}

// print an error
func printError(msg string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", msg)
}

// print an info msg
func printInfo(msg string) {
	fmt.Fprintf(os.Stdout, "%s\n", msg)
}
