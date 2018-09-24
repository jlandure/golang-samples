// Copyright 2017 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"time"
	"log"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
)

func addDocAsMap(ctx context.Context, client *firestore.Client) error {
	// [START fs_add_simple_doc_as_map]
	_, err := client.Collection("cities").Doc("LA").Set(ctx, map[string]interface{}{
		"name":    "Los Angeles",
		"state":   "CA",
		"country": "USA",
	})
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_add_simple_doc_as_map]
	return err
}

func addDocDataTypes(ctx context.Context, client *firestore.Client) error {
	// [START fs_add_doc_data_types]
	doc := make(map[string]interface{})
	doc["stringExample"] = "Hello world!"
	doc["booleanExample"] = true
	doc["numberExample"] = 3.14159265
	doc["dateExample"] = time.Now()
	doc["arrayExample"] = []interface{}{5, true, "hello"}
	doc["nullExample"] = nil
	doc["objectExample"] = map[string]interface{}{
		"a": 5,
		"b": true,
	}

	_, err := client.Collection("data").Doc("one").Set(ctx, doc)
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_add_doc_data_types]
	return err
}

func addDocWithID(ctx context.Context, client *firestore.Client) error {
	var data = make(map[string]interface{})
	// [START fs_add_doc_with_id]
	_, err := client.Collection("cities").Doc("new-city-id").Set(ctx, data)
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_add_doc_with_id]
	return err
}

func addDocWithoutID(ctx context.Context, client *firestore.Client) error {
	// [START fs_add_doc_auto_id]
	_, _, err := client.Collection("cities").Add(ctx, map[string]interface{}{
		"name":    "Tokyo",
		"country": "Japan",
	})
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_add_doc_auto_id]
	return err
}

func addDocAsEntity(ctx context.Context, client *firestore.Client) error {
	// [START fs_add_simple_doc_as_entity]
	city := City{
		Name:    "Los Angeles",
		Country: "USA",
	}
	_, err := client.Collection("cities").Doc("LA").Set(ctx, city)
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_add_simple_doc_as_entity]
	return err
}

func addDocAfterAutoGeneratedID(ctx context.Context, client *firestore.Client) error {
	data := City{
		Name:    "Sydney",
		Country: "Australia",
	}

	// [START fs_add_doc_data_after_auto_id]
	ref := client.Collection("cities").NewDoc()

	// later...
	_, err := ref.Set(ctx, data)
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_add_doc_data_after_auto_id]
	return err
}

func updateDoc(ctx context.Context, client *firestore.Client) error {
	// [START fs_update_doc]
	_, err := client.Collection("cities").Doc("DC").Set(ctx, map[string]interface{}{
		"capital": true,
	}, firestore.MergeAll)
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_update_doc]
	return err
}

func updateDocCreateIfMissing(ctx context.Context, client *firestore.Client) error {
	// [START fs_update_create_if_missing]
	_, err := client.Collection("cities").Doc("BJ").Set(ctx, map[string]interface{}{
		"capital": true,
	}, firestore.MergeAll)

	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_update_create_if_missing]
	return err
}

func updateDocMultiple(ctx context.Context, client *firestore.Client) error {
	_, err := client.Collection("cities").Doc("Delhi").Set(ctx, City{Name: "Delhi"})
	if err != nil {
		return err
	}

	// [START fs_update_multiple_fields]
	_, err = client.Collection("cities").Doc("Delhi").Set(ctx, map[string]interface{}{
		"capital":           true,
		"country":           "India",
		"population":        16787941,
		"areaInSquareMiles": 573.0,
	}, firestore.MergeAll)
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_update_multiple_fields]
	return err
}

func updateDocNested(ctx context.Context, client *firestore.Client) error {
	// [START fs_update_nested_fields]
	initialData := map[string]interface{}{
		"name": "Frank",
		"age":  12,
		"favorites": map[string]interface{}{
			"food":    "Pizza",
			"color":   "Blue",
			"subject": "recess",
		},
	}

	// [START_EXCLUDE]
	_, preErr := client.Collection("users").Doc("frank").Set(ctx, initialData)
	if preErr != nil {
		return preErr
	}
	// [END_EXCLUDE]

	_, err := client.Collection("users").Doc("frank").Set(ctx, map[string]interface{}{
		"age": 13,
		"favorites": map[string]interface{}{
			"color": "Red",
		},
	}, firestore.MergeAll)
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_update_nested_fields]
	return err
}

