var Web3 = require('web3');
var web3 = new Web3();
web3.setProvider(new web3.providers.HttpProvider('http://54.245.138.237:8545'));

var maxGas=4700000;

var abi= [{"constant":false,"inputs":[],"name":"sold","outputs":[],"payable":false,"type":"function"},
{"constant":false,"inputs":[],"name":"notForSale","outputs":[],"payable":false,"type":"function"},
{"constant":false,"inputs":[{"name":"value","type":"uint256"}],"name":"setAppraisalValue","outputs":[],"payable":false,"type":"function"},
{"constant":true,"inputs":[],"name":"askingPrice","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},
{"constant":true,"inputs":[],"name":"isForSale","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},
{"constant":false,"inputs":[{"name":"price","type":"uint256"}],"name":"forSale","outputs":[],"payable":false,"type":"function"},
{"constant":false,"inputs":[{"name":"changedMortgageValue","type":"uint256"}],"name":"changeMortgageValue","outputs":[],"payable":false,"type":"function"},
{"constant":false,"inputs":[],"name":"kill","outputs":[],"payable":false,"type":"function"},
{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"bids","outputs":[{"name":"bidder","type":"address"},
{"name":"bidValue","type":"uint256"},
{"name":"accepted","type":"bool"},
{"name":"bidNumber","type":"uint256"}],"payable":false,"type":"function"},
{"constant":true,"inputs":[],"name":"inspected","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},
{"constant":true,"inputs":[],"name":"streetAddress","outputs":[{"name":"","type":"bytes32"}],"payable":false,"type":"function"},
{"constant":true,"inputs":[],"name":"appraisalValue","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},
{"constant":true,"inputs":[{"name":"biddingPrice","type":"uint256"}],"name":"checkBid","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},
{"constant":false,"inputs":[{"name":"bidderAddress","type":"address"}],"name":"rejectBid","outputs":[],"payable":false,"type":"function"},
{"constant":true,"inputs":[],"name":"bidIndex","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},
{"constant":true,"inputs":[],"name":"houseOwner","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},
{"constant":false,"inputs":[{"name":"bidder","type":"address"}],"name":"acceptBid","outputs":[],"payable":false,"type":"function"},
{"constant":true,"inputs":[],"name":"buyer","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},
{"constant":true,"inputs":[],"name":"bidValue","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},
{"constant":false,"inputs":[{"name":"status","type":"uint256"}],"name":"inspectionStatus","outputs":[],"payable":false,"type":"function"},
{"constant":false,"inputs":[{"name":"biddingPrice","type":"uint256"}],"name":"placeBid","outputs":[],"payable":false,"type":"function"},
{"constant":false,"inputs":[{"name":"newPrice","type":"uint256"}],"name":"changePrice","outputs":[],"payable":false,"type":"function"},
{"constant":false,"inputs":[{"name":"value","type":"uint256"}],"name":"rejectMortgage","outputs":[],"payable":false,"type":"function"},
{"constant":true,"inputs":[],"name":"state","outputs":[{"name":"","type":"uint8"}],"payable":false,"type":"function"},
{"constant":false,"inputs":[{"name":"requestedMortgage","type":"uint256"}],"name":"applyMortgage","outputs":[],"payable":false,"type":"function"},
{"constant":true,"inputs":[],"name":"acceptedBid","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},
{"constant":true,"inputs":[],"name":"bidder","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},
{"constant":true,"inputs":[],"name":"soldPrice","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},
{"inputs":[{"name":"stAddress","type":"bytes32"}],"payable":false,"type":"constructor"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"seller","type":"address"},
{"indexed":true,"name":"askingPrice","type":"uint256"},
{"indexed":true,"name":"stAddress","type":"bytes32"}],"name":"HouseForSale","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"seller","type":"address"},
{"indexed":true,"name":"stAddress","type":"bytes32"}],"name":"HouseNotForSale","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"seller","type":"address"},
{"indexed":true,"name":"askingPrice","type":"uint256"},
{"indexed":true,"name":"stAddress","type":"bytes32"}],"name":"PriceChanged","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"bidder","type":"address"},
{"indexed":true,"name":"biddingPrice","type":"uint256"}],"name":"BidPlaced","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"bidder","type":"address"}],"name":"BidRejected","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"bidder","type":"address"}],"name":"BidAccepted","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"bidder","type":"address"},
{"indexed":true,"name":"biddingPrice","type":"uint256"}],"name":"BidExists","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"inspector","type":"address"},
{"indexed":true,"name":"status","type":"uint256"}],"name":"Inspected","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"lender","type":"address"},
{"indexed":true,"name":"value","type":"uint256"}],"name":"MortgageApplied","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"lender","type":"address"},
{"indexed":true,"name":"value","type":"uint256"}],"name":"MortgageApproved","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"lender","type":"address"},
{"indexed":true,"name":"value","type":"uint256"}],"name":"MortgageRejected","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"buyer","type":"address"},
{"indexed":true,"name":"price","type":"uint256"}],"name":"HouseSold","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"mortgageValue","type":"uint256"}],"name":"ChangedMortgageValue","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"biddingPrice","type":"uint256"}],"name":"BidAlreadyExists","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"bidNumber","type":"uint256"}],"name":"BidAlreadyAccepted","type":"event"},
{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"}],"name":"Killed","type":"event"}];


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

function callCreateHouse() {
    // let's assume that we have a private key to coinbase ;)
    web3.eth.defaultAccount = web3.eth.coinbase;
    var houseAddress = document.getElementById('houseAddress').value;

    document.getElementById('housecreation_frm').style.visibility = 'hidden';
    document.getElementById('status').innerText = "Transaction to create contract for house @ address " + houseAddress + " sent, waiting for confirmation";
    console.log("Creating contract for a house at address " + houseAddress);

    web3.personal.unlockAccount(web3.eth.accounts[0], "Molodoi1", 1000);
    web3.eth.contract(abi).new(houseAddress, {
        data: code,
        gas: maxGas
    },
 function(err, c) {
        if (err) {
            console.error(err);
            return;

            // callback fires twice, we only want the second call when the contract is deployed
        } else if (c.address) {

            contract = c;
            console.log('address: ' + contract.address);
            document.getElementById('status').innerText = "Contract for House at Address " + houseAddress + " created...";
            document.getElementById('houseForSale_frm').style.visibility = 'visible';
            //	    	    document.getElementById('housePriceChange_frm').style.visibility = 'visible';
            document.getElementById('houseState_frm').style.visibility = 'visible';

            housePriceChangeEvent = contract.PriceChanged(housePriceChanged);
            houseForSaleEvent = contract.HouseForSale(houseForSale);
            houseNotForSaleEvent = contract.HouseNotForSale(houseNotForSale);
            houseBidPlacedEvent = contract.BidPlaced(houseBidPlaced);
            houseBidRejectedEvent = contract.BidRejected(houseBidRejected);
            houseBidAcceptedEvent = contract.BidAccepted(houseBidAccepted);
            houseBidExistsEvent = contract.BidExists(houseBidExists);
            houseInspectedEvent = contract.Inspected(houseInspected);
            houseMorgagePlacedEvent = contract.MortgageApplied(houseMortgagePlaced);
            houseMortgageApprovedEvent = contract.MortgageApproved(houseMortgageApproved);
            houseMortgageRejectedEvent = contract.MortgageRejected(houseMortgageRejected);
            houseMortgageHigherThanBidEvent = contract.MortgageHigherThanBid(houseMortgageHigherThanBid);
            houseSoldEvent = contract.HouseSold(houseSold);
        }
    });
}

function houseForSale(err, event) {
    var houseAddress = web3.toUtf8(event.args.stAddress);
    var askingPrice = event.args.askingPrice;
    document.getElementById('status').innerText = "House @ address " + houseAddress + " is on the market for a price of $" + askingPrice;
    document.getElementById('houseNotForSale').style.visibility = 'visible';
    document.getElementById('houseForSale_frm').style.visibility = 'hidden';
    document.getElementById('housePriceChange_frm').style.visibility = 'visible';
    document.getElementById('houseBid_frm').style.visibility = 'visible';
    if (!contract.inspected()) {
        document.getElementById('houseState_frm').style.visibility = 'visible';
    } else {
        document.getElementById('houseState_frm').style.visibility = 'hidden';
    }
}

function housePriceChanged(err, event) {
    var houseAddress = web3.toUtf8(event.args.stAddress);
    var askingPrice = event.args.askingPrice;
    document.getElementById('status').innerText = "House @ address " + houseAddress + " has a new price of $" + askingPrice;
}

function houseNotForSale(err, event) {
    var houseAddress = web3.toUtf8(event.args.stAddress);
    document.getElementById('status').innerText = "House @ address " + houseAddress + " has been removed from the market";
    document.getElementById('houseForSale_frm').style.visibility = "visible";
    document.getElementById('housePriceChange_frm').style.visibility = "hidden";
    document.getElementById('houseNotForSale').style.visibility = "hidden";
    document.getElementById('houseBid_frm').style.visibility = "hidden";
}

function houseBidPlaced(err, event) {
    var bidder = event.args.bidder;
    var biddingPrice = event.args.biddingPrice;
    var houseAddress = document.getElementById('houseAddress').value;

    document.getElementById('status').innerText = "Bidder " + bidder + " has placed a bid with a price of $" + biddingPrice + " on house @ address " + houseAddress;

    document.getElementById('houseNotForSale').style.visibility = 'hidden';
    document.getElementById('houseForSale_frm').style.visibility = 'hidden';
    document.getElementById('housePriceChange_frm').style.visibility = 'hidden';
    document.getElementById('houseBid_frm').style.visibility = 'hidden';
    document.getElementById('houseBidButtons_tbl').style.visibility = 'visible';
    document.getElementById('houseBidAccept_btn').style.visibility = 'visible';
    document.getElementById('houseBidReject_btn').style.visibility = 'visible';
    document.getElementById('houseListBids_btn').style.visibility = 'hidden';
}

function houseBidAccepted(err, event) {
    var bidder = event.args.bidder;
    var biddingPrice = event.args.biddingPrice;
    var houseAddress = document.getElementById('houseAddress').value;

    document.getElementById('status').innerText = "Bid from bidder " + bidder + " with a price of $" + biddingPrice + " has been accepted";

    document.getElementById('houseNotForSale').style.visibility = 'hidden';
    document.getElementById('houseForSale_frm').style.visibility = 'hidden';
    document.getElementById('housePriceChange_frm').style.visibility = 'hidden';
    document.getElementById('houseBid_frm').style.visibility = 'hidden';
    document.getElementById('houseBidButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseBidAccept_btn').style.visibility = 'hidden';
    document.getElementById('houseBidReject_btn').style.visibility = 'hidden';
    document.getElementById('houseListBids_btn').style.visibility = 'hidden';
    document.getElementById('houseMortgage_frm').style.visibility = 'visible';
    document.getElementById('houseMortgageValue').style.visibility = 'visible';
    document.getElementById('lenderAddress').style.visibility = 'visible';
    document.getElementById('houseMortgageApply_btn').style.visibility = 'visible';
    document.getElementById('houseMortgageButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseAcceptMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgage_btn').style.visibility = 'hidden';
}

function houseBidExists(err, event) {
    var bidder = event.args.bidder;
    var biddingPrice = event.args.biddingPrice;
    document.getElementById('status').innerText = "Bid from bidder " + bidder + " with a price of $" + biddingPrice + " already exists";

    document.getElementById('houseNotForSale').style.visibility = 'hidden';
    document.getElementById('houseForSale_frm').style.visibility = 'hidden';
    document.getElementById('housePriceChange_frm').style.visibility = 'hidden';
    document.getElementById('houseBidButtons_tbl').style.visibility = 'visible';
    document.getElementById('houseBidAccept_btn').style.visibility = 'visible';
    document.getElementById('houseBidReject_btn').style.visibility = 'visible';
    document.getElementById('houseListBids_btn').style.visibility = 'visible';
}

function houseBidRejected(err, event) {
    var bidder = event.args.bidder;
    var biddingPrice = event.args.biddingPrice;
    var houseAddress = document.getElementById('houseAddress').value;

    document.getElementById('status').innerText = "Bid from bidder " + bidder + " with a price of $" + biddingPrice + " has been rejected";

    document.getElementById('houseNotForSale').style.visibility = 'hidden';
    document.getElementById('houseForSale_frm').style.visibility = 'hidden';
    document.getElementById('housePriceChange_frm').style.visibility = 'hidden';
    document.getElementById('houseBid_frm').style.visibility = 'visible';
    document.getElementById('houseBidButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseBidAccept_btn').style.visibility = 'hidden';
    document.getElementById('houseBidReject_btn').style.visibility = 'hidden';
    document.getElementById('houseListBids_btn').style.visibility = 'hidden';
    document.getElementById('houseMortgage_frm').style.visibility = 'hidden';
    document.getElementById('houseMortgageButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseMortgageValue').style.visibility = 'hidden';
    document.getElementById('lenderAddress').style.visibility = 'hidden';
    document.getElementById('houseMortgageApply_btn').style.visibility = 'hidden';
    document.getElementById('houseAcceptMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgageAppraisal_btn').style.visibility = 'hidden';
}

function houseListBids(err, event) {
    var bidder = event.args.bidder;
    var biddingPrice = event.args.biddingPrice;
    var houseAddress = document.getElementById('houseAddress').value;
    var msg = "Here are the list of bids on house at address " + houseAddress + "\n";

    document.getElementById('status').innerText = msg;
}

function houseInspected(err, event) {
    var inspector = event.args.inspector;
    var state = event.args.status;
    var houseAddress = document.getElementById('houseAddress').value;

    document.getElementById('status').innerText = "House at address " + houseAddress + " has been inspected by " + inspector + " and its state is " + state;
    document.getElementById('houseState_frm').style.visibility = 'hidden';
}


function houseMortgagePlaced(err, event) {
    var lender = event.args.lender;
    var mortgageValue = event.args.value;
    var houseAddress = document.getElementById('houseAddress').value;
    var bidderAddress = document.getElementById('bidderAddress').value;

    document.getElementById('status').innerText = "Bidder " + bidderAddress + " has applied for a mortgage of $" + mortgageValue + " to lender " + lender + " for house @ address " + houseAddress;

    document.getElementById('houseNotForSale').style.visibility = 'hidden';
    document.getElementById('houseForSale_frm').style.visibility = 'hidden';
    document.getElementById('housePriceChange_frm').style.visibility = 'hidden';
    document.getElementById('houseBid_frm').style.visibility = 'hidden';
    document.getElementById('houseBidButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseBidAccept_btn').style.visibility = 'hidden';
    document.getElementById('houseBidReject_btn').style.visibility = 'hidden';
    document.getElementById('houseListBids_btn').style.visibility = 'hidden';
    document.getElementById('houseMortgage_frm').style.visibility = 'hidden';
    document.getElementById('houseMortgageValue').style.visibility = 'hidden';
    document.getElementById('lenderAddress').style.visibility = 'hidden';
    document.getElementById('houseMortgageApply_btn').style.visibility = 'hidden';
    document.getElementById('houseMortgageButtons_tbl').style.visibility = 'visible';
    document.getElementById('houseAcceptMortgage_btn').style.visibility = 'visible';
    document.getElementById('houseRejectMortgage_btn').style.visibility = 'visible';
    document.getElementById('houseRejectMortgageAppraisal_btn').style.visibility = 'visible';

}

function houseMortgageApproved(err, event) {
    var lender = event.args.lender;
    var value = event.args.value;
    var houseAddress = document.getElementById('houseAddress').value;

    document.getElementById('houseNotForSale').style.visibility = 'hidden';
    document.getElementById('houseForSale_frm').style.visibility = 'hidden';
    document.getElementById('housePriceChange_frm').style.visibility = 'hidden';
    document.getElementById('houseBidButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseBidAccept_btn').style.visibility = 'hidden';
    document.getElementById('houseBidReject_btn').style.visibility = 'hidden';
    document.getElementById('houseListBids_btn').style.visibility = 'hidden';
    document.getElementById('houseMortgageButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseMortgageValue').style.visibility = 'hidden';
    document.getElementById('lenderAddress').style.visibility = 'hidden';
    document.getElementById('houseMortgageApply_btn').style.visibility = 'hidden';
    document.getElementById('houseAcceptMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgageAppraisal_btn').style.visibility = 'hidden';
    document.getElementById('status').innerText = "Mortgage for house at address " + houseAddress + " has been accepted";

    var houseInspected = contract.inspected();
    console.log("house Inspected = " + houseInspected);
    if (contract.inspected() == 1) {
        contract.sold();
    } else {
        document.getElementById('houseState_frm').style.visibility = 'visible';
        document.getElementById('houseMortgageButtons_tbl').style.visibility = 'hidden';
        document.getElementById('houseMortgageValue').style.visibility = 'hidden';
        document.getElementById('lenderAddress').style.visibility = 'hidden';
        document.getElementById('houseMortgageApply_btn').style.visibility = 'hidden';
        document.getElementById('status').innerText = "Before the house can be sold, it needs to be inspected, please update the inspection form";
    }
}

function houseMortgageRejected(err, event) {
    var lender = event.args.lender;
    var value = event.args.value;
    var houseAddress = document.getElementById('houseAddress').value;

    document.getElementById('houseNotForSale').style.visibility = 'hidden';
    document.getElementById('houseForSale_frm').style.visibility = 'hidden';
    document.getElementById('housePriceChange_frm').style.visibility = 'hidden';
    document.getElementById('houseBidButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseBidAccept_btn').style.visibility = 'hidden';
    document.getElementById('houseBidReject_btn').style.visibility = 'hidden';
    document.getElementById('houseListBids_btn').style.visibility = 'hidden';
    document.getElementById('houseMortgage_frm').style.visibility = 'visible';
    document.getElementById('houseMortgageValue').style.visibility = 'visible';
    document.getElementById('lenderAddress').style.visibility = 'visible';
    document.getElementById('houseMortgageApply_btn').style.visibility = 'hidden';
    document.getElementById('houseAppraisal_frm').style.visibility = 'hidden';
    document.getElementById('houseMortgageButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseAcceptMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgageAppraisal_btn').style.visibility = 'hidden';
    document.getElementById('status').innerText = "Mortgage for house at address " + houseAddress + " has been rejected";

}

function houseMortgageHigherThanBid(err, event) {
    var lender = event.args.lender;
    var value = event.args.value;
    var houseAddress = document.getElementById('houseAddress').value;

    document.getElementById('houseNotForSale').style.visibility = 'hidden';
    document.getElementById('houseForSale_frm').style.visibility = 'hidden';
    document.getElementById('housePriceChange_frm').style.visibility = 'hidden';
    document.getElementById('houseBidButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseBidAccept_btn').style.visibility = 'hidden';
    document.getElementById('houseBidReject_btn').style.visibility = 'hidden';
    document.getElementById('houseListBids_btn').style.visibility = 'hidden';
    document.getElementById('houseMortgage_frm').style.visibility = 'visible';
    document.getElementById('houseMortgageValue').style.visibility = 'visible';
    document.getElementById('lenderAddress').style.visibility = 'visible';
    document.getElementById('houseMortgageApply_btn').style.visibility = 'visible';
    document.getElementById('houseMortgageButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseAcceptMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgageAppraisal_btn').style.visibility = 'hidden';
    document.getElementById('status').innerText = "Mortgage for house at address " + houseAddress + " is equal or higher than the bid value, rejecting it...";

}

function houseSold(err, event) {
    var buyer = event.args.buyer;
    var price = event.args.price;
    var houseAddress = document.getElementById('houseAddress').value;

    document.getElementById('status').innerText = "House at address " + houseAddress + " has been sold to " + buyer + " for the price of $" + price;
    document.getElementById('houseNotForSale').style.visibility = 'hidden';
    document.getElementById('houseForSale_frm').style.visibility = 'visible';
    document.getElementById('housePriceChange_frm').style.visibility = 'hidden';
    document.getElementById('houseBidButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseBidAccept_btn').style.visibility = 'hidden';
    document.getElementById('houseBidReject_btn').style.visibility = 'hidden';
    document.getElementById('houseListBids_btn').style.visibility = 'hidden';
    document.getElementById('houseMortgage_frm').style.visibility = 'hidden';
    document.getElementById('houseMortgageValue').style.visibility = 'hidden';
    document.getElementById('lenderAddress').style.visibility = 'hidden';
    document.getElementById('houseMortgageApply_btn').style.visibility = 'hidden';
    document.getElementById('houseMortgageButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseAcceptMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgageAppraisal_btn').style.visibility = 'hidden';
    document.getElementById('houseState_frm').style.visibility = 'visible';
}


//        var update = function (err, x) {
//            document.getElementById('result').textContent = JSON.stringify(x, null, 2);
//        };


function callHouseForSale(c,ap,s) {
	queryContract = web3.eth.contract(abi);
	contract = queryContract.at(c);
	web3.personal.unlockAccount(s, "password01");
	contract.forSale(ap, {from: s});
}

function callHousePriceChange() {
    var houseNewPrice = document.getElementById('houseNewPrice').value;
    document.getElementById('status').innerText = "Waiting for the transaction to change the price of the house to be mined...";
    web3.personal.unlockAccount(web3.eth.accounts[0], "Molodoi1", 1000);
    contract.changePrice(houseNewPrice);
}

function callHouseNotForSale(c,s) {
	queryContract = web3.eth.contract(abi);
	contract = queryContract.at(c);
	web3.personal.unlockAccount(s, "password01");
	contract.notForSale({from: s});
}

function callHousePlaceBid(c,b,u) {
	queryContract = web3.eth.contract(abi);
	contract = queryContract.at(c);
	// Now that a bid has been placed, the price of the house can't be changed
	// and the house can't be removed from the market
	web3.personal.unlockAccount(u, "password01", 1000);
	contract.placeBid(b, { from: u, gas: maxGas});
}

function callHouseAcceptBid(c,u) {
	queryContract = web3.eth.contract(abi);
	contract = queryContract.at(c);
	address="0xbc006b353770becc7fdecfd11eff9633a3ea651f"
	web3.personal.unlockAccount(address, "password01");
	contract.acceptBid(u,{from: address, gas: maxGas});
}

function callHouseRejectBid(c,b,u) {
	queryContract = web3.eth.contract(abi);
	contract = queryContract.at(c);
	address="0xbc006b353770becc7fdecfd11eff9633a3ea651f"
	web3.personal.unlockAccount(address, "password01");
	contract.rejectBid(u, b, {from: address, gas: maxGas});
}

function callHouseListBids() {
    var houseAddress = document.getElementById('houseAddress').value;
    document.getElementById('status').innerText = "Waiting to get the list of all bids on house at address " + houseAddress;


}

function callHouseApplyMortgage(c,a) {
	queryContract = web3.eth.contract(abi);
	contract = queryContract.at(c);
	l="0xf82335bf229a2eeee898108125937b34eaddc457"

    	web3.personal.unlockAccount(l, "password01", 1000);
    	contract.applyMortgage(a, { from: l, gas: maxGas});
}

function callAcceptMortgage() {
    var houseAddress = document.getElementById('houseAddress').value;
    var lenderAddress = document.getElementById('lenderAddress').value;
    var mortgageValue = document.getElementById('houseMortgageValue').value;

    document.getElementById('status').innerText = "Waiting for the transaction to accept the mortgage from " + lenderAddress + " for a value of $" + mortgageValue;
    web3.personal.unlockAccount(lenderAddress, "password01", 1000);
    contract.approveMortgage(mortgageValue);
}

function callRejectMortgage() {
    var houseAddress = document.getElementById('houseAddress').value;
    var lenderAddress = document.getElementById('lenderAddress').value;
    var mortgageValue = document.getElementById('houseMortgageValue').value;

    document.getElementById('status').innerText = "Waiting for the transaction to accept the mortgage from " + lenderAddress + " for a value of $" + mortgageValue;
    web3.personal.unlockAccount(lenderAddress, "password01", 1000);
    contract.rejectMortgage(mortgageValue);
}

function callHouseAppraisalYes() {
    var houseAddress = document.getElementById('houseAddress').value;
    var bidder = document.getElementById('bidderAddress').value;

    document.getElementById('status').innerText = "Bidder " + bidder + " accepts his mortgage despite the mortgage value being higher than the appraised value of the house at address " + houseAddress;

    document.getElementById('houseNotForSale').style.visibility = 'hidden';
    document.getElementById('houseForSale_frm').style.visibility = 'hidden';
    document.getElementById('housePriceChange_frm').style.visibility = 'hidden';
    document.getElementById('houseBid_frm').style.visibility = 'hidden';
    document.getElementById('houseBidButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseBidAccept_btn').style.visibility = 'hidden';
    document.getElementById('houseBidReject_btn').style.visibility = 'hidden';
    document.getElementById('houseListBids_btn').style.visibility = 'hidden';
    document.getElementById('houseMortgage_frm').style.visibility = 'hidden';
    document.getElementById('houseMortgageButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseMortgageValue').style.visibility = 'hidden';
    document.getElementById('lenderAddress').style.visibility = 'hidden';
    document.getElementById('houseMortgageApply_btn').style.visibility = 'hidden';
    document.getElementById('houseAppraisal_frm').style.visibility = 'hidden';
    document.getElementById('houseAcceptMortgage_btn').style.visibility = 'visible';
    document.getElementById('houseRejectMortgage_btn').style.visibility = 'visible';
    document.getElementById('houseRejectMortgageAppraisal_btn').style.visibility = 'visible';
}

function callHouseAppraisalNo() {
    var houseAddress = document.getElementById('houseAddress').value;
    var bidder = document.getElementById('bidderAddress').value;

    document.getElementById('status').innerText = "Bidder " + bidder + " needs to apply for a new mortgage";

    document.getElementById('houseNotForSale').style.visibility = 'hidden';
    document.getElementById('houseForSale_frm').style.visibility = 'hidden';
    document.getElementById('housePriceChange_frm').style.visibility = 'hidden';
    document.getElementById('houseBid_frm').style.visibility = 'hidden';
    document.getElementById('houseBidButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseBidAccept_btn').style.visibility = 'hidden';
    document.getElementById('houseBidReject_btn').style.visibility = 'hidden';
    document.getElementById('houseListBids_btn').style.visibility = 'hidden';
    document.getElementById('houseMortgage_frm').style.visibility = 'visible';
    document.getElementById('houseMortgageButtons_tbl').style.visibility = 'visible';
    document.getElementById('houseMortgageValue').style.visibility = 'visible';
    document.getElementById('lenderAddress').style.visibility = 'visible';
    document.getElementById('houseMortgageApply_btn').style.visibility = 'visible';
    document.getElementById('houseAppraisal_frm').style.visibility = 'hidden';
    document.getElementById('houseAcceptMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgageAppraisal_btn').style.visibility = 'hidden';
}

function callRejectMortgageAppraisal() {
    var houseAddress = document.getElementById('houseAddress').value;

    document.getElementById('houseNotForSale').style.visibility = 'hidden';
    document.getElementById('houseForSale_frm').style.visibility = 'hidden';
    document.getElementById('housePriceChange_frm').style.visibility = 'hidden';
    document.getElementById('houseBidButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseBidAccept_btn').style.visibility = 'hidden';
    document.getElementById('houseBidReject_btn').style.visibility = 'hidden';
    document.getElementById('houseListBids_btn').style.visibility = 'hidden';
    document.getElementById('houseMortgage_frm').style.visibility = 'hidden';
    document.getElementById('houseMortgageValue').style.visibility = 'hidden';
    document.getElementById('lenderAddress').style.visibility = 'hidden';
    document.getElementById('houseMortgageApply_btn').style.visibility = 'hidden';
    document.getElementById('houseAppraisal_frm').style.visibility = 'visible';
    document.getElementById('houseMortgageButtons_tbl').style.visibility = 'hidden';
    document.getElementById('houseAcceptMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgage_btn').style.visibility = 'hidden';
    document.getElementById('houseRejectMortgageAppraisal_btn').style.visibility = 'hidden';
    document.getElementById('status').innerText = "Mortgage for house at address " + houseAddress + " has been rejected, because the requested mortgage value is higher than the appraised value of the house";

}

function callHouseAppraisal(c,a) {
	queryContract = web3.eth.contract(abi);
	contract = queryContract.at(c);
	address="0x863c91fbf1b60de5d63f199a5e1ec4eec96b3857"
    	web3.personal.unlockAccount(address, "password01", 1000);
    	contract.setAppraisalValue(a, { from: address, gas: maxGas});
}

function callHouseInspected() {
    var houseState;
    var radios = document.getElementsByName('houseState');
    var inspectorAddress = document.getElementById('inspectorAddress').value;
    for (var i = 0, length = radios.length; i < length; i++) {
        if (radios[i].checked) {
            houseState = radios[i].value;
        }
    }
    document.getElementById('status').innerText = "Waiting for the transaction for the house inspection...";
    web3.personal.unlockAccount(inspectorAddress, "password01", 1000);
    contract.inspectionStatus(houseState, {
        from: inspectorAddress
    });
}
