var maxGas=4700000;
var abi = [ { "constant": false, "inputs": [], "name": "sold", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [], "name": "getBidIndex", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [], "name": "notForSale", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "value", "type": "uint256" } ], "name": "setAppraisalValue", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [], "name": "askingPrice", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "isForSale", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "appraised", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "price", "type": "uint256" } ], "name": "forSale", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "bidNumber", "type": "uint256" } ], "name": "getBid", "outputs": [ { "name": "", "type": "address" }, { "name": "", "type": "uint256" }, { "name": "", "type": "bool" }, { "name": "", "type": "uint256" }, { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "changedMortgageValue", "type": "uint256" } ], "name": "changeMortgageValue", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [], "name": "kill", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "bids", "outputs": [ { "name": "bidder", "type": "address" }, { "name": "bidValue", "type": "uint256" }, { "name": "accepted", "type": "bool" }, { "name": "bidNumber", "type": "uint256" }, { "name": "status", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "inspected", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "streetAddress", "outputs": [ { "name": "", "type": "bytes32" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "appraisalValue", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "biddingPrice", "type": "uint256" } ], "name": "checkBid", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "bidderAddress", "type": "address" } ], "name": "rejectBid", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [], "name": "bidIndex", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "houseOwner", "outputs": [ { "name": "", "type": "address" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "newBidder", "type": "address" } ], "name": "acceptBid", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [], "name": "buyer", "outputs": [ { "name": "", "type": "address" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "bidValue", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "status", "type": "uint256" } ], "name": "inspectionStatus", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "bidNumber", "type": "uint256" } ], "name": "deleteBid", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "biddingPrice", "type": "uint256" } ], "name": "placeBid", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "newPrice", "type": "uint256" } ], "name": "changePrice", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "value", "type": "uint256" } ], "name": "rejectMortgage", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [], "name": "state", "outputs": [ { "name": "", "type": "uint8" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "requestedMortgage", "type": "uint256" } ], "name": "applyMortgage", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [], "name": "acceptedBid", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [], "name": "deleteAllBids", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [], "name": "bidder", "outputs": [ { "name": "", "type": "address" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "soldPrice", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "inputs": [ { "name": "stAddress", "type": "bytes32" } ], "payable": false, "stateMutability": "nonpayable", "type": "constructor" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "ownerAddress", "type": "address" } ], "name": "InitHouse", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "ownerAddress", "type": "address" } ], "name": "HouseInitialized", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "seller", "type": "address" }, { "indexed": true, "name": "askingPrice", "type": "uint256" }, { "indexed": true, "name": "stAddress", "type": "bytes32" } ], "name": "HouseForSale", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "seller", "type": "address" }, { "indexed": true, "name": "stAddress", "type": "bytes32" } ], "name": "HouseNotForSale", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "seller", "type": "address" }, { "indexed": true, "name": "askingPrice", "type": "uint256" }, { "indexed": true, "name": "stAddress", "type": "bytes32" } ], "name": "PriceChanged", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "bidder", "type": "address" }, { "indexed": true, "name": "biddingPrice", "type": "uint256" } ], "name": "BidPlaced", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "bidder", "type": "address" } ], "name": "BidRejected", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "bidder", "type": "address" }, { "indexed": true, "name": "acceptedBid", "type": "uint256" } ], "name": "BidAccepted", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "bidder", "type": "address" }, { "indexed": true, "name": "biddingPrice", "type": "uint256" } ], "name": "BidExists", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "inspector", "type": "address" }, { "indexed": true, "name": "status", "type": "uint256" } ], "name": "Inspected", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "lender", "type": "address" }, { "indexed": true, "name": "value", "type": "uint256" } ], "name": "MortgageApplied", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "lender", "type": "address" }, { "indexed": true, "name": "value", "type": "uint256" } ], "name": "MortgageApproved", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "lender", "type": "address" }, { "indexed": true, "name": "value", "type": "uint256" } ], "name": "MortgageRejected", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "buyer", "type": "address" }, { "indexed": true, "name": "price", "type": "uint256" } ], "name": "HouseSold", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "mortgageValue", "type": "uint256" } ], "name": "ChangedMortgageValue", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "biddingPrice", "type": "uint256" } ], "name": "BidAlreadyExists", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "bidNumber", "type": "uint256" } ], "name": "BidAlreadyAccepted", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "acceptedBid", "type": "uint256" } ], "name": "NoBidsAccepted", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "bidder", "type": "address" }, { "indexed": true, "name": "bidNumber", "type": "uint256" } ], "name": "FoundBid", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "bidder", "type": "address" }, { "indexed": true, "name": "bidNumber", "type": "uint256" } ], "name": "UpdatedBid", "type": "event" }, { "anonymous": false, "inputs": [], "name": "AllBidsDeleted", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "requestor", "type": "address" } ], "name": "GetBidIndexRequest", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "requestor", "type": "address" }, { "indexed": true, "name": "bindNumber", "type": "uint256" } ], "name": "GetBidRequest", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "appraiser", "type": "address" }, { "indexed": true, "name": "value", "type": "uint256" } ], "name": "AppraisalValueSet", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "bidNumber", "type": "uint256" } ], "name": "DeleteBids", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "bidNumber", "type": "uint256" }, { "indexed": true, "name": "acceptedBid", "type": "uint256" } ], "name": "BidsDeleted", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "from", "type": "address" } ], "name": "Killed", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "from", "type": "address" }, { "indexed": true, "name": "bidNumber", "type": "uint256" } ], "name": "BidDeleted", "type": "event" } ]

