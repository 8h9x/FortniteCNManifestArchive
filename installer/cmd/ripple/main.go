package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/8h9x/FortniteCNManifestArchive/installer/internal/manifest"
)

func main() {
	// default manifest file path
	_ = flag.String("out", "", "")
	_ = flag.String("host", "", "")
	_ = flag.String("remote-manifest", "", "")

	res, err := http.Get("https://wegame-manifest-mirror.vaultnite.com/2000196/100_verify.json")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var depot manifest.Depot
	if err := json.NewDecoder(res.Body).Decode(&depot); err != nil {
		log.Fatal(err)
	}

	for _, file := range depot.Files {
		parts := strings.Split(file.Name, ".")

		for _, part := range parts {
			if part == "7z" {
			}
		}

		if parts[len(parts)-2] == "7z" { // get file type without the archive index suffix
			log.Println(file.Name, "queued for download")
		}
	}

	// log.Println(depot)
}

// opts: {manifest-filepath} --out {filepath} --host {wgdl.qq.com|dlied1.qq.com|down.qq.com} --remote-manifest {url}
