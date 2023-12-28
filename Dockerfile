ARG GO_VERSION
ARG OS
FROM golang:${GO_VERSION}-${OS}

ARG GOOS
ARG GOARCH

ENV IS_CONTAINER ${IS_CONTAINER}
ENV CONF_PATH ${CONF_PATH}

ENV DB_NAME ${POSTGRES_DB}
ENV DB_USER ${POSTGRES_USER}
ENV DB_PASSWORD ${POSTGRES_PASSWORD}
ENV PORT ${PORT}

WORKDIR /usr/src/configbay

COPY go.mod go.sum Makefile ./
RUN make install

COPY . ./

RUN make clean-api
RUN make build-api

CMD ["make", "run-api"]

EXPOSE ${PORT}