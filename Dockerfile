FROM golang:1.17.0-alpine3.14
# アップデートとgitのインストール
ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
# appディレクトリの作成
# RUN mkdir /go/src/app
# ワーキングディレクトリの設定
# WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行

# COPY . /go/src/app
COPY go.mod .
RUN go mod download

COPY ./main.go  ./