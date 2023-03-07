package aws

import (
	"bytes"
	"context"
	"net/http"

	"github.com/duyledat197/go-gen-tools/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Client struct {
	session    *session.Session
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
	remover    *s3manager.BatchDelete
	Storage    config.Storage
	Bucket     string
}

func (c *S3Client) Connect(ctx context.Context) error {
	session, err := session.NewSession(&aws.Config{
		Region:           aws.String(c.Storage.Region),
		Credentials:      credentials.NewStaticCredentials(c.Storage.AccessKey, c.Storage.SecretKey, ""),
		Endpoint:         aws.String(c.Storage.Endpoint),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return err
	}
	c.uploader = s3manager.NewUploader(session)
	c.downloader = s3manager.NewDownloader(session)
	c.session = session
	return nil
}

func (c *S3Client) Stop(ctx context.Context) error {
	return nil
}

func (c *S3Client) GetImage(ctx context.Context, key string) ([]byte, error) {
	buf := aws.NewWriteAtBuffer([]byte{})
	if _, err := c.downloader.DownloadWithContext(ctx, buf, &s3.GetObjectInput{
		Bucket: &c.Bucket,
		Key:    aws.String(key),
	}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *S3Client) UploadImage(ctx context.Context, path string, file []byte) (string, error) {
	if _, err := c.uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket:      aws.String(c.Bucket),
		Key:         aws.String(path),
		Body:        bytes.NewReader(file),
		ContentType: aws.String(http.DetectContentType(file)),
	}); err != nil {
		return "", err
	}
	return "", nil
}

func (c *S3Client) DeleteImage(ctx context.Context, key string) error {
	if err := c.remover.Delete(ctx, &s3manager.DeleteObjectsIterator{Objects: []s3manager.BatchDeleteObject{
		{
			Object: &s3.DeleteObjectInput{
				Bucket: aws.String(c.Bucket),
				Key:    aws.String(key),
			},
		},
	}}); err != nil {
		return err
	}
	return nil
}
