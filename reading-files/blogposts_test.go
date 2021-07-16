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
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: Post 2")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	got := posts[0]
	want := blogposts.Post{Title: "Post 1"}

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %v , but got %v posts", want, got)
	}
}
