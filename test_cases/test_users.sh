#!/bin/bash

curl http://127.0.0.1:3000/users
curl http://127.0.0.1:3000/users?limit=10
curl http://127.0.0.1:3000/users?limit=10&tags=eiei
curl http://127.0.0.1:3000/users?tags=eiei
curl http://127.0.0.1:3000/users?tags=eiei&tags=eueu

curl -X POST http://127.0.0.1:3000/users -H 'Content-Type: application/json' -d '{"name": "hey", "tag": "eiei"}'

curl http://127.0.0.1:3000/users/10
curl http://127.0.0.1:3000/users/51