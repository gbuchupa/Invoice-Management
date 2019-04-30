# Invoice Management on Hyperledger Fabric (*** Only Issue, Accept, Pay, View and History Details of Invoice implemented ***)
Trade finance application on Hyperledger Fabric

*** Use sudo prefix to commands if you get permission denied error while executing any command, assumption is you already have  required software to run Hyperledger fabric network and node SDK *** 

## Start the Hyperledger Fabric Network 

1. cd hlf-trade-finance
2. ./start.sh (with this you will start docker-compose.yml up -d )

## Setup the Hyperledger Fabric Network

1. cd hlf-trade-finance
2. ./setup.sh (With this you will create the channel genesis block, add the peer0 to the channel created and instantiate tfbc chaincode.) 

*** In this usecase CA's are already generated. 

We **do not have to run** the following again:

1. "generate --config=crypto-config.yaml"
2. "TFBCOrgOrdererGenesis -outputBlock ./config/genesis.block" 
3. "TFBCOrgChannel -outputCreateChannelTx ./config/tfbcchannel.tx -channelID tfbcchannel". 

These three statements are part of the "generate.sh" file here. 


## Setup API users 

1. cd hlf-trade-finance/tfbc-api
2. npm install
3. rm hfc-key-store/*
4. node enrollBankUser.js
5. node enrollBuyerUser.js
6. node enrollSellerUser.js

## Run Node APIs

1. cd hlf-trade-finance/tfbc-api
2. npm start

## Execute APIs on Swagger UI 
(Swagger allows you to describe the structure of your APIs so that machines can read them. The ability of APIs to describe their own structure is the root of all awesomeness in Swagger. Why is it so great? Well, by reading your API’s structure, we can automatically build beautiful and interactive API documentation. We can also automatically generate client libraries for your API in many languages and explore other possibilities like automated testing. Swagger does this by asking your API to return a YAML or JSON that contains a detailed description of your entire API.)

- To check this application specific swagger file go to:  hlf-trade-finance/tfbc-api/swagger.json 

http://localhost:3000/api-docs

## API definition if you want to run APIs using some rest client like Postman etc. 

### Issue Invoice
  1. URL -> http://localhost:3000/tfbc/issueInvoice
  2. Http Method -> Post
  3. content-type: application/json
  4. Input->
  {
  	"invoiceId": "KB200",
  	"invoiceDate": "20-04-2019",
  	"supplier": "Kill Bill",
  	"customer": "Talent Sprint",
  	"paymentTerms": "days 30",
  	"amount": "50000",
  	"notes": "good payment"
  }
  5. Output-> 
  {
    "code": "200",
    "message": "Invoice issued successsfully."
  }
### Accept Invoice

 1. URL -> http://localhost:3000/tfbc/acceptInvoice
 2. Http Method -> Post
 3. content-type: application/json
 4. Input->
  {
	"invoiceId": "KB200"
  }
 5. Output-> 
  {
    "code": "200",
    "message": "Invoice accepted successsfully."
  }
### Pay Invoice
 1. URL -> http://localhost:3000/tfbc/payInvoice
 2. Http Method -> Post
 3. content-type: application/json
 4. Input->
  {
	"invoiceId": "KB200"
  }
 5. Output-> 
  {
    "code": "200",
    "message": "Invoice paid successsfully."
  }
### Get Invoice Details 
 1. URL -> http://localhost:3000/tfbc/getInvoice
 2. Http Method -> Get
 3. content-type: application/json
 4. Input->
 {
	"invoiceId": "KB200"
  }
 5. Output-> 
  {
    "amount": 50000,
    "customer": "Talent Sprint",
    "invoiceDate": "20-04-2019",
    "invoiceId": "KB200",
    "notes": "good payment",
    "paymentTerms": "days 30",
    "status": "Paid",
    "supplier": "Kill Bill"
  }
### Get Invoice History 
 1. URL -> http://localhost:3000/tfbc/getInvoiceHistory
 2. Http Method -> Get
 3. content-type: application/json
 4. Input->
  {
	"invoiceId": "KB200"
  }
 5. Output-> 
  {
  "code": "200",
  "data": [
    {
      "TxId": "462a18ab161c55214d1121f43ffaa05150fb3b5ec5cedb33210c820cb2a0eefe",
      "Value": {
        "invoiceId": "KB200",
        "invoiceDate": "20-04-2019",
        "supplier": "Kill Bill",
        "customer": "Talent Sprint",
        "paymentTerms": "days 30",
        "amount": 50000,
        "notes": "good payment",
        "status": "Issued"
      },
      "Timestamp": "2019-04-20 12:54:20.334 +0000 UTC",
      "IsDelete": "false"
    },
    {
      "TxId": "0e5bd5d278cf8a836e3612350866404e2e8584e27b2ea6bd32b0d2506fc108a8",
      "Value": {
        "invoiceId": "KB200",
        "invoiceDate": "20-04-2019",
        "supplier": "Kill Bill",
        "customer": "Talent Sprint",
        "paymentTerms": "days 30",
        "amount": 50000,
        "notes": "good payment",
        "status": "Accepted"
      },
      "Timestamp": "2019-04-20 12:54:49.991 +0000 UTC",
      "IsDelete": "false"
    },
    {
      "TxId": "32a8594fd991420ef6053ed18b83a77b76268352c4662d6f632cee1a7360d4ba",
      "Value": {
        "invoiceId": "KB200",
        "invoiceDate": "20-04-2019",
        "supplier": "Kill Bill",
        "customer": "Talent Sprint",
        "paymentTerms": "days 30",
        "amount": 50000,
        "notes": "good payment",
        "status": "Paid"
      },
      "Timestamp": "2019-04-20 12:55:18.913 +0000 UTC",
      "IsDelete": "false"
    }
  ]
}

## Stop the network

1. cd hlf-trade-finance
2. ./stop.sh


#invoice-management_v1.0
