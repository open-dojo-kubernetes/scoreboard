package dbclient

import (
	"github.com/open-dojo-kubernetes/scoreboard/model"
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"strconv"
	"encoding/json"
)

type IBoltClient interface {
	OpenBoltDb()
	ListAllGames() (chan model.Game)
	StartAGame()
}

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("games.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("GameBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

func (bc *BoltClient) ListAllGames() (gameChannel chan model.Game) {
	game := model.Game{}
	bc.boltDB.View(func(tx *bolt.Tx) error {
		tx.Bucket([]byte("GameBucket")).ForEach(func(k, gameBytes []byte) error {
			json.Unmarshal(gameBytes, &game)
			fmt.Printf("%v\n", game)
			gameChannel <- game
			return nil
		})
		return nil
	})
	return nil
}

func (bc *BoltClient) startAGame() {

	total := 100
	for i := 0; i < total; i++ {

		// Generate a key 10000 or larger
		key := strconv.Itoa(10000 + i)

		// Create an instance of our Account struct
		acc := model.Game{
			Id: key,
			LeftSideScore: model.Score{ 0, 0 },
			RightSideScore: model.Score{ 0,0 },
		}

		// Serialize the struct to JSON
		jsonBytes, _ := json.Marshal(acc)

		// Write the data to the AccountBucket
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake accounts...\n", total)
}