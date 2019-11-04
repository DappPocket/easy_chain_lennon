const web3utils = require('web3-utils');
import _ from 'lodash';
import moment from 'moment';

export default {
  hexToUtf8: (state) => (input) => {
    return web3utils.hexToUtf8(input)
  },
  utf8ToHex: (state) => (input) => {
    return web3utils.utf8ToHex(input)
  },
  etherscanLink: (state) => (hash) => {
    return `${state.ether_net_domain}/tx/${hash}`
  },
  getComments: (state) => {
    let comments = _.orderBy(state.comments, ['unix_time'], ['desc'])
    comments = _.uniqBy(comments, (e) => e["hash"]);
    return comments
  },
  pickColorByNoce: (state) => (nonce, block_number) => {
    const cbase = nonce + block_number
    if (cbase % 4 == 0) {
      return "is-four"
    } else if (cbase % 3 == 0) {
      return "is-three"
    } else if (cbase % 2 == 0) {
      return "is-two"
    } else if (cbase == 0 || cbase % 1 == 0) {
      return "is-one"
    }
  },
  convertData: (state, getters) => (o, msg) => {
    if(!msg){
      msg = getters.hexToUtf8(o.input)
    }
    const timebase = moment(o.timestamp, 'YYYY-MM-DDTH:mm a')
    return {
      hash: o.hash,
      message: msg,
      nonce: o.nonce,
      bg_color: getters.pickColorByNoce(o.nonce, o.block_number),
      block_number: o.block_number,
      confirmed_status: +o.block_number > 0,
      unix_time: timebase.unix(),
      timestamp: timebase.format('YYYY/MM/DD hh:mm a'),
    }
  },
  getHashList: (state) => {
    return _.map(state.comments, (e) => e.hash)
  },
  nonComfirmedList: (state) => {
    return _.map(state.comments, (e) => e.block_number === 0)
  }
}
