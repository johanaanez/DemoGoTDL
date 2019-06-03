#!/bin/bash

function get_data() {
  echo -e "{\"AccountNumber\":1,\"Amount\":2}" 
}

echo -e "secuential endpoint" 
curl -X POST -H 'Content-Type:application/json' --data "$(get_data)" http://localhost:8080/api/v1/bank/buy -v
echo -e "concurrent endpoint"
curl -X POST -H 'Content-Type:application/json' --data "$(get_data)" http://localhost:8080/api/v1/bank/buyc -v