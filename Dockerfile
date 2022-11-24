FROM golang:1.19.2 AS BUILD_IMAGE
WORKDIR /app
COPY . ./
RUN go build -o /server

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=BUILD_IMAGE /server /server
EXPOSE 9098
USER nonroot:nonroot

ENTRYPOINT [ "/server" ]