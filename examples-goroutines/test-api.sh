#!/bin/bash

function get_data() {
  echo -e "{\"AccountNumber\":1,\"Amount\":2}" 
}

function get_data_error() {
  echo -e "{\"AccountNumber\":-10,\"Amount\":2}" 
}

echo -e "secuential endpoint - status OK" 
time curl -X POST -H 'Content-Type:application/json' --data "$(get_data)" http://localhost:8080/api/v1/bank/buy -v
echo -e "secuential endpoint - status 500" 
time curl -X POST -H 'Content-Type:application/json' --data "$(get_data_error)" http://localhost:8080/api/v1/bank/buy -v
echo -e "concurrent endpoint - status OK"
time curl -X POST -H 'Content-Type:application/json' --data "$(get_data)" http://localhost:8080/api/v1/bank/buyc -v
echo -e "concurrent endpoint - status 500"
time curl -X POST -H 'Content-Type:application/json' --data "$(get_data_error)" http://localhost:8080/api/v1/bank/buyc -v