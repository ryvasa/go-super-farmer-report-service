package minio

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/env"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/logrus"
)

func NewMinioClient(env *env.Env) *minio.Client {
	endpoint := env.MinIO.EndPoint
	MinIOID := env.MinIO.ID
	MinIOSecret := env.MinIO.Secret
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinIOID, MinIOSecret, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	logrus.Log.Info("coba")
	log.Printf("%#v\n", minioClient) // minioClient is now set up

	// Test connection by listing buckets
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		log.Fatalln("Error connecting to MinIO:", err)
		panic(err)
	}

	logrus.Log.Info("Successfully connected to MinIO")
	logrus.Log.Info("Available buckets:")
	for _, bucket := range buckets {
		logrus.Log.Info("- ", bucket.Name)
	}

	return minioClient
}
