package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/samber/lo"
)

var latestReleaseJsonUrl = "https://api.github.com/repos/mayswind/AriaNg/releases/latest"

type Asset struct {
	Name        string `json:"name"`
	DownloadUrl string `json:"browser_download_url"`
}

type Release struct {
	Name   string  `json:"name"`
	Assets []Asset `json:"assets"`
}

func main() {
	release := getLatestRelease()
	asset := findAllIneOneZipAsset(release.Assets)
	log.Println("Downloading", asset.Name)
	zipFile := lo.Must(http.Get(asset.DownloadUrl))
	defer zipFile.Body.Close()
	bb := lo.Must(io.ReadAll(zipFile.Body))
	zipReader := lo.Must(zip.NewReader(bytes.NewReader(bb), zipFile.ContentLength))
	Unzip(zipReader, "index.html", "aria-ng.html")
	log.Println("Latest release:", release.Name)
}

func findAllIneOneZipAsset(assets []Asset) Asset {
	for _, asset := range assets {
		if strings.HasSuffix(asset.Name, "AllInOne.zip") {
			return asset
		}
	}
	log.Panicln("No zip asset found")
	return Asset{}
}

func getLatestRelease() Release {
	var release Release
	res := lo.Must(http.Get(latestReleaseJsonUrl))
	defer res.Body.Close()
	lo.Must0(json.NewDecoder(res.Body).Decode(&release))
	return release
}

func Unzip(r *zip.Reader, wanted, as string) {
	for _, f := range r.File {
		if f.Name != wanted {
			continue
		}
		source := lo.Must(f.Open())
		defer source.Close()
		dest := lo.Must(os.OpenFile(as, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode()))
		defer dest.Close()
		lo.Must(io.Copy(dest, source))
	}
}
