<template lang="pug">
v-container(fluid)
  v-layout(warp text-center)
    v-row(class="mb-6" justify="center")
      v-col(lg="4" mb="6" xs="12")
        v-card(:loading="cardloading")
          v-list-item
            v-list-item-content
              v-list-item-title 將想說的話寫入區塊鏈
              v-form(ref="form" v-model="valid" lazy-validation)
                v-text-field(v-model="name" :counter="20" :rules="nameRules" label="名稱" required)
                v-textarea(outlined :counter="50" name="input-7-4" v-model="message" :rules="messageRules" label="內容" required)
          v-card-actions
            v-btn(color="success" class="mr-4" @click="submitform" :disabled="!valid || !(name && message)") 送出
            router-link(:to="{name: 'Comments'}" v-slot="{href}")
              v-btn(:href="href" class="mr-4" ) 取消
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
import Web3 from 'web310';

export default {
  data: () => {
    return {
      dialog: true,
      web3help: undefined,
      cardloading: false,
      accounts: [],
      valid: false,
      name: undefined,
      nameRules: [
        v => !!v || 'Name is required',
        v => (v && v.length <= 20) || 'Name must be less than 20 characters',
      ],
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
    // Modern dapp browsers...
    if (window.ethereum) {
      console.log("Modern dapp browsers")
      window.web3 = new Web3(ethereum);
      try {
        this.accounts = await ethereum.enable()
        // Request account access if needed
        // await ethereum.enable();
        // Acccounts now exposed
        web3.eth.sendTransaction({/* ... */});
      } catch (error) {
          // User denied account access...
      }
      this.dialog = false
    }
    // Legacy dapp browsers...
    else if (window.web3) {
      console.log("Legacy dapp browsers")
      window.web3 = new Web3(Web3.currentProvider);
      try {
        this.accounts = await ethereum.enable()

      } catch (error) {
        // Handle error. Likely the user rejected the login
        console.error(error)
      }
      this.dialog = false
    } else {
      this.dialog = true
    }
  },
  mounted(){
    this.valid = false
    this.$http.secured.get("http://localhost:3000/trigger_query/query").then(e => console.log).catch(e => console.error)
    // callback for change accounts
    // web3.currentProvider.publicConfigStore.on('update', async()=>{
    //   this.accounts = await ethereum.enable();
    // });
    this.$http.secured.get("/trigger_query/query").then(e => console.log).catch(e => console.error)
  },
  methods: {
    async submitform(){
      let data = web3.utils.utf8ToHex(`${this.name}\t${this.message}`)
      console.log(ethereum, this.accounts)
      let tx_params = {
        from: this.accounts[0],
        to: "0x99992213Adf6471e52ED09EF47B36Faf7b769600",
        value: "0x0",
        data,
      }
      let gas = 0
      await web3.eth.estimateGas(tx_params)
      .then(g => gas = g);
      let gasPrice = await web3.eth.getGasPrice()
      tx_params = {
        gas,
        gasPrice,
        ...tx_params,
      }
      console.log(tx_params)
      web3.eth.sendTransaction({
        ...tx_params
      })
      .on('transactionHash', (hash) => {
        this.snackbar = true
        this.notices = "簽名訊息送出,等待區塊確認.."
        this.cardloading = true
      })
      .on('receipt', (receipt) => {
        console.log('receipt', receipt)
      })
      .on('confirmation', async (confirmationNumber, receipt) => {
        this.snackbar = true
        this.cardloading = false
        this.notices = "區塊確認!"
        await this.$http.secured.get("/trigger_query/query").then(e => console.log).catch(e => console.error)
        console.log(confirmationNumber, receipt)
      })
      .on('error', (error) => {
        console.log(error)
        if(error.message.match("User denied transaction signature")){
          this.snackbar = true
          this.cardloading = false
          this.notices = `已取消請求`
        }else{
          this.snackbar = true
          this.cardloading = false
          this.notices = `區塊確認失敗.請檢查或是重試: ${error}`
        }
       // Like a typical promise, returns an error on rejection.
      })
    }
  }
}
</script>


<style lang="scss">
</style>
