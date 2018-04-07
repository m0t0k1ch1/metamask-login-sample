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

          return web3.app.signTypedData([
            {
              type: 'string',
              name: 'message',
              value: 'Hi, Alice!',
            },
            {
              type: 'uint32',
              name: 'value',
              value: 42,
            },
          ], accounts[0]) // TODO
        })
        .then((result) => {
          $this.$message(result) // TODO
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

window.addEventListener('load', () => {
  // Is MetaMask installed?
  if (typeof web3 === 'undefined') {
    app.warn('Please install MetaMask')
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

  app.isLoginButtonDisabled = false
})
