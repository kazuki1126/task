package db

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

type Task struct {
	Key   int
	Value string
}

var taskBucket = []byte("tasks")
var completedTaskBucket = []byte("completed")
var db *bolt.DB

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(taskBucket); err != nil {
			return err
		} else if _, err := tx.CreateBucketIfNotExists(completedTaskBucket); err != nil {
			return err
		}
		return nil
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		return b.Put(itob(id), []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}
func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func StoreCompleteTask(taskValue string) error {
	key := []byte(taskValue)
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(completedTaskBucket)
		value := time.Now().Format(time.RFC3339Nano)
		return b.Put(key, []byte(value))
	})
}

func AllCompleteTasks() ([]string, error) {
	var doneTasks []string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(completedTaskBucket)
		c := b.Cursor()
		oneDay := []byte(time.Now().Add(-(24 * time.Hour)).Format(time.RFC3339Nano))
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if bytes.Compare(v, oneDay) < 0 {
				continue
			} else {
				doneTasks = append(doneTasks, string(k))
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return doneTasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