func updateDocServerTimestamp(ctx context.Context, client *firestore.Client) error {
	_, preErr := client.Collection("objects").Doc("some-id").Set(ctx, map[string]interface{}{
		"timestamp": 0,
	})
	if preErr != nil {
		return preErr
	}

	// [START fs_update_server_timestamp]
	_, err := client.Collection("objects").Doc("some-id").Set(ctx, map[string]interface{}{
		"timestamp": firestore.ServerTimestamp,
	}, firestore.MergeAll)
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_update_server_timestamp]
	return err
}

func deleteDoc(ctx context.Context, client *firestore.Client) error {
	// [START fs_delete_doc]
	_, err := client.Collection("cities").Doc("DC").Delete(ctx)
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_delete_doc]
	return err
}

func deleteField(ctx context.Context, client *firestore.Client) error {
	// [START fs_delete_field]
	_, err := client.Collection("cities").Doc("BJ").Update(ctx, []firestore.Update{
		{
			Path:  "capital",
			Value: firestore.Delete,
		},
	})
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_delete_field]

	// Use Set once this feature is implemented:
	// https://github.com/GoogleCloudPlatform/google-cloud-go/issues/832
	// Set(ctx, map[string]interface{}{
	//	"capital": firestore.Delete,
	//})
	return err
}

// [START fs_delete_collection]
func deleteCollection(ctx context.Context, client *firestore.Client,
		ref *firestore.CollectionRef, batchSize int) error {

	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}

// [END fs_delete_collection]

func runSimpleTransaction(ctx context.Context, client *firestore.Client) error {
	_, preErr := client.Collection("cities").Doc("SF").Set(ctx, map[string]interface{}{
		"population": 860000,
	})

	if preErr != nil {
		return preErr
	}

	// [START fs_run_simple_transaction]
	ref := client.Collection("cities").Doc("SF")
	err := client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(ref) // tx.Get, NOT ref.Get!
		if err != nil {
			return err
		}
		pop, err := doc.DataAt("population")
		if err != nil {
			return err
		}
		return tx.Set(ref, map[string]interface{}{
			"population": pop.(int64) + 1,
		}, firestore.MergeAll)
	})
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_run_simple_transaction]
	return err
}

func infoTransaction(ctx context.Context, client *firestore.Client) error {
	// [START fs_return_info_transaction]
	ref := client.Collection("cities").Doc("SF")
	err := client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(ref)
		if err != nil {
			return err
		}
		pop, err := doc.DataAt("population")
		if err != nil {
			return err
		}
		newpop := pop.(int64) + 1
		if newpop <= 1000000 {
			return tx.Set(ref, map[string]interface{}{
				"population": pop.(int64) + 1,
			}, firestore.MergeAll)
		}
		return errors.New("population is too big")
	})
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_return_info_transaction]
	return err
}

func batchWrite(ctx context.Context, client *firestore.Client) error {
	// [START fs_batch_write]
	// Get a new write batch.
	batch := client.Batch()

	// Set the value of "NYC".
	nycRef := client.Collection("cities").Doc("NYC")
	batch.Set(nycRef, map[string]interface{}{
		"name": "New York City",
	})

	// Update the population of "SF".
	sfRef := client.Collection("cities").Doc("SF")
	batch.Set(sfRef, map[string]interface{}{
		"population": 1000000,
	}, firestore.MergeAll)

	// Delete the city "LA".
	laRef := client.Collection("cities").Doc("LA")
	batch.Delete(laRef)

	// Commit the batch.
	_, err := batch.Commit(ctx)
	if err != nil {
		// Handle the error correctly here.
		log.Printf("An error has occurred: %s", err)
	}
	// [END fs_batch_write]
	return err
}
