# Documentation

This file contains documentation fetched by DocFetch.

---

## Home - ALATbyWema - developer portal

ALAT Open APIs...
Solving everyday business
problems with digital solutions
Explore the best
possibilities with ALAT APIs
Start building your own investment products with powerful experiences as quickly as possible.
Account Creation
Open higher tiers of accounts for your customers and get them to do more.
Get Statement Service
Use transaction statement for informed analytics.
Wallet Services
Create customized transaction wallets for your users with full banking capabilities.
How does it
work?
Register
Register on the developer portal to access and review products and APIs that meet your needs.
Build and Test
Build your solutions following the specifications outlined in the respective API references.
Go Live
Hit the market with customized solutions to for everyday needs.
©ALAT 2024. All Rights Reserved
Privacy
Terms of use

---

## Home - ALATbyWema - developer portal

ALAT Open APIs...
Solving everyday business
problems with digital solutions
Explore the best
possibilities with ALAT APIs
Start building your own investment products with powerful experiences as quickly as possible.
Account Creation
Open higher tiers of accounts for your customers and get them to do more.
Get Statement Service
Use transaction statement for informed analytics.
Wallet Services
Create customized transaction wallets for your users with full banking capabilities.
How does it
work?
Register
Register on the developer portal to access and review products and APIs that meet your needs.
Build and Test
Build your solutions following the specifications outlined in the respective API references.
Go Live
Hit the market with customized solutions to for everyday needs.
©ALAT 2024. All Rights Reserved
Privacy
Terms of use

---

## Product List - ALATbyWema - developer portal

Products
Build fast, reliable and secure financial solutions with ALAT OpenAPIs.
Get started by subscribing to a product of choice.
All Products
Account Creation - Address Verification
Get your customers to transact more volumes by opening higher tiers of accounts.
ALAT Rewards
Programmatically create deals and verify discount codes presented by customers for redemption.
Get Statement Service
Use transaction statements for informed analytics.
Pay with Bank Account
A fully digitized payment gateway to help businesses in collections.
Wallet Services
Create customized transaction wallets for your users with full banking capabilities.

---

## product-wallet-services - ALATbyWema - developer portal

Wallet Services
Create customized transaction wallets for your users with full banking capabilities.
Overview
This is a collection of APIs which can be used by partners of Wema bank to open and manage wallet accounts. They include:
Wallet Creation API
Credit Wallet API
Debit Wallet API
Bills Payment API
Transaction Notification API
Account Management API
Note
: Unlike the Virtual Accounts, these wallets operate, hold funds and transact like tier 1 accounts. However, National Identification Number (NIN) is required to operate these wallets.
Go Live
How does it work?
Partners can request creation of wallets on behalf of their customers and would have autonomy on use and management of these wallets. However, wallets operation are similar to tier 1 accounts, with minimum KYC requirements.
Partners can receive notifications on all debit and credit transactions that occur on the wallets they generate. These wallets can also be linked to debit cards for transactions.
Who can use this?
Possible use cases for wallet services include:
Promoting subscription based payment options on service applications.
Store of funds by customers used in playing games in lottery and betting applications.
Payout of winnings by lottery and betting agencies.
Issuance of branded gift cards by agencies or stores.
API Reference
Wallet Creation API
API to Generate customized wallets for client application.
Debit Wallet API
API to perform interbank and intrabank transactions on wallets.
Account Management API
API to retrieve details and transactional details on wallets.
Credit Wallet API
API used by clients to fund wallets generated in their books.
Notification API
Get real time transaction notifications on custom wallets.
Bills Payment
API to process bills payment using wallets as source accounts

---

## Product - Get Statement Service - ALATbyWema - developer portal

Get Statement Service
This service provides access to the transactional activities of Wema/ALAT customers.
Overview
Account transactional data is captured from transactions. It is a record of the amount, time, and channel by which a transaction occurred. It also records the type, as well as other quantities and qualities associated with the transaction.
This service requires getting customers' consent before transaction data are shared.
How does it work?
Wema/ALAT bank account holders must give consent before data can be shared. Customers are to give their consent using the ALAT mobile application.
Who can use this?
Possible users of this service include but not limited to:
Embassies.
Credit eligibility checks.
Financial management institutions
Go Live
API Reference
Get Statement API
Use transaction statement for informed analytics

