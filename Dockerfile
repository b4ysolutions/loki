FROM golang
WORKDIR /app/src/github.com/b4ysolutions/loki
ENV GOPATH=/app
COPY . /app/src/github.com/b4ysolutions/loki
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/JonathanMH/goClacks/echo
RUN go get -u github.com/labstack/echo
RUN go get -u github.com/labstack/echo/middleware
RUN go get -u github.com/labstack/gommon/log
RUN go build -o main .
CMD [ "./main" ]