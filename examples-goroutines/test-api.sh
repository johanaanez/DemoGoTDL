#!/bin/bash

RED="\e[31m"
GREEN="\e[32m"
CYAN="\e[96m"
YELLOW="\e[93m"
BOLD="\033[1m"
DEFAULT="\e[0m"
EXEC='./tp1'


function header() {
  echo -e "$CYAN#####################################$DEFAULT"
  echo -e "$CYAN$1$DEFAULT"
  echo -e "$CYAN#####################################$DEFAULT"
}

function msg_true () {
  echo -e "$GREEN\0PASSED $DEFAULT:\n$1 = $GREEN $2 $DEFAULT"
}

function msg_false () {
  echo -e "$RED\0NOT EQUAL $DEFAULT:\n$1 = $YELLOW $2 $DEFAULT"
  echo -e "EXPECTED:\n$1"
}

function msg() {
  echo -e "  $BOLD$1$DEFAULT"
}

function success_msg() {
  echo -e "  $GREEN$1$DEFAULT"
}

function error_msg() {
  echo -e "  $RED$1$DEFAULT"
}


EXPECTED_OUTPUT_SECUENTIAL=("200")

EXPECTED_OUTPUT_CONCURRENT=("200")

function test_endpoint_secuential(){
  header "TEST ENDPOINT SECUENTIAL"

	commands=("curl -X POST -H \"Content-Type:application/json\" -d '{\"AccountNumber\":1,\"Amount\":2}' http://localhost:8080/api/v1/bank/buy --write-out %{http_code} --output /dev/null")

	for i in "${commands[@]}"
	do

		msg "$EXEC $i"

		ACTUAL_OUTPUT=$($EXEC $i 2>&1)


    if [[ "$EXPECTED_OUTPUT_SECUENTIAL" == "$ACTUAL_OUTPUT" ]]; then
      msg_true "$EXPECTED_OUTPUT_SECUENTIAL" "$ACTUAL_OUTPUT"
    else
      msg_false "$EXPECTED_OUTPUT_SECUENTIAL" "$ACTUAL_OUTPUT"
    fi

	done
}

function test_endpoint_secuential(){
  header "TEST ENDPOINT CONCURRENT"

	commands=("curl -X POST -H \"Content-Type:application/json\" -d '{\"AccountNumber\":1,\"Amount\":2}' http://localhost:8080/api/v1/bank/buyc --write-out %{http_code} --output /dev/null")

	for i in "${commands[@]}"
	do

		msg "$EXEC $i"

		ACTUAL_OUTPUT=$($EXEC $i 2>&1)


    if [[ "$EXPECTED_OUTPUT_CONCURRENT" == "$ACTUAL_OUTPUT" ]]; then
      msg_true "$EXPECTED_OUTPUT_CONCURRENT" "$ACTUAL_OUTPUT"
    else
      msg_false "$EXPECTED_OUTPUT_CONCURRENT" "$ACTUAL_OUTPUT"
    fi

	done
}

test_endpoint_secuential