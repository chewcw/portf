FROM golang AS builder

WORKDIR /app
COPY . .
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
RUN go build -a -o portf

FROM scratch
COPY --from=builder /app/portf .
ENTRYPOINT ["/portf"]
