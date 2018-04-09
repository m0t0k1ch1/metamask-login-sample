function AppError(message) {
  this.message = message
}
Object.setPrototypeOf(AppError, Error)
AppError.prototype = Object.create(Error.prototype)
AppError.prototype.name = "AppError"
AppError.prototype.message = ""
AppError.prototype.constructor = AppError
