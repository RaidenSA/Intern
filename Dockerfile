# syntax=docker/dockerfile:1
FROM golang:1.17-alpine
WORKDIR /app
ADD https://github.com/RaidenSA/Intern.git ./
RUN go mod download
RUN go build ./cmd -o /dockerShortener
EXPOSE 8080
EXPOSE 8088
CMD [ "/dockerShortener" ]

