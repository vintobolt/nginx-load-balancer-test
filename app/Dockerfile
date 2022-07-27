FROM golang:alpine as BUILDER
WORKDIR /app
COPY . .
RUN go build -o /build
FROM alpine:latest
WORKDIR /app
COPY --from=BUILDER /build /app
EXPOSE 5000
CMD [ "/build" ]