const Web3 = require('web3');
const INFURA_ENDPOINT = "https://ropsten.infura.io/v3" + process.env.INFURA_ACCESS_TOKEN;
const web3 = new Web3(INFURA_ENDPOINT);

exports.handler = async function(event) {
  const promise = new Promise(function(resolve, reject) {
    https.get(url, (res) => {
        resolve(res.statusCode)
      }).on('error', (e) => {
        reject(Error(e))
      })
    })
  return promise
}
