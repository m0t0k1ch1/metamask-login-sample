'use strict';

let AppError = function(message, code = 0) {
  this.message = message;
  this.code    = code;
};
AppError.prototype.name = 'AppError';

module.exports = AppError;
