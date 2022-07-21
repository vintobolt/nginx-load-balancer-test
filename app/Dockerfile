FROM golang:alpine
EXPOSE 5000
WORKDIR /app
COPY . .
RUN go build -o /build
CMD [ "/build" ]