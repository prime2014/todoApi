FROM golang:1.24.3-bookworm

RUN apt-get update -y && apt-get install -y netcat-openbsd && rm -rf /var/lib/apt/lists/*


# Create a non-root user with home directory
RUN addgroup --system prime && \
    adduser --system --ingroup prime --home /todoApi prime

WORKDIR /todoApi

RUN chown -R prime:prime /todoApi/

USER prime

COPY --chown=prime:prime ./go.mod ./go.sum ./

RUN go mod download 

COPY --chown=prime:prime . .

RUN chmod +x /todoApi/entrypoint.sh \
    && go build -o todoApi .

EXPOSE 8080

ENTRYPOINT ["/todoApi/entrypoint.sh"]

CMD ["./todoApi"]
