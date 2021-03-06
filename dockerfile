FROM golang AS build

RUN git clone https://github.com/alanwade2001/spa-customer-api.git /app
WORKDIR /app
RUN mkdir -p /go/src/app
RUN cp go.mod go.sum /go/src/app/

WORKDIR /go/src/app/
RUN go mod download

RUN cp -r /app /go/src/
RUN ls -l /go/src/app
RUN go install 

FROM scratch
COPY --from=build /go/bin/spa-customer-api /
ENTRYPOINT ["/spa-customer-api"]