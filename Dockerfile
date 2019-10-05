# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /product-search
RUN cd /product-search && go build /product-search/cmd/api.go

# final stage
FROM alpine
WORKDIR /app/cmd
COPY --from=build-env /product-search/api /app/cmd/
COPY --from=build-env /product-search/config /app/config
COPY --from=build-env /product-search/templates /app/templates
EXPOSE 8080
ENTRYPOINT ./api