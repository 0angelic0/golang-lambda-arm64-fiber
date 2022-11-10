#!/bin/bash

curl http://127.0.0.1:3000/roles
curl http://127.0.0.1:3000/roles?limit=10
curl http://127.0.0.1:3000/roles?limit=10&tags=eiei
curl http://127.0.0.1:3000/roles?tags=eiei
curl http://127.0.0.1:3000/roles?tags=eiei&tags=eueu

curl -X POST http://127.0.0.1:3000/roles -H 'Content-Type: application/json' -d '{"name": "hey", "tag": "eiei"}'

curl http://127.0.0.1:3000/roles/10
curl http://127.0.0.1:3000/roles/51