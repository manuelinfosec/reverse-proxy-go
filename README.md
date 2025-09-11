# reverse-proxy-go

A minimal http reverse proxy, written in Go intended to run as a Google Cloud Function.
Configuration can be done by environment variables:

enableHttps - boolean - default: false - When set to true the proxy can be called with https.
