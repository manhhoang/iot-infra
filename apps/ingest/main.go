package main

import (
	"fmt"
	"context"
	"google.golang.org/api/iterator"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"cloud.google.com/go/bigquery"
	"io"
)

func main() {

}

func comsume() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}

// Item represents a row item.
type Item struct {
	Name string
	Age  int
}

// Save implements the ValueSaver interface.
func (i *Item) Save() (map[string]bigquery.Value, string, error) {
	return map[string]bigquery.Value{
		"full_name": i.Name,
		"age":       i.Age,
	}, "", nil
}

// insertRows demonstrates inserting data into a table using the streaming insert mechanism.
func insertRows(projectID, datasetID, tableID string) error {
	// projectID := "my-project-id"
	// datasetID := "mydataset"
	// tableID := "mytable"
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	inserter := client.Dataset(datasetID).Table(tableID).Inserter()
	items := []*Item{
		// Item implements the ValueSaver interface.
		{Name: "Phred Phlyntstone", Age: 32},
		{Name: "Wylma Phlyntstone", Age: 29},
	}
	if err := inserter.Put(ctx, items); err != nil {
		return err
	}
	return nil
}

func insertData() {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		// TODO: Handle error.
	}

	q := client.Query(`
    SELECT year, SUM(number) as num
    FROM ` + "`bigquery-public-data.usa_names.usa_1910_2013`" + `
    WHERE name = "William"
    GROUP BY year
    ORDER BY year
`)
	it, err := q.Read(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

// queryBasic demonstrates issuing a query and reading results.
func queryBasic(w io.Writer, projectID string) error {
	// projectID := "my-project-id"
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	q := client.Query(
		"SELECT name FROM `bigquery-public-data.usa_names.usa_1910_2013` " +
			"WHERE state = \"TX\" " +
			"LIMIT 100")
	// Location must match that of the dataset(s) referenced in the query.
	q.Location = "US"
	// Run the query and print results when the query job is completed.
	job, err := q.Run(ctx)
	if err != nil {
		return err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		return err
	}
	if err := status.Err(); err != nil {
		return err
	}
	it, err := job.Read(ctx)
	for {
		var row []bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Fprintln(w, row)
	}
	return nil
}