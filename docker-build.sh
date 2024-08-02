docker build -t print-pdf-builder .

docker create --name extract print-pdf-builder

docker cp extract:/root/print-pdf .

docker rm extract