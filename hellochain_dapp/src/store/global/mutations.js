export default {
  setPerPage(state, per_page){
    state.per_page = per_page
  },
  setPagging(state, page){
    state.currentpage = page
  },
  setComments(state, comments){
    state.comments = comments
  },
  appendComments(state, comment){
    state.comments.push(comment)
  },
  updateCommentByIndex(state, payload){
    let comments = _.filter(state.comments, (v,i) => i !== payload["indx"])
    comments.push(payload["data"])
    state.comments = comments
  },
  setBlockwsMessage(state, msg){
    state.blockwsMessage = msg
  }
}
