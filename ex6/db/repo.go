package db

import (
	"encoding/binary"
	"github.com/boltdb/bolt"
)

func AllTasks() ([]Task, error) {
	var tasks []Task
	var err error
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{Key: btoi(k), Value: string(v)})
		}
		return nil
	})
	return tasks, err
}

func CreateTask(t string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		return b.Put(itob(id), []byte(t))
	})
	if err != nil {
		return id, err
	}
	return id, nil
}

func DeleteTask(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		return b.Delete(itob(id))
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