---

## API Account Creation Face Biometric Authentication - ALATbyWema - developer portal

Account Creation - Face Biometric Authentication API
The Account Creation API enables clients to create wallets for their users by verifying their identities through facial biometrics using either BVN or NIN. The account creation journey differs slightly based on the account tier and the identity document used.
How does it work?
The
Account Creation API
allows clients request the creation of customized wallets. These wallets can only be managed and operated using the channel or application of choice by the requesting client.
The client performs face biometric authentication using end users' BVN or NIN using the
face biometric authentication web app.
The client makes a request on behalf of the customer to generate an account.
The client gets a
PENDING
status of while awaiting account generation.
Client receives a notification of wallet details via a callback URL.
Prerequisites
A valid
x-api-key
assigned by the bank.
A face biometric authentication UI flow that integrates with the provided facial prompt URL.
A callback webhook URL configured to receive account details.
https://face-verification-dev.azurewebsites.net/?{identifierType}={identifierValue}&x_tk{x-api-key}
Go Live
Face Biometric Authentication Journey
Access the Face Biometric Authentication web app using the
BaseURL
above. A successful face biometric authentication by end user via the web app would return a correlationId to the client, with which to initiate account generation request.
The Face Biometric Authentication web app makes use of several key parameters to note:
identifierType
= Either BVN or NIN
identifierValue
= Value of either BVN or NIN.
x_tk
= Value of
x-api-key
assigned by the bank to client.
Additional/Optional Parameters
cb_uri
= A callback URI to be passed by the client if details of
correlationId
is to be delivered to the client's server-side application.
For callback URL ->
{{BASE_URL }}/?nin={{ NIN }}&x_tk={{ x-api-key }}&cb_uri={{ CALLBACK_URL}}
Where callback url is provided, a POST request would be made to the callback with the payload.
POST - {value of callback URI passed by client}
{
"success": true,
"c_id": {value of correlationID},
"id": {value of BVN_OR_NIN used},
"id_type": "bvn" //or "nin"
}
rd_uri
=
A redirect URI to be passed by the client if details of
correlationId
is to be delivered by a return to the client's web application.
For redirect URL ->
{{ BASE_URL }}/?nin={{ NIN }}&x_tk={{ x-api-key }}&rd_uri={{ REDIRECT_URL }}
On success of face biometric authentication, redirect to
{{REDIRECT_URL }}/?success=true&&c_id=${{ CorrelationID_HERE }}&id=${{ value of BVN_OR_NIN }}&id_type=${{ "bvn" or "nin" }}
Note
Values of
cb_uri
and
rd_uri
should maintain the format
https://{domain_name_of_client_web_application}
On success of face biometric authentication, if no redirect or callback URLs are provided redirect would occur at
{{ BASE_URL }}/?success=true&&c_id=${{ ID_HERE }}&id=${{ BVN_OR_NIN }}&id_type=${{ "bvn" or "nin" }}

---

## API - Account Creation with Address Verification - ALATbyWema - developer portal

