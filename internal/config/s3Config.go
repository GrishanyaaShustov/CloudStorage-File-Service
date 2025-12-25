package config

type S3Config struct {
	AccessKeyID     string `env:"S3_ACCESS_KEY_ID,required"`
	SecretAccessKey string `env:"S3_SECRET_KEY,required"`

	Region   string `env:"S3_REGION_NAME,required"`
	Endpoint string `env:"S3_ENDPOINT_NAME,required"`

	Bucket string `env:"S3_BUCKET_NAME,required"`
}
