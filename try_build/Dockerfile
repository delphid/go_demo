FROM registry.leyantech.com/sre/golang:1.17.9-1
WORKDIR /lain/app
ENV git config --global --add url."git@git.leyantech.com:".insteadOf "https://git.leyantech.com/"
COPY go.mod go.sum ./
RUN \
   go env && \
   go mod download -x
