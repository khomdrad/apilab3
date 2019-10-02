FROM golang:alpine
LABEL maintainer="Khomdrad Boontae<khomdradboon@gmail.com>"
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
CMD ["/app/main"]