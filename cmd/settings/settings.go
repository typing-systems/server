package settings

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type settings struct {
	DEBUG bool
}

var Values settings

func Load(env string) {
	filename := env + "_settings.json"

	if _, err := os.Stat(filepath.Clean(filename)); errors.Is(err, os.ErrNotExist) {
		// filename does not exist
		Values = generate(env)
	} else if err != nil {
		log.Fatalf("error detecting if settings file exists: %v", err)
	} else {
		f, err := os.Open(filepath.Clean(filename))
		if err != nil {
			log.Fatalf("error opening settings file: %v", err)
		}

		decodeJSON(f)

		if err := f.Close(); err != nil {
			log.Fatalf("error closing settings file: %v", err)
		}
	}
}

func decodeJSON(f *os.File) {
	decoder := json.NewDecoder(f)
	err := decoder.Decode(&Values)
	if err != nil {
		log.Fatalf("error decoding json from settings file: %v", err)
	}
}

func generate(env string) settings {
	// default settings
	s := settings{
		DEBUG: false,
	}

	jsonCfg, err := json.MarshalIndent(s, "", "	")
	if err != nil {
		log.Fatalf("error generating settings file json: %v", err)
	}

	// 0600 file perm means read/write by owner only
	filename := env + "_settings.json"
	err = ioutil.WriteFile(filepath.Clean(filename), jsonCfg, 0600)
	if err != nil {
		log.Fatalf("error writing settings file: %v", err)
	}

	return s
}
