const DisasterToken = artifacts.require('DisasterToken'); 
const ethutil = require('ethereumjs-util');
const ethabi = require('ethereumjs-abi');

function logEvents(receipt) {
    if(receipt.logs.length == 0) return;
    console.log("\n------------------Events------------------\n");
    for(let i = 0; i < receipt.logs.length; i++) {
        console.log(receipt.logs[i].event + ": " + JSON.stringify(receipt.logs[i].args) + "\n");
    }
    console.log("------------------------------------------\n");
}

module.exports = async (callback) => {
    try {
        const accounts = await web3.eth.getAccounts();
        const netId = await web3.eth.net.getId();

        let privKeys = [];
        //develop環境
        if(netId == 5777) {
            privKeys.push("0x29c26a0b39614c5a98bcd65a8f13039102c51d0e9e0058ff918d105202c83426");
            privKeys.push("0x30d7497ac8065286e30d2ef0abde7ad270b1621b173d5a50b236d9a2bb53844c");
            privKeys.push("0x0e70d2d612c986a50cce5fdfdac4e1c018d65b765e4ac9673eec946d3715aa99");
        } else {
            privKeys.push("0x" + web3.currentProvider.wallets[accounts[0].toLowerCase()]._privKey.toString("hex"));
            privKeys.push("0x" + web3.currentProvider.wallets[accounts[1].toLowerCase()]._privKey.toString("hex"));
        }

        console.log("setting up accounts");
        const account1 = web3.eth.accounts.privateKeyToAccount(privKeys[0]);
        const account2 = web3.eth.accounts.privateKeyToAccount(privKeys[1]);

        console.log("accounts: " + accounts);
        const token = await DisasterToken.deployed();
        console.log("token: " + token.address);

        const balance = await token.balanceOf(account1.address);
        console.log("balance: " + balance);
        const seed = `${account1.address.toLowerCase()}`;
        const id = web3.utils.keccak256(seed);
        console.log("keccak seed: " + seed);
        console.log("id: " + id);

        const title = "Earthquake";
        const seriousness = "A";
        const category = "earthquake";
        const date = "2016-04-14T10:44:00+0000";
        const mintReceipt = await token.mintWithCertificate
            .sendTransaction(
                account1.address,
                id,
                title,
                seriousness,
                category,
                date
            )
            .once('receipt', (receipt) => {
                logEvents(receipt);
            });

        const certificate = await token.getCertificate.call(id);
        console.log(certificate);
    } catch(err) {
        console.error(err);
    }
    return callback();
}
