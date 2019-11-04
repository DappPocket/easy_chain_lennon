export default {
  setPerPage({commit}, per_page){
    commit('setPerPage', per_page)
  },
  setPagging({commit}, page){
    commit('setPagging', page)
  },
  plusPagging({commit, state}){
    const page = state.currentpage+1
    commit('setPagging', page)
  },
  dividePagging({commit, state}){
    const page = state.currentpage-1
    if(page > 0){
      commit('setPagging', page)
    }
  },
  setComments({commit}, comments){
    commit('setComments', comments)
  },
  addOneComments({commit, state}, comment){
    // filter hash that already added
    const findMatched = _.filter(state.comments, (c) => comment.hash === c.hash)
    // if fined matched result, do nothing here
    if(findMatched.length === 0){
      commit('appendComments', comment)
    }
  },
  checkAndUpdateComments({commit, state, getters}, newcomment){
    let hashList = getters.getHashList
    // find data not yet put into list
    let newdata = _.filter(newcomment, (d) => {
      return !_.includes(hashList, d["hash"])
    })
    for(var i=0; i<state.comments.length; i++){
      // update data that comfired block
      if(state.comments[i].block_number == 0){
        let matchedresult = _.filter(newcomment, (e) => e["hash"] === state.comments[i]["hash"])
        if(matchedresult && matchedresult.length !== 0){
          commit('updateCommentByIndex', {indx: i, data: getters.convertData(matchedresult[0])})
          if(matchedresult[0]["block_number"] != 0){
            commit('setBlockwsMessage', `${state.comments[i]["hash"]} - 交易已確認`)
          } else {
            commit('setBlockwsMessage', `${state.comments[i]["hash"]} - 等待確認`)
          }
        }
      }
    }
    // merge new data
    if(newdata && newdata.length !== 0){
      _.each(newdata, (o) => {
        const convertedData = getters.convertData(o)
        commit('appendComments', convertedData)
        commit('setBlockwsMessage', `列表已更新`)
      })
    }
  }
}
