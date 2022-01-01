FROM ubuntu:18.04 as base

ENV GOLANG_VERSION 1.17.5
ENV GIT_USERNAME sakshisharma84
RUN apt-get update
RUN apt-get install -y wget git gcc

# Install wget
RUN apt update && apt install -y build-essential wget
# Install Go
RUN wget https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz
RUN rm -f go${GOLANG_VERSION}.linux-amd64.tar.gz

ENV PATH "$PATH:/usr/local/go/bin"

WORKDIR /usr/local/
RUN git clone https://github.com/sakshisharma84/go-postgres.git

WORKDIR /usr/local/go-postgres
COPY .env /usr/local/go-postgres
RUN go build

COPY go-postgres /bin/ 

EXPOSE 8080

ENTRYPOINT ["/bin/go-postgres"]
