FROM golang

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o movie-suggester cmd/main.go

CMD ./movie-suggester
EXPOSE 3001