Account Creation - Address Verification
Create customized accounts for customers. These accounts are upgraded upon successful validation of BVN and NIN.
How does it work?
The
Account Creation API
allows clients request the creation of accounts, which can be upgraded to higher tiers upon successful physical verification of the home addresses of the customers. These accounts can only be managed and operated using the channel or application of choice by the requesting client.
The client makes a request on behalf of the customer to generate an account.
The client gets a
PENDING
status of while awaiting account generation.
Client receives a notification of account details via a callback URL.
Generated accounts are upgraded to higher tiers upon successful verification of customer's home address. Verification of addresses is managed independently by the bank.
Go Live
Step 1 - Generate Account
Request URL
POST - {baseURL}/partnership-address-verification/api/CustomerAccount/GeneratePartnershipAccount
Client initiates request on behalf of customer using the request model below.
Client is required to pass a valid
x-api-key
value in the header of their request. This value would be provided by the bank.
Customer's NIN, BVN and address are mandatory.
A successful call to this endpoint would
check validity of NIN and BVN provided. Note, failed validation of NIN and BVN would terminate the process.
send SMS OTP to the phone number on NIN to ascertain customer's consent.
a
trackingId
is returned to the client to be used for validation of NIN in step 2 below.
Client proceeds to step 2, in order to validate NIN.
Step 2 - Validate NIN
Request URL
POST - {baseURL}/partnership-address-verification/api/CustomerAccount/ValidateOtpAndGeneratePartnershipAccount
Client provides a medium/interface for customer to input OTP received from step 1.
Client initiates request on behalf of customer by sending
otp, trackingId
and
phoneNumber
used in step 1 above.
Client is required to pass a valid
x-api-key
value in the header of their request. Same as in step 1
A successful call to this endpoint would return a
PENDING
status of wallet generation. The wait time for wallet generation is a maximum of 10 minutes.
Client receives a callback once account generation process is completed. The callback model can be seen below.
Note
A tier 1 account would be opened immediately OTP validation is done.
Account would remain operational as tier 1 pending upgrade.
Upgrade to higher tiers would take place after successful physical verification of the customer's address by the bank.
Failed address verification would require the client to resubmit address details using step 4.
Step 3 - Get Partnership Account Details
Request URL
POST - {baseURL}/partnership-address-verification/api/CustomerAccount/GetPartnershipAccountDetails[?phoneNumber]
This is a GET endpoint that helps client retrieve details of the generated account for customers.
It accepts the customer's
phoneNumber
as query parameter.
Expected response model can be seen below.

---

## Product - ALAT Rewards - ALATbyWema - developer portal

ALAT Rewards
ALAT Rewards is a platform that offers merchants an opportunity to showcase products and offers deals to users on ALAT.
Overview
ALAT Rewards is a product of the ALATbyWema platform that serves as a trusted and secured marketplace for merchants and franchises to drive sales of their products to millions of users (buyers or subscribers).
Merchants can have their deals and offerings published programmatically in near real time visibility to customers on ALAT.
Merchant deals would undergo a verification process by the bank before they are published to customers on ALAT Rewards.
Contact the business team at
retailpartnership@wemabank.com
.
How does it work?
Wema/ALAT bank account holders must give consent before data can be shared. Customers are to give their consent using the ALAT mobile application.
Businesses and partners are registered by the bank as merchants on ALAT Rewards.
Merchants create deals and discount on the merchant portal and is published.
Merchants can programmatically verify discount codes using the
Verify Discount Code API
.
Who can use this?
Possible merchants to use this service include but not limited to:
Franchises - Superstores, Restaurants, etc.
Subscription Based Web Platforms
Go Live
API Reference
Verify Discount Code API
Programmatically verify discount codes for redemption.

---

## Product - Pay with Bank Account - ALATbyWema - developer portal

Pay with Bank Account
A digitized payment option using a Wema/ALAT account number.
Overview
This service offers an additional mode of online payments and collections to merchants and e-commerce platforms.
How does it work?
Merchant is profiled by the bank to use this service.
Merchant profile is setup to with charge type - flat fee or rate. (To be agreed with the bank)
Charge is debited from the amount paid by customer on merchant website.
Wema/ALAT bank account holders must authorize a transaction using a one-time-password sent via SMS.
Who can use this?
All merchants or e-commerce platforms willing to diversify payment options available to customers.
Go Live
API Documentation
Pay with Bank Account
Digital payment/collection using account number.

---

## API - Debit Wallet - ALATbyWema - developer portal

