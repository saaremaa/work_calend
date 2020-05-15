# Start from golang base image
FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="krb@klgd.su"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

ARG UUID=1000

ARG UGID=1000 

# Create a group and user
RUN addgroup -g ${UGID} -S docker && adduser --uid ${UUID} -S docker -G docker

# Tell docker that all future commands should run as the appuser user
USER docker

WORKDIR /home/docker

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download


# Build the Go app
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN go build -ldflags "-s -w" -o work-calend .

# Start a new stage from scratch
FROM alpine:latest

#change timezone on Europe/Kaliningrad
RUN apk add --no-cache tzdata && apk add --no-cache curl

ENV TZ=Europe/Kaliningrad

RUN date

RUN apk --no-cache add ca-certificates

# Create a group and user

ARG UUID=1000

ARG UGID=1000 

RUN addgroup -g ${UGID} -S docker && adduser -u ${UUID} -S docker -G docker 

# Tell docker that all future commands should run as the docker user
USER docker

ENV PATH="/home/docker:${PATH}"

WORKDIR /home/docker

RUN mkdir ./data

# Copy the Pre-built binary file from the previous stage.
# Observe we also copied the .env file
# COPY --from=builder /home/docker/data/data.csv ./data/

COPY --from=builder /home/docker/work-calend .

COPY --from=builder /home/docker/.env .

EXPOSE 2081

HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD curl -f http://127.0.0.1:2081/api/v1/ping || exit 1

ENTRYPOINT ./work-calend

#Command to run the executable
#CMD ["./main"]
