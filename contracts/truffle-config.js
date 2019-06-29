require('babel-polyfill')
const HDWalletProvider = require("truffle-hdwallet-provider");
const mnemonic = process.env.ROPSTEN_MNEMONIC_1;
const ropsten_endpoint = process.env.GINCO_ROPSTEN_ENDPOINT;
const infura_endpoint = "https://ropsten.infura.io/v3" + process.env.INFURA_ACCESS_TOKEN;

module.exports = {
    compilers: {
        solc: {
            version: "^0.5.6",
            settings: {
                optimizer: {
                    enabled: true,
                    runs: 1000, // Optimize for how many times you intend to run the code
                },
                evmVersion: "constantinople" // Default: "byzantium"
            }
        }
    },
    networks: {
        ropsten: {
            provider: () => new HDWalletProvider(mnemonic, ropsten_endpoint, 0, 5),
            network_id: "3",
            websockets: true,
        },
        infura: {
            provider: () => new HDWalletProvider(mnemonic, infura_endpoint, 0, 5),
            network_id: "3",
        },
    },
    license: "MIT"
};
