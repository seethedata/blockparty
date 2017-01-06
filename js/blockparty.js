var Web3 = require('web3');
var web3 = new Web3();
web3.setProvider(new web3.providers.HttpProvider('http://54.245.138.237:8545'));

var code = "60606040526040516020806112d2833981016040528080519060200190919050505b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff02191690836c010000000000000000000000009081020402179055505b33600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff02191690836c01000000000000000000000000908102040217905550806004600050819055506000600260006101000a81548160ff02191690837f01000000000000000000000000000000000000000000000000000000000000009081020402179055505b506111dd806100f56000396000f36060604052361561013a576000357c01000000000000000000000000000000000000000000000000000000009004806302c7e7af1461013f5780630ba062b5146101535780631a7ac7261461016757806320b08dc21461018f5780632831f2f4146101b95780633c8c7714146101d657806341c0e1b514610200578063446b13f41461021457806344ec56401461023157806347d905761461025b57806364bb6270146102875780637150d8ae146102c557806372e92e2214610303578063809d79021461032b5780638175e54314610351578063955a5a761461036e5780639979ef4514610394578063a2b40d19146103b1578063bcdea56e146103ce578063c19d93fb146103eb578063c88c721114610413578063d9fa1c5d14610430578063f496d88214610458578063f746905b146104965761013a565b610002565b346100025761015160048050506104be565b005b3461000257610165600480505061055d565b005b3461000257610179600480505061061b565b6040518082815260200191505060405180910390f35b34610002576101a16004805050610624565b60405180821515815260200191505060405180910390f35b34610002576101d46004808035906020019091905050610637565b005b34610002576101e860048050506106fb565b60405180821515815260200191505060405180910390f35b34610002576102126004805050610717565b005b346100025761022f6004808035906020019091905050610813565b005b346100025761024360048050506108e3565b60405180821515815260200191505060405180910390f35b346100025761026d60048050506108f6565b604051808260001916815260200191505060405180910390f35b346100025761029960048050506108ff565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34610002576102d76004805050610925565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3461000257610315600480505061094b565b6040518082815260200191505060405180910390f35b346100025761034f6004808035906020019091908035906020019091905050610954565b005b346100025761036c60048080359060200190919050506109e8565b005b34610002576103926004808035906020019091908035906020019091905050610bdd565b005b34610002576103af6004808035906020019091905050610d5b565b005b34610002576103cc6004808035906020019091905050610f49565b005b34610002576103e96004808035906020019091905050610fcf565b005b34610002576103fd60048050506110a0565b6040518082815260200191505060405180910390f35b346100025761042e60048080359060200190919050506110b3565b005b346100025761044260048050506111a5565b6040518082815260200191505060405180910390f35b346100025761046a60048050506111ae565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34610002576104a860048050506111d4565b6040518082815260200191505060405180910390f35b600860149054906101000a900460ff1680156104e65750600660159054906101000a900460ff165b1561055a57600360005054600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f510ab7a94b701fa3bd6d2ff90fd1094a7d19c89d6d2aedd62176f3ac4313a4d460405180905060405180910390a35b5b565b6000600260006101000a81548160ff02191690837f0100000000000000000000000000000000000000000000000000000000000000908102040217905550600060016000508190555060046000505460001916600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f5ea8dd485da0261bde5fc64a9c4fee1392475c0beaaf650d785f0ce7eef05a4360405180905060405180910390a35b565b60016000505481565b600260009054906101000a900460ff1681565b6001600260006101000a81548160ff02191690837f01000000000000000000000000000000000000000000000000000000000000009081020402179055508060016000508190555060046000505460001916600160005054600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f2a938dc7ef6cf47e9419d8487fa433c934396b54bc95e115ed57d319c0325e6f60405180905060405180910390a45b50565b6000600860149054906101000a900460ff169050610714565b90565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561081057600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f4b0bc4f25f8d0b92d2e12b686ba96cd75e4e69325e6cf7b1f3119d14eaf2cbdf60405180905060405180910390a2600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b33600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff02191690836c010000000000000000000000009081020402179055506001600860146101000a81548160ff02191690837f010000000000000000000000000000000000000000000000000000000000000090810204021790555080600960005081905550803373ffffffffffffffffffffffffffffffffffffffff167f3510d209ed8ef435dd4023b99b4b4e4f8e7310278825a2037fd2566adcc959ef60405180905060405180910390a35b50565b600660159054906101000a900460ff1681565b60046000505481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600b6000505481565b6000600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff02191690836c010000000000000000000000009081020402179055506000600b60005081905550808273ffffffffffffffffffffffffffffffffffffffff167e363fe256a1f2c23b39e0cfe65d691c725e83e534d368d718337b9d46b705ae60405180905060405180910390a35b5050565b6000811415610a34576000600660146101000a81548160ff02191690837f0100000000000000000000000000000000000000000000000000000000000000908102040217905550610b17565b6001811415610a80576001600660146101000a81548160ff02191690837f0100000000000000000000000000000000000000000000000000000000000000908102040217905550610b16565b6002811415610acc576002600660146101000a81548160ff02191690837f0100000000000000000000000000000000000000000000000000000000000000908102040217905550610b15565b6003811415610b14576003600660146101000a81548160ff02191690837f01000000000000000000000000000000000000000000000000000000000000009081020402179055505b5b5b5b33600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff02191690836c010000000000000000000000009081020402179055506001600660156101000a81548160ff02191690837f0100000000000000000000000000000000000000000000000000000000000000908102040217905550803373ffffffffffffffffffffffffffffffffffffffff167f7fd571334fe4d478d78872a01f7d126d0fe95e28d1e81c46ed18161a6c89f8ad60405180905060405180910390a35b50565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610d0e576000600b60005054141515610cc1578060036000508190555081600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff02191690836c01000000000000000000000000908102040217905550808273ffffffffffffffffffffffffffffffffffffffff167f68113ca7b748f7d17637d79856900bc9336b88166e726bc9b55c62a39408b40760405180905060405180910390a3610d09565b808273ffffffffffffffffffffffffffffffffffffffff167f15349387d7f320310878a8474c7e50498ab4b154479bc9c888b1797ba9c2d82e60405180905060405180910390a35b610d56565b808273ffffffffffffffffffffffffffffffffffffffff167f15349387d7f320310878a8474c7e50498ab4b154479bc9c888b1797ba9c2d82e60405180905060405180910390a35b5b5050565b600260009054906101000a900460ff1615610f45573373ffffffffffffffffffffffffffffffffffffffff16600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610e5a5733600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff02191690836c0100000000000000000000000090810204021790555080600b60005081905550803373ffffffffffffffffffffffffffffffffffffffff167f3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d60405180905060405180910390a3610f44565b80600b60005054141515610efb5733600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff02191690836c0100000000000000000000000090810204021790555080600b60005081905550803373ffffffffffffffffffffffffffffffffffffffff167f3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d60405180905060405180910390a3610f43565b803373ffffffffffffffffffffffffffffffffffffffff167f854a7fa16e4b8b5d5b5405473dc7f0b09bf4895e71d969e01053a7e35519992f60405180905060405180910390a35b5b5b5b50565b8060016000508190555060046000505460001916600160005054600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fe4cc3e7461adbcd283d6b2dcfe79c630f880de8f53b5bb6fbf01744d11804d0a60405180905060405180910390a45b50565b33600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff02191690836c010000000000000000000000009081020402179055506000600860146101000a81548160ff02191690837f01000000000000000000000000000000000000000000000000000000000000009081020402179055506000600960005081905550803373ffffffffffffffffffffffffffffffffffffffff167f8c736f7d8b2be59ea96ee5d817b3faee2e1a60ad2756eff1b8c098b5a09ff47260405180905060405180910390a35b50565b600660149054906101000a900460ff1681565b80600960005081905550600b600050548110156111535733600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff02191690836c01000000000000000000000000908102040217905550803373ffffffffffffffffffffffffffffffffffffffff167fccb7d77cb44936ccd10f7fabd49cdbeb2c2273d72e38cefa3dac8c8533c6d8a660405180905060405180910390a36111a1565b600b60005054813373ffffffffffffffffffffffffffffffffffffffff167f63e1e1d035eb975aeb19b4c7b9d46de5d738e75e655be000e2b3db26cd9045a360405180905060405180910390a45b5b50565b600c6000505481565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6003600050548156";
var abi = [{
    "constant": false,
    "inputs": [],
    "name": "sold",
    "outputs": [],
    "payable": false,
    "type": "function"
}, {
    "constant": false,
    "inputs": [],
    "name": "notForSale",
    "outputs": [],
    "payable": false,
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "askingPrice",
    "outputs": [{
        "name": "",
        "type": "uint256"
    }],
    "payable": false,
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "isForSale",
    "outputs": [{
        "name": "",
        "type": "bool"
    }],
    "payable": false,
    "type": "function"
}, {
    "constant": false,
    "inputs": [{
        "name": "price",
        "type": "uint256"
    }],
    "name": "forSale",
    "outputs": [],
    "payable": false,
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "checkMortgageStatus",
    "outputs": [{
        "name": "mortgageStatus",
        "type": "bool"
    }],
    "payable": false,
    "type": "function"
}, {
    "constant": false,
    "inputs": [],
    "name": "kill",
    "outputs": [],
    "payable": false,
    "type": "function"
}, {
    "constant": false,
    "inputs": [{
        "name": "value",
        "type": "uint256"
    }],
    "name": "approveMortgage",
    "outputs": [],
    "payable": false,
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "inspected",
    "outputs": [{
        "name": "",
        "type": "bool"
    }],
    "payable": false,
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "streetAddress",
    "outputs": [{
        "name": "",
        "type": "bytes32"
    }],
    "payable": false,
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "houseOwner",
    "outputs": [{
        "name": "",
        "type": "address"
    }],
    "payable": false,
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "buyer",
    "outputs": [{
        "name": "",
        "type": "address"
    }],
    "payable": false,
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "bidValue",
    "outputs": [{
        "name": "",
        "type": "uint256"
    }],
    "payable": false,
    "type": "function"
}, {
    "constant": false,
    "inputs": [{
        "name": "bidderAddress",
        "type": "address"
    }, {
        "name": "biddingPrice",
        "type": "uint256"
    }],
    "name": "rejectBid",
    "outputs": [],
    "payable": false,
    "type": "function"
}, {
    "constant": false,
    "inputs": [{
        "name": "status",
        "type": "uint256"
    }],
    "name": "inspectionStatus",
    "outputs": [],
    "payable": false,
    "type": "function"
}, {
    "constant": false,
    "inputs": [{
        "name": "bidderAddress",
        "type": "address"
    }, {
        "name": "biddingPrice",
        "type": "uint256"
    }],
    "name": "acceptBid",
    "outputs": [],
    "payable": false,
    "type": "function"
}, {
    "constant": false,
    "inputs": [{
        "name": "biddingPrice",
        "type": "uint256"
    }],
    "name": "placeBid",
    "outputs": [],
    "payable": false,
    "type": "function"
}, {
    "constant": false,
    "inputs": [{
        "name": "newPrice",
        "type": "uint256"
    }],
    "name": "changePrice",
    "outputs": [],
    "payable": false,
    "type": "function"
}, {
    "constant": false,
    "inputs": [{
        "name": "value",
        "type": "uint256"
    }],
    "name": "rejectMortgage",
    "outputs": [],
    "payable": false,
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "state",
    "outputs": [{
        "name": "",
        "type": "uint8"
    }],
    "payable": false,
    "type": "function"
}, {
    "constant": false,
    "inputs": [{
        "name": "value",
        "type": "uint256"
    }],
    "name": "applyMortgage",
    "outputs": [],
    "payable": false,
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "acceptedBid",
    "outputs": [{
        "name": "",
        "type": "uint256"
    }],
    "payable": false,
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "bidder",
    "outputs": [{
        "name": "",
        "type": "address"
    }],
    "payable": false,
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "soldPrice",
    "outputs": [{
        "name": "",
        "type": "uint256"
    }],
    "payable": false,
    "type": "function"
}, {
    "inputs": [{
        "name": "stAddress",
        "type": "bytes32"
    }],
    "type": "constructor"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "seller",
        "type": "address"
    }, {
        "indexed": true,
        "name": "askingPrice",
        "type": "uint256"
    }, {
        "indexed": true,
        "name": "stAddress",
        "type": "bytes32"
    }],
    "name": "HouseForSale",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "seller",
        "type": "address"
    }, {
        "indexed": true,
        "name": "stAddress",
        "type": "bytes32"
    }],
    "name": "HouseNotForSale",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "seller",
        "type": "address"
    }, {
        "indexed": true,
        "name": "askingPrice",
        "type": "uint256"
    }, {
        "indexed": true,
        "name": "stAddress",
        "type": "bytes32"
    }],
    "name": "PriceChanged",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "bidder",
        "type": "address"
    }, {
        "indexed": true,
        "name": "biddingPrice",
        "type": "uint256"
    }],
    "name": "BidPlaced",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "bidder",
        "type": "address"
    }, {
        "indexed": true,
        "name": "biddingPrice",
        "type": "uint256"
    }],
    "name": "BidRejected",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "bidder",
        "type": "address"
    }, {
        "indexed": true,
        "name": "biddingPrice",
        "type": "uint256"
    }],
    "name": "BidAccepted",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "bidder",
        "type": "address"
    }, {
        "indexed": true,
        "name": "biddingPrice",
        "type": "uint256"
    }],
    "name": "BidExists",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "inspector",
        "type": "address"
    }, {
        "indexed": true,
        "name": "status",
        "type": "uint256"
    }],
    "name": "Inspected",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "lender",
        "type": "address"
    }, {
        "indexed": true,
        "name": "value",
        "type": "uint256"
    }],
    "name": "MortgageApplied",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "lender",
        "type": "address"
    }, {
        "indexed": true,
        "name": "value",
        "type": "uint256"
    }],
    "name": "MortgageApproved",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "lender",
        "type": "address"
    }, {
        "indexed": true,
        "name": "value",
        "type": "uint256"
    }],
    "name": "MortgageRejected",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "buyer",
        "type": "address"
    }, {
        "indexed": true,
        "name": "price",
        "type": "uint256"
    }],
    "name": "HouseSold",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "bidderAddress",
        "type": "address"
    }, {
        "indexed": true,
        "name": "biddingPrice",
        "type": "uint256"
    }],
    "name": "WrongBid",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "sender",
        "type": "address"
    }, {
        "indexed": true,
        "name": "mortgageValue",
        "type": "uint256"
    }, {
        "indexed": true,
        "name": "biddingPrice",
        "type": "uint256"
    }],
    "name": "MortgageHigherThanBid",
    "type": "event"
}, {
    "anonymous": false,
    "inputs": [{
        "indexed": true,
        "name": "from",
        "type": "address"
    }],
    "name": "Killed",
    "type": "event"
}];

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
        gas: 4700000
    }, function(err, c) {
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


function callHouseForSale(c,ap) {
	queryContract = web3.eth.contract(abi);
	contract = queryContract.at(c);
	address="0xbc006b353770becc7fdecfd11eff9633a3ea651f"
	web3.personal.unlockAccount(address, "password01");
	contract.forSale(parseInt(ap), {from: address});
}

function callHousePriceChange() {
    var houseNewPrice = document.getElementById('houseNewPrice').value;
    document.getElementById('status').innerText = "Waiting for the transaction to change the price of the house to be mined...";
    web3.personal.unlockAccount(web3.eth.accounts[0], "Molodoi1", 1000);
    contract.changePrice(houseNewPrice);
}

function callHouseNotForSale(c) {
	queryContract = web3.eth.contract(abi);
	contract = queryContract.at(c);
	address="0xbc006b353770becc7fdecfd11eff9633a3ea651f"
	web3.personal.unlockAccount(address, "password01");
	contract.notForSale({from: address});
}

function callHousePlaceBid() {
    var houseAddress = document.getElementById('houseAddress').value;
    var bidPrice = document.getElementById('houseBidPrice').value;
    var bidderAddress = document.getElementById('bidderAddress').value;

    // Now that a bid has been placed, the price of the house can't be changed
    // and the house can't be removed from the market
    document.getElementById('status').innerText = "Waiting for the transaction to place bid on the house...";
    web3.personal.unlockAccount(bidderAddress, "password01", 1000);
    contract.placeBid(bidPrice, {
        from: bidderAddress
    });
}

function callHouseAcceptBid() {
    var houseAddress = document.getElementById('houseAddress').value;
    var bidderAddress = document.getElementById('bidderAddress').value;
    var biddingPrice = document.getElementById('houseBidPrice').value;
    document.getElementById('status').innerText = "Waiting for the transaction to accept the bid from " + bidderAddress + " on the house...";

    web3.personal.unlockAccount(web3.eth.accounts[0], "Molodoi1", 1000);
    contract.acceptBid(bidderAddress, biddingPrice);
}

function callHouseRejectBid() {
    var houseAddress = document.getElementById('houseAddress').value;
    var bidderAddress = document.getElementById('bidderAddress').value;
    var biddingPrice = document.getElementById('houseBidPrice').value;
    document.getElementById('status').innerText = "Waiting for the transaction to reject the bid from " + bidderAddress + " on the house...";

    web3.personal.unlockAccount(web3.eth.accounts[0], "Molodoi1", 1000);
    contract.rejectBid(bidderAddress, biddingPrice);
}

function callHouseListBids() {
    var houseAddress = document.getElementById('houseAddress').value;
    document.getElementById('status').innerText = "Waiting to get the list of all bids on house at address " + houseAddress;


}

function callHouseApplyMortgage() {
    var houseAddress = document.getElementById('houseAddress').value;
    var mortgageValue = document.getElementById('houseMortgageValue').value;
    var lenderAddress = document.getElementById('lenderAddress').value;

    document.getElementById('status').innerText = "Waiting for the application for a mortgage value of $" + mortgageValue + " to lender " + lenderAddress + " to be mined";
    web3.personal.unlockAccount(lenderAddress, "password01", 1000);
    contract.applyMortgage(mortgageValue, {
        from: lenderAddress
    });
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
