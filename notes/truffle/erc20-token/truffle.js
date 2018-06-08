module.exports = {
    // See <http://truffleframework.com/docs/advanced/configuration>
    // to customize your Truffle configuration!
    solc: {
      optimizer: {
        enabled: true,
      }
    },
    networks: {
      prodwork: {
        host: "127.0.0.1",
        port: 8541,
        network_id: "*",
        address: "0x9b4eabea5d69a3c434c40f84f65282f6b4d9b232",
        gasPrice: 18000000000,
        gas: 0x47b760,
      }
    }
};