package main

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/log/model"
)

func main() {
	s := gocron.NewScheduler(time.Now().Local().Location())
	s.Every(1).Second().Do(collectData)
	s.StartAsync()
	s.StartBlocking()
}

func collectData() {
	data := []model.LogData{
		{
			Healthcare: "hospital_siloam",
			DBName:     "database_2",
			TBName:     "patient_info",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_mayapada ",
			DBName:     "database_3",
			TBName:     "appointments",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_roem",
			DBName:     "database_9",
			TBName:     "invoices",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_eldelweis",
			DBName:     "database_11",
			TBName:     "invoices",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_hasan",
			DBName:     "database_7",
			TBName:     "patient_name",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		}, {
			Healthcare: "hospital_hasani",
			DBName:     "database_10",
			TBName:     "room_info",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_internasional",
			DBName:     "database_69",
			TBName:     "rawat_jalan",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_nasi_padang",
			DBName:     "database_14",
			TBName:     "rawat_inap",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_raos",
			DBName:     "database 17",
			TBName:     "rawat_rujukan",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_nasional",
			DBName:     "database_19",
			TBName:     "rawatan_inap",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_bermuda",
			DBName:     "database_21",
			TBName:     "rawat_saja",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_bertua",
			DBName:     "database_1",
			TBName:     "rawat_sembuh",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_nonstop",
			DBName:     "database_23",
			TBName:     "rawat_sekarang",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_uks",
			DBName:     "database_25",
			TBName:     "rawat_untuk",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_telkom",
			DBName:     "database_27",
			TBName:     "rawat_opname",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_univ",
			DBName:     "database_30",
			TBName:     "rawat_kos",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_militer",
			DBName:     "database_33",
			TBName:     "rawat_intensif",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_umum",
			DBName:     "database_36",
			TBName:     "rawat_vip",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_daerah",
			DBName:     "database_39",
			TBName:     "rawat_klinik",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
		{
			Healthcare: "hospital_gigi",
			DBName:     "database_40",
			TBName:     "rawat_gigi",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		},
	}
	err := SendDataToProducer(data)

	if err != nil {
		log.Println(err)
		return
	}
}
