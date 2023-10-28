FROM amazonlinux:2023

ENV GOARCH=amd64 \
	GOROOT=/usr/local/go \
	GOPATH=/go \
	APP_PATH=/app \
	PATH=/usr/local/go/bin:/go/bin:$PATH

WORKDIR $APP_PATH

RUN dnf update -y && \
    dnf install -y wget tar gzip git make gcc java-11-amazon-corretto-headless git-clang-format rsync graphviz unzip
RUN dnf install -y https://rpm.nodesource.com/pub_18.x/el/9/x86_64/nodejs-18.16.0-1nodesource.x86_64.rpm && \
  dnf install -y nodejs-18.16.0

	# Install Go
RUN mkdir /go && \
		wget https://dl.google.com/go/go1.20.7.linux-amd64.tar.gz && \
		tar -xzf go1.20.7.linux-amd64.tar.gz && \
		cp -R go /usr/local/ && \
		rm -r go

    # Install protoc
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.13.0/protoc-3.13.0-linux-x86_64.zip && \
    unzip protoc-3.13.0-linux-x86_64.zip -d protoc && rm protoc-3.13.0-linux-x86_64.zip && \
    cp -r protoc/bin/* /usr/local/bin/ && cp -r protoc/include/* /usr/local/include/ && rm -r protoc


#COPY go.mod ./
#RUN go mod download || true

COPY api ./api
COPY Makefile ./Makefile
RUN make install
