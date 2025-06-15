FROM golang:1.24.4 AS builder

WORKDIR /docs.piquel.fr

# Setup env
RUN export PATH="$PATH:$(go env GOPATH)/bin"

# Setup go dependencies
COPY go.mod .
RUN go mod download

# Copy everything else
COPY . .

# Build the binary
RUN CGO_ENABLED=0 go build -o ./bin/main ./main.go

# Now for run env
FROM alpine:latest

WORKDIR /docs.piquel.fr

RUN apk add --no-cache git

# Copy static files and configuration
COPY --from=builder /docs.piquel.fr/bin/main .

CMD [ "./main" ]
