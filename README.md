Step#1: Install Below Mentioned Pre-requisites on the machine
1. Install Go1.9 
	• curl -O https://storage.googleapis.com/golang/go1.11.2.linux-amd64.tar.gz
	• tar -xvf go1.11.2.linux-amd64.tar.gz
	• sudo mv go /usr/local
	• sudo nano ~/.profile
  	export GOPATH=$HOME/work
  	export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

2. Install npm 
○ Check if node is already installed by node -v command
○ Follow below sequence to install NodeJs on Ubuntu 16.04( For other operating systems installations steps please follow Starter Kit ReadMe file)
	§ sudo apt install curl
	§ curl -sL https://deb.nodesource.com/setup_10.x | sudo bash -
	§ sudo apt install nodejs

3. Install Python
	○ sudo apt-get install python

4. Install Docker
	○ sudo apt-get update
	○ sudo apt-get install \
    	apt-transport-https \
    	ca-certificates \
    	curl \
    	software-properties-common
	○ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
	○ sudo add-apt-repository \
   	"deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   	$(lsb_release -cs) \
   	stable"
	○ sudo apt-get update
	○ sudo apt-get install docker-ce
	○ cat /etc/group | grep docker
	○ sudo usermod -aG docker ubuntu
	○ cat /etc/group | grep docker
	○ Logout and log back in to machine

5. Install Docker Compose
	○ sudo apt-get install docker-compose

6. Clone Trade Finance from github
	• sudo git clone https://github.com/gbuchupa/Invoice-Management.git
	• sudo chown -R ubuntu:ubuntu hlf-trade-finance/
	• cd  Invoice-Management
	• git checkout release

7. Create ".env" file under project directory, if not already existing.
	• cd /Invoice-Management
	• vim .env
	COMPOSE_PROJECT_NAME=net
	IMAGE_TAG=latest

# Invoice Management on Hyperledger Fabric (*** Only Issue, Accept, Pay, View and History Details of Invoice implemented ***)
Trade finance application on Hyperledger Fabric

*** Use sudo prefix to commands if you get permission denied error while executing any command, assumption is you already have  required software to run Hyperledger fabric network and node SDK *** 

Step#2: Start the Hyperledger Fabric Network 
	cd hlf-trade-finance
	./start.sh (with this you will start docker-compose.yml up -d )

Step#3: Setup the Hyperledger Fabric Network
	cd hlf-trade-finance
	./setup.sh (With this you will create the channel genesis block, add the peer0 to the channel created and instantiate tfbc 		chaincode.) 

//*** In this usecase CA's are already generated. 
We **do not have to run** the following again:

	1. "generate --config=crypto-config.yaml"
	2. "TFBCOrgOrdererGenesis -outputBlock ./config/genesis.block" 
	3. "TFBCOrgChannel -outputCreateChannelTx ./config/tfbcchannel.tx -channelID tfbcchannel". 
These three statements are part of the "generate.sh" file here. //


Step#5: Setup API users 
	○ cd Invoice-Management/tfbc-api
	○ sudo apt install build-essential g++
	○ npm install -g
	○ rm hfc-key-store/*
	○ node enrollBankUser.js
	○ node enrollBuyerUser.js
	○ node enrollSellerUser.js


Step#6: Run Node APIs
	cd hlf-trade-finance/tfbc-api
	npm start

Step#7: Execute APIs on Swagger UI 
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
