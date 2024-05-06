# LocalStack S3 Example with Terraform and Go

## 1. Start LocalStack
```bash
docker compose up
```

## 2. Provision S3 Bucket
```bash
docker build terraform/ -t localstack-terraform
docker run --rm -v ./terraform:/app localstack-terraform tflocal init
docker run --net=host --rm -v ./terraform:/app localstack-terraform tflocal apply -auto-approve
```

3. Upload/Download a file to S3
```bash
docker build app/ -t localstack-app
docker run --net=host --rm -v ./app:/usr/src/app localstack-app -op upload my-bucket example.txt
rm app/example.txt
docker run --net=host --rm -v ./app:/usr/src/app localstack-app -op download my-bucket example.txt
```
```
