// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/pkg/errors"

	"github.com/elastic/package-registry/util"
)

var (
	imageRe            = regexp.MustCompile(`image::[^\[]+`)
	imageTitleReplacer = strings.NewReplacer("_", " ", "-", " ", "/", "")
)

type imageContent struct {
	source string
}

func createImages(beatDocsPath, modulePath string) ([]imageContent, error) {
	var images []imageContent

	moduleDocsPath := path.Join(modulePath, "_meta", "docs.asciidoc")
	moduleDocsFile, err := ioutil.ReadFile(moduleDocsPath)
	if err != nil && !os.IsNotExist(err) {
		return nil, errors.Wrapf(err, "reading module docs file failed (path: %s)", moduleDocsPath)
	} else if os.IsNotExist(err) {
		log.Printf("\tNo docs found (path: %s), skipped", moduleDocsPath)
	} else {
		log.Printf("\tDocs found (path: %s)", moduleDocsPath)
		images = append(images, extractImages(beatDocsPath, moduleDocsFile)...)
	}

	datasetDirs, err := ioutil.ReadDir(modulePath)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read module directory %s", modulePath)
	}

	for _, datasetDir := range datasetDirs {
		if !datasetDir.IsDir() {
			continue
		}
		datasetName := datasetDir.Name()

		if datasetName == "_meta" {
			continue
		}

		log.Printf("\t%s: dataset found", datasetName)

		datasetDocsPath := path.Join(modulePath, datasetName, "_meta", "docs.asciidoc")
		datasetDocsFile, err := ioutil.ReadFile(datasetDocsPath)
		if err != nil && !os.IsNotExist(err) {
			return nil, errors.Wrapf(err, "reading dataset docs file failed (path: %s)", datasetDocsPath)
		} else if os.IsNotExist(err) {
			log.Printf("\t%s: no docs found (path: %s), skipped", datasetName, datasetDocsPath)
			continue
		}

		log.Printf("\t%s: docs found (path: %s)", datasetName, datasetDocsPath)
		images = append(images, extractImages(beatDocsPath, datasetDocsFile)...)
	}

	return images, nil
}

func extractImages(beatDocsPath string, docsFile []byte) []imageContent {
	matches := imageRe.FindAll(docsFile, -1)

	var contents []imageContent
	for _, match := range matches {
		contents = append(contents, imageContent{
			source: path.Join(beatDocsPath, string(match[7:])), // skip: image::
		})
	}
	return contents
}

func createManifestImages(images []imageContent) ([]util.Image, error) {
	var imgs []util.Image
	for _, image := range images {
		i := strings.LastIndex(image.source, "/")
		sourceFileName := image.source[i:]

		imageSize, err := readImageSize(image.source)
		if err != nil {
			return nil, errors.Wrapf(err, "reading image size failed")
		}

		imageType, err := extractImageType(image.source)
		if err != nil {
			return nil, errors.Wrapf(err, "extracting image type failed")
		}

		imgs = append(imgs, util.Image{
			Src:   path.Join("img", sourceFileName),
			Title: toImageTitle(sourceFileName),
			Size:  imageSize,
			Type:  imageType,
		})
	}
	return imgs, nil
}

func toImageTitle(fileName string) string {
	i := strings.LastIndex(fileName, ".")
	return strings.Title(imageTitleReplacer.Replace(fileName[:i]))
}

func readImageSize(imagePath string) (string, error) {
	f, err := os.Open(imagePath)
	if err != nil {
		return "", errors.Wrapf(err, "opening image failed (path: %s)", imagePath)
	}

	img, _, err := image.DecodeConfig(f)
	if err != nil {
		return "", errors.Wrapf(err, "opening image failed (path: %s)", imagePath)
	}
	return fmt.Sprintf("%dx%d", img.Width, img.Height), nil
}

func extractImageType(imagePath string) (string, error) {
	if strings.HasSuffix(imagePath, ".png") {
		return "image/png", nil
	} else if strings.HasSuffix(imagePath, ".jpg") {
		return "image/jpg", nil
	}
	return "", fmt.Errorf("unknown image type (path: %s)", imagePath)
}
