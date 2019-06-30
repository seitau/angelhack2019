const infuraEndpoint = "https://ropsten.infura.io/v3" + process.env.INFURA_ACCESS_TOKEN;
const disasterTokenContractAddress = "0xeec79725e5665e9578fbb05a9489ccfaa0542e08";
const Web3 = require('web3');
const web3 = new Web3(infuraEndpoint);
const balanceOfMethodId = "0x70a08231";
const getCertificateMethodId = "0x51640fee";

exports.handler = async (event, context, callback) => {
    let body = JSON.parse(event.body);
    //let body = event.body
    
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
        for(let i = 0; i < balance; i++) {
            const seed = `${body.address.toLowerCase()}`;
            console.log("keccak seed: " + seed);
            const id = web3.utils.keccak256(seed);
            console.log("id: " + id);
            console.log(web3.utils.keccak256(`${body.address.toLowerCase()}`));

            const getCertificateTxData = getCertificateMethodId + web3.eth.abi.encodeParameter('uint256', id).slice(2);
            console.log('txData: ' + getCertificateTxData);

            const certificateAbi = await web3.eth.call({
                to: disasterTokenContractAddress,
                data: getCertificateTxData
            });
            const certificate = web3.eth.abi.decodeParameters(
                ['string', 'string', 'string', 'string', 'string'],
                certificateAbi
            )
            console.log("certificate: " + JSON.stringify(certificate));
            certificates.push({
                title: certificates[0],
                seriousness: certificates[1],
                imageRef: certificates[2],
                category: certificates[3],
                date: certificates[4],
                estimatedMoney: certificates[5],
            });
            return {
                statusCode: 200,
                body: JSON.stringify({
                    message: certificate,
                    input: event,
                }, null, 2),
            };
        }
    } catch(err) {
        console.error(err);
        const response = {
            message: err.message
        }
        return response;
    }
}
