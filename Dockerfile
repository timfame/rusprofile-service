FROM golang AS builder

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN export GOBIN=$GOPATH/bin && go get ./cmd/service
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o /go/bin/service ./cmd/service/

FROM alpine:latest

ENV CACHE_EXPIRATION_DURATION_SECONDS=240
ENV CACHE_CLEARING_INTERVAL_SECONDS=120

ENV RUSPROFILE_BASE_URL="https://www.rusprofile.ru"
ENV RUSPROFILE_SEARCH_URL="https://www.rusprofile.ru/search?query="

ENV RUSPROFILE_MAIN_DIV_ID="main"
ENV RUSPROFILE_SEARCH_AMBIGUOUS_RESULT_DIV_CLASS="company-main search-result__main"
ENV RUSPROFILE_SEARCH_EMPTY_RESULT_DIV_CLASS="main-content search-result emptyresult"
ENV RUSPROFILE_SEARCH_COMPANY_ITEM_DIV_CLASS="company-item"
ENV RUSPROFILE_SEARCH_COMPANY_ITEM_TITLE_DIV_CLASS="company-item__title"

ENV RUSPROFILE_COMPANY_NAME_DIV_CLASS="company-name"
ENV RUSPROFILE_COMPANY_KPP_SPAN_ID="clip_kpp"
ENV RUSPROFILE_COMPANY_INFO_TITLE_CLASS="company-info__title"
ENV RUSPROFILE_COMPANY_INFO_TITLE_DIRECTOR_TEXT="Руководитель"
ENV RUSPROFILE_COMPANY_INFO_TEXT_CLASS="company-info__text"

ENV GRPC_PORT=8088
ENV GRPC_GATEWAY_PORT=8099

RUN apk add --no-cache bash
COPY --from=builder /go/bin/service /app/

WORKDIR /app
EXPOSE 8088 8099
CMD ["./service"]