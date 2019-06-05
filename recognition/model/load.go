package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
)

func Load(filename string) VectorModel {

	fileContent, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Panic(err)
	}

	var model VectorModel

	err = json.Unmarshal(fileContent, &model.Polygons)

	if err != nil {
		log.Panic(err)
	}

	sort.Slice(model.Polygons, func(i, j int) bool {
		return !model.Polygons[i].Transparent && model.Polygons[j].Transparent
	})

	model.Dimensions = FillModelDimensions(model)
	return model
}
