ARG GO_VERSION
FROM golang:${GO_VERSION}

WORKDIR /usr/src/configbay

COPY go.mod go.sum Makefile ./
RUN make install

COPY . ./

ARG GOOS
ARG GOARCH

ENV CONF_PATH ${CONF_PATH}
ENV DB_NAME ${POSTGRES_DB}
ENV DB_USER ${POSTGRES_USER}
ENV DB_PASSWORD ${POSTGRES_PASSWORD}
ENV IS_CONTAINER ${IS_CONTAINER}

RUN make clean-api
RUN make build-api

CMD ["make", "run-api"]