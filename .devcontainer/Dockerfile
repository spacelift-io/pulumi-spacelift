ARG UBUNTU_VERSION=22.04
ARG TARGETPLATFORM=linux/amd64
FROM --platform=$TARGETPLATFORM ubuntu:${UBUNTU_VERSION} 

# Update apt-get and install various needed utilities
RUN apt-get update && \
  apt-get install -y curl && \
  apt-get install -y wget && \
  apt-get install -y xz-utils && \
  apt-get install -y make && \
  apt-get install -y gcc && \
  apt-get install -y git

# Install bridged provider prerequisites
# See README.md

# Install go 
ARG GO_VERSION=1.21.3
RUN rm -rf /usr/local/go && \
  wget -O ${GO_VERSION}.tar.gz https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
  tar -C /usr/local -xzf ${GO_VERSION}.tar.gz && \
  rm ${GO_VERSION}.tar.gz 

ENV GOPATH=/root/go
ENV PATH=$PATH:/usr/local/go/bin

# Install go linter
RUN mkdir -p $GOPATH/bin && \
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOPATH/bin v1.46.2

ENV PATH=$PATH:$GOPATH/bin

# Install pulumictl
ARG PULUMICTL_VERSION=v0.0.45
RUN rm -rf /usr/local/bin/pulumictl && \
  wget -O pulumictl.${PULUMICTL_VERSION}.tar.gz https://github.com/pulumi/pulumictl/releases/download/${PULUMICTL_VERSION}/pulumictl-${PULUMICTL_VERSION}-linux-amd64.tar.gz && \
  tar -C /usr/local/bin -xzf pulumictl.${PULUMICTL_VERSION}.tar.gz

# Install nodejs
ARG NODEJS_VERSION=v18.20.4
ARG NODEJS_PKG=node-${NODEJS_VERSION}-linux-x64
ARG NODEJS_TARBALL=${NODEJS_PKG}.tar.xz
RUN rm -rf /usr/local/node && \
  wget -O ${NODEJS_TARBALL} https://nodejs.org/dist/${NODEJS_VERSION}/${NODEJS_TARBALL} && \
  tar -C /usr/local -xf ${NODEJS_TARBALL} && \
  mv /usr/local/${NODEJS_PKG} /usr/local/node

ENV PATH=$PATH:/usr/local/node/bin

# Install yarn
RUN npm install --global yarn

# Install python and related items
RUN apt-get install -y python3 && \
  apt-get install -y python3-setuptools

# Install .NET
RUN wget https://packages.microsoft.com/config/ubuntu/22.04/packages-microsoft-prod.deb -O packages-microsoft-prod.deb && \
  dpkg -i packages-microsoft-prod.deb && \
  rm packages-microsoft-prod.deb && \
  apt-get update && \
  apt-get install -y dotnet-sdk-6.0

# Install Pulumi
RUN curl -fsSL https://get.pulumi.com | sh
ENV PATH=$PATH:/root/.pulumi/bin