Debit Wallet API
Make fund transfers from customized wallets/accounts (source).
Test Source Account:
0410530975
Test Destination Account:
0280100171
How does it work?
The
Debit Wallet API
allows clients to make fund transfers from customized wallets/accounts generated by them. This API works as both an
intrabank and interbank transfer service
; destination accounts can belong to any bank in Nigeria.
The client is profiled to perform a funds transfer. This activity is done in the backend by the bank, and mapped to the clients
access(x-api)
key.
The client makes an enquiry on the
sourceAccountNumber.
This is to ensure the client has the permission to debit said wallet/account.
The client makes an enquiry to confirm the destination account information.
Based on the amount band, the client makes an enquiry to confirm NIP charges. This only applies if the destination account number is non-Wema. (
Note
:
This NIP charge will be taken care of by the bank; there is no need to add it to your request.)
The client initiates a payout request and gets a
PENDING
status.
A callback is made to the client to authorize the transaction - debit to the source account and credit to the destination account/wallet.
The bank proceeds to fulfill the transfer request.
Client receives a notification of transaction status.
Note:
You will need to profile your Authentication Callback URL for this endpoint to avoid getting the authentication failed error while testing.
Go Live
Step 1: Name Enquiry - Source Account
Request URL
POST - {baseURL}/debit-wallet/api/Shared/AccountNameEnquiry/Wallet/{accountNumber}
The client is required to pass a valid
access(x-api)
value in the header of their request. This value would be provided by the bank.
Client validates the source account number using the request URL.
A successful call to this endpoint would return the
accountName
name of the beneficiary.
Step 2: Get List of Banks
Request URL
GET - {baseURL}/debit-wallet/api/Shared/GetAllBanks
This endpoint returns a list of all banks in Nigeria. Clients are to select destination bank name and code from this list.
Step 3: Name Enquiry - Destination Account
Request URL
GET - {baseURL}/debit-wallet/api/Shared/AccountNameEnquiry/{bankCode}/{accountNumber}[?channelId]
Clients are to pass the bankCode retrieved from the list of banks, along with the destinationAccountNumber, as query parameters to perform a name enquiry on the destination account.
A successful call to this endpoint would return details of the destination account number, as well as a list of charge fees based on amount bands.
Step 4: Get NIP Charges
Request URL
GET - {baseURL}/debit-wallet/api/Shared/GetNIPCharges
This endpoint returns a list of all charges base on transaction type and amount.
Step 5: Process Funds Transfer
Request URL
POST - {baseURL}/debit-wallet/api/Shared/ProcessClientTransfer
The client initiates a request for transfer to
destinationAccountNumber
, specifying
amount
and custom
transactionReference
and narration (set
useCustomNarration
as
true
).
The client is expected to pass a
securityInfo
. This should be an encrypted value whose signature is known only by the client. The client should be the only one with the key to decrypt its value. This value serves as the mandate for authorization of debit.
The bank makes an authentication request to the client's authentication URL to confirm the
securityInfo
of the transaction. The client is expected to validate this authentication request for the transaction to be executed. Authentication request and response models can be seen below.
The client gets a
PENDING
transaction status while the transaction is being executed.
Client receives final transaction status via a callback URL.
Note
: Transaction statuses are
PENDING
,
SUCCESSFUL
or
FAILED
.
Note:
Your security information is unique to you and can be an alphanumeric string of any length. It helps secure your transactions and authenticate every transaction initiated by you, preventing unauthorized access by third parties.
Authentication Model
Request Model
The bank makes this request to the client to authorize this transaction request.
The authorization callback URL is profiled by the bank and mapped to the client's
access(x-api)
key.
The bank makes an authentication request to the client by passing the
securityInfo
and custom
transactionReference
specified in the payout request.
Clients should ensure that the encryption used for the
securityInfo
parameter is highly secure, dynamic, and unique to each transaction request.
This endpoint should not take any headers as it is generic to all clients using this service.
Expected Response Model
Client responds with the
transactionReference
and
authorized
status of the authentication request.
Authentication status
authorized = true;
if
securityInfo
is valid
authorized = false;
if
securityInfo
could not be validated.
The Callback Model
Clients are to build their callback following the specification below. Your callback URL would be profiled to your assigned
channelId
.

---

