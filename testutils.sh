#!/usr/bin/env bash
HOST="http://localhost:8000"

function getToken() {
	res=$(curl -X POST \
		-H "Content-Type: json/application" \
		"$HOST/login" \
		-d '{"user":"admin","pass":"admin","email":"example@gmail.com"}')
	TOKEN=$(echo "$res" | tr -d '"{}' | sed 's/bearer://g')
	echo "$TOKEN"
}

function fakeOrders() {
	for i in $(seq 10); do
		curl -X POST \
			-H "Content-Type: json/application" \
			-H "Authorization: Bearer ${TOKEN}" \
			"$HOST/orders" \
			-d '{"description":"product '"${i}"'"}'
	done
}

function completeOrders() {
        for i in $(seq 10); do
                curl -X POST \
                        -H "Content-Type: json/application" \
                        -H "Authorization: Bearer ${TOKEN}" \
                        "$HOST/orders/first/done"
        done
}

function calcelOrders() {
	for i in $(seq 1 9); do
		curl -X POST \
			-H "Content-Type: json/application" \
			-H "Authorization: Bearer ${TOKEN}" \
			"$HOST/orders/${i}/cancel"
	done
}

function fakeRefunds() {
	 for i in $(seq 10); do
		curl -X POST \
			-H "Content-Type: json/application" \
			-H "Authorization: Bearer ${TOKEN}" \
			"$HOST/refunds/${i}"
	done
}

function completeRefunds() {
	 for i in $(seq 10); do
		curl -X POST \
			-H "Content-Type: json/application" \
			-H "Authorization: Bearer ${TOKEN}" \
			"$HOST/refunds/first/done"
	done
}
