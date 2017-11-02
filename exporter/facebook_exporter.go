package exporter

import (
	"encoding/csv"
	"fmt"
	"github.com/schmidtjohannes/social-media-data-mining/model"
	"os"
	"strconv"
	"time"
)

func ExportFacebookData(data model.FacebookStatistics) error {
	ts := getCurrentTimestamp()
	file, err := os.Create(fmt.Sprintf("stats-%s.csv", ts))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	//header
	err = writer.Write([]string{"Group-ID", "Post", "Likes", "Comments"})
	if err != nil {
		return err
	}

	for grpId, fbGroupStatistics := range data.GroupData {
		for detailIdx := range fbGroupStatistics.Details {
			detail := fbGroupStatistics.Details[detailIdx]
			s := []string{grpId, detail.Post, strconv.Itoa(int(detail.Likes)), strconv.Itoa(detail.Comments)}
			err = writer.Write(s)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getCurrentTimestamp() string {
	current := time.Now()
	return current.Format("2006-01-02_15:04:05")
}
