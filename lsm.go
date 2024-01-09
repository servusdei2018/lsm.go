package lsm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

// LSM represents a Linux Software Map (version) entry
//
// Mandatory fields: Title, Version, Entered-date, Description, Author, Primary-site
type LSM struct {
	// Title: the name of the package
	Title string
	// Version: version number or other designation
	Version string
	// EnteredDate: Date in format YYYY-MM-DD of when the LSM entry was last modified
	EnteredDate string
	// Description: short description of the package
	Description string
	// Keywords: short list of keywords that describe the package
	Keywords string
	// Author: original author(s) of the package, in RFC 822 format
	Author string
	// MaintainedBy: maintainer(s) of the package, in RFC 822 format
	MaintainedBy string
	// PrimarySite: A specification of on which site, in which directory, and which files are part of the package
	//
	// Example:
	// 		Primary-site: sunsite.unc.edu /pub/Linux/docs
	//			10kB lsm-1994.01.01.tar.gz
	//			997  lsm-template
	//			22 M /pub/Linux/util/lsm-util.tar.gz
	PrimarySite string
	// AlternateSite: One alternate site may be given, in the same format as PrimarySite
	AlternateSite string
	// OriginalSite: The original package, if this is a port to Linux, in the same format as PrimarySite
	OriginalSite string
	// Platforms: Software or hardware that is required, if unusual
	Platforms string
	// CopyingPolicy: Copying policy.
	//
	// Use "GPL" for GNU General Public License, "BSD" for Berkeley style of copyright, "Shareware"  for shareware,
	// "MIT" for MIT License, and some other description for other styles of copyrights.
	CopyingPolicy string
}

// Parse parses a LSM from an io.Reader
func Parse(r io.Reader) (*LSM, error) {
	rd := bufio.NewReader(r)

	// Check for begin token
	if line, err := rd.ReadString('\n'); err != nil {
		return nil, err
	} else if line != "Begin4\n" {
		return nil, fmt.Errorf("error, LSM doesn't begin with Begin4 token")
	}

	var lsm LSM
	var last *string
	EOF := false
	for !EOF {
		line, err := rd.ReadString('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return nil, err
			}
			EOF = true
		}

		// Blank lines
		if len(line) == 0 {
			continue
		}

		// Continuation of current field
		if startsWithWhitespace(line) {
			*last += line

			// New field
		} else {
			// Get field
			words := strings.Split(line, " ")

			// EOF
			if words[0] == "End" {
				break
			}

			// Field names must end with a colon
			if !strings.HasSuffix(words[0], ":") {
				return nil, fmt.Errorf(`error, field name "%v" doesn't end with a colon`, words[0])
			}

			// Select field
			switch words[0] {
			case "Title:":
				last = &lsm.Title
			case "Version:":
				last = &lsm.Version
			case "Entered-date:":
				last = &lsm.EnteredDate
			case "Description:":
				last = &lsm.Description
			case "Keywords:":
				last = &lsm.Keywords
			case "Author:":
				last = &lsm.Author
			case "Maintained-by:":
				last = &lsm.MaintainedBy
			case "Primary-site:":
				last = &lsm.PrimarySite
			case "Alternate-site:":
				last = &lsm.AlternateSite
			case "Original-site:":
				last = &lsm.OriginalSite
			case "Platforms:":
				last = &lsm.Platforms
			case "Copying-policy:":
				last = &lsm.CopyingPolicy
			default:
				return nil, fmt.Errorf(`error, unknown field "%v"`, words[0])
			}

			// Update field
			*last += line[len(words[0]):]
		}
	}

	return &lsm, nil
}

// startsWithWhitespace returns whether a string begins with whitespace
func startsWithWhitespace(s string) bool {
	for _, char := range s {
		if !unicode.IsSpace(char) {
			return false
		}
		//lint:ignore SA4004 this unconditional termination is intended
		break
	}
	return true
}
