# agetaro

### Terraform

```sh
export PROJECT_ID=geek-agetaro
gcloud auth login
export ACCOUNT_ID=$(gcloud beta billing accounts list | grep True | cut -d ' ' -f1)
gcloud projects create $PROJECT_ID
gcloud beta billing projects link $PROJECT_ID --billing-account=$ACCOUNT_ID

# enable apis
gcloud services enable \
    cloudapis.googleapis.com \
    cloudresourcemanager.googleapis.com \
    container.googleapis.com \
    containerregistry.googleapis.com \
    iam.googleapis.com \
    redis.googleapis.com \
    servicenetworking.googleapis.com \
    sqladmin.googleapis.com \
    run.googleapis.com \
    vpcaccess.googleapis.com

# Create a service account for terraform
gcloud iam service-accounts create terraform \
    --description="Terraform Service Account" \
    --display-name="Terraform"

gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member serviceAccount:terraform@$PROJECT_ID.iam.gserviceaccount.com \
  --role roles/owner

gcloud iam service-accounts keys create CREDENTIALS_FILE.json --iam-account=terraform@$PROJECT_ID.iam.gserviceaccount.com --project $PROJECT_ID

# Create Bucket
gsutil mb gs://terraform-tfstate-agetaro

# Deploy using Terraform.
terraform init
terraform plan
terraform apply
```

#### Cloud Runで用いるためのコンテナを更新するためのコマンド

```sh
gcloud builds submit --tag gcr.io/agetaro-2517d/agetaro-kun
```

### CloudSQLの初期化

```sh
# 後で更新する
```