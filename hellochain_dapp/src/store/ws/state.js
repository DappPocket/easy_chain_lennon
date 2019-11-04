export default {
  wsurl: `${process.env.VUE_APP_WS_HOST}/chain_watch`,
  // wsurl: 'ws://192.168.2.104:3005/mqtt',
  conn: undefined,
  message: undefined,
  reconenct: true,
  // after 3 sec retry connection
  waitTime: 3000,
}
