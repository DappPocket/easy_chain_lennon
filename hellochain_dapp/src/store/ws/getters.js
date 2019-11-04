export default {
  CheckMessage: (state) => {
    let data = state.message
    if(data && data.length !== 0){
      for(var i=0;i<=data.length;i++){
        let datatmp = data[i]
        if(datatmp == undefined || datatmp.topic != 'datafeed'){
          continue
        }
        return JSON.parse(datatmp.data)
      }
    }
    return undefined
  }
}
