package internal

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
)

type Source interface {
	String() string
	Read() (string, error)
}

func NewSource(uri string) (Source, error) {
	switch Scheme(uri) {
	case "spanner":
		return NewSpannerSource(uri)
	case "file", "":
		return NewFileSource(uri)
	}
	return nil, errors.New("invalid source")
}

type SpannerSource struct {
	uri    string
	client *Client
}

func NewSpannerSource(uri string) (*SpannerSource, error) {
	client, err := NewClient(context.Background(), uri)
	if err != nil {
		return nil, err
	}
	return &SpannerSource{uri: uri, client: client}, nil
}

func (s *SpannerSource) String() string {
	return s.uri
}

func (s *SpannerSource) Read() (string, error) {
	return s.client.GetDatabaseDDL(context.Background())
}

func (s *SpannerSource) Apply(ddls []DDL) error {
	return s.client.ApplyDatabaseDDL(context.Background(), ddls)
}

type FileSource struct {
	uri  string
	path string
}

func NewFileSource(uri string) (*FileSource, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to parse uri: %w", err)
	}
	return &FileSource{uri: uri, path: u.Path}, nil
}

func (s *FileSource) String() string {
	return s.uri
}

func (s *FileSource) Read() (string, error) {
	ddls, err := ioutil.ReadFile(s.path)
	if err != nil {
		return "", err
	}
	return string(ddls), nil
}
