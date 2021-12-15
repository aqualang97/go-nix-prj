package filesystem

import (
	"encoding/json"
	"io"
	"my-uuid/repositories/models"
	"os"
)

type UserFileRepository struct {
}

func (ufr UserFileRepository) GetByEmail(_ string) (user models.User) {
	data := []byte{}
	file, err := os.Open("./datastore/files/user_1.json")
	if err != nil {
		return models.User{}
	}
	defer file.Close()
	for {
		var chunk []byte
		_, err := file.Read(chunk)
		if err == io.EOF {
			//
			break
			//
		}
		data = append(data, chunk...)
	}

	json.Unmarshal(data, &user)
	return user
}
