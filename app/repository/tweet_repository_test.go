package repository

import (
	"reflect"
	"testing"

	"github.com/mercari/Week6/Week6/NamikoToriyama/model"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func TestList(t *testing.T) {
	opt := aetest.Options{StronglyConsistentDatastore: true} //データストアに即反映
	instance, err := aetest.NewInstance(&opt)
	if err != nil {
		t.Fatalf("Failed to create aetest instance: %v", err)
	}
	defer instance.Close()

	// Contextが必要なので、ダミーのhttp.Request
	req, err := instance.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	c := appengine.NewContext(req)

	// 更新される購読者エンティティを用意しておく
	want := []*model.Tweet{
		{
			ID:       "001",
			Username: "foo",
			Tweet:    "Good morning",
		},
		{
			ID:       "002",
			Username: "hoge",
			Tweet:    "Good afternoon",
		},
		{
			ID:       "003",
			Username: "bar",
			Tweet:    "Good evening",
		},
	}
	key := datastore.NewKey(c, "Tweet", want[0].ID, 0, nil)
	if _, err = datastore.Put(c, key, &want[0]); err != nil {
		t.Fatal(err)
	}
	key = datastore.NewKey(c, "Tweet", want[1].ID, 0, nil)
	if _, err = datastore.Put(c, key, &want[1]); err != nil {
		t.Fatal(err)
	}
	key = datastore.NewKey(c, "Tweet", want[2].ID, 0, nil)
	if _, err = datastore.Put(c, key, &want[2]); err != nil {
		t.Fatal(err)
	}

	target := &tweetRepository{}
	// Execute test.
	got, err := target.List(c)
	if err != nil {
		t.Fatal(err)
	}

	// データストアにグループ名が書き込まれていること
	for i := 0; i < 3; i++ {
		if !reflect.DeepEqual(want[i], got[i]) {
			t.Fatalf("want = %v, got = %v", want[i], got[i])
		}
	}
}
func TestPut(t *testing.T) {
	opt := aetest.Options{StronglyConsistentDatastore: true} //データストアに即反映
	instance, err := aetest.NewInstance(&opt)
	if err != nil {
		t.Fatalf("Failed to create aetest instance: %v", err)
	}
	defer instance.Close()

	// Contextが必要なので、ダミーのhttp.Request
	req, err := instance.NewRequest("POST", "/tweet", nil)
	if err != nil {
		t.Fatal(err)
	}
	c := appengine.NewContext(req)

	// 更新される購読者エンティティを用意しておく
	entity := &model.Tweet{
		ID:       "123456",
		Username: "Tom",
		Tweet:    "Hello",
	}

	target := &tweetRepository{}
	// Execute test.
	if _, err := target.Put(c, entity); err != nil {
		t.Fatal(err)
	}

	// データストアにグループ名が書き込まれていること
	actualEntity := new(model.Tweet)
	key := datastore.NewKey(c, "Tweet", entity.ID, 0, nil)
	if err = datastore.Get(c, key, actualEntity); err != nil {
		t.Fatal(err)
	}
	if actualEntity.ID != entity.ID {
		t.Fatalf("want = %v, got = %v", entity.ID, actualEntity.ID)
	}
	if actualEntity.Username != entity.Username {
		t.Fatalf("want = %v, got = %v", entity.Username, actualEntity.Username)
	}
	if actualEntity.Tweet != entity.Tweet {
		t.Fatalf("want = %v, got = %v", entity.Tweet, actualEntity.Tweet)
	}
}

func TestGet(t *testing.T) {
	opt := aetest.Options{StronglyConsistentDatastore: true} //データストアに即反映
	instance, err := aetest.NewInstance(&opt)
	if err != nil {
		t.Fatalf("Failed to create aetest instance: %v", err)
	}
	defer instance.Close()

	// Contextが必要なので、ダミーのhttp.Request
	req, err := instance.NewRequest("GET", "/tweet", nil)
	if err != nil {
		t.Fatal(err)
	}
	c := appengine.NewContext(req)

	// 更新される購読者エンティティを用意しておく
	entity := &model.Tweet{
		ID:       "123456",
		Username: "Tom",
		Tweet:    "Hello",
	}
	key := datastore.NewKey(c, "Tweet", entity.ID, 0, nil)
	if _, err = datastore.Put(c, key, &entity); err != nil {
		t.Fatal(err)
	}

	target := &tweetRepository{}
	// Execute test.
	actualEntity, err := target.Get(c, entity.ID)
	if err != nil {
		t.Fatal(err)
	}

	// データストアにグループ名が書き込まれていること
	if actualEntity.ID != entity.ID {
		t.Fatalf("want = %v, got = %v", entity.ID, actualEntity.ID)
	}
	if actualEntity.Username != entity.Username {
		t.Fatalf("want = %v, got = %v", entity.Username, actualEntity.Username)
	}
	if actualEntity.Tweet != entity.Tweet {
		t.Fatalf("want = %v, got = %v", entity.Tweet, actualEntity.Tweet)
	}
}

func TestDelete(t *testing.T) {
	opt := aetest.Options{StronglyConsistentDatastore: true} //データストアに即反映
	instance, err := aetest.NewInstance(&opt)
	if err != nil {
		t.Fatalf("Failed to create aetest instance: %v", err)
	}
	defer instance.Close()

	// Contextが必要なので、ダミーのhttp.Request
	req, err := instance.NewRequest("DELETE", "/tweet", nil)
	if err != nil {
		t.Fatal(err)
	}
	c := appengine.NewContext(req)

	// 更新される購読者エンティティを用意しておく
	entity := &model.Tweet{
		ID:       "123456",
		Username: "Tom",
		Tweet:    "Hello",
	}
	key := datastore.NewKey(c, "Tweet", entity.ID, 0, nil)
	if _, err = datastore.Put(c, key, &entity); err != nil {
		t.Fatal(err)
	}

	target := &tweetRepository{}
	// Execute test.
	if err := target.Delete(c, entity.ID); err != nil {
		t.Fatal(err)
	}

	tweets := []*model.Tweet{}
	if _, err := datastore.NewQuery("Tweet").Order("-registerDate").GetAll(c, &tweets); err != nil {
		t.Fatal(err)
	}
	if len(tweets) != 0 {
		t.Fatalf("Failed to delete entity.")
	}

}
