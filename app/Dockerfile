FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o /build
EXPOSE 5000
CMD [ "/build" ]