export default {
  Connect({state, commit, dispatch}){
    if ('WebSocket' in window) {
      if(state.conn == undefined){
        let conn = new WebSocket(state.wsurl)
        conn.onopen = function(){ console.log('websocket connected') }
        conn.onmessage = function(e){
          console.log('receviedMessage', e.data)
          dispatch('receviedMessage', e.data)
        }
        conn.onclose = () => {
          commit('setConenction', undefined)
          commit('setMessage', undefined)
          dispatch('reconnectWs')
        }
        conn.onerror = (err) => console.err
        commit('setConenction', conn)
      } else {
        console.log("connection alreay connected")
      }
    } else {
      console.log("not support websocket")
    }
  },
  receviedMessage({state, commit}, msg){
    commit('setMessage', msg)
  },
  reconnectWs({dispatch, state}){
    console.log("reconnect to websocket")
    if(state.reconenct){
      setTimeout(() => {
        dispatch('Connect')
      }, state.waitTime);
    }
  }
}
