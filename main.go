package auth

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// package constants
const (
	// Iterations -- the number of times to hash the password
	Iterations = 2 << 13
	// KeyLength -- the length of the token to store.
	KeyLength = 2 << 5
	// SaltSize -- the number of bytes of salt entropy to include
	SaltSize = 2 << 4
	// LineSeparator separates the lines when dumping or reading from the
	// string format. The separator substrings are completely arbitrary, the
	// only requirement for Line/ColSeparator characters is
	// =~ /[\dabcdef$]*/
	// that is, it must not be any of: numeric digits, the letters A through F,
	// or the $ symbol. Anything else is fair game.
	LineSeparator = "\n"
	// ColSeparator separates the columns when dumping or reading from the
	// string format
	ColSeparator = "-|-"
	WHITESPACE   = "\n\t "
	// wordListURL is where to download the list of words to choose from
	wordListURL = "http://svnweb.freebsd.org/csrg/share/dict/words" +
		"?view=co&content-type=text/plain"
)

// ConfigLocation is where the usernames and (encrypted) passwords are stored
var ConfigLocation string

// wordListLocation is where the word list should be stored
var wordListLocation string

func init() {
	AllUsers = make(map[User]*AuthToken)
	wordListLocation = path.Join(ConfigLocation, "..", "wordlist.txt")
}

// Initialize global variables
func Initialize(config string) error {
	ConfigLocation = config
	info, err := os.Stat(ConfigLocation)
	if os.IsExist(err) || (err == nil && info.Size() > 0) {
		AllUsers, err = ReadFrom(config)
		if err != nil {
			return err
		}
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}
	if len(AllUsers) < 1 {
		if err = PromptForSingleUser(); err != nil {
			return err
		}
	}
	if err = SyncAllUsers(); err != nil {
		return err
	}
	readValues, err := ReadFrom(config)
	if err != nil {
		return err
	}
	for k, v := range AllUsers {
		if AllUsers[k] != v {
			return fmt.Errorf(
				"got mismatch after write:\n%#+v\n%#+v\nAT %s: %#+v != %#+v",
				AllUsers,
				readValues,
				k,
				readValues[k],
				v,
			)
		}
	}
	return nil
}

func ReadFrom(config string) (map[User]*AuthToken, error) {
	txt, err := ioutil.ReadFile(config)
	if err != nil {
		return map[User]*AuthToken{}, err
	}
	text := string(txt)
	return FromStringToValues(text)
}
