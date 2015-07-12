package main
import (
	"github.com/codegangsta/cli"
	"os"
	"encoding/csv"
	"log"
)

func csvRender(c *cli.Context) {
	var dataPath = c.String("data-path");
	if dataPath == "" {
		 log.Fatal("data path is empty")
	}

	file, err := os.Open(dataPath)
	if err != nil {
		 log.Fatalf("cannot open data %s", err)
	}

	reader := csv.NewReader(file)
	reader.Comma = ','
	headerLine, err := reader.Read()

	if err != nil {
		log.Fatalf("read header failed %s", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("can not read all %s", err)
	}

	for _, record := range records {
		fileName := record[0]
		data := map[string]string{}
		err = constructData(&data, c.Args())
		if err != nil {
			log.Fatalf("%s", err)
		}
		addCsvData(&data, headerLine, record)
		t := getTemplate(c)
		output, err := os.Create(fileName)
		if err != nil {
			log.Fatalf("%s", err)
		}

		render(t, data, output)
	}
}

func addCsvData(data *map[string]string, header []string, record []string) {
	for index, key := range header {
		(*data)[key] = record[index]
	}
}