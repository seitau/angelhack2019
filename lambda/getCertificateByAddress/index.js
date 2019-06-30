const infuraEndpoint = "https://ropsten.infura.io/v3" + process.env.INFURA_ACCESS_TOKEN;
const disasterTokenContractAddress = "0xdc7414410f683472553316Dadcb2c8763e07De8D";
const Web3 = require('web3');
const web3 = new Web3(infuraEndpoint);
const balanceOfMethodId = "0x70a08231";
const getCertificateMethodId = "0x70a08231";

exports.handler = async (event, context, callback) => {
    let body = event.body;
    console.log(body);
    console.log(JSON.parse(body));
    body = JSON.parse(body);
    
    console.log(body.address);
    if(!body.hasOwnProperty("address")) {
        console.error('Address property must be specified in the request');
        const response = {
            message: 'Address property must be specified in the request'
        }
        return response;
    } 
    try {
        const balanceOfTxData = balanceOfMethodId + web3.eth.abi.encodeParameter('address', body.address).slice(2);
        const balanceHex = await web3.eth.call({
            to: disasterTokenContractAddress,
            data: balanceOfTxData
        });
        const balance = web3.utils.hexToNumber(balanceHex);
        console.log("balance: " + balance);

        if(balance === 0) {
            console.error('No certificate is issued to this address');
            const response = {
                message: 'No certificate is issued to this address'
            }
            return response;
        }

        let certificates = new Array();
        for(let i = 0; i < balance-1; i++) {
            const id = web3.utils.keccak256(body.address + (balance-i));
            console.log("id: " + id);

            const getCertificateTxData = getCertificateMethodId + web3.eth.abi.encodeParameter('uint256', id).slice(2);
            const certificate = await web3.eth.call({
                to: disasterTokenContractAddress,
                data: getCertificateTxData
            });
            console.log("certificate: " + certificate);
            certificates.push(certificates);
        }
        callback(null, certificates);
    } catch(err) {
        console.error(err);
        const response = {
            message: err.message
        }
        return response;
    }
}
