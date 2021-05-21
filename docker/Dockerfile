FROM golang:1.11.6

RUN apt-get update -y && \
    apt-get install lsb-release -y

RUN export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)" && \
    echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
    apt-get update -y && \
    apt-get install google-cloud-sdk google-cloud-sdk-app-engine-python google-cloud-sdk-app-engine-go google-cloud-sdk-datastore-emulator -y

RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.15.0

WORKDIR /go/src/github.com/mercari/Week6/Week6
