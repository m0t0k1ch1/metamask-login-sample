function AppError(message) {
  this.message = message
}
Object.setPrototypeOf(AppError, Error)
AppError.prototype = Object.create(Error.prototype)
AppError.prototype.name = "AppError"
AppError.prototype.message = ""
AppError.prototype.constructor = AppError

let client = axios.create()
client.interceptors.response.use((response) => {
  let data = response.data
  if (data.state === "error") {
    return Promise.reject(new AppError(data.result.message))
  }
  return response
})

new Vue({
  el: '#app',
  data: {
    isLoginButtonDisabled: true,
    user: null,
  },
  created: function() {
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

          return client.post('/auth/challenge', params)
        })
        .then((response) => {
          let result = response.data.result

          let typedData = [{
            type: 'string',
            name: 'challenge',
            value: response.data.result.challenge,
          }]

          return web3.app.signTypedData(typedData, accounts[0])
        })
        .then((result) => {
          let params = new URLSearchParams()
          params.append('address', accounts[0])
          params.append('signature', result)

          return client.post('/auth/authorize', params)
        })
        .then((response) => {
          let result = response.data.result

          client.defaults.headers.common['Authorization'] = 'Bearer ' + result.token

          return client.get('/api/users/' + accounts[0])
        })
        .then((response) => {
          let result = response.data.result

          $this.user = result
        })
        .catch((e) => {
          $this.handleError(e)
        })
    },
    updateUser: function() {
      let $this = this

      let params = new URLSearchParams()
      params.append('name', $this.user.name)

      client.put('/api/users/' + $this.user.address, params)
        .then((response) => {
          $this.info('success')
        })
        .catch((e) => {
          $this.handleError(e)
        })
    },
    deleteUser: function() {
      let $this = this

      client.delete('/api/users/' + $this.user.address)
        .then((response) => {
          $this.logout()
        })
        .catch((e) => {
          $this.handleError(e)
        })
    },
    handleError: function(e) {
      if (e instanceof AppError) {
        this.warn(e.message)
      }
      else if (e.message.match(/User denied message signature\./)) {
        this.warn('Please accept the signature request')
      }
      else {
        throw e
      }
    },
    info: function(message) {
      this.$message({
        showClose: true,
        message: message,
        type: 'info',
      })
    },
    warn: function(message) {
      this.$message({
        showClose: true,
        message: message,
        type: 'warning',
      })
    },
    logout: function() {
      delete client.defaults.headers.common['Authorization']
      this.user = null
    },
  },
})
