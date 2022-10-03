package main

import "errors"

const (
	AmazonS3Type = iota
	GoogleCloudStorageType
	AzureCloudStorageType
)

type CloudStorage struct {
	Repository Repository
}

type Repository interface {
	Load(key string, data interface{}) error
	Delete(key string) error
	Store(key string, data interface{}) error
	CloneData(in Repository) error
}

type AmazonS3 struct{}

func (s *AmazonS3) Load(key string, data interface{}) error  { return nil }
func (s *AmazonS3) Delete(key string) error                  { return nil }
func (s *AmazonS3) Store(key string, data interface{}) error { return nil }
func (s *AmazonS3) CloneData(in Repository) error            { return nil }

type GoogleCloudStorage struct{}

func (g *GoogleCloudStorage) Load(key string, data interface{}) error  { return nil }
func (g *GoogleCloudStorage) Delete(key string) error                  { return nil }
func (g *GoogleCloudStorage) Store(key string, data interface{}) error { return nil }
func (g *GoogleCloudStorage) CloneData(in Repository) error            { return nil }

type AzureCloudStorage struct{}

func (g *AzureCloudStorage) Load(key string, data interface{}) error  { return nil }
func (g *AzureCloudStorage) Delete(key string) error                  { return nil }
func (g *AzureCloudStorage) Store(key string, data interface{}) error { return nil }
func (g *AzureCloudStorage) CloneData(in Repository) error            { return nil }

func (c CloudStorage) SetStorageType(storageType int, in Repository) error {
	switch storageType {
	case AmazonS3Type:
		s3 := &AmazonS3{}
		err := s3.CloneData(in)
		if err != nil {
			return err
		}
		c.Repository = s3
	case GoogleCloudStorageType:
		gcp := &GoogleCloudStorage{}
		err := gcp.CloneData(in)
		if err != nil {
			return err
		}
		c.Repository = gcp
	case AzureCloudStorageType:
		azure := &AzureCloudStorage{}
		err := azure.CloneData(in)
		if err != nil {
			return err
		}
		c.Repository = azure
	default:
		return errors.New("invalid storageTyp")
	}
	return nil
}

/*
Стратегия — это поведенческий паттерн, который выносит набор алгоритмов в собственные классы и делает их взаимозаменимыми.
+ Стратегия позволяет подменять схожие и взаимозаменяемые алгоритмы, инкапсулируя их в отдельный объект.
*/
