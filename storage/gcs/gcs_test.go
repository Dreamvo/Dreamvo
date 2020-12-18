package gcs

import (
	"bytes"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dreamvo/gilfoyle/storage"
)

func Test(t *testing.T) {
	ctx := context.Background()
	s, err := NewMockedStorage(ctx, "test-bucket")
	if err != nil {
		t.Fatal(err)
	}

	if _, err = s.Stat(ctx, "doesnotexist"); !errors.Is(err, storage.ErrNotExist) {
		t.Errorf("expected not exists, got %v", err)
	}

	before := time.Now()
	if err := s.Save(ctx, bytes.NewBufferString("hello"), "world"); err != nil {
		t.Fatal(err)
	}
	now := time.Now()

	stat, err := s.Stat(ctx, "world")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if stat.Size != 5 {
		t.Errorf("expected size to be %d, got %d", 5, stat.Size)
	}
	if stat.ModifiedTime.Before(before) {
		t.Errorf("expected modtime to be after %v, got %v", before, stat.ModifiedTime)
	}
	if stat.ModifiedTime.After(now) {
		t.Errorf("expected modtime to be before %v, got %v", now, stat.ModifiedTime)
	}
}
