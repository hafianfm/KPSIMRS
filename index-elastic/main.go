package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/uuid"
)

type LogData struct {
	Healthcare string
	DBName     string
	TBName     string
	Status     string
	DateTime   time.Time
	CreatedAt  time.Time
	RecordId   uuid.UUID
}

func main() {
	cfg := elasticsearch.Config{
		CloudID: "SIMRS:dXMtY2VudHJhbDEuZ2NwLmNsb3VkLmVzLmlvJDhjMmViMTcwNGEwNjQ2MzdiNDhlYmVmMjI5NjJhZDA5JDEwNjA0ZDhlNzZiNTQzNTFhNzgzMTVmZTVlMjQ4M2U4",
		APIKey:  "aE16VUtvd0JvWGM2ZGdVOUpVRlU6SVlLTDNGZlVSZy1pNGItbjFsZVh2QQ==",
		
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	data := []LogData{
		{
			Healthcare: "rs_siloam",
			DBName:     "database_2",
			TBName:     "patient_info",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_mayapada ",
			DBName:     "database_3",
			TBName:     "appointments",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_roem",
			DBName:     "database_9",
			TBName:     "invoices",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_eldelweis",
			DBName:     "database_11",
			TBName:     "invoices",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_hasan",
			DBName:     "database_7",
			TBName:     "patient_name",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		}, {
			Healthcare: "rs_hasani",
			DBName:     "database_10",
			TBName:     "room_info",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_internasional",
			DBName:     "database_69",
			TBName:     "rawat_jalan",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_kemayoran",
			DBName:     "database_14",
			TBName:     "rawat_inap",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_sahid",
			DBName:     "database 17",
			TBName:     "rawat_rujukan",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_nasional",
			DBName:     "database_19",
			TBName:     "rawatan_inap",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_bermuda",
			DBName:     "database_21",
			TBName:     "rawat_saja",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_nusa",
			DBName:     "database_1",
			TBName:     "rawat_sembuh",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_advent",
			DBName:     "database_23",
			TBName:     "rawat_sekarang",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_sentosa",
			DBName:     "database_25",
			TBName:     "rawat_untuk",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_telkom",
			DBName:     "database_27",
			TBName:     "rawat_opname",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_univ",
			DBName:     "database_30",
			TBName:     "rawat_kos",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_militer",
			DBName:     "database_33",
			TBName:     "rawat_intensif",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_umum",
			DBName:     "database_36",
			TBName:     "rawat_vip",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_daerah",
			DBName:     "database_39",
			TBName:     "rawat_klinik",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_gigi",
			DBName:     "database_40",
			TBName:     "rawat_gigi",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
	}

	for _, d := range data {
		// Generate a random time between a certain range
		min := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
		max := time.Now().Unix()
		createdAtTimestamp := rand.Int63n(max-min) + min
		createdAt := time.Unix(createdAtTimestamp, 0)

		body := fmt.Sprintf(`{
		"Healthcare": "%s",
		"DBName": "%s",
		"TBName": "%s",
		"Status": "%s",
		"DateTime": "%s",
		"CreatedAt": "%s",
		"RecordId": "%s"
	}`, d.Healthcare, d.DBName, d.TBName, d.Status, d.DateTime.Format(time.RFC3339), createdAt.Format(time.RFC3339), d.RecordId)

		req := esapi.IndexRequest{
			Index:      "search-logs2", 
			DocumentID: d.RecordId.String(),
			Body:       strings.NewReader(body),
			Refresh:    "true",
		}

		res, err := req.Do(context.Background(), client)
		if err != nil {
			log.Printf("Error indexing document: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Printf("Error indexing document: %s", res.String())
		} else {
			log.Printf("Indexed document with ID: %s", d.RecordId.String())
		}
	}

}
