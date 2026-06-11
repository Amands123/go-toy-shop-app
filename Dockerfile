FROM golang:1.24 as base

WORKDIR /app

 # kind of requirements.txt in python
COPY go.mod .

# kind of pip install -r requirements.txt
RUN go mod download

COPY . .

RUN go build -o main .

# Final stage -Distroless image for smaller size and better security

FROM gcr.io/distroless/base

COPY --from=base /app/main .

COPY --from=base /app/static ./static

COPY --from=base /app/templates ./templates

EXPOSE 8080

CMD ["/main"]