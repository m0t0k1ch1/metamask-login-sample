function AppError(message) {
  this.message = message
}
Object.setPrototypeOf(AppError, Error)
AppError.prototype = Object.create(Error.prototype)
AppError.prototype.name = "AppError"
AppError.prototype.message = ""
AppError.prototype.constructor = AppError

let app = new Vue({
  el: '#app',
  data: {
    isLoginButtonDisabled: true,
  },
  created: function () {
    // Is MetaMask installed?
    if (typeof web3 === 'undefined') {
      this.warn('Please install MetaMask')
      return
    }

    window.web3 = new Web3(web3.currentProvider)
    window.web3.extend({
      property: 'app',
      methods: [{
        name: 'signTypedData',
        call: 'eth_signTypedData',
        params: 2,
      }],
    })

    this.isLoginButtonDisabled = false
  },
  methods: {
    login: function() {
      let $this = this

      let accounts = []

      web3.eth.getAccounts()
        .then((result) => {
          // Are there available accounts?
          if (result.length <= 0) {
            throw new AppError('Please unlock MetaMask account')
          }

          accounts = result

          return web3.eth.net.getId()
        })
        .then((result) => {
          // Does MetaMask connect to Ropeten?
          if (result !== 3) {
            throw new AppError('Please connect MetaMask to Ropsten Test Network')
          }

          let params = new URLSearchParams()
          params.append('address', accounts[0])

          return axios.post('/challenge', params)
        })
        .then((result) => {
          let data = result.data
          if (data.state === "error") {
            throw new AppError(data.result.message)
          }

          let typedData = [{
            type: 'string',
            name: 'message',
            value: data.result.challenge,
          }]

          return web3.app.signTypedData(typedData, accounts[0])
        })
        .then((result) => {
          let params = new URLSearchParams()
          params.append('address', accounts[0])
          params.append('signature', result)

          return axios.post('/login', params)
        })
        .then((result) => {
          console.log(result)
        })
        .catch((e) => {
          if (e instanceof AppError) {
            $this.warn(e.message)
          }
          else if (e.message.match(/User denied message signature\./)) {
            $this.warn('Please accept the signature request')
          }
          else {
            throw e
          }
        })
    },
    warn: function(message) {
      this.$message({
        showClose: true,
        message: message,
        type: 'warning',
      })
    },
  },
})
