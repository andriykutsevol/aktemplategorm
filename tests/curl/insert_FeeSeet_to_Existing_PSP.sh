#!/bin/bash


# curl -X 'POST' \
#   'http://0.0.0.0:8080/pspfee/calculatesinglefee' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#     "amount": "CM",
#     "psp_code": "CG-MTN-MTN_3",
#     }'







# # management.POST(":psp_code", r.mobileMoneyWithdrawalFeeHandler.AddMobileProviderFeeSet)
# curl -X 'POST' \
#   'http://0.0.0.0:8080/pspfee/set/management/CG-MTN-MTN_1' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#   "psp": {
#     "psp_country_code": "string",
#     "psp_code": "CG-MTN-MTN_1",
#     "psp_short_name": "string"
#   },
#   "currency": "string"
# }'