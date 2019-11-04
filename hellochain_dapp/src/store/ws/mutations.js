export default {
  setConenction(state, conn){
    state.conn = conn
  },
  setMessage(state, msg){
    if(msg !== null && msg !== undefined && msg !== ""){
      try {
        let parsedData = JSON.parse(msg)
        state.message = parsedData
      } catch(e) {
        console.error(e)
      }
    }
  }
}
