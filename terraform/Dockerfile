FROM python:3.11-slim

RUN apt-get update && \
    apt-get install -y curl unzip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV TERRAFORM_VERSION=1.7.5
RUN curl -LO https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_arm64.zip && \
    unzip terraform_${TERRAFORM_VERSION}_linux_arm64.zip -d /usr/local/bin/ && \
    rm terraform_${TERRAFORM_VERSION}_linux_arm64.zip

RUN pip3 install terraform-local

WORKDIR /app
