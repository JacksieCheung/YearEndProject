FROM golang:1.14.9
ENV GO111MODULE "on" 
ENV GOPROXY "https://goproxy.cn"
ADD . $GOPATH/src/github.com/JacksieCheung/YearEndProject
WORKDIR $GOPATH/src/github.com/JacksieCheung/YearEndProject
RUN make
EXPOSE 8899
CMD ["./main", "-c", "service/conf/config.yaml"]
