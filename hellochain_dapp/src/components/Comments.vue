<template lang="pug">
v-container(fluid)
  v-layout(wrap)
    v-row
      v-flex.xs12.sm6.mb3.lg2.card-flex-col(:key="i" v-for="co, i in comments")
        a(:href="etherscanLink(co.hash)" class="etherscan_link" target="_blank")
          v-card.post-card(:hover="true" :id="co.bg_color")
            v-list-item
              v-list-item-content
                .overline.mb-4
                  .v-list-datetime {{co.timestamp}}
                v-list-item-title(:id="vListItemTitleFontSize(co.message)") {{co.message}}
            .v-list-item-block__status
              span(class="confirmed_blcok_status" v-if="co.confirmed_status")
                v-icon mdi-check-circle
              span(class="pending_blcok_status" v-if="!co.confirmed_status")
                v-icon mdi-alert-circle
    .no-data__display(v-if="!loading && (!comments || comments.length === 0)")
      div 無資料
    div(v-if="loading")
      v-progress-circular(indeterminate color="primary" class="bottom")
</template>

<script>
import _ from 'lodash';
import moment from 'moment';
import {mapGetters, mapState, mapActions} from 'vuex';

export default {
  name: "Comments",
  data: () => {
    return {
      bottom: false,
      loading: false,
    }
  },
  created() {
    window.addEventListener('scroll', () => {
      console.log("scroll")
      this.bottom = this.bottomVisible()
    })
    this.addComment()
  },
  mounted() {
  },
  methods: {
    compute_data(e){
      return _.chain(e.data).map(o => {
        let msg = ""
        try {
          let tmp = this.hexToUtf8(o.input)
          msg = tmp
        }catch(e){
          console.log(e)
        }
        return this.convertData(o, msg)
      }).filter(o => {
        return o.message && o.message != ""
      }).value()
    },
    bottomVisible() {
      return Math.ceil(window.innerHeight + window.scrollY) >= document.body.offsetHeight
    },
    addComment(){
      this.loading = true
      this.$http.secured.get(`/transactions?page=${this.page}&per_page=${this.per_page}`)
      .then(async e => {
        let data = await this.compute_data(e)
        if(data && data.length !== 0){
          this.plusPagging()
          let comments = this.comments || []
          _.each(data, d => comments.push(d))
          this.comments = comments
          if (this.bottomVisible()) {
            this.addComment()
          }
        }
        this.loading = false
      })
      .catch(e => {
        this.loading = false
        console.error(e)
      })
    },
    vListItemTitleFontSize(msg){
      if(msg.length < 10){
        return "text-lg"
      } else if (msg.length < 25){
        return "text-md"
      } else {
        return "text-sm"
      }
    },
    ...mapActions('global', [
      'setComments', 'setPerPage', 'setPagging', 'plusPagging'
    ])
  },
  computed: {
    comments: {
      get() {
        return this.getComments
      },
      async set(comments) {
        await this.setComments(comments)
      },
    },
    ...mapState('global', {
      page: (state) => state.currentpage,
      per_page: (state) => state.per_page,
    }),
    ...mapGetters('global', ['hexToUtf8', 'getComments', 'etherscanLink', 'pickColorByNoce', 'convertData']),
  },
  watch:{
    bottom(bottom) {
      if (bottom) {
        this.addComment()
      }
    }
  }
}
</script>

<style lang="scss">
.v-content__wrap {
  background-color: #062737;
  .etherscan_link {
    text-decoration: none;
    cursor: pointer;
    margin: auto;
  }
  .card-flex-col {
    padding: .5rem;
  }
  .v-card.post-card {
    position: relative;
    text-decoration: none;
    min-height: 260px;
    &#is-odd {
      background-color: #ffd825;
    }
    &#is-even {
      background-color: #fdeb7f;
    }
    &#is-one {
      background-color: #fdeb7f;
    }
    &#is-two {
      background-color: #aad7f8;
    }
    &#is-three {
      background-color: #d9ef8f;
    }
    &#is-four {
      background-color: #f8b0e0;
    }
    .v-list-item__content {
      .v-list-datetime {
        font-size: .8rem !important;
        color: #466b91;
      }
      .v-list-item__title {
        white-space: normal;
        &#text-sm {
          font-size: 1.8rem;
          @media (max-width: 992px) {
            font-size: 1.8rem*1.5;
          }
          @media (max-width: 768px) {
            font-size: 1.8rem*2;
          }
        }
        &#text-md {
          font-size: 2.3rem;
          @media (max-width: 992px) {
            font-size: 2.3rem*1.5;
          }
          @media (max-width: 768px) {
            font-size: 2.3rem*2;
          }
        }
        &#text-lg {
          font-size: 3rem;
          @media (max-width: 992px) {
            font-size: 3rem*1.5;
          }
          @media (max-width: 768px) {
            font-size: 3rem*2;
          }
        }
      }
    }
    .v-list-item-block__status {
      position: absolute;
      bottom: 0px;
      right: 0px;
      padding-right: .5rem;
      padding-bottom: .5rem;
      .confirmed_blcok_status i {
        color: #578853;
        font-size: 2rem;
      }
      .pending_blcok_status i {
        color: orange;
        font-size: 2rem;
      }
    }
  }
  .no-data__display {
    color: white;
    text-align: center;
    width: 100%;
  }
}
</style>
