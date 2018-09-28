FROM alpine:latest
MAINTAINER Søren Mathiasen <sorenm@mymessages.dk>

ADD migrations/ migrations

# UI stuff
ADD public/dist /

# Add binary
ADD seneferu /seneferu
ENTRYPOINT ["/seneferu"]
