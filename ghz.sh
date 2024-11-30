

/e/Programs/ghz/ghz --insecure \
  --proto=./proto/greeter.proto \
  --call greeter.Greeter.Hello \
  -c 50 \
  -n 10000 \
  --metadata="{\"content-type\":\"application/json\"}" \
  -d '{"name":"Joe"}' \
 --keepalive=10s \
  0.0.0.0:8081






