FROM saichler/builder:latest AS build

COPY main.go /home/src/github.com/saichler/build/
RUN go mod init
RUN GOPROXY=direct GOPRIVATE=github.com go mod tidy
RUN go build -o spring

FROM saichler/erp-postgres:latest AS final

COPY --from=build /home/src/github.com/saichler/build/spring /home/run/spring

ENTRYPOINT ["/home/run/spring"]
