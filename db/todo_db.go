package todo_db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func InitDB(dbpath string) error {
	var err error
	db, err = bolt.Open(dbpath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

// all the task will be salted by a string at the end
// if its an task in todo then $actv is added
// elif its an removed task then $remv
// elif the task is done then $done
func AddTask(task string) (int, error) {
	var id int
	task += "$actv"
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func get_tasks(tag string) (map[int]Task, error) {
	// tag can be "$actv" or "$remv" or "$done"
	var tasks_map = map[int]Task{}
	var tasks_map_to_display = map[int]Task{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		var idx int = 1
		for k, v := c.First(); k != nil; k, v = c.Next() {
			t_k := btoi(k)
			t_v := string(v)
			if t_v[len(t_v)-5:] != tag {
				continue
			}
			t_add := Task{
				Key:   t_k,
				Value: t_v,
			}
			tasks_map[idx] = t_add
			tasks_map_to_display[idx] = Task{
				Key:   t_k,
				Value: t_v[:len(t_v)-5],
			}
			idx++
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	// fmt.Printf("All tasks (%s) : %v\n", tag, tasks_map)
	return tasks_map_to_display, nil
}

func Get_active_task() (map[int]Task, error) {

	return get_tasks("$actv")
}

func Get_removed_task() (map[int]Task, error) {

	return get_tasks("$remv")
}

func Get_done_task() (map[int]Task, error) {

	return get_tasks("$done")
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

// Doesn't removes the task but the tag is changed from $actv to $remv so it won't be listed in the active list
func Remove_task(task_key int, task_value string) error {

	tag := "$remv"
	task_value += tag
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)

		key := itob(task_key)
		return b.Put(key, []byte(task_value))
	})
	if err != nil {
		return err
	}
	return nil
}

// Doesn't removes the task but the tag is changed from $actv to $done so it won't be listed in the active list
func Done_task(task_key int, task_value string) error {

	tag := "$done"
	task_value += tag
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)

		key := itob(task_key)
		return b.Put(key, []byte(task_value))
	})
	if err != nil {
		return err
	}
	return nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
