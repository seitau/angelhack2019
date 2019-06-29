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

        console.log("deploying DisasterToken");
        const name = "DisasterToken";
        const symbol = "DST";
        const token = await DisasterToken.new(name, symbol, { from: accounts[0] })
            .once('transactionHash', (hash) => {
                console.log('transactionHash: ' + hash);
            })
            .once('receipt', (receipt) => {
                console.log('status: ' + receipt.status);
                logEvents(receipt);
            });

        const balance = await token.balanceOf(account1.address);
        const id = web3.utils.keccak256(account1.address + balance);
        console.log("id: " + id);

        const title = "earthquake";
        const seriousness = "A";
        const category = "earthquake";
        const date = "earthquake";
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
