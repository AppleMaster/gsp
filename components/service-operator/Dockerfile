# Build options

# setup go build env
FROM golang:1.12.5 as builder
WORKDIR /workspace

# install kubebuilder and bundled control-plane for tests
ENV KUBEBUILDER_VERSION=2.0.0
RUN wget -nv https://github.com/kubernetes-sigs/kubebuilder/releases/download/v${KUBEBUILDER_VERSION}/kubebuilder_${KUBEBUILDER_VERSION}_linux_amd64.tar.gz \
	&& tar xvf kubebuilder_${KUBEBUILDER_VERSION}_linux_amd64.tar.gz \
	&& mv kubebuilder_${KUBEBUILDER_VERSION}_linux_amd64 /usr/local/kubebuilder
ENV PATH=$PATH:/usr/local/kubebuilder/bin

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile
RUN go mod download \
	&& make controller-gen \
	&& go install github.com/onsi/ginkgo/ginkgo

# Copy the go source
COPY . .

# Build and test environment
ENV CGO_ENABLED="0"
ENV GOOS="linux"
ENV GOARCH="amd64"
ENV GO111MODULE="on"
ENV GOPROXY=https://proxy.golang.org,direct

# Build args for integration testing
ARG AWS_INTEGRATION="false"
ENV AWS_INTEGRATION="${AWS_INTEGRATION}"
ARG AWS_ACCESS_KEY_ID=""
ENV AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}"
ARG AWS_SECRET_ACCESS_KEY=""
ENV AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}"
ARG AWS_SESSION_TOKEN=""
ENV AWS_SESSION_TOKEN="${AWS_SESSION_TOKEN}"
ARG AWS_RDS_SECURITY_GROUP_ID
ENV AWS_RDS_SECURITY_GROUP_ID="${AWS_RDS_SECURITY_GROUP_ID}"
ARG AWS_RDS_SUBNET_GROUP_NAME
ENV AWS_RDS_SUBNET_GROUP_NAME="${AWS_RDS_SUBNET_GROUP_NAME}"
ARG AWS_PRINCIPAL_SERVER_ROLE_ARN
ENV AWS_PRINCIPAL_SERVER_ROLE_ARN="${AWS_PRINCIPAL_SERVER_ROLE_ARN}"
ARG AWS_PRINCIPAL_PERMISSIONS_BOUNDARY_ARN
ENV AWS_PRINCIPAL_PERMISSIONS_BOUNDARY_ARN="${AWS_PRINCIPAL_PERMISSIONS_BOUNDARY_ARN}"

# run tests
RUN make test

# build binary
RUN go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:latest
WORKDIR /
COPY --from=builder /workspace/manager .
ENTRYPOINT ["/manager"]
