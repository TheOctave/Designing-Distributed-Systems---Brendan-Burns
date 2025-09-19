vcl 4.1;

backend default {
    .host = "dictionary-server-service.default.svc.cluster.local";
    .port = "80";
}