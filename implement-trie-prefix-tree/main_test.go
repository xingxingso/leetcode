package implement_trie_prefix_tree

import "testing"

type TrieInterface interface {
	Insert(word string)
	Search(word string) bool
	StartsWith(prefix string) bool
}

func TestTrie(t *testing.T) {
	tree := Constructor()
	testEx1(&tree, t)

	tree2 := Constructor()
	testEx2(&tree2, t)
}

func testEx1(s TrieInterface, t *testing.T) {
	s.Insert("apple")

	if got, want := s.Search("apple"), true; got != want {
		t.Errorf("s.Search() = %v, want %v", got, want)
	}

	if got, want := s.Search("app"), false; got != want {
		t.Errorf("s.Search() = %v, want %v", got, want)
	}

	if got, want := s.StartsWith("app"), true; got != want {
		t.Errorf("s.StartsWith() = %v, want %v", got, want)
	}

	s.Insert("app")

	if got, want := s.Search("app"), true; got != want {
		t.Errorf("s.Search() = %v, want %v", got, want)
	}
}

func testEx2(s TrieInterface, t *testing.T) {
	s.Insert("app")
	s.Insert("apple")
	s.Insert("beer")
	s.Insert("add")
	s.Insert("jam")
	s.Insert("rental")

	if got, want := s.Search("apps"), false; got != want {
		t.Errorf("s.Search() = %v, want %v", got, want)
	}

	if got, want := s.Search("app"), true; got != want {
		t.Errorf("s.Search() = %v, want %v", got, want)
	}

	if got, want := s.Search("ad"), false; got != want {
		t.Errorf("s.Search() = %v, want %v", got, want)
	}

	if got, want := s.Search("applepie"), false; got != want {
		t.Errorf("s.Search() = %v, want %v", got, want)
	}

	if got, want := s.Search("rest"), false; got != want {
		t.Errorf("s.Search() = %v, want %v", got, want)
	}

	if got, want := s.Search("jan"), false; got != want {
		t.Errorf("s.Search() = %v, want %v", got, want)
	}

	if got, want := s.Search("rent"), false; got != want {
		t.Errorf("s.Search() = %v, want %v", got, want)
	}

	if got, want := s.Search("beer"), true; got != want {
		t.Errorf("s.Search() = %v, want %v", got, want)
	}

	if got, want := s.Search("jam"), true; got != want {
		t.Errorf("s.Search() = %v, want %v", got, want)
	}

	if got, want := s.StartsWith("apps"), false; got != want {
		t.Errorf("s.StartsWith() = %v, want %v", got, want)
	}

	if got, want := s.StartsWith("app"), true; got != want {
		t.Errorf("s.StartsWith() = %v, want %v", got, want)
	}

	if got, want := s.StartsWith("ad"), true; got != want {
		t.Errorf("s.StartsWith() = %v, want %v", got, want)
	}

	if got, want := s.StartsWith("applepie"), false; got != want {
		t.Errorf("s.StartsWith() = %v, want %v", got, want)
	}

	if got, want := s.StartsWith("rest"), false; got != want {
		t.Errorf("s.StartsWith() = %v, want %v", got, want)
	}

	if got, want := s.StartsWith("jan"), false; got != want {
		t.Errorf("s.StartsWith() = %v, want %v", got, want)
	}

	if got, want := s.StartsWith("rent"), true; got != want {
		t.Errorf("s.StartsWith() = %v, want %v", got, want)
	}

	if got, want := s.StartsWith("beer"), true; got != want {
		t.Errorf("s.StartsWith() = %v, want %v", got, want)
	}

	if got, want := s.StartsWith("jam"), true; got != want {
		t.Errorf("s.StartsWith() = %v, want %v", got, want)
	}
}
