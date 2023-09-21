FROM golang:1.21

WORKDIR /usr/src/configbay

COPY go.mod go.sum Makefile ./
RUN make install

COPY . ./
ARG GOOS
ARG GOARCH
ARG IS_CONTAINER

ENV GOOS $GOOS
ENV GOARCH $GOARCH
ENV IS_CONTAINER $IS_CONTAINER

RUN make clean-api
RUN make build-api

CMD ["make", "run-api"]