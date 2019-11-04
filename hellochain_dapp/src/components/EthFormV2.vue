<template lang="pug">
v-container.eth_form_container(fluid)
  v-layout(warp text-center)
    v-row(class="mb-6" justify="center")
      v-col(lg="4" mb="6" xs="12")
        .v-card-help
          v-card(:loading="cardloading")
            v-list-item
              v-list-item-content
                v-list-item-title
                  span 將想說的話寫入區塊鏈
                  img(src="/ethereum-icon.png")
                v-form(ref="form" v-model="valid" lazy-validation)
                  v-textarea(outlined :counter="50" name="input-7-4" v-model="message" :rules="messageRules" label="內容" required)
            v-card-actions
              v-btn(color="success" class="mr-4" @click="submitform" :disabled="!valid || !message") 送出
              router-link(:to="{name: 'Comments'}" v-slot="{href}")
                v-btn(:href="href" class="mr-4") 取消
  v-row
    v-col
      v-dialog(v-model='dialog' persistent='' max-width='290')
        v-card
          v-card-title.headline 請開啟 MataMask
          v-card-text
            div
              v-img(src="/maxfox.png" aspect-ratio="1" class="grey lighten-2" max-width="500" max-height="300")
            div *請確認MataMask是否正確開啟
          v-card-actions
            v-spacer
            v-btn(color='green darken-1' text='' @click='dialog = false') OK
  v-snackbar(v-model="snackbar")
    | {{notices}}
</template>

<script>
const Web3 = require('web3');
import moment from 'moment';
import {mapGetters, mapState, mapActions} from 'vuex';

export default {
  data: () => {
    return {
      dialog: true,
      web3help: undefined,
      cardloading: false,
      accounts: [],
      valid: false,
      message: undefined,
      messageRules: [
        v => !!v || 'Message is required',
        v => (v && v.length <= 50) || 'Name must be less than 50 characters',
      ],
      snackbar: false,
      notices: '',
    }
  },
  async created(){
    // https://medium.com/metamask/https-medium-com-metamask-breaking-change-injecting-web3-7722797916a8
    // Legacy dapp browsers...
    if (window.web3) {
      console.log("Legacy dapp browsers")
      try {
        web3 = new Web3(web3.currentProvider);
      //   this.accounts = await web3.eth.accounts
      } catch (error) {
        // Handle error. Likely the user rejected the login
        console.error(error)
      }
      await ethereum.enable()
      this.dialog = false
    } else {
      this.dialog = true
    }
  },
  mounted(){
    this.valid = false
    this.update_recent_transcations()
    this.$http.secured.get("/trigger_query/query").then(e => console.log).catch(e => console.error)
  },
  methods: {
    update_recent_transcations(){
      // disabled
      // this.$http.secured.get("http://localhost:3000/update_recent_transcations").then(e => console.log).catch(e => console.error)
    },
    async submitform(){
      let data = this.utf8ToHex(`${this.message}`)
      let from_addr = web3.eth.accounts[0]
      let tx_params = {
        from: from_addr,
        to: this.to_address,
        value: "0x0",
        data,
      }
      let gas = 0
      let gasPrice = 0
      let self = this
      await web3.eth.estimateGas(tx_params, async function(err, estgas) {
        console.log("estimateGas", estgas);
        gas = await estgas
        await web3.eth.getGasPrice(async function(err, gPrice) {
          console.log("gasPrice", gPrice);
          gasPrice = await gPrice;
          tx_params = await {
            gas,
            gasPrice,
            ...tx_params,
          }
          console.log(tx_params)
          await web3.eth.sendTransaction({
            ...tx_params
          }, function(err, transactionHash) {
            if (!err){
              console.log(transactionHash); // "0x7f9fade1c0d57a7af66ab4ead7c2eb7b11a91385"
              // submit hash to backend
              setTimeout(() => {
                self.$http.secured.post('/api/v1/forces_insert_one', {"hash": transactionHash, input: data})
                .then(e => console.log(e))
                .catch(console.error)
              }, 500)

              // a submmited tx on frontend. waiting for backend response
              self.snackbar = true
              self.notices = "簽名訊息送出,等待區塊確認.."
              const currentTimestamp = Math.floor(moment.now()/1000)
              self.addOneComments({
                hash: transactionHash,
                from: from_addr,
                to: self.to_address,
                value: "0x0",
                input: data,
                gas: gas,
                is_error: "0",
                message: self.message,
                nonce: -1,
                bg_color: "is-none",
                block_number: 0,
                confirmed_status: false,
                unix_time: currentTimestamp,
                timestamp: moment(moment.now()).format('YYYY-MM-DDTH:mm a'),
              })
              self.$router.push({ name: 'Comments' })
            } else {
              self.notices = err.message
              self.snackbar = true
            }
          })
        })
      })
    },
    testsb(){
      const self = this
      const currentTimestamp = Math.floor(moment.now()/1000)
      self.addOneComments({
        hash: "0x0",
        from: "0x0",
        to: self.to_address,
        value: "0x0",
        input: "dd",
        gas: "0",
        is_error: "0",
        message: "testtest",
        nonce: -1,
        bg_color: "is-none",
        block_number: 0,
        confirmed_status: false,
        unix_time: currentTimestamp,
        timestamp: moment(moment.now()).format('YYYY/MM/DD h:mm a'),
      })
      self.$router.push({ name: 'Comments' })
    },
    ...mapActions('global', ['addOneComments'])
  },
  computed: {
    ...mapGetters('global', ['hexToUtf8', 'utf8ToHex', 'etherscanLink']),
    ...mapState('global', ['to_address']),
  },
}
</script>
