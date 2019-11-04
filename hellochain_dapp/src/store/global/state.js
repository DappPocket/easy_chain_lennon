export default {
  // ether_net_domain: "https://etherscan.io",
  ether_net_domain: process.env.VUE_APP_ETHERSCAN_HOST,
  to_address: process.env.VUE_APP_TARGET_ADDR,
  web3util: undefined,
  comments: [],
  currentpage: 1,
  per_page: 6,
  // display popout messages
  blockwsMessage: undefined,
}
