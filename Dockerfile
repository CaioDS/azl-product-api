FROM golang:1.20
RUN mkdir azl-votes-api
WORKDIR /azl-votes-api
COPY . .
RUN go build -o cmd/workload cmd/api/main.go
RUN ["chmod", "+x", "/azl-votes-api/cmd/workload/"]
CMD [ "./cmd/workload/main" ]