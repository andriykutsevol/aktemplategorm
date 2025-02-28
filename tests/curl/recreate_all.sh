#!/bin/bash

#-------------------------------------------------------
# Create PSP



curl -X 'POST' \
  'http://localhost:8080/pspfee' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "psp_country_code": "CM",
    "psp_code": "CG-MTN-MTN_1",
    "psp_short_name": "CG-MTN-MTN_1 short name"
}'


curl -X 'POST' \
  'http://localhost:8080/pspfee' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "psp_country_code": "CM",
    "psp_code": "CG-MTN-MTN_2",
    "psp_short_name": "CG-MTN-MTN_2 short name"
}'


curl -X 'POST' \
  'http://localhost:8080/pspfee' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "psp_country_code": "CM",
    "psp_code": "CG-MTN-MTN_3",
    "psp_short_name": "CG-MTN-MTN_3 short name"
}'



#-------------------------------------------------------
# List of all PSP

# By code
curl -X 'GET' \
  'http://0.0.0.0:8080/pspfee?psp_code=CG-MTN-MTN_2' \
  -H 'accept: application/json'


# By id
curl -X 'GET' \
  'http://0.0.0.0:8080/pspfee?psp_id=1' \
  -H 'accept: application/json'

# List All
curl -X 'GET' \
  'http://0.0.0.0:8080/pspfee' \
  -H 'accept: application/json'


#-------------------------------------------------------
# Add FeeSet to PSP

TODO: Maybe we have to initialize First active FeeSet on PSP Creation?
Because now It try to check active feeset to deactivate it
And when there is no active feeset it fails.
curl -X 'POST' \
  'http://localhost:8080/pspfee/feeset' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "psp_id": 1
}'

curl -X 'POST' \
  'http://localhost:8080/pspfee/feeset' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "psp_id": 1
}'

curl -X 'POST' \
  'http://localhost:8080/pspfee/feeset' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "psp_id": 1
}'




#-------------------------------------------------------
# Add FeeRange to FeeSet



curl -X 'POST' \
  'http://localhost:8080/pspfee/feerange' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "feeset_id": 5,
    "from": 0,
    "to": 1000,
    "fee_fixed": 50.00,
    "fee_percentage": 0.0,
    "max_total_fee": 999999999000
}'



curl -X 'POST' \
  'http://localhost:8080/pspfee/feerange' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "feeset_id": 5,
    "from": 1000.01,
    "to": 2000.00,
    "fee_fixed": 50.00,
    "fee_percentage": 0.0,
    "max_total_fee": 999999999000
}'



curl -X 'POST' \
  'http://localhost:8080/pspfee/feerange' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "feeset_id": 5,
    "from": 2000.01,
    "to": 3000.00,
    "fee_fixed": 60.00,
    "fee_percentage": 0.0,
    "max_total_fee": 999999999000
}'




#-------------------------------------------------------
# Calculate Single Fee


curl -X 'GET' \
  'http://0.0.0.0:8080/pspfee/calculate-single-fee?psp_code=CG-MTN-MTN_1&amount=3000.01' \
  -H 'accept: application/json'



#-------------------------------------------------------
# Calculate Bulk Fee

curl -X 'POST' \
  'http://0.0.0.0:8080/pspfee/calculate-bulk-fee' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[
  {
    "ref_id": "1",
    "psp_code": "CG-MTN-MTN_1",
    "amount": "1000.00"
  },
  {
    "ref_id": "1",
    "psp_code": "CG-MTN-MTN_1",
    "amount": "2000.00"
  },
  {
    "ref_id": "1",
    "psp_code": "CG-MTN-MTN_1",
    "amount": "2500.00"
  }
]'




























