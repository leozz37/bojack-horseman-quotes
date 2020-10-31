FROM golang:alpine
ENV CGO_ENABLED=0

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o bojack-horseman .

EXPOSE $PORT

CMD [ "./bojack-horseman" ]