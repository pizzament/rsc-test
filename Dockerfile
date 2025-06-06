FROM golang:1.23.4-alpine as builder

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /rsc ./cmd/server/main.go

FROM scratch
COPY --from=builder /rsc /bin/rsc
COPY configs/values_local_docker.yaml /bin/config/values_local_docker.yaml

ENV CONFIG_FILE=/bin/config/values_local_docker.yaml

ENTRYPOINT ["/bin/rsc"]