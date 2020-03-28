 go mod vendor && \
 gcloud functions deploy HandlerEDX --runtime go113 --memory=128MB --trigger-http