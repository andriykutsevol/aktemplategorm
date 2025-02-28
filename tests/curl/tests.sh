#!/bin/bash

#-------------------------------------------------------
# Create PSP



# curl -X 'POST' \
#   'http://localhost:8080/pspfee' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#     "psp_country_code": "CM",
#     "psp_code": "CG-MTN-MTN_1",
#     "psp_short_name": "CG-MTN-MTN_1 short name"
# }'


# curl -X 'POST' \
#   'http://localhost:8080/pspfee' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#     "psp_country_code": "CM",
#     "psp_code": "CG-MTN-MTN_2",
#     "psp_short_name": "CG-MTN-MTN_2 short name"
# }'


# curl -X 'POST' \
#   'http://localhost:8080/pspfee' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#     "psp_country_code": "CM",
#     "psp_code": "CG-MTN-MTN_3",
#     "psp_short_name": "CG-MTN-MTN_3 short name"
# }'



#-------------------------------------------------------
# List of all PSP

# curl -X 'GET' \
#   'http://0.0.0.0:8080/pspfee' \
#   -H 'accept: application/json'





#-------------------------------------------------------
# Add FeeSet to PSP

# curl -X 'POST' \
#   'http://localhost:8080/pspfee/feeset' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#     "psp_id": 2
# }'

# curl -X 'POST' \
#   'http://localhost:8080/pspfee/feeset' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#     "psp_id": 1
# }'

# curl -X 'POST' \
#   'http://localhost:8080/pspfee/feeset' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#     "psp_id": 1
# }'




#-------------------------------------------------------
# Add FeeRange to FeeSet



# curl -X 'POST' \
#   'http://localhost:8080/pspfee/feerange' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#     "feeset_id": 1,
#     "from": 0,
#     "to": 1000,
#     "fee_fixed": 50.00,
#     "fee_percentage": 0.0,
#     "max_total_fee": 999999999000
# }'



# curl -X 'POST' \
#   'http://localhost:8080/pspfee/feerange' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#     "feeset_id": 1,
#     "from": 1000.01,
#     "to": 2000,
#     "fee_fixed": 50.00,
#     "fee_percentage": 0.0,
#     "max_total_fee": 999999999000
# }'



# curl -X 'POST' \
#   'http://localhost:8080/pspfee/feerange' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#     "feeset_id": 1,
#     "from": 2000.01,
#     "to": 3000,
#     "fee_fixed": 60.00,
#     "fee_percentage": 0.0,
#     "max_total_fee": 999999999000
# }'







#-------------------------------------------------------
# Calculate Single Fee


# curl -X 'GET' \
#   'http://0.0.0.0:8080/pspfee/calculate-single-fee?psp_code=CG-MTN-MTN_1&amount=2000.01' \
#   -H 'accept: application/json'



# #-------------------------------------------------------
# # Calculate Bulk Fee

# curl -X 'POST' \
#   'http://0.0.0.0:8080/pspfee/calculate-bulk-fee' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '[
#   {
#     "psp_id": "1",
#     "psp_code": "CG-MTN-MTN_1",
#     "amount": "1000.00"
#   },
#   {
#     "psp_id": "1",
#     "psp_code": "CG-MTN-MTN_1",
#     "amount": "2000.00"
#   },
#   {
#     "psp_id": "1",
#     "psp_code": "CG-MTN-MTN_1",
#     "amount": "2000.01"
#   }   
# ]'











