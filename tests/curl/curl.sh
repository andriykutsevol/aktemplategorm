#!/bin/bash

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
#   "currency": "string",
#   "auditDbRecord": {
#     "created_at": "2023-09-21T15:30:00Z",
#     "updated_at": "string",
#     "deleted_at": "string",
#     "created_by_guid": "string",
#     "updated_by_guid": "string",
#     "deleted_by_guid": "string"
#   }
# }'


# # management.POST(":psp_code", r.mobileMoneyWithdrawalFeeHandler.AddMobileProviderFeeSet)
# curl -X 'POST' \
#   'http://0.0.0.0:8080/pspfee/set/management/CG-MTN-MTN_2' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#   "psp": {
#     "psp_country_code": "string",
#     "psp_code": "CG-MTN-MTN_2",
#     "psp_short_name": "string"
#   },
#   "currency": "string",
#   "auditDbRecord": {
#     "created_at": "2023-09-21T15:30:00Z",
#     "updated_at": "string",
#     "deleted_at": "string",
#     "created_by_guid": "string",
#     "updated_by_guid": "string",
#     "deleted_by_guid": "string"
#   }
# }'



# # management.POST(":psp_code", r.mobileMoneyWithdrawalFeeHandler.AddMobileProviderFeeSet)
# curl -X 'POST' \
#   'http://0.0.0.0:8080/pspfee/set/management/CG-MTN-MTN_3' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#   "psp": {
#     "psp_country_code": "string",
#     "psp_code": "CG-MTN-MTN_3",
#     "psp_short_name": "string"
#   },
#   "currency": "string",
#   "auditDbRecord": {
#     "created_at": "2023-09-21T15:30:00Z",
#     "updated_at": "string",
#     "deleted_at": "string",
#     "created_by_guid": "string",
#     "updated_by_guid": "string",
#     "deleted_by_guid": "string"
#   }
# }'



# #====================================================================

# # As  system module, 
#     # I support the configuration of the following type of withdrawal fees

#         # Fixed fee per range
#         # % fee per range
#         # A combination of Fixed fee & % fees
#         # max  total fee

#     # As  system module, 
#         # I expose an API that allows a 3rd party service/application 
#         # to get the withdrawal fees


#     # from: lower bound this fee is configured to
#     # to:   upper bound this fee is configured to

# # management.POST(":pspfeeset_id", r.mobileMoneyWithdrawalFeeHandler.AddMobileProviderFeeRange)
# curl -X 'POST' \
#   'http://0.0.0.0:8080/pspfee/range/management/2' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#   "set_id": "2",
#   "from": 0,
#   "to": 1000.00,
#   "fee_fixed": 50,
#   "fee_percentage": 0,
#   "max_total_fee": 999999999000,
#   "auditDbRecord": {
#     "created_at": "2023-09-21T15:30:00Z",
#     "updated_at": "string",
#     "deleted_at": "string",
#     "created_by_guid": "string",
#     "updated_by_guid": "string",
#     "deleted_by_guid": "string"
#   }
# }'


# # management.POST(":pspfeeset_id", r.mobileMoneyWithdrawalFeeHandler.AddMobileProviderFeeRange)
# curl -X 'POST' \
#   'http://0.0.0.0:8080/pspfee/range/management/2' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#   "set_id": "2",
#   "from": 1000.01,
#   "to": 2000.00,
#   "fee_fixed": 50,
#   "fee_percentage": 0,
#   "max_total_fee": 999999999000,
#   "auditDbRecord": {
#     "created_at": "2023-09-21T15:30:00Z",
#     "updated_at": "string",
#     "deleted_at": "string",
#     "created_by_guid": "string",
#     "updated_by_guid": "string",
#     "deleted_by_guid": "string"
#   }
# }'


# # management.POST(":pspfeeset_id", r.mobileMoneyWithdrawalFeeHandler.AddMobileProviderFeeRange)
# curl -X 'POST' \
#   'http://0.0.0.0:8080/pspfee/range/management/2' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#   "set_id": "2",
#   "from": 2000.01,
#   "to": 3000.00,
#   "fee_fixed": 50,
#   "fee_percentage": 0,
#   "max_total_fee": 999999999000,
#   "auditDbRecord": {
#     "created_at": "2023-09-21T15:30:00Z",
#     "updated_at": "string",
#     "deleted_at": "string",
#     "created_by_guid": "string",
#     "updated_by_guid": "string",
#     "deleted_by_guid": "string"
#   }
# }'


# # management.POST(":pspfeeset_id", r.mobileMoneyWithdrawalFeeHandler.AddMobileProviderFeeRange)
# curl -X 'POST' \
#   'http://0.0.0.0:8080/pspfee/range/management/2' \
#   -H 'accept: application/json' \
#   -H 'Content-Type: application/json' \
#   -d '{
#   "set_id": "2",
#   "from": 3000.01,
#   "to": 4000.00,
#   "fee_fixed": 60,
#   "fee_percentage": 0,
#   "max_total_fee": 999999999000,
#   "auditDbRecord": {
#     "created_at": "2023-09-21T15:30:00Z",
#     "updated_at": "string",
#     "deleted_at": "string",
#     "created_by_guid": "string",
#     "updated_by_guid": "string",
#     "deleted_by_guid": "string"
#   }
# }'




# #====================================================================
# # ListFeeSetRange
# #====================================================================


# # management.GET(":pspfeeset_id", r.mobileMoneyWithdrawalFeeHandler.ListFeeSetRange)
# curl -X 'GET' \
#   'http://0.0.0.0:8080/pspfee/range/management/2' \
#   -H 'accept: application/json'




# #====================================================================
# # Calculate and return the withdrawal fees for given amount
# #====================================================================

# # "psp_code": "CG-MTN-MTN_2",
# # "set_id": "2", We have only one active set at a time for the PSP


# # functional.GET(":psp_code/:amount", r.mobileMoneyWithdrawalFeeHandler.CalculateFeeForAmount)
# curl -X 'GET' \
#   'http://0.0.0.0:8080/pspfee/set/functional/CG-MTN-MTN_2/2500' \
#   -H 'accept: application/json'



#====================================================================
# Patch
#====================================================================

# management.PATCH(":pspfeeset_id/:feerange_id", r.mobileMoneyWithdrawalFeeHandler.PatchFeeRange)
curl -X 'PATCH' \
  'http://0.0.0.0:8080/pspfee/range/management/1/2' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "from": 0,
  "to": 0,
  "fee_fixed": 0,
  "fee_percentage": 0,
  "auditDbRecord": {
    "created_at": "string",
    "updated_at": "string",
    "deleted_at": "string",
    "created_by_guid": "string",
    "updated_by_guid": "string",
    "deleted_by_guid": "string"
  }
}'


