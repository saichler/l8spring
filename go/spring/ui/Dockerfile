FROM saichler/builder:latest AS build

COPY ./main/main.go /home/src/github.com/saichler/build/
RUN go mod init
RUN GOPROXY=direct GOPRIVATE=github.com go mod tidy
RUN go build -o springweb

FROM saichler/erp-security:latest AS final
COPY --from=build /home/src/github.com/saichler/build/springweb /home/run/springweb
COPY ./web /home/run/web

ENTRYPOINT ["/home/run/springweb"]
