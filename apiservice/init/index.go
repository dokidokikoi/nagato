package inittask

import (
	"context"
	"nagato/apiservice/internal/db/data"
	"nagato/apiservice/internal/db/data/es"
	"nagato/apiservice/internal/model"
)

func initIndex() {
	store, err := data.GetStoreDBFactory()
	if err != nil {
		panic(err)
	}

	userList, err := store.Users().List(context.TODO(), &model.User{}, nil)
	if err != nil {
		panic(nil)
	}

	for i := range userList {
		err := store.Blanks().CreateIndices(userList[i].ID, es.BlankIndex)
		if err != nil {
			panic(err)
		}

		err = store.Matters().CreateIndices(userList[i].ID, es.ResourceIndex)
		if err != nil {
			panic(err)
		}
	}
}
