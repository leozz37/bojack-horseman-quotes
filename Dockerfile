FROM golang:alpine

COPY . .

RUN go mod download
RUN go build -o bojack-horseman

CMD [ "./bojack-horseman" ]