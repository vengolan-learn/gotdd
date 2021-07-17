package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/vengolan-learn/gotdd/reading-files"
)

type StubFailingsFS struct{}

func (r *StubFailingsFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Oh no... i always faile")
}

func TestFailure(t *testing.T) {
	fs := &StubFailingsFS{}
	_, err := blogposts.NewPostsFromFS(fs)
	if err == nil {
		t.Fatal("expected an error here....")
	}
}

func TestNewBlogPosts(t *testing.T) {

	const (
		firstBody = `Title: Post 1 
Description: Description 1
Tags: tdd, go
---
Hello 
World!`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	got := posts[0]
	want := blogposts.Post{Title: "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World!`,
	}

	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, got, want)
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v-, but got %v-", want, got)
	}
}
