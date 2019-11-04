<template>
  <v-app>
    <v-app-bar app>
      <v-toolbar-title class="headline text-uppercase">
        <router-link :to="{name: 'Comments'}" v-slot="{href}">
          <a :href="href">
            <span>鏈上連儂牆</span>
          </a>
        </router-link>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <router-link :to="{name: 'EthForm'}" v-slot="{href}">
        <v-btn id="gcreate_note" :href="href" color="primary" fab small dark>
          <v-icon>mdi-plus</v-icon>
        </v-btn>
      </router-link>
    </v-app-bar>
    <v-dialog v-model="dialog" persistent>
      <v-card>
        <v-card-title class="headline">鏈上連儂牆</v-card-title>
        <v-card-text>
          <p class="v-card-text_subject">
            在這邊你可以暢所欲言。不論是想評論時事、想跟喜歡的對象告白、想吐槽公司主管，或是對於區塊鏈的未來有什麼看法，都歡迎你來留言！
          </p>

          <p class="notice-head">
            <v-icon>mdi-alert-octagon</v-icon>
            注意事項：
          </p>
          <p>1. 切勿人身攻擊及不雅字眼，鏈上連儂牆主張暢所欲言，但如有人身攻擊及不雅字眼，可能還會予以屏蔽（不過雖然被屏蔽，鏈上還是可以查得到所有留言，只是不會在瀏覽器上顯示）</p>
          <p>2. 此連儂牆每則留言資訊皆會上傳至以太坊區塊鏈上。</p>
          <p>3. 點擊留言的下方勾勾，即會顯示此留言在以太坊上的狀態。</p>
          <p>4. 每則留言都需要以太坊<v-icon>mdi-ethereum</v-icon>的燃料費（目前來說不到3塊錢台幣<v-icon>mdi-coin-outline</v-icon>）。</p>
          <p>5. 有任何使用上的問題以及意見，可以至鏈上生活粉專詢問（你要直接在牆上留言給意見也沒問題！）。</p>
          <p>6. 此牆留言與鏈上生活立場無關，純屬個人意見及立場。</p>
        </v-card-text>
        <v-card-actions>
          <v-btn color="success notice-ok" text @click="dialog = false">
            <v-icon>mdi-gavel</v-icon>
            OK
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-content>
      <router-view></router-view>
    </v-content>
    <v-snackbar v-model="snackbar" :top="true" :left="true" :multi-line="true">
      {{blockwsMessage}}
    </v-snackbar>
  </v-app>
</template>

<script>
import {mapGetters, mapState, mapActions} from 'vuex';
export default {
  name: 'App',
  created(){
    this.Connect()
  },
  data: () => ({
    dialog: true,
    snackbar: false,
  }),
  methods: {
    ...mapActions('ws', ['Connect']),
    ...mapActions('global', ['checkAndUpdateComments']),
  },
  computed: {
    ...mapGetters('ws', ['CheckMessage']),
    ...mapState('global', ['blockwsMessage']),
  },
  watch: {
    CheckMessage(data){
      console.log('CheckMessage', data)
      this.checkAndUpdateComments(data)
    },
    blockwsMessage(d){
      this.snackbar = true
    }
  }
};
</script>

<style lang="scss">
header.v-sheet.v-sheet--tile {
  background-color: #083140;
  .v-toolbar__title.headline {
    a {
      font-weight: bold;
      color: white;
      font-family: 'Noto Sans JP', sans-serif !important;
      text-decoration: none;
    }
  }
  #gcreate_note {
    background-color: #feb904 !important;
    border-color: #feb904 !important;
  }
}
</style>