var address;
var contract;
var houseForSaleEvent;
var houseNotForSaleEvent;
var housePriceChangeEvent;
var houseBidPlacedEvent;
var houseBidAcceptedEvent;
var houseBidRejectEvent;
var houseInspectedEvent;
var houseMortgagePlacedEvent;
var houseMortgageAcceptedEvent;
var houseMortgageRejectEvent;
var houseSoldEvent;
var houseMortgageHigherThanBidEvent;


function callHouseForSale(c,ap,s) {
    queryContract = web3.eth.contract(abi);
    contract = queryContract.at(c);
    contract.forSale(ap, {from: s});
}

function callHouseNotForSale(c,s) {
    queryContract = web3.eth.contract(abi);
    contract = queryContract.at(c);
    contract.notForSale({from: s});
}

function callHousePlaceBid(c,b,u) {
    queryContract = web3.eth.contract(abi);
    contract = queryContract.at(c);
    // Now that a bid has been placed, the price of the house can't be changed
    // and the house can't be removed from the market
    contract.placeBid(b, { from: u});
}

function callHouseAcceptBid(c,u) {
    queryContract = web3.eth.contract(abi);
    contract = queryContract.at(c);
    address = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
    contract.acceptBid(u,{from: address});
}

function callHouseRejectBid(c,b,u) {
    queryContract = web3.eth.contract(abi);
    contract = queryContract.at(c);
    address = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
    contract.rejectBid(u, b, {from: address});
}

function callHouseApplyMortgage(c,a) {
    queryContract = web3.eth.contract(abi);
    contract = queryContract.at(c);
    l = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee";
    contract.applyMortgage(a, { from: l});
}


function callHouseChangeMortgageValue(c,a) {
    queryContract = web3.eth.contract(abi);
    contract = queryContract.at(c);
    l = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee";
    contract.changeMortgageValue(a, { from: l});
}

function callRejectMortgage() {
    var houseAddress = document.getElementById('houseAddress').value;
    var lenderAddress = document.getElementById('lenderAddress').value;
    var mortgageValue = document.getElementById('houseMortgageValue').value;

    document.getElementById('status').innerText = "Waiting for the transaction to accept the mortgage from " + lenderAddress + " for a value of $" + mortgageValue;
    contract.rejectMortgage(mortgageValue);
}

function callHouseAppraisal(c,a) {
    queryContract = web3.eth.contract(abi);
    contract = queryContract.at(c);
    address = "0xdddddddddddddddddddddddddddddddddddddddd";
    contract.setAppraisalValue(a, { from: address});
}

function listBids(c) {
	queryContract = web3.eth.contract(abi);
	contract = queryContract.at(c);
	numberBids = contract.getBidIndex();
	var bids = {"bids" : []};
	bids["numberBids"] = numberBids;
	for(i = 1; i < parseInt(numberBids); i++) {
		bid = {};
		res = contract.getBid(i);
		bid["bidder"] = res[0];
		bid["amount"] = res[1];
		bid["accepted"] = res[2];
		bid["index"] = res[3];
		bids["bids"].push(bid);
	}
	return bids;
}