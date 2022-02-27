FROM golang:latest

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go install github.com/spf13/cobra/cobra@latest && \
    go install github.com/golang/mock/mockgen@latest

RUN apt-get update && apt-get install sqlite3 -y

CMD ["tail", "-f", "/dev/null"]