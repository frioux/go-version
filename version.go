package version // import "go.hein.dev/go-version"

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

// Info creates a formattable struct for output
type Info struct {
	Version string `json:"Version,omitempty"`
	Commit  string `json:"Commit,omitempty"`
	Date    string `json:"Date,omitempty"`
}

// New will create a pointer to a new version object
func New(version string, commit string, date string) *Info {
	return &Info{
		Version: version,
		Commit:  commit,
		Date:    date,
	}
}

// Func will add the versioning code
func Func(out io.Writer, shortened bool, version, commit, date string) func(*cobra.Command, []string) {
	return func(_ *cobra.Command, _ []string) {
		var response string
		versionOutput := New(version, commit, date)

		if shortened {
			response = versionOutput.ToShortened()
		} else {
			response = versionOutput.ToJSON()
		}
		fmt.Fprintf(out, "%+v", response)
		return
	}
}

// ToJSON converts the Info into a JSON String
func (v *Info) ToJSON() string {
	bytes, _ := json.Marshal(v)
	return string(bytes) + "\n"
}

// ToShortened converts the Info into a JSON String
func (v *Info) ToShortened() (str string) {
	var version, commit, date string
	if v.Version != "" {
		version = "Version: " + v.Version
	}
	if v.Commit != "" {
		commit = "Commit: " + v.Commit
	}
	if v.Date != "" {
		date = "Date: " + v.Date
	}
	values := []string{version, commit, date}
	values = deleteEmpty(values)
	str = strings.Join(values, "\n")
	return str + "\n"
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